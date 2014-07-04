package main

import "testing"

// Test whether a valid board is created correctly.
func TestValidNewBoard(t *testing.T) {
	in := make([]string, 4)
	in[0] = "5 5"
	in[1] = "1 3 3"
	in[2] = "3 2 2"
	in[3] = "4 1 1"

	out := &Board{5, 5, []*Tile{&Tile{1, 3, 3}, &Tile{ 3, 2, 2}, &Tile{4, 1, 1}}}

	testOut, err := newBoard(in)

	if err != nil {
		t.Errorf("Error should be nil, but was \"%v\".", err)
		return
	}

	if out.height != testOut.height {
		t.Errorf("Height = %v, want %v.", testOut.height, out.height)
	}

	if out.width != testOut.width {
		t.Errorf("Width = %v, want %v.", testOut.width, out.width)
	}

	for tileNo, value := range testOut.stack {
		if value.amount != out.stack[tileNo].amount || value.width != out.stack[tileNo].width || value.height != out.stack[tileNo].height {
			t.Errorf("Tile %v = %v, %v, %v, want %v, %v, %v.", tileNo, value.amount, value.width, value.height, out.stack[tileNo].amount, out.stack[tileNo].width, out.stack[tileNo].height)
		}
	}
}

// Test whether an invalid board returns the correct error.
func TestInvalidBoards(t *testing.T) {
	inSmall := make([]string, 4)
	inSmall[0] = "9 9"
	inSmall[1] = "1 3 3"
	inSmall[2] = "3 2 2"

	inBig := make([]string, 4)
	inBig[0] = "5 5"
	inBig[1] = "5 3 3"
	inBig[2] = "5 2 2"
	inBig[3] = "5 1 1"

	outSmall, errSmall := newBoard(inSmall)
	if errSmall == nil {
		t.Errorf("No error occurred while there were not enough tiles.")
	} else if berr, ok := errSmall.(*BoardError); ok {
		if berr.totalArea <= berr.tileArea {
			t.Errorf("The tile area should be smaller than %v, but it was %v.", berr.totalArea, berr.tileArea)
		}
	}
	if outSmall != nil {
		t.Errorf("Board should be nil, but was %v.", outSmall)
	}

	outBig, errBig := newBoard(inBig)
	if errBig == nil {
		t.Errorf("No error occurred while there were too many tiles.")
	} else if berr, ok := errBig.(*BoardError); ok {
		if berr.totalArea >= berr.tileArea {
			t.Errorf("The tile area should be smaller than %v, but it was %v.", berr.totalArea, berr.tileArea)
		}
	}
	if outBig != nil {
		t.Errorf("Board should be nil, but was %v.", outBig)
	}
}