package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"signalwkfldata"
)

func main() {
	log.Printf("%sGo worker starting..%s", signalwkfldata.ColorGreen, signalwkfldata.ColorReset)

	// The client and Worker are heavyweight objects that should be created once per process.
	clientOptions, err := signalwkfldata.LoadClientOption()
	if err != nil {
		log.Fatalf("Failed to load Temporal Cloud environment: %v", err)
	}
	c, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, signalwkfldata.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(signalwkfldata.SignalWorkflow)
	w.RegisterActivity(&signalwkfldata.Activities{})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

	log.Printf("%sGo worker stopped.%s", signalwkfldata.ColorGreen, signalwkfldata.ColorReset)
}
