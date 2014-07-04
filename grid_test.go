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
		t.Errorf("Grid area is %v, but should be %v.", gridArea, dimension * dimension)
	}
}