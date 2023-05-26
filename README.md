# temporal-signal-workflow-data
Sample app to show use of Signals and Queries to access Workflow variables  

Based on code from [subscription go example](https://github.com/temporalio/subscription-workflow-go)


## Sample output:

Run full subscription term, 12 periods:
```c
$ go run worker/main.go
2023/05/26 14:27:01 Go worker starting..
2023/05/26 14:27:01 INFO  Started Worker
2023/05/26 14:27:08 INFO   SignalWorkflow started for customer: Id-1
2023/05/26 14:27:08 DEBUG ExecuteActivity ActivityID 5 ActivityType SendWelcomeEmail
2023/05/26 14:27:08 INFO   Activity-SendWelcomeEmail:  Sending welcome email to customer Id-1
2023/05/26 14:27:08 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:27:08 DEBUG NewTimer TimerID 11 Duration 30s
2023/05/26 14:27:35 INFO   SignalWorkflow:  Received Query - QueryBillingPeriodNumberName: 1
2023/05/26 14:27:35 INFO   SignalWorkflow:  Received Query - QueryBillingPeriodChargeAmountName: 20
2023/05/26 14:27:35 INFO   SignalWorkflow:  Received Query - QuerySubscriptionActive: true
2023/05/26 14:27:38 INFO   SignalWorkflow: Period: 1 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:27:38 DEBUG ExecuteActivity ActivityID 16 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:27:38 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 1 for customer Id-1
2023/05/26 14:27:38 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:27:38 DEBUG NewTimer TimerID 22 Duration 30s
2023/05/26 14:28:08 INFO   SignalWorkflow: Period: 2 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:28:08 DEBUG ExecuteActivity ActivityID 31 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:28:08 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 2 for customer Id-1
2023/05/26 14:28:09 INFO   SignalWorkflow:  Received Signal - billingperiodcharge: 25
2023/05/26 14:28:09 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:28:09 DEBUG NewTimer TimerID 37 Duration 30s
2023/05/26 14:28:39 INFO   SignalWorkflow: Period: 3 Customer: Id-1 (Active: true , Amount: 25 )
2023/05/26 14:28:39 DEBUG ExecuteActivity ActivityID 42 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:28:39 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 3 for customer Id-1
2023/05/26 14:28:39 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:28:39 DEBUG NewTimer TimerID 48 Duration 30s
2023/05/26 14:29:09 INFO   SignalWorkflow:  Received Signal - cancelsubscription: true
2023/05/26 14:29:09 INFO   SignalWorkflow: Period: 4 Customer: Id-1 (Active: false , Amount: 25 )
2023/05/26 14:29:09 INFO   SignalWorkflow: Cancelled for customer Id-1
2023/05/26 14:29:09 DEBUG ExecuteActivity ActivityID 57 ActivityType SendCancellationEmail
2023/05/26 14:29:09 INFO   Activity-SendCancellationEmail:  Sending cancellation email to customer Id-1
2023/05/26 14:29:09 INFO   SignalWorkflow: Complete.
2023/05/26 14:29:33 INFO   SignalWorkflow started for customer: Id-1
2023/05/26 14:29:33 DEBUG ExecuteActivity ActivityID 5 ActivityType SendWelcomeEmail
2023/05/26 14:29:33 INFO   Activity-SendWelcomeEmail:  Sending welcome email to customer Id-1
2023/05/26 14:29:33 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:29:33 DEBUG NewTimer TimerID 11 Duration 30s
2023/05/26 14:30:03 INFO   SignalWorkflow: Period: 1 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:30:03 DEBUG ExecuteActivity ActivityID 16 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:30:03 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 1 for customer Id-1
2023/05/26 14:30:03 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:30:03 DEBUG NewTimer TimerID 22 Duration 30s
2023/05/26 14:30:33 INFO   SignalWorkflow: Period: 2 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:30:33 DEBUG ExecuteActivity ActivityID 27 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:30:33 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 2 for customer Id-1
2023/05/26 14:30:34 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:30:34 DEBUG NewTimer TimerID 33 Duration 30s
2023/05/26 14:31:04 INFO   SignalWorkflow: Period: 3 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:31:04 DEBUG ExecuteActivity ActivityID 38 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:31:04 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 3 for customer Id-1
2023/05/26 14:31:04 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:31:04 DEBUG NewTimer TimerID 44 Duration 30s
2023/05/26 14:31:34 INFO   SignalWorkflow: Period: 4 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:31:34 DEBUG ExecuteActivity ActivityID 49 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:31:34 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 4 for customer Id-1
2023/05/26 14:31:34 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:31:34 DEBUG NewTimer TimerID 55 Duration 30s
2023/05/26 14:32:04 INFO   SignalWorkflow: Period: 5 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:32:04 DEBUG ExecuteActivity ActivityID 60 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:32:04 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 5 for customer Id-1
2023/05/26 14:32:04 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:32:04 DEBUG NewTimer TimerID 66 Duration 30s
2023/05/26 14:32:35 INFO   SignalWorkflow: Period: 6 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:32:35 DEBUG ExecuteActivity ActivityID 71 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:32:35 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 6 for customer Id-1
2023/05/26 14:32:35 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:32:35 DEBUG NewTimer TimerID 77 Duration 30s
2023/05/26 14:33:05 INFO   SignalWorkflow: Period: 7 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:33:05 DEBUG ExecuteActivity ActivityID 82 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:33:05 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 7 for customer Id-1
2023/05/26 14:33:05 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:33:05 DEBUG NewTimer TimerID 88 Duration 30s
2023/05/26 14:33:35 INFO   SignalWorkflow: Period: 8 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:33:35 DEBUG ExecuteActivity ActivityID 93 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:33:35 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 8 for customer Id-1
2023/05/26 14:33:35 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:33:35 DEBUG NewTimer TimerID 99 Duration 30s
2023/05/26 14:34:06 INFO   SignalWorkflow: Period: 9 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:34:06 DEBUG ExecuteActivity ActivityID 104 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:34:06 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 9 for customer Id-1
2023/05/26 14:34:06 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:34:06 DEBUG NewTimer TimerID 110 Duration 30s
2023/05/26 14:34:36 INFO   SignalWorkflow: Period: 10 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:34:36 DEBUG ExecuteActivity ActivityID 115 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:34:36 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 10 for customer Id-1
2023/05/26 14:34:36 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:34:36 DEBUG NewTimer TimerID 121 Duration 30s
2023/05/26 14:35:06 INFO   SignalWorkflow: Period: 11 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:35:06 DEBUG ExecuteActivity ActivityID 126 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:35:06 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 11 for customer Id-1
2023/05/26 14:35:06 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:35:06 DEBUG NewTimer TimerID 132 Duration 30s
2023/05/26 14:35:36 INFO   SignalWorkflow: Period: 12 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 14:35:36 DEBUG ExecuteActivity ActivityID 137 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 14:35:37 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 12 for customer Id-1
2023/05/26 14:35:37 SignalWorkflow: Subscription Period Ended for customer Id-1
2023/05/26 14:35:37 DEBUG ExecuteActivity ActivityID 143 ActivityType SendSubscriptionEndedEmail
2023/05/26 14:35:37 INFO   Activity-SendSubscriptionEndedEmail:  Sending subscription ended email to customer Id-1
2023/05/26 14:35:37 INFO   SignalWorkflow: Complete.
```

Run with update amount signal, query and then cancel early:
```c
2023/05/26 14:59:30 INFO   SignalWorkflow started for customer: Id-1
2023/05/26 14:59:30 DEBUG ExecuteActivity ActivityID 5 ActivityType SendWelcomeEmail
2023/05/26 14:59:30 INFO   Activity-SendWelcomeEmail:  Sending welcome email to customer Id-1
2023/05/26 14:59:30 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 14:59:30 DEBUG NewTimer TimerID 11 Duration 30s
2023/05/26 15:00:00 INFO   SignalWorkflow: Period: 1 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 15:00:00 DEBUG ExecuteActivity ActivityID 16 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 15:00:00 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 1 for customer Id-1
2023/05/26 15:00:00 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 15:00:00 DEBUG NewTimer TimerID 22 Duration 30s
2023/05/26 15:00:02 INFO   SignalWorkflow:  Received Query - QueryBillingPeriodNumberName: 2
2023/05/26 15:00:03 INFO   SignalWorkflow:  Received Query - QueryBillingPeriodChargeAmountName: 20
2023/05/26 15:00:03 INFO   SignalWorkflow:  Received Query - QuerySubscriptionActive: true
2023/05/26 15:00:30 INFO   SignalWorkflow: Period: 2 Customer: Id-1 (Active: true , Amount: 20 )
2023/05/26 15:00:30 DEBUG ExecuteActivity ActivityID 31 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 15:00:30 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 2 for customer Id-1
2023/05/26 15:00:31 INFO   SignalWorkflow:  Received Signal - billingperiodcharge: 25
2023/05/26 15:00:31 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 15:00:31 DEBUG NewTimer TimerID 37 Duration 30s
2023/05/26 15:00:36 INFO   SignalWorkflow:  Received Query - QueryBillingPeriodNumberName: 3
2023/05/26 15:00:36 INFO   SignalWorkflow:  Received Query - QueryBillingPeriodChargeAmountName: 25
2023/05/26 15:00:37 INFO   SignalWorkflow:  Received Query - QuerySubscriptionActive: true
2023/05/26 15:01:01 INFO   SignalWorkflow: Period: 3 Customer: Id-1 (Active: true , Amount: 25 )
2023/05/26 15:01:01 DEBUG ExecuteActivity ActivityID 42 ActivityType ChargeCustomerForBillingPeriod
2023/05/26 15:01:01 INFO   Activity-ChargeCustomerForBillingPeriod:  Charging for billing period 3 for customer Id-1
2023/05/26 15:01:01 INFO   SignalWorkflow: Sleeping for a BillingPeriod ( 30s )..
2023/05/26 15:01:01 DEBUG NewTimer TimerID 48 Duration 30s
2023/05/26 15:01:31 INFO   SignalWorkflow:  Received Signal - cancelsubscription: true
2023/05/26 15:01:31 INFO   SignalWorkflow: Period: 4 Customer: Id-1 (Active: false , Amount: 25 )
2023/05/26 15:01:31 INFO   SignalWorkflow: Cancelled for customer Id-1
2023/05/26 15:01:31 DEBUG ExecuteActivity ActivityID 57 ActivityType SendCancellationEmail
2023/05/26 15:01:31 INFO   Activity-SendCancellationEmail:  Sending cancellation email to customer Id-1
2023/05/26 15:01:31 INFO   SignalWorkflow: Complete.
```

cli commands for above output:  
```c
[cli]$ ./start
2023/05/26 14:59:30 Started workflow WorkflowID SignalWorkflowId-1 RunID 0639e16b-6f83-44ef-b8fd-62115e3d762a

[cli]$ ./query
2023/05/26 15:00:03 Workflow: Id SignalWorkflowId-1
2023/05/26 15:00:03   Billing Period 2
2023/05/26 15:00:03   Billing Period Charge 20
2023/05/26 15:00:03   Subscription Active? true

[cli]$ ./update

[cli]$ ./query
2023/05/26 15:00:37 Workflow: Id SignalWorkflowId-1
2023/05/26 15:00:37   Billing Period 3
2023/05/26 15:00:37   Billing Period Charge 25
2023/05/26 15:00:37   Subscription Active? true

[cli]$ ./cancel

[cli]$ ./query
2023/05/26 15:02:28 Workflow: Id SignalWorkflowId-1
2023/05/26 15:02:28   Billing Period 4
2023/05/26 15:02:28   Billing Period Charge 25
2023/05/26 15:02:28   Subscription Active? false
```

Note: the query on a completed workflow gives the end values of the workflow variables.  Using temporal logger stops the log messages outputing on the reparsing the workflow to provide these query values to the remote client (cli)

