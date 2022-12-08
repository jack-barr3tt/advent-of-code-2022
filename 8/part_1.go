package main

import (
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Strings are basically byte arrays so by splitting the file by lines we essentially get a 2x2 byte matrix
	grid := strings.Split(string(data), "\n")

	// Map where the key is an encoding of the position (purely to avoid casting) and the value is whether the tree can be seen
	// The encoding is just x * 128 + y (128 is purely used to ensure that two different positions don't have the same key)
	// Because the same key will only appear once in the map, trees that can be seen from multiple directions are only counted once
	found := make(map[int]int)

	// Consider left and right visibility
	for x := range grid {
		// Store the tallest tree for left and right checks
		var maxL byte = 0
		var maxR byte = 0
		// For each row, loop across it
		for y := range grid[x] {
			if grid[x][y] > maxL {
				maxL = grid[x][y]
				found[128*x + y] = 1
			}

			// y is the value to use for checking to the right, so we use yL for checking to the left
			yL := len(grid[0]) - 1 - y

			if grid[x][yL] > maxR {
				maxR = grid[x][yL]
				found[128*x + yL] = 1
			}
		}
	}

	// Consider up and down visibility
	for y := range grid[0] {
		// Store the tallest tree for the up and down checks
		var maxU byte = 0
		var maxD byte = 0
		// Loop across each column
		for x := range grid {
			if grid[x][y] > maxU {
				maxU = grid[x][y]
				found[128*x + y] = 1
			}

			// x is the value to use for checking down, so we use xD for checking up
			xD := len(grid) - 1 - x

			if grid[xD][y] > maxD {
				maxD = grid[xD][y]
				found[128*xD + y] = 1
			}
		}
	} 

	// The number of entries in the map is the number of trees that are visible from the outside
	println(len(found))
}