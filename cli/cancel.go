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

	// Signal all Workflow Executions and cancel the subscription
	for i := 1; i < signalwkfldata.NumberOfCustomers+1; i++ {
		err = c.SignalWorkflow(context.Background(),
			"SignalWorkflowId-"+strconv.Itoa(i), "", "cancelsubscription", true)
		if err != nil {
			log.Fatalln("Unable to signal workflow", err)
		}
	}
}

