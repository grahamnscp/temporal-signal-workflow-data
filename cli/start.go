package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.temporal.io/sdk/client"

	"signalwkfldata"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	clientOptions, err := signalwkfldata.LoadClientOption()
        if err != nil {
                log.Fatalf("Failed to load Temporal Cloud environment: %v", err)
        }
        c, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// default subscription
	sub := signalwkfldata.Subscription{
			BillingPeriod:       time.Duration(signalwkfldata.PeriodLengthSeconds)*time.Second,
			MaxBillingPeriods:   12,
			BillingPeriodCharge: 20,
			Active:              true,
		}

	// create Workflow Execution for NumberOfCustomers
	for i := 1; i < (signalwkfldata.NumberOfCustomers+1); i++ {
		cust := signalwkfldata.Customer{
			FirstName:    "First Name-" + strconv.Itoa(i),
			LastName:     "Last Name-" + strconv.Itoa(i),
			Email:        "someemail-" + strconv.Itoa(i),
			Subscription: sub,
			Id:           "Id-" + strconv.Itoa(i),
		}

		workflowOptions := client.StartWorkflowOptions{
			ID:                 "SignalWorkflow" + cust.Id,
			TaskQueue:          signalwkfldata.TaskQueueName,
			WorkflowRunTimeout: time.Minute * 10,
		}

		we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, signalwkfldata.SignalWorkflow, cust)
		if err != nil {
			log.Fatalln("Unable to execute workflow", err)
		}

		log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	}
}

