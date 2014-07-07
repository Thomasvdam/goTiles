package goTiles

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test whether newPath returns a path of the correct lenght and capacity.
func TestNewPath(t *testing.T) {
	const size = 8

	testPath := newPath(size)

	assert.Equal(t, 0, len(testPath), "The lenght of a new path should be 0.")
	assert.Equal(t, size, cap(testPath), "The capacity should be 8.")
}

func TestPlace(t *testing.T) {
	const x, y = 1, 4
	tile := &Tile{1, 3, 3}

	testPath := newPath(8)
	testPath = append(testPath, Placement{x, y, tile})

	assert.Equal(t, x, testPath[0].x, "They should be equal.")
	assert.Equal(t, y, testPath[0].y, "They should be equal.")
	assert.Equal(t, tile, testPath[0].tile, "They should be equal.")
}

// Test whether the pretty printing of a path works.
func TestPrettyPath(t *testing.T) {
	const prettyPath = "3x3, 2x2, 2x2, 2x2, 1x1, 1x1, 1x1, 1x1, "

	bigTile := &Tile{1, 3, 3}
	medTile := &Tile{3, 2, 2}
	smaTile := &Tile{4, 1, 1}

	testPath := newPath(8)
	testPath = append(testPath, Placement{1, 1, bigTile})
	testPath = append(testPath, Placement{1, 1, medTile})
	testPath = append(testPath, Placement{1, 1, medTile})
	testPath = append(testPath, Placement{1, 1, medTile})
	testPath = append(testPath, Placement{1, 1, smaTile})
	testPath = append(testPath, Placement{1, 1, smaTile})
	testPath = append(testPath, Placement{1, 1, smaTile})
	testPath = append(testPath, Placement{1, 1, smaTile})

	testPrettyPath := testPath.String()
	testInterface := fmt.Sprint(testPath)

	assert.Equal(t, prettyPath, testPrettyPath, "The paths should be the same.")
	assert.Equal(t, prettyPath, testInterface, "The interface should work.")
}
