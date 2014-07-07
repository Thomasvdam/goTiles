package goTiles

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

// Return the size of the board and a 2-dimensional slice
// with the numerical values extracted from the strings.
func convertLines(lines []string) (width, height int, stack [][]int) {
	stack = make([][]int, len(lines)-1)

	tmpDimension := strings.Split(lines[0], " ")
	width, _ = strconv.Atoi(tmpDimension[0])
	height, _ = strconv.Atoi(tmpDimension[1])

	for i := 1; i < len(lines); i++ {
		tmpTile := strings.Split(lines[i], " ")
		stack[i-1] = make([]int, 3)

		for index, value := range tmpTile {
			stack[i-1][index], _ = strconv.Atoi(value)
		}
	}

	return width, height, stack
}
