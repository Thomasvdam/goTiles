package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test whether a text file can be opened and read correctly.
func TestImportLines(t *testing.T) {
	testOut, err := importLines("./test.txt")
	
	if assert.Nil(t, err) {
		for index, value := range boardStringSlice {
			assert.Equal(t, testOut[index], value, "They should be equal.")
		}
	}
}

// Test wheter string slices are parsed correctly.
func TestConvertLines(t *testing.T) {
	const width, height = 5, 5

	testWidth, testHeight, testOut := convertLines(boardStringSlice)

	assert.Equal(t, testWidth, width, "They should be equal.")
	assert.Equal(t, testHeight, height, "They should be equal.")

	for tile, slice := range boardIntSlice {
		for index, value := range slice {
			assert.Equal(t, testOut[tile][index], value, "They should be equal.")
		}
	}
}