package main

import (
	"fmt"
	"github.com/Arcania0311/goTiles/goTiles"
)

func main() {

	results := goTiles.NewSolver("tiles.txt")

	for value := range results {
		fmt.Println(value)
	}
}
