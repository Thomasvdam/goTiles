package goTiles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test whether a valid tile is created correctly.
func TestValidNewTile(t *testing.T) {
	const amount, width, height = 3, 4, 4

	testTile, err := newTile(amount, width, height)
	if assert.Nil(t, err) {
		assert.Equal(t, amount, testTile.amount, "They should be equal.")
		assert.Equal(t, width, testTile.width, "They should be equal.")
		assert.Equal(t, height, testTile.height, "They should be equal.")
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
			assert.Equal(t, value.amount, testStack[index].amount)
			assert.Equal(t, value.width, testStack[index].width)
			assert.Equal(t, value.height, testStack[index].height)
		}
	}
}

// Test whether the stack size is returned correctly.
func TestStackSize(t *testing.T) {
	boardLiteral := boardLiteral()

	assert.Equal(t, 8, boardLiteral.stack.size(), "They should be equal.")
}
