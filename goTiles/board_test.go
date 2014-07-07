package goTiles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test whether a valid board is created correctly.
func TestValidNewBoard(t *testing.T) {
	boardLiteral := boardLiteral()

	testOut, err := NewBoard(boardStringSlice())

	if assert.Nil(t, err) {

		assert.Equal(t, boardLiteral.grid.width, testOut.grid.width, "They should be equal.")
		assert.Equal(t, boardLiteral.grid.height, testOut.grid.height, "They should be equal.")
		assert.Empty(t, testOut.path, "No tiles should have been placed yet.")

		for tileNo, value := range boardLiteral.stack {
			assert.Equal(t, value.amount, testOut.stack[tileNo].amount, "They should be equal.")
			assert.Equal(t, value.width, testOut.stack[tileNo].width, "They should be equal.")
			assert.Equal(t, value.height, testOut.stack[tileNo].height, "They should be equal.")
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

	outSmall, errSmall := NewBoard(inSmall)
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

	outBig, errBig := NewBoard(inBig)
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
	boardLiteral := boardLiteral()

	testBoard, err := NewBoard(boardStringSlice())

	if assert.Nil(t, err) {
		assert.True(t, testBoard.placeTile(testBoard.stack[0]))
		assert.Equal(t, boardLiteral.stack[0].amount, testBoard.stack[0].amount+1, "Amount should decrease by 1.")
		assert.Equal(t, 1, len(testBoard.path), "Pathlenght should increase by one.")

		literalGrid := emptyGrid()
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				literalGrid[i][j] = true
			}
		}

		// Compare grids.
		for x, row := range literalGrid {
			for y, value := range row {
				assert.Equal(t, value, testBoard.grid[x][y], "They should be equal.")
			}
		}

		// Place a second tile.
		assert.True(t, testBoard.placeTile(testBoard.stack[1]))
		assert.Equal(t, boardLiteral.stack[1].amount, testBoard.stack[1].amount+1, "Amount should decrease by 1.")
		assert.Equal(t, 2, len(testBoard.path), "Pathlenght should increase by one.")

		// Place a 2x2 tile under the 3x3 tile in the comparison grid.
		literalGrid[0][4], literalGrid[1][4], literalGrid[0][3], literalGrid[1][3] = true, true, true, true

		// Compare grids.
		for x, row := range literalGrid {
			for y, value := range row {
				assert.Equal(t, value, testBoard.grid[x][y], "They should be equal.")
			}
		}
	}
}

// Test whether invalid tile placement is handled correctly.
func TestPlaceTilesInvalid(t *testing.T) {

	literalGrid := emptyGrid()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			literalGrid[i][j] = true
		}
	}

	testBoard, err := NewBoard(boardStringSlice())
	if assert.Nil(t, err) {
		testBoard.stack[0].amount++

		assert.True(t, testBoard.placeTile(testBoard.stack[0]))
		// There should be a nice before/after wrapper in testify.
		amountBefore := testBoard.stack.size
		lenghtBefore := len(testBoard.path)
		assert.False(t, testBoard.placeTile(testBoard.stack[0]))
		amountAfter := testBoard.stack.size
		lenghtAfter := len(testBoard.path)
		assert.Equal(t, amountBefore, amountAfter, "They amount of tiles should not change.")
		assert.Equal(t, lenghtBefore, lenghtAfter, "The path should not change.")

		// Compare grids.
		for x, row := range literalGrid {
			for y, value := range row {
				assert.Equal(t, value, testBoard.grid[x][y], "They should be equal.")
			}
		}
	}
}

// Test whether backtracking is handled correctly.
func TestBacktrack(t *testing.T) {
	boardLiteral := boardLiteral()

	testBoard, err := NewBoard(boardStringSlice())

	if assert.Nil(t, err) {
		assert.True(t, testBoard.placeTile(testBoard.stack[0]))

		testBoard.backtrack()

		assert.Empty(t, testBoard.path, "Path should be empty.")
		assert.Equal(t, boardLiteral.stack.size(), testBoard.stack.size(), "Both stacks should be equal.")
		assert.Equal(t, boardLiteral.stack[0].amount, testBoard.stack[0].amount, "They should be equal.")

		for x, row := range boardLiteral.grid {
			for y, value := range row {
				assert.Equal(t, value, testBoard.grid[x][y], "They should be equal.")
			}
		}
	}
}
