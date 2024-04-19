package Handlers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/twilio/twilio-go"
)

func DeleteMessageLogs() {
	client := InitClient()
	MessageRecords := getMessageRecords(client)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for _, message := range MessageRecords {
		select {
		case <-interrupt:
			fmt.Println("Interrupt received. Cancelling deletion...")
			close(interrupt)
			return
		default:
			err := client.Api.DeleteMessage(*message.Sid, &twilioApi.DeleteMessageParams{})
			if err != nil {
				fmt.Printf("Error deleting Message log %s: %v\n", *message.Sid, err)
			} else {
				fmt.Printf("Deleted Message log %s\n", *message.Sid)
			}
		}
	}
}

func getMessageRecords(client *twilio.RestClient) []twilioApi.ApiV2010Message {
	parameters := &twilioApi.ListMessageParams{}
	parameters.SetPageSize(1000)
	MessageRecords, err := client.Api.ListMessage(parameters)
	checkError(err)
	return MessageRecords
}
