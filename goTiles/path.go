package goTiles

import (
	"fmt"
)

type Path []Placement

type Placement struct {
	x    int
	y    int
	tile *Tile
}

// Create a new path with the capacityh and length set correctly.
func newPath(size int) Path {
	path := make([]Placement, 0, size)
	return path
}

// Pretty print the tiles in the order they were placed.
func (p Path) String() string {
	prettyPath := ""
	for _, value := range p {
		prettyPath += fmt.Sprintf("%vx%v:(%v,%v) - ", value.tile.width, value.tile.height, value.x, value.y)
	}
	return prettyPath
}
