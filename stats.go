package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Stats Model
type statsModel struct{}

// Initialize Stats Model
func initialStatsModel() statsModel {
	return statsModel{}
}

// Implement Init for Stats Model
func (s statsModel) Init() tea.Cmd {
	return nil
}

// Stats Model Update
func (s statsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit

		case "enter", " ":
			return initialMenuModel(), nil // Go back to the main menu
		}
	}
	return s, nil
}

// Stats Model View
func (s statsModel) View() string {
	return fmt.Sprintf(
		"Your stats are here...\n"+
			"--------------------------------------\n"+
			"Last Easy :   %d.\n--\n"+
			"Last Medium : %d.\n--\n"+
			"Last Hard :   %d.\n"+
			"--------------------------------------\n",
		easyScore, mediumScore, hardScore,
	)
}
