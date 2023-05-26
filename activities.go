package signalwkfldata

import (
	"context"

	"go.temporal.io/sdk/activity"
)

type Activities struct {
}

func (a *Activities) SendWelcomeEmail(ctx context.Context, customer Customer) (string, error) {
        activity.GetLogger(ctx).Info(ColorYellow, "Activity-SendWelcomeEmail:", 
		ColorBlue, "Sending welcome email to customer", customer.Id, ColorReset)
	return "Sending welcome email completed for " + customer.Id, nil
}

func (a *Activities) ChargeCustomerForBillingPeriod(ctx context.Context, customer Customer, billingPeriodNum int) (string, error) {
        activity.GetLogger(ctx).Info(ColorYellow, "Activity-ChargeCustomerForBillingPeriod:", 
		ColorBlue, "Charging for billing period", billingPeriodNum, "for customer", customer.Id, ColorReset)
	return "Charging for billing period completed for " + customer.Id, nil
}

func (a *Activities) SendCancellationEmail(ctx context.Context, customer Customer) (string, error) {
	activity.GetLogger(ctx).Info(ColorYellow, "Activity-SendCancellationEmail:",
		ColorBlue, "Sending cancellation email to customer", customer.Id, ColorReset)
	return "Sending cancellation email to customer " + customer.Id, nil
}

func (a *Activities) SendSubscriptionEndedEmail(ctx context.Context, customer Customer) (string, error) {
	activity.GetLogger(ctx).Info(ColorYellow, "Activity-SendSubscriptionEndedEmail:",
		ColorBlue, "Sending subscription ended email to customer", customer.Id, ColorReset)
	return "Sending subscription ended email completed for " + customer.Id, nil
}
