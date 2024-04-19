package Handlers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/twilio/twilio-go"
	studio "github.com/twilio/twilio-go/rest/studio/v2"
)

func DeleteFlowLogs() {
	client := InitClient()
	flowRecords := getFlowRecords(client)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for _, flow := range flowRecords {
		select {
		case <-interrupt:
			fmt.Println("Interrupt received. Cancelling deletion...")
			close(interrupt)
			return
		default:
			err := client.StudioV2.DeleteFlow(*flow.Sid)
			if err != nil {
				fmt.Printf("Error deleting flow log %s: %v\n", *flow.Sid, err)
			} else {
				fmt.Printf("Deleted flow log %s\n", *flow.Sid)
			}
		}
	}
}

func getFlowRecords(client *twilio.RestClient) []studio.StudioV2Flow {
	parameters := &studio.ListFlowParams{}
	parameters.SetPageSize(1000)
	flowRecords, err := client.StudioV2.ListFlow(parameters)
	checkError(err)
	return flowRecords
}
