package main

import (
	"os"
	"fmt"
	"bufio"
)

var (

)

func main() {
	

	// Open the file.
	lines, err := importLines("./tiles.txt")
	if err != nil {
		fmt.Println("readLines: %s", err)
	}

	fmt.Println(lines)
}

// Open the specified file path and return a slice of its lines.
func importLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}