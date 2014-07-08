package main

import (
	"fmt"
	"github.com/Arcania0311/goTiles/goTiles"
)

func main() {

	results := goTiles.NewSolver(nil)

	for value := range results {
		fmt.Println(value)
	}
}
