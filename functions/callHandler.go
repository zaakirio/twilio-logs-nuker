package Handlers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/twilio/twilio-go"
)

func DeleteCallLogs() {
	client := InitClient()
	callRecords := getCallRecords(client)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for _, call := range callRecords {
		select {
		case <-interrupt:
			fmt.Println("Interrupt received. Cancelling deletion...")
			close(interrupt)
			return
		default:
			err := client.Api.DeleteCall(*call.Sid, &twilioApi.DeleteCallParams{})
			if err != nil {
				fmt.Printf("Error deleting call log %s: %v\n", *call.Sid, err)
			} else {
				fmt.Printf("Deleted call log %s\n", *call.Sid)
			}
		}
	}
}

func getCallRecords(client *twilio.RestClient) []twilioApi.ApiV2010Call {
	parameters := &twilioApi.ListCallParams{}
	parameters.SetPageSize(1000)
	callRecords, err := client.Api.ListCall(parameters)
	checkError(err)
	return callRecords
}
