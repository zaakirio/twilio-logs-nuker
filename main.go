package main

import (
	"fmt"
	"log"
	"os"
	"twilio-logs-nuker/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	from_phone_number := os.Getenv("FROM_PHONE_NUMBER")
	to_phone_number := os.Getenv("TO_PHONE_NUMBER")

	if from_phone_number == "" || to_phone_number == "" {
		log.Fatal("FROM_PHONE_NUMBER and TO_PHONE_NUMBER must be set in .env file")
	}

	// Move as an option later in bt
	Generate(from_phone_number, to_phone_number)
	m := ui.LoadModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
