package tui

import "math/rand"

func randomState(width, height int) [][]bool {
	var state [][]bool
	for range height {
		row := make([]bool, width)
		state = append(state, row)
	}

	for r, row := range state {
		for c := range row {
			// have a probability of 20%
			// this had modulo bias but its good enough for now
			isAlive := (rand.Int() % 5) == 0
			if isAlive {
				state[r][c] = true
			}
		}
	}

	return state
}

func nextGeneration(state [][]bool) {
	// do checks here:
	// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	// 2. Any live cell with two or three live neighbours lives on to the next generation.
	// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
	// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
}
