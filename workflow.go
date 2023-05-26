package signalwkfldata

import (
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

/* SignalWorkflow */
func SignalWorkflow(ctx workflow.Context, customer Customer) (string, error) {

	workflowCustomer := customer
	subscriptionCancelled := false
	billingPeriodNum := 1
	actResult := ""

	QueryCustomerIdName := "customerid"
	QueryBillingPeriodNumberName := "billingperiodnumber"
	QueryBillingPeriodChargeAmountName := "billingperiodchargeamount"
	QuerySubscriptionActive := "subscriptionactive"

	logger := workflow.GetLogger(ctx)

	logger.Info(ColorGreen, "SignalWorkflow started for customer:", customer.Id, ColorReset)

	// Define query handlers

	//   QueryCustomerIdName
	err := workflow.SetQueryHandler(ctx, QueryCustomerIdName, func() (string, error) {
		logger.Info(ColorGreen, "SignalWorkflow:", ColorCyan, "Received Query - QueryCustomerIdName:",
			workflowCustomer.Id, ColorReset)
		return workflowCustomer.Id, nil
	})
	if err != nil {
		logger.Info("Workflow: SetQueryHandler: QueryCustomerIdName handler failed.", "Error", err)
		return "Error", err
	}

	//   QueryBillingPeriodNumberName
	err = workflow.SetQueryHandler(ctx, QueryBillingPeriodNumberName, func() (int, error) {
		logger.Info(ColorGreen, "SignalWorkflow:", ColorCyan, "Received Query - QueryBillingPeriodNumberName:",
			billingPeriodNum, ColorReset)
		return billingPeriodNum, nil
	})
	if err != nil {
		logger.Info("Workflow: SetQueryHandler: QueryBillingPeriodNumberName handler failed.", "Error", err)
		return "Error", err
	}

	//   QueryBillingPeriodChargeAmountName
	err = workflow.SetQueryHandler(ctx, QueryBillingPeriodChargeAmountName, func() (int, error) {
		logger.Info(ColorGreen, "SignalWorkflow:", ColorCyan, "Received Query - QueryBillingPeriodChargeAmountName:",
			workflowCustomer.Subscription.BillingPeriodCharge, ColorReset)
		return workflowCustomer.Subscription.BillingPeriodCharge, nil
	})
	if err != nil {
		logger.Info("Workflow: SetQueryHandler: QueryBillingPeriodChargeAmountName handler failed.", "Error", err)
		return "Error", err
	}

	//   QuerySubscriptionActive
	err = workflow.SetQueryHandler(ctx, QuerySubscriptionActive, func() (bool, error) {
		logger.Info(ColorGreen, "SignalWorkflow:", ColorCyan, "Received Query - QuerySubscriptionActive:",
			workflowCustomer.Subscription.Active, ColorReset)
		return workflowCustomer.Subscription.Active, nil
	})
	if err != nil {
		logger.Info("Workflow: SetQueryHandler: QuerySubscriptionActive handler failed.", "Error", err)
		return "Error", err
	}
	// end defining query handlers

	// Define signal channels

	//   billing period charge change signal
	chargeSelector := workflow.NewSelector(ctx)
	signalCh := workflow.GetSignalChannel(ctx, "billingperiodcharge")
	chargeSelector.AddReceive(signalCh, func(ch workflow.ReceiveChannel, _ bool) {
		// do this when signal received

		// read contents from signal
		var chargeSignal int
		ch.Receive(ctx, &chargeSignal)

		logger.Info(ColorGreen, "SignalWorkflow:", ColorCyan, "Received Signal - billingperiodcharge:",
			chargeSignal, ColorReset)

		// update workflow variable value
		workflowCustomer.Subscription.BillingPeriodCharge = chargeSignal
	})

	//   cancel subscription signal
	cancelSelector := workflow.NewSelector(ctx)
	cancelCh := workflow.GetSignalChannel(ctx, "cancelsubscription")
	cancelSelector.AddReceive(cancelCh, func(ch workflow.ReceiveChannel, _ bool) {
		// do this when signal received

		// read contents from signal
		var cancelSubSignal bool
		ch.Receive(ctx, &cancelSubSignal)

		logger.Info(ColorGreen, "SignalWorkflow:", ColorCyan, "Received Signal - cancelsubscription:",
			cancelSubSignal, ColorReset)

		// update workflow variable value
		subscriptionCancelled = cancelSubSignal
		workflowCustomer.Subscription.Active = false
	})
	// end defining signal channels

	// Activities
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	var activities *Activities

	// Send welcome email to customer
	err = workflow.ExecuteActivity(ctx, activities.SendWelcomeEmail, workflowCustomer).Get(ctx, &actResult)
	if err != nil {
		log.Fatalln("Failure executing SendWelcomeEmail", err)
	}

	// Start billing loop until we reach the max billing periods for the subscription or has been cancelled
	for workflowCustomer.Subscription.Active {

		// Wait 1 billing period to charge customer or if they cancel subscription
		// whichever comes first
		logger.Info(ColorGreen, "SignalWorkflow: Sleeping for a BillingPeriod (", 
			workflowCustomer.Subscription.BillingPeriod, ")..", ColorReset)

		//workflow.AwaitWithTimeout(ctx, workflowCustomer.Subscription.BillingPeriod, cancelSelector.HasPending)
		workflow.Sleep(ctx, workflowCustomer.Subscription.BillingPeriod)

		// Check if cancel signal received during period
		for cancelSelector.HasPending() {
			cancelSelector.Select(ctx)
		}
		logger.Info(ColorGreen, "SignalWorkflow: Period:", billingPeriodNum, "Customer:", 
			customer.Id, "(Active:", workflowCustomer.Subscription.Active, ", Amount:", 
			workflowCustomer.Subscription.BillingPeriodCharge, ")", ColorReset)

		// Check if cancelled while in timer
		if subscriptionCancelled {
			logger.Info(ColorGreen, "SignalWorkflow: Cancelled for customer", customer.Id, ColorReset)
			workflowCustomer.Subscription.Active = false

			// send a cancellation email to customer
			err = workflow.ExecuteActivity(ctx, activities.SendCancellationEmail, workflowCustomer).Get(ctx, &actResult)
			if err != nil {
				log.Fatalln("Failure executing SendCancellationEmail", err)
			}
			break

		} else {
			// Subscription still active:
			// Charge customer for end of current billing period
			err = workflow.ExecuteActivity(ctx, activities.ChargeCustomerForBillingPeriod, workflowCustomer, billingPeriodNum).Get(ctx, &actResult)
			if err != nil {
				log.Fatalln("Failure executing ChargeCustomerForBillingPeriod", err)
			}

			if billingPeriodNum == workflowCustomer.Subscription.MaxBillingPeriods {
				workflowCustomer.Subscription.Active = false
			} else {
				// If period charge was changed, set new value for next period
				for chargeSelector.HasPending() {
					chargeSelector.Select(ctx)
				}

				// Increment to next billing period
				billingPeriodNum++
			}
		}
	}

	// Subscription period is now over or cancelled

	// notify the customer to buy a new subscription
	if !subscriptionCancelled {
		log.Printf("%sSignalWorkflow: Subscription Period Ended for customer %s%s\n", ColorGreen, customer.Id, ColorReset)
		err = workflow.ExecuteActivity(ctx, activities.SendSubscriptionEndedEmail, workflowCustomer).Get(ctx, &actResult)
		if err != nil {
			log.Fatalln("Failure executing SendSubscriptionEndedEmail", err)
		}
	}

	logger.Info(ColorGreen, "SignalWorkflow: Complete.", ColorReset)

	return "Completed Subscription Workflow", err
}
