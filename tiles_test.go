package main

import "testing"

// Test whether a valid tile is created correctly.
func TestValidNewTile(t *testing.T) {
	const amount, width, height = 3, 4, 4

	testOut, err := newTile(amount, width, height)
	if err != nil {
		t.Errorf("Error should be nil, but was \"%v\".", err)
		return
	}

	if testOut.amount != amount {
		t.Errorf("Amount = %v, want %v.")
	}

	if testOut.width != width {
		t.Errorf("Width = %v, want %v.")
	}

	if testOut.height != height {
		t.Errorf("Height = %v, want %v.")
	}
}

// Test whether incorrect tiles return errors.
func TestInvalidNewTiles(t *testing.T) {
	const wrong, good = -1, 1

	testOut, err := newTile(wrong, good, good)
	if err == nil {
		t.Errorf("No error occurred while invalid parameters were passed: amount %v, width %v, height %v.", wrong, good, good)
	}
	if testOut != nil {
		t.Errorf("Tile should be nil, but was %v.", testOut)
	}

	testOut, err = newTile(good, wrong, good)
	if err == nil {
		t.Errorf("No error occurred while invalid parameters were passed: amount %v, width %v, height %v.", good, wrong, good)
	}
	if testOut != nil {
		t.Errorf("Tile should be nil, but was %v.", testOut)
	}

	testOut, err = newTile(good, good, wrong)
	if err == nil {
		t.Errorf("No error occurred while invalid parameters were passed: amount %v, width %v, height %v.", good, good, wrong)
	}
	if testOut != nil {
		t.Errorf("Tile should be nil, but was %v.", testOut)
	}
}

// Test whether a valid stack is created correctly.
func TestValidNewStack(t *testing.T) {
	in := make([][]int, 3)
	in[0] = []int{1, 3, 3}
	in[1] = []int{3, 2, 2}
	in[2] = []int{4, 1, 1}

	testOut, err := newStack(in)
	if err != nil {
		t.Errorf("Error should be nil, but was \"%v\".", err)
	}

	for index, value := range testOut {
		if value.amount != in[index][0] || value.width != in[index][1] || value.height != in[index][2] {
			t.Errorf("Tile %v = %v, %v, %v, want %v, %v, %v.", index, value.amount, value.width, value.height, in[index][0], in[index][1], in[index][2])
		}
	}
}
