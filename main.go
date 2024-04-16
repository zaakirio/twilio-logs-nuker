package main

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func initClient() *twilio.RestClient {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return client
}

func main() {
	fmt.Println("Twilio Log Deletion Tool")
	fmt.Println("----------------------")

	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Delete Message Logs (by phone number)")
		fmt.Println("2. Delete Studio Flow Logs (by Studio Flow SID)")
		fmt.Println("3. Delete Conversation Logs (by phone number)")
		fmt.Println("4. Delete Call Logs (by phone number)")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			deleteMessageLogs()
		case 2:
			deleteStudioFlowLogs()
		case 3:
			deleteConversationLogs()
		case 4:
			deleteCallLogs()
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func deleteMessageLogs() {
}

func deleteStudioFlowLogs() {
}

func deleteConversationLogs() {
}

func deleteCallLogs() {
	client := initClient()
	callRecords := GetCallRecords(client)
	for _, call := range callRecords {
		err := client.Api.DeleteCall(*call.Sid, &twilioApi.DeleteCallParams{})
		if err != nil {
			fmt.Printf("Error deleting call log %s: %v\n", *call.Sid, err)
		} else {
			fmt.Printf("Deleted call log %s\n", *call.Sid)
		}
	}

	fmt.Println("Call logs deleted successfully.")
}

func GetCallRecords(client *twilio.RestClient) []twilioApi.ApiV2010Call {
	parameters := &twilioApi.ListCallParams{}
	parameters.SetPageSize(1000)
	callRecords, err := client.Api.ListCall(parameters)
	checkError(err)
	return callRecords
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
