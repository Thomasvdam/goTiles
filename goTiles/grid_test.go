package goTiles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test whether a new Grid has the correct dimensions and whether the
// 2-dimensional array has the right size and is set to false.
func TestNewGrid(t *testing.T) {
	const dimension = 5

	testGrid := newGrid(dimension, dimension)

	assert.Equal(t, dimension, testGrid.width(), "They should be equal.")
	assert.Equal(t, dimension, testGrid.height(), "They should be equal.")

	gridArea := 0
	for x, row := range testGrid {
		for y, value := range row {
			assert.False(t, value, "(%v, %v) should be false.", x, y)
			gridArea++
		}
	}

	assert.Equal(t, dimension*dimension, gridArea, "They should be equal.")
}

// Test whether the leftmost upper corner can be found.
func TestFindPlace(t *testing.T) {
	emptyGrid := emptyGrid()

	// Empty grid.
	const emptyX, emptyY = 0, 0
	testEmptyX, testEmptyY := emptyGrid.findPlace()
	assert.Equal(t, emptyX, testEmptyX, "They should be equal.")
	assert.Equal(t, emptyY, testEmptyY, "They should be equal.")

	// 3x3 tile on the grid.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			emptyGrid[i][j] = true
		}
	}

	const oneTileX, oneTileY = 0, 3
	testOneTileX, testOneTileY := emptyGrid.findPlace()
	assert.Equal(t, oneTileX, testOneTileX, "They should be equal.")
	assert.Equal(t, oneTileY, testOneTileY, "They should be equal.")

	// Completely filled grid.
	for x, row := range emptyGrid {
		for y, _ := range row {
			emptyGrid[x][y] = true
		}
	}

	var fullX, fullY = emptyGrid.width(), emptyGrid.height()
	testFullX, testFullY := emptyGrid.findPlace()
	assert.Equal(t, fullX, testFullX, "They should be equal.")
	assert.Equal(t, fullY, testFullY, "They should be equal.")
}
