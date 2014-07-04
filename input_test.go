package main

import "testing"

// Test whether a text file can be opened and read correctly.
func TestImportLines(t *testing.T) {
	out := make([]string, 4)
	out[0] = "5 5"
	out[1] = "1 3 3"
	out[2] = "3 2 2"
	out[3] = "4 1 1"

	testOut, err := importLines("./test.txt")
	if err != nil {
		t.Errorf("readLines: %s.", err)
		return
	}

	for index, value := range out {
		if value != testOut[index] {
			t.Errorf("importLines() line %v = %v, want %v.", index, testOut[index], value)
		}
	}
}

// Test wheter string slices are parsed correctly.
func TestConvertLines(t *testing.T) {
	in := make([]string, 4)
	in[0] = "5 5"
	in[1] = "1 3 3"
	in[2] = "3 2 2"
	in[3] = "4 1 1"

	const width, height = 5, 5
	out := make([][]int, 3)
	out[0] = []int{1, 3, 3}
	out[1] = []int{3, 2, 2}
	out[2] = []int{4, 1, 1}

	testWidth, testHeight, testOut := convertLines(in)

	if testWidth != width {
		t.Errorf("convertLines() width = %v, want %v.", testWidth, width)
	}

	if testHeight != height {
		t.Errorf("convertLines() height = %v, want %v.", testHeight, height)		
	}

	for tile, slice := range testOut {
		for index, value := range slice {
			if value != out[tile][index] {
				t.Errorf("convertLines() tile %v = %v, want %v.", tile, value, out[tile][index])
			}
		}
	}
}