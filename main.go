package main

import (
	"fmt"
	"github.com/Arcania0311/goTiles/goTiles"
)

func main() {

	board, err := goTiles.NewBoard(nil)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	fmt.Println(board)
}
