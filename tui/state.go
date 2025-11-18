package tui

import (
	"math/rand"
)

func randomState(width, height int) [][]int {
	var state [][]int
	for range height {
		row := make([]int, width)
		state = append(state, row)
	}

	for r, row := range state {
		for c := range row {
			// this might has modulo bias but its good enough for now
			isAlive := (rand.Int() % 5) == 0
			if isAlive {
				state[r][c] = 1
			}
		}
	}

	return state
}

func nextGeneration(state [][]int) [][]int {
	height := len(state)
	width := len(state[0])
	var newState [][]int
	for range height {
		newRow := make([]int, width)
		newState = append(newState, newRow)
	}

	for r, row := range state {
		for c := range row {
			neighbours := neighbourCount(state, r, c)
			switch state[r][c] {
			case 0:
				if neighbours == 3 {
					newState[r][c] = 1
				}
			case 1:
				switch neighbours {
				case 2, 3:
					newState[r][c] = state[r][c]
				default:
					newState[r][c] = 0
				}
			}
		}
	}

	return newState
}

func neighbourCount(state [][]int, row, col int) int {
	width := len(state[0])
	height := len(state)
	count := 0

	if row-1 >= 0 && col-1 >= 0 {
		count += state[row-1][col-1]
	}

	if row-1 >= 0 {
		count += state[row-1][col]
	}

	if row-1 >= 0 && col+1 < width {
		count += state[row-1][col+1]
	}

	if col-1 >= 0 {
		count += state[row][col-1]
	}

	if col+1 < width {
		count += state[row][col+1]
	}

	if row+1 < height && col-1 >= 0 {
		count += state[row+1][col-1]
	}

	if row+1 < height {
		count += state[row+1][col]
	}

	if row+1 < height && col+1 < width {
		count += state[row+1][col+1]
	}

	return count
}
