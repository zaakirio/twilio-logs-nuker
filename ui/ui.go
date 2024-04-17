package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	label string
	value string
}

type Model struct {
	currentView string
	credentials bool
	services    []item
	selectedIdx int
}

func (m model) Init() tea.Cmd {
	// Check for credentials in environment variables
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	m.credentials = accountSID != "" && authToken != ""

	if !m.credentials {
		m.currentView = "promptCredentials"
	} else {
		m.currentView = "mainMenu"
		m.services = []item{
			{"Delete Calls", "calls"},
			{"Delete Messages", "messages"},
			{"Delete Conversations", "conversations"},
			{"Delete Studio Flows", "studio_flows"},
		}
	}

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.currentView == "mainMenu" {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter":
				if m.selectedIdx >= 0 && m.selectedIdx < len(m.services) {
					// Call delete function based on selected service
					fmt.Printf("Deleting %s...\n", m.services[m.selectedIdx].label)
					// Replace with your actual delete logic using Twilio API
					// ...
					m.currentView = "mainMenu"
				}
			case "up", "down":
				if msg.String() == "up" {
					m.selectedIdx--
					if m.selectedIdx < 0 {
						m.selectedIdx = len(m.services) - 1
					}
				} else {
					m.selectedIdx++
					if m.selectedIdx >= len(m.services) {
						m.selectedIdx = 0
					}
				}
			}
		} else if m.currentView == "promptCredentials" {
			if msg.String() == "enter" {
				m.currentView = "mainMenu"
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.currentView {
	case "mainMenu":
		s := lipgloss.NewStyle().Bold(true)
		menu := s.Render("Select Service:") + "\n\n"
		for i, service := range m.services {
			prefix := "  "
			if i == m.selectedIdx {
				prefix = s.Render("> ")
			}
			menu += fmt.Sprintf("%s%s\n", prefix, service.label)
		}
		return menu
	case "promptCredentials":
		return "Please enter your Twilio Account SID and Auth Token in your .bashrc\n" +
			"and restart this program.\n"
	default:
		return "Unknown view"
	}
}
func Run(m Model) error {
	p := tea.NewProgram(m)
	return p.Start()
}
