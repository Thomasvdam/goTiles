package main

import (
	"fmt"
)

func main() {
	
	board, err := newBoard(nil)
	if err != nil {
		fmt.Println("readLines:", err)
		return
	}

	fmt.Println(board)
}