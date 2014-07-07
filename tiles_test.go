package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test whether a valid tile is created correctly.
func TestValidNewTile(t *testing.T) {
	const amount, width, height = 3, 4, 4

	testTile, err := newTile(amount, width, height)
	if assert.Nil(t, err) {
		assert.Equal(t, testTile.amount, amount, "They should be equal.")
		assert.Equal(t, testTile.width, width, "They should be equal.")
		assert.Equal(t, testTile.height, height, "They should be equal.")
	}
}

// Test whether incorrect tiles return errors.
func TestInvalidNewTiles(t *testing.T) {
	const wrong, good = -1, 1

	testTilewgg, err := newTile(wrong, good, good)
	assert.NotNil(t, err)
	assert.Nil(t, testTilewgg)

	testTilegwg, err := newTile(good, wrong, good)
	assert.NotNil(t, err)
	assert.Nil(t, testTilegwg)

	testTileggw, err := newTile(good, good, wrong)
	assert.NotNil(t, err)
	assert.Nil(t, testTileggw)
}

// Test whether a valid stack is created correctly.
func TestValidNewStack(t *testing.T) {
	boardLiteral := boardLiteral()

	testStack, err := newStack(boardIntSlice())
	if assert.Nil(t, err) {
		for index, value := range boardLiteral.stack {
			assert.Equal(t, testStack[index].amount, value.amount)
			assert.Equal(t, testStack[index].width, value.width)
			assert.Equal(t, testStack[index].height, value.height)
		}
	}
}
