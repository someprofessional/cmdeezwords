package main

import (
	"fmt"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type playModel struct {
	word       string
	userInput  string
	remaining  time.Duration
	timer      *time.Timer
	isGameOver bool
	difficulty string
	score      int
}

type tickMsg struct{}

func initialPlayModel(word string, difficulty string) playModel {
	return playModel{
		word:       word,
		remaining:  60 * time.Second,
		timer:      time.NewTimer(60 * time.Second),
		isGameOver: false,
		difficulty: difficulty,
		score:      0,
	}
}

func (p playModel) Init() tea.Cmd {
	return tea.Batch(
		tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return tickMsg{}
		}),
	)
}

func (p *playModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if p.isGameOver {
			return initialMenuModel(), nil // Transition to menuModel
		}

		switch keyMsg := msg.String(); keyMsg {
		case "enter":
			if p.userInput == p.word {
				p.score++
				p.userInput = ""
				p.word = getRandomWord(p.difficulty) // Update word with a new one
				if p.difficulty == "easy" {
					easyScore++
				} else if p.difficulty == "medium" {
					mediumScore++
				} else if p.difficulty == "hard" {
					hardScore++
				}
			}
		case "backspace":
			if len(p.userInput) > 0 {
				p.userInput = p.userInput[:len(p.userInput)-1]
			}
		default:
			if len(keyMsg) == 1 { // Check if it's a single character
				p.userInput += keyMsg
			}
		}

	case tickMsg:
		if !p.isGameOver {
			p.remaining -= time.Second
			if p.remaining <= 0 {
				p.remaining = 0
				p.isGameOver = true
				return initialMenuModel(), tea.Quit // This will stop the current program execution
			}
			return p, tea.Tick(time.Second, func(t time.Time) tea.Msg {
				return tickMsg{}
			})
		}
	}

	return p, nil
}

func (p playModel) View() string {
	if p.isGameOver {
		clearScreen()
		fmt.Print("game over")
		initialMenuModel()
	}

	return fmt.Sprintf("Time left: %s\nWord: %s\nYour input: %s\n\nYour score : %d",
		formatDuration(p.remaining),
		p.word,
		p.userInput,
		p.score,
	)
}

// Get a random word based on difficulty
func getRandomWord(difficulty string) string {
	wordList, exists := words[difficulty]
	if !exists || len(wordList) == 0 {
		return "error"
	}
	return wordList[rand.Intn(len(wordList))]
}
