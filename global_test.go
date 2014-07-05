// Global testing variables that are useful for multiple tests.
package main

// Slice containing strings of a valid board .txt file.
var boardStringSlice = []string{
	"5 5",
	"1 3 3",
	"3 2 2",
	"4 1 1",
}

// Slice containing slices of ints of a valid board.
var boardIntSlice = [][]int{
	[]int{1, 3, 3},
	[]int{3, 2, 2},
	[]int{4, 1, 1},
}

// All false 5x5 2-dimensional bool slice.
var emptyGrid = Grid{
	[]bool{false, false, false, false, false},
	[]bool{false, false, false, false, false},
	[]bool{false, false, false, false, false},
	[]bool{false, false, false, false, false},
	[]bool{false, false, false, false, false},
}

// Valid empty board.
var	boardLiteral = &Board{
	emptyGrid,
	[]*Tile{
		&Tile{1, 3, 3},
		&Tile{3, 2, 2},
		&Tile{4, 1, 1},
	},
}
