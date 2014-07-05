package main

type Grid [][]bool

// Creates a new 2-dimensional array of false bools based on the
// dimensions passed.
func newGrid(width, height int) Grid {
	grid := make([][]bool, width)

	for index, _ := range grid {
		grid[index] = make([]bool, height)
	}

	return grid
}

// Return the width of the grid as an int.
func (g Grid) width() int {
	return len(g)
}

// Return the height of the grid as an int.
func (g Grid) height() int {
	return len(g[0])
}

// Find the leftmost upper corner that is free.
func (g Grid) findPlace() (int, int) {
	for x, collumn := range g {
		for y, value := range collumn {
			if !value {
				return x, y
			}
		}
	}
	
	return len(g), len(g[0])
}