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

	// Array containing the current positions of all knots
	knots := make([][2]int, 0)

	// Map of visited positions for each knot
	visited := make([]map[string]int, 0)

	// Number of knots to simulate (including head)
	numKnots := 10

	for i := 0; i < numKnots; i++ {
		// Every knot starts at 0,0
		knots = append(knots, [2]int{0,0})

		// Every knot needs a map of visited positions
		visited = append(visited, make(map[string]int))

		// Instantly add the starting position to the map
		visited[i]["0,0"] = 1
	}

	// Map of moves as bytes to their vector transformation
	moves := make(map[byte][]int)
	moves['R'] = []int{1,0}
	moves['L'] = []int{-1,0}
	moves['U'] = []int{0,1}
	moves['D'] = []int{0,-1}

	for l := range lines {
		line := lines[l]
		parts := strings.Split(line, " ")
		mag, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		// In this part I decided to recalculate for every individual movement rather than doing each line in one step
		// For whatever reason my previous approach was causing my answer to be too low
		for c := 0; c < mag; c++ {
			// This represents one step in the direction required by the instruction
			knots[0][0] += moves[line[0]][0]
			knots[0][1] += moves[line[0]][1]
			
			for i := 1; i < len(knots); i++ {
				for true {
					// Work out the component distances
					xSep := knots[i-1][0] - knots[i][0]
					ySep := knots[i-1][1] - knots[i][1]

					// Use pythagorus to work out how far the head and tail are from each other after the head moves
					dist := math.Sqrt(math.Pow(float64(xSep), 2) + math.Pow(float64(ySep), 2))

					// If they are more than 2 units apart, the tail is pulled toward the head
					if dist < 2 {break}

					// Use abs / actual to find the movement needed in the x and y direction and apply it
					// Only do this if the component distance is not 0 otherwise we get divide by 0 problems
					if xSep != 0 {knots[i][0] += int(math.Abs(float64(xSep)) / float64(xSep))}
					if ySep != 0 {knots[i][1] += int(math.Abs(float64(ySep)) / float64(ySep))}

					// I have to actually convert to string properly here because x and y could be negative
					// which doesn't play nice with ascii
					visited[i][strconv.Itoa(knots[i][0]) + "," + strconv.Itoa(knots[i][1])] = 1
				}
			}
		}
	}

	// The answer is the length of the last map in the array of maps
	println(len(visited[len(visited) - 1]))
}