package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	p := tea.NewProgram(newGame(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
