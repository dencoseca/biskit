package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}
