package main

import "testing"

// Test whether a new Grid has the correct dimensions and whether the
// 2-dimensional array has the right size and is set to false.
func TestNewGrid(t *testing.T) {
	const dimension = 5

	testOut := newGrid(5, 5)

	if testOut.height != dimension {
		t.Errorf("Height = %v, want %v.", testOut.height, dimension)
	}

	if testOut.width != dimension {
		t.Errorf("Width = %v, want %v.", testOut.width, dimension)
	}

	gridArea := 0
	for indexRow, row := range testOut.grid {
		for indexCollumn, value := range row {
			if value {
				t.Errorf("Gridpoint (%v,%v) should be false but is true.", indexRow, indexCollumn)
			}
			gridArea++
		}
	}

	if gridArea != dimension * dimension {
		t.Errorf("Grid area should be %v, but is %v.", gridArea, dimension * dimension)
	}
}

// Test whether the leftmost upper corner can be found.
func TestFindPlace(t *testing.T) {
	grid := newGrid(5, 5)

	// Empty grid.
	outX, outY := 0, 0
	testX, testY := grid.findPlace()
	if testX != outX || testY != outY {
		t.Errorf("Lefmost upper corner should be (%v, %v), but is (%v, %v).", outX, outY, testX, testY)
	}

	// 3x3 tile on the grid.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid.grid[i][j] = true
		}
	}

	outX, outY = 0, 3
	testX, testY = grid.findPlace()
	if testX != outX || testY != outY {
		t.Errorf("Lefmost upper corner should be (%v, %v), but is (%v, %v).", outX, outY, testX, testY)
	}

	// Completely filled grid.
	for x, row := range grid.grid {
		for y, _ := range row {
			grid.grid[x][y] = true
		}
	}

	outX, outY = grid.width, grid.height
	testX, testY = grid.findPlace()
	if testX != outX || testY != outY {
		t.Errorf("Lefmost upper corner should be (%v, %v), but is (%v, %v).", outX, outY, testX, testY)
	}

}