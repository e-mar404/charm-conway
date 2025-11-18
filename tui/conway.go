package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	loaded bool
	timer  timer.Model
	state  [][]int
}

var (
	liveCell = lipgloss.NewStyle().Background(lipgloss.Color("69"))
	deadCell = lipgloss.NewStyle().Background(lipgloss.Color("0"))
)

func newGame() model {
	return model{timer: timer.NewWithInterval(time.Minute, time.Millisecond*150)}
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.state = randomState(msg.Width, msg.Height-1)

		m.loaded = true

	case timer.TickMsg:
		m.state = nextGeneration(m.state)
		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}

// View implements tea.Model.
func (m model) View() string {
	if !m.loaded {
		return "loading..."
	}

	var board string

	for _, row := range m.state {
		for _, cell := range row {
			if cell == 1 {
				board += liveCell.Render(" ")
			} else {
				board += deadCell.Render(" ")
			}
		}
		board += "\n"
	}

	board += m.timer.View()

	return board
}
