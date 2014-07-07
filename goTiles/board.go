package goTiles

import (
	"fmt"
)

// Board struct which contains all the information about the problem.
type Board struct {
	grid  Grid    // Current state of the grid.
	path  []*Tile // List of tiles in the order they were placed.
	stack Stack   // Stack of available tiles.
}

// An error type that details whether the board has too many
// or too few tiles for its total area.
type BoardError struct {
	totalArea int // Height of the board * width of the board.
	tileArea  int // Combined area of all tiles.
}

// Implement the Error() method for the error interface.
func (e *BoardError) Error() string {
	if e.totalArea > e.tileArea {
		return fmt.Sprintf("There are not enough tiles. Board area = %v, tile area = %v.", e.totalArea, e.tileArea)
	} else {
		return fmt.Sprintf("There are too many tiles. Board area = %v, tile area = %v.", e.totalArea, e.tileArea)
	}
}

// Creates a new board from the strings passed and checks whether the
// board is valid. If no slice is given as an argument the tiles.txt
// file will be used instead. Error is nil when the board is valid.
func NewBoard(lines []string) (board *Board, err error) {
	// Open tiles.txt file when no slice is passed as an argument.
	if lines == nil {
		lines, err = importLines("./tiles.txt")
		if err != nil {
			return nil, err
		}
	}

	// Convert the strings to dimensions and a stack of tiles.
	width, height, stackSlice := convertLines(lines)
	stack, err := newStack(stackSlice)
	if err != nil {
		return nil, err
	}

	// Create the board and check for errors.
	board = &Board{newGrid(width, height), make([]*Tile, stack.size()), stack}
	if err = board.valid(); err != nil {
		return nil, err
	}

	return board, nil
}

// Check whether the area of all tiles combined is equal to the area of
// the board. Returns nil when the board is valid, else it returns an
// error describing the issue.
func (board *Board) valid() error {
	totalArea := board.grid.width() * board.grid.height()

	// Calculate total area of all tiles.
	tileArea := 0
	for _, tile := range board.stack {
		tileArea += tile.amount * tile.width * tile.height
	}

	if totalArea != tileArea {
		return error(&BoardError{totalArea, tileArea})
	}

	return nil
}

// Attempt to place the passed tile in the leftmost upper corner. Returns
// false and revert the board to it's original state when not successful,
// returns true and updates the board when successful.
func (b *Board) placeTile(tile *Tile) bool {
	x, y := b.grid.findPlace()

	// Check whether the tile will fit.
	if x+tile.width > b.grid.width() || y+tile.height > b.grid.height() {
		return false
	}

	// Check whether the tile will not overlap.
	for i := 0; i < tile.width; i++ {
		for j := 0; j < tile.height; j++ {
			if b.grid[x+i][y+j] {
				return false
			}
		}
	}

	// Place tile and remove it from the stack.
	tile.amount--
	for i := 0; i < tile.width; i++ {
		for j := 0; j < tile.height; j++ {
			b.grid[x+i][y+j] = true
		}
	}

	return true
}
