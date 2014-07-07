package goTiles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test whether a text file can be opened and read correctly.
func TestImportLines(t *testing.T) {
	boardStringSlice := boardStringSlice()

	testOut, err := importLines("../testFiles/test.txt")

	if assert.Nil(t, err) {
		for index, value := range boardStringSlice {
			assert.Equal(t, value, testOut[index], "They should be equal.")
		}
	}
}

// Test wheter string slices are parsed correctly.
func TestConvertLines(t *testing.T) {
	boardIntSlice := boardIntSlice()
	const width, height = 5, 5

	testWidth, testHeight, testOut := convertLines(boardStringSlice())

	assert.Equal(t, width, testWidth, "They should be equal.")
	assert.Equal(t, height, testHeight, "They should be equal.")

	for tile, slice := range boardIntSlice {
		for index, value := range slice {
			assert.Equal(t, value, testOut[tile][index], "They should be equal.")
		}
	}
}
