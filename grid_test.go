package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test whether a new Grid has the correct dimensions and whether the
// 2-dimensional array has the right size and is set to false.
func TestNewGrid(t *testing.T) {
	const dimension = 5

	testGrid := newGrid(dimension, dimension)

	assert.Equal(t, testGrid.width(), dimension, "They should be equal.")
	assert.Equal(t, testGrid.height(), dimension, "They should be equal.")


	gridArea := 0
	for x, row := range testGrid {
		for y, value := range row {
			assert.False(t, value, "(%v, %v) should be false.", x, y)
			gridArea++
		}
	}

	assert.Equal(t, gridArea, dimension * dimension, "They should be equal.")
}

// Test whether the leftmost upper corner can be found.
func TestFindPlace(t *testing.T) {

	// Empty grid.
	const emptyX, emptyY = 0, 0
	testEmptyX, testEmptyY := emptyGrid.findPlace()
	assert.Equal(t, testEmptyX, emptyX, "They should be equal.")
	assert.Equal(t, testEmptyY, emptyY, "They should be equal.")

	// 3x3 tile on the grid.
	var literalGrid = Grid{
		[]bool{true, true, true, false, false},
		[]bool{true, true, true, false, false},
		[]bool{true, true, true, false, false},
		[]bool{false, false, false, false, false},
		[]bool{false, false, false, false, false},
	}

	const oneTileX, oneTileY = 0, 3
	testOneTileX, testOneTileY := literalGrid.findPlace()
	assert.Equal(t, testOneTileX, oneTileX, "They should be equal.")
	assert.Equal(t, testOneTileY, oneTileY, "They should be equal.")

	// Completely filled grid.
	for x, row := range literalGrid {
		for y, _ := range row {
			literalGrid[x][y] = true
		}
	}

	var fullX, fullY = literalGrid.width(), literalGrid.height()
	testFullX, testFullY := literalGrid.findPlace()
	assert.Equal(t, testFullX, fullX, "They should be equal.")
	assert.Equal(t, testFullY, fullY, "They should be equal.")
}