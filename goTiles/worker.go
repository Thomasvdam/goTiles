package goTiles

import (
	"fmt"
)

// Start a batch of workers that attempt to solve the board concurrently.
func NewSolver(lines []string) chan Path {
	results := make(chan Path, 10)

	boardOne, err := newBoard(lines)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return nil
	}

	done := make(chan bool, len(boardOne.stack))

	go newWorker(boardOne, 0, results, done)
	for i := 1; i < len(boardOne.stack); i++ {
		board, err := newBoard(lines)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return nil
		}
		go newWorker(board, i, results, done)
	}

	go func() {
		counter := 0
		for _ = range done {
			counter++
			if counter == len(boardOne.stack) {
				close(results)
			}
		}
	}()

	return results
}

// A worker that attempts to solve the board, always starting with
// the tile number passed. This way the workers are less likely to
// find the same results.
func newWorker(b *Board, tileNo int, results chan Path, done chan bool) {
	done <- true
}
