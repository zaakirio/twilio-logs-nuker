package Handlers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/twilio/twilio-go"
	conversation "github.com/twilio/twilio-go/rest/conversations/v1"
)

func DeleteConversationsLogs() {
	client := InitClient()
	ConversationsRecords := getConversationsRecords(client)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for _, Conversations := range ConversationsRecords {
		select {
		case <-interrupt:
			fmt.Println("Interrupt received. Cancelling deletion...")
			close(interrupt)
			return
		default:
			err := client.ConversationsV1.DeleteConversation(*Conversations.Sid, &conversation.DeleteConversationParams{})
			if err != nil {
				fmt.Printf("Error deleting Conversations log %s: %v\n", *Conversations.Sid, err)
			} else {
				fmt.Printf("Deleted Conversations log %s\n", *Conversations.Sid)
			}
		}
	}
}

func getConversationsRecords(client *twilio.RestClient) []conversation.ConversationsV1Conversation {
	parameters := &conversation.ListConversationParams{}
	parameters.SetPageSize(1000)
	ConversationsRecords, err := client.ConversationsV1.ListConversation(parameters)
	checkError(err)
	return ConversationsRecords
}
