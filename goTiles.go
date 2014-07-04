package main

import (
	"fmt"
)

func main() {
	
	board, err := newBoard(nil)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	fmt.Println(board)
}