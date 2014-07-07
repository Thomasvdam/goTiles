package goTiles

import (
	"fmt"
)

type Stack []*Tile

// Tile struct containing an amount, a width and a height.
type Tile struct {
	amount int
	width  int
	height int
}

// Create a new Tile based on the integers passed. Returns an error when
// integers lower than 1 are passed.
func newTile(amount, width, height int) (tile *Tile, err error) {
	if amount < 1 || width < 1 || height < 1 {
		err = fmt.Errorf("Invalid tile, all integers should be greater than 0, but the values were: amount %v, width %v, height %v.", amount, width, height)
		return nil, err
	}

	tile = &Tile{amount, width, height}
	return tile, nil
}

// Create a stack of tiles from the slice of integer slices passed.
// Returns an error when a tile can't be created.
func newStack(stackSlice [][]int) (stack Stack, err error) {
	stack = make([]*Tile, len(stackSlice))
	for index, tileSlice := range stackSlice {
		stack[index], err = newTile(tileSlice[0], tileSlice[1], tileSlice[2])
		if err != nil {
			return nil, err
		}
	}

	return
}

// Return the total amount of tiles in a stack.
func (s Stack) size() (total int) {
	for _, value := range s {
		total += value.amount
	}
	return
}
