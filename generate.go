package main

import (
	"fmt"
	Handlers "twilio-logs-nuker/functions"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	conversations "github.com/twilio/twilio-go/rest/conversations/v1"
	studio "github.com/twilio/twilio-go/rest/studio/v2"
)

func Generate(from string, to string) {
	client := Handlers.InitClient()
	generateRandomCallLog(client, from, to)
	generateRandomConversationLog(client)
	generateRandomMessageLog(client, from, to)
	generateRandomStudioFlowLog(client, from, to)
}

func generateRandomCallLog(client *twilio.RestClient, from string, to string) {
	params := &openapi.CreateCallParams{}
	params.SetFrom(from)
	params.SetTo(to)
	params.SetTwiml("<Response><Say>Hello, this is a test call!</Say></Response>")
	call, err := client.Api.CreateCall(params)
	if err != nil {
		fmt.Println("Error creating call log:", err)
		return
	}
	fmt.Println("Created call log with SID:", call.Sid)
}

func generateRandomMessageLog(client *twilio.RestClient, from string, to string) {
	body := "Hello, this is a test message!"
	params := &openapi.CreateMessageParams{
		Body: &body,
		To:   &to,
		From: &from,
	}
	message, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error creating message log:", err)
		return
	}
	fmt.Println("Created message log with SID:", message.Sid)
}

func generateRandomConversationLog(client *twilio.RestClient) {
	params := &conversations.CreateConversationParams{}
	params.SetFriendlyName("Test Conversation")

	resp, err := client.ConversationsV1.CreateConversation(params)
	if err != nil {
		fmt.Println("Error creating conversation log:", err)
		return
	}
	fmt.Println("Created conversation log with SID:", resp.Sid)
}

func generateRandomStudioFlowLog(client *twilio.RestClient, from string, to string) {
	// Example: executing a Studio Flow
	// Fetch the Studio Flow SID from the Twilio Console
	params := &studio.ListFlowParams{}
	params.SetLimit(20)
	flows, err := client.StudioV2.ListFlow(params)
	if err != nil {
		fmt.Println("Error fetching studio flows:", err)
	}

	for _, flow := range flows {
		params := &studio.CreateExecutionParams{}
		params.SetFrom(from)
		params.SetTo(to)
		execution, err := client.StudioV2.CreateExecution(*flow.Sid, params)
		if err != nil {
			fmt.Println("Error creating studio flow log:", err)
			return
		}
		fmt.Println("Created studio flow execution log with SID:", execution.Sid, "in flow ", *flow.Sid)
	}
}
