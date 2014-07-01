package main

import "testing"

func TestImportLines(t *testing.T) {
	out := make([]string, 4)
	out[0] = "5 5"
	out[1] = "1 3 3"
	out[2] = "3 2 2"
	out[3] = "4 1 1"

	x, err := importLines("./test.txt")
	if err != nil {
		t.Errorf("readLines: %s", err)
		return
	}

	for index, value := range out {
		if value != x[index] {
			t.Errorf("importLines() line %v = %v, want %v ", index, x[index], value)
		}
	}
}