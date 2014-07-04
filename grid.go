package main

// Grid struct which contains the dimensions and a 2-dimensional slice
// of bools.
type Grid struct {
	width int
	height int

	grid [][]bool
}

// Creates a new 2-dimensional array of false bools based on the
// dimensions passed.
func newGrid(width, height int) *Grid {
	grid := &Grid{width, height, make([][]bool, width)}

	for index, _ := range grid.grid {
		grid.grid[index] = make([]bool, height)
	}

	return grid
}