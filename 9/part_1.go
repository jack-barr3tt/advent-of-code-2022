package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	// Position of the head
	xH := 0
	yH := 0
	// Position of the tail
	xT := 0
	yT := 0

	// Map of moves as bytes to their vector transformation
	moves := make(map[byte][]int)
	moves['R'] = []int{1,0}
	moves['L'] = []int{-1,0}
	moves['U'] = []int{0,1}
	moves['D'] = []int{0,-1}

	// Map of visited positions
	visited := make(map[string]int)

	// Instantly add the starting position to the map
	visited["0,0"] = 1

	for l := range lines {
		line := lines[l]
		parts := strings.Split(line, " ")
		mag, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		xH += moves[line[0]][0] * mag
		yH += moves[line[0]][1] * mag
		
		for true {
			// Work out the component distances
			xSep := xH - xT
			ySep := yH - yT

			// Use pythagorus to work out how far the head and tail are from each other after the head moves
			dist := math.Sqrt(math.Pow(float64(xSep), 2) + math.Pow(float64(ySep), 2))

			// If they are more than 2 units apart, the tail is pulled toward the head
			if dist < 2 {
				break
			}

			// Use abs / actual to find the movement needed in the x and y direction and apply it
			// Only do this if the component distance is not 0 otherwise we get divide by 0 problems
			if xSep != 0 {xT += int(math.Abs(float64(xSep)) / float64(xSep))}
			if ySep != 0 {yT += int(math.Abs(float64(ySep)) / float64(ySep))}

			// I have to actually convert to string properly here because x and y could be negative
			// which doesn't play nice with ascii
			visited[strconv.Itoa(xT) + "," + strconv.Itoa(yT)] = 1
		}
	}

	println(len(visited))
}