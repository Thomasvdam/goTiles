package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test whether a valid board is created correctly.
func TestValidNewBoard(t *testing.T) {

	testOut, err := newBoard(boardStringSlice)

	if assert.Nil(t, err) {

		assert.Equal(t, testOut.grid.width, boardLiteral.grid.width, "They should be equal.")
		assert.Equal(t, testOut.grid.height, boardLiteral.grid.height, "They should be equal.")

		for tileNo, value := range boardLiteral.stack {
			assert.Equal(t, testOut.stack[tileNo].amount, value.amount, "They should be equal.")
			assert.Equal(t, testOut.stack[tileNo].width, value.width, "They should be equal.")
			assert.Equal(t, testOut.stack[tileNo].height, value.height, "They should be equal.")
		}
	}
}

// Test whether an invalid board returns the correct error.
func TestInvalidBoards(t *testing.T) {
	var inSmall = []string{
		"9 9",
		"1 3 3",
		"3 2 2",
	}

	outSmall, errSmall := newBoard(inSmall)
	if assert.NotNil(t, errSmall) {
		if berr, ok := errSmall.(*BoardError); ok {
			// assert.Condition might be worthwhile here, but can't get it to work.
			if berr.totalArea <= berr.tileArea {
				t.Errorf("The tile area should be smaller than %v, but it was %v.", berr.totalArea, berr.tileArea)
			}
		}
	}
	assert.Nil(t, outSmall)

	var inBig = []string{
		"5 5",
		"5 3 3",
		"5 2 2",
		"5 1 1",
	}

	outBig, errBig := newBoard(inBig)
	if assert.NotNil(t, errBig) {
		if berr, ok := errBig.(*BoardError); ok {
			// assert.Condition might be worthwhile here, but can't get it to work.
			if berr.totalArea >= berr.tileArea {
				t.Errorf("The tile area should be smaller than %v, but it was %v.", berr.totalArea, berr.tileArea)
			}
		}
	}
	assert.Nil(t, outBig)
}

// Test whether valid tile placement is handled correctly.
func TestPlaceTilesValid(t *testing.T) {
	testBoard, err := newBoard(boardStringSlice)

	if assert.Nil(t, err) {
		assert.True(t, testBoard.placeTile(testBoard.stack[0]))
		assert.Equal(t, testBoard.stack[0].amount + 1, boardLiteral.stack[0].amount, "Amount should decrease by 1.")

		var literalGrid = Grid{
			[]bool{true, true, true, false, false},
			[]bool{true, true, true, false, false},
			[]bool{true, true, true, false, false},
			[]bool{false, false, false, false, false},
			[]bool{false, false, false, false, false},
		}

		// Compare grids.
		for x, row := range literalGrid {
			for y, value := range row {
				assert.Equal(t, testBoard.grid[x][y], value, "They should be equal.")
			}
		}

		// Place a second tile.
		assert.True(t, testBoard.placeTile(testBoard.stack[1]))
		assert.Equal(t, testBoard.stack[1].amount + 1, boardLiteral.stack[1].amount, "Amount should decrease by 1.")

		// Place a 2x2 tile under the 3x3 tile.
		literalGrid[0][4], literalGrid[1][4], literalGrid[0][3], literalGrid[1][3] = true, true, true, true

		// Compare grids.
		for x, row := range literalGrid {
			for y, value := range row {
				assert.Equal(t, testBoard.grid[x][y], value, "They should be equal.")
			}
		}
	}
}

// Test whether invalid tile placement is handled correctly.
func TestPlaceTilesInvalid(t *testing.T) {

	var literalGrid = Grid{
		[]bool{true, true, true, false, false},
		[]bool{true, true, true, false, false},
		[]bool{true, true, true, false, false},
		[]bool{false, false, false, false, false},
		[]bool{false, false, false, false, false},
	}

	testBoard, err := newBoard(boardStringSlice)
	if assert.Nil(t, err) {
		testBoard.stack[0].amount++

		assert.True(t, testBoard.placeTile(testBoard.stack[0]))
		// There should be a nice before/after wrapper in testify.
		amountBefore := testBoard.stack[0].amount
		assert.False(t, testBoard.placeTile(testBoard.stack[0]))
		amountAfter := testBoard.stack[0].amount
		assert.Equal(t, amountBefore, amountAfter, "They should be equal.")

		// Compare grids.
		for x, row := range literalGrid {
			for y, value := range row {
				if value != testBoard.grid[x][y] {
					assert.Equal(t, testBoard.grid[x][y], value, "They should be equal.")
				}
			}
		}
	}
}
