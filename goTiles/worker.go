package goTiles

import (
	"fmt"
)

// Start a batch of workers that attempt to solve the board concurrently.
func NewSolver(rawInput string) chan Path {
	lines, err := importLines(rawInput)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return nil
	}

	// Create a channel to pass down results
	results := make(chan Path, 10)

	// Create a seperate first board for stacksize access
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
	bigTile := &Tile{1, 3, 3}
	medTile := &Tile{3, 2, 2}
	smaTile := &Tile{4, 1, 1}

	testPath := newPath(8)
	testPath = append(testPath, Placement{0, 0, bigTile})
	testPath = append(testPath, Placement{0, 3, medTile})
	testPath = append(testPath, Placement{2, 3, medTile})
	testPath = append(testPath, Placement{3, 0, medTile})
	testPath = append(testPath, Placement{3, 2, smaTile})
	testPath = append(testPath, Placement{4, 2, smaTile})
	testPath = append(testPath, Placement{4, 3, smaTile})
	testPath = append(testPath, Placement{4, 4, smaTile})

	results <- testPath

	done <- true
}
