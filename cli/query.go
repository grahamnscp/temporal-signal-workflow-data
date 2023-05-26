package main

import (
	"context"
	"log"
	"strconv"

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

	// Query all Workflow Executions to get current billing information
	for i := 1; i < signalwkfldata.NumberOfCustomers+1; i++ {
		bpnresp, err := c.QueryWorkflow(context.Background(), "SignalWorkflowId-"+strconv.Itoa(i), "", "billingperiodnumber")
		if err != nil {
			log.Fatalln("Unable to query workflow", err)
		}
		var result interface{}
		if err := bpnresp.Get(&result); err != nil {
			log.Fatalln("Unable to decode query result", err)
		}

		bpcresp, err := c.QueryWorkflow(context.Background(), "SignalWorkflowId-"+strconv.Itoa(i), "", "billingperiodchargeamount")
		if err != nil {
			log.Fatalln("Unable to query workflow", err)
		}
		var result2 interface{}
		if err := bpcresp.Get(&result2); err != nil {
			log.Fatalln("Unable to decode query result", err)
		}

		scresp, err := c.QueryWorkflow(context.Background(), "SignalWorkflowId-"+strconv.Itoa(i), "", "subscriptionactive")
		if err != nil {
			log.Fatalln("Unable to query workflow", err)
		}
		var result3 interface{}
		if err := scresp.Get(&result3); err != nil {
			log.Fatalln("Unable to decode query result", err)
		}

		log.Println("Workflow:", "Id", "SignalWorkflowId-"+strconv.Itoa(i))
		log.Println("  Billing Period", result)
		log.Println("  Billing Period Charge", result2)
		log.Println("  Subscription Active?", result3)
	}
}

