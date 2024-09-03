package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var choice int
var easyScore, mediumScore, hardScore = 0, 0, 0
var words map[string][]string

func init() {
	// Load words from JSON file
	file, err := os.Open("words.json")
	if err != nil {
		fmt.Printf("Error opening words.json: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading words.json: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &words)
	if err != nil {
		fmt.Printf("Error parsing words.json: %v\n", err)
		os.Exit(1)
	}
}

func formatDuration(d time.Duration) string {
	seconds := int(d.Seconds())
	return fmt.Sprintf("%02d:%02d", seconds/60, seconds%60)
}

func difficulty() {
	fmt.Println("Choose your difficulty :")
	fmt.Println("0 -- Go back to menu")
	fmt.Println("1 - easy")
	fmt.Println("2 - medium")
	fmt.Println("3 - hard")

	fmt.Scan(&choice)

	var difficultyStr string
	switch choice {
	case 1:
		difficultyStr = "easy"
	case 2:
		difficultyStr = "medium"
	case 3:
		difficultyStr = "hard"
	case 0:
		clearScreen()
		return
	default:
		fmt.Println("Choose something else please")
		difficulty()
		return
	}

	// Initialize a new playModel with a random word and the chosen difficulty
	word := getRandomWord(difficultyStr)
	model := initialPlayModel(word, difficultyStr)
	p := tea.NewProgram(&model)
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting the program: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	clearScreen()
	fmt.Println("Welcome to 寿司打! But without sushi ...")

	model := initialMenuModel()
	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting the program: %v\n", err)
		os.Exit(1)
	}
}
