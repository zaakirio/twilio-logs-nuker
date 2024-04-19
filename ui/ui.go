package ui

import (
	Handlers "bubble-tea-test/functions"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Item struct {
	title string
	desc  string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

type model struct {
	List     list.Model
	Choice   string
	Quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i.title)
			}
			// Don't return tea.Quit here
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.Choice != "" {
		Handlers.DeleteCallLogs()
		return docStyle.Render(fmt.Sprintf("%s Deleted", m.Choice))
	}
	if m.Quitting {
		return docStyle.Render("Exitting, nothing to be deleted")
	}
	return "\n" + m.List.View()
	// return docStyle.Render(m.list.View())
}

func LoadModel() model {
	// Load items from the database
	items := []list.Item{
		Item{title: "Call Logs", desc: "Fetches all call logs and removes them from twilio account"},
		Item{title: "Message Logs", desc: "Fetches all message logs and removes them from twilio account"},
		Item{title: "Studio Flow Log", desc: "Fetches all studio flow logs and removes them from twilio account"},
		Item{title: "Conversations Logs", desc: "Fetches all conversations logs and removes them from twilio account"},
		Item{title: "Everything!", desc: "Nukes all logs from the twilio account"},
	}

	model := model{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	model.List.Title = "Twilio Logs Nuker"
	return model
}
