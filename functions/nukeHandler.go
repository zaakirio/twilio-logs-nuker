package Handlers

import (
	"sync"
)

func DeleteAll() {
	var wg sync.WaitGroup
	wg.Add(4) // Number of concurrent goroutines

	go func() {
		defer wg.Done()
		DeleteMessageLogs()
	}()
	go func() {
		defer wg.Done()
		DeleteCallLogs()
	}()
	go func() {
		defer wg.Done()
		DeleteConversationsLogs()
	}()
	go func() {
		defer wg.Done()
		DeleteFlowLogs()
	}()

	wg.Wait() // Wait for all goroutines to finish
}
