package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Menu Model
type menuModel struct {
	choice []string
	cursor int
}

// Initialize Menu Model
func initialMenuModel() menuModel {
	return menuModel{
		choice: []string{"Choose a difficulty", "See my stats", "Exit the game"},
	}
}

// Implement Init for Menu Model
func (m menuModel) Init() tea.Cmd {
	return nil
}

// Menu Model Update
func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choice)-1 {
				m.cursor++
			}

		case "enter", " ":
			switch m.cursor {
			case 0:
				clearScreen()
				difficulty()
			case 1:
				clearScreen()
				return initialStatsModel(), nil
			case 2:
				clearScreen()
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

// Menu Model View
func (m menuModel) View() string {
	var view string
	for i, choice := range m.choice {
		cursor := " " // Default cursor is not visible
		if m.cursor == i {
			cursor = ">" // Cursor for the selected item
		}
		view += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return view
}
