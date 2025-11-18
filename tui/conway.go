package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{
	loaded bool
	state [][]bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.state = randomState(msg.Width, msg.Height)
		
		m.loaded = true

	case tea.KeyMsg:
		switch  msg.String() {
		case "q":
			return m, tea.Quit
		}
	
	default:
		nextGeneration(m.state)	
	}
	return m, nil
}

// View implements tea.Model.
func (m model) View() string {
	if !m.loaded {
		return "loading..."
	}

	var board string

	for _, row := range m.state {
		for _, alive := range row {
			if alive {
				board += "x"
			} else {
				board += "."
			}
		}
		board += "\n"
	}

	return board
}
