package main

import (
	"fmt"
)

type Board struct {
	width int
	height int

	stack [][]int
}

// An error type that details whether the board has too many
// or too few tiles for its total area.
type BoardError struct {
	totalArea int
	tileArea int
}

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
func newBoard(lines []string) (board *Board, err error) {

	if lines == nil {
		// Open the tiles.txt file.
		lines, err = importLines("./tiles.txt")
		if err != nil {
			return nil, err
		}
	}

	width, height, stack := convertLines(lines)
	board = &Board{width, height, stack}

	if err = board.valid(); err != nil {
		return nil, err
	}

	return board, nil
}

// Check whether the area of all tiles combined is equal to the area of
// the board. Returns nil when the board is valid, else it returns an
// error describing the issue.
func (board *Board) valid() error {
	totalArea := board.width * board.height

	tileArea := 0
	for _, value := range board.stack {
		tileArea += value[0] * value[1] * value[2]
	}

	if totalArea != tileArea {
		return error(&BoardError{totalArea, tileArea})
	}

	return nil
}