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

	// This will store the location with the best scenic score
	bestScore := 0

	// Loop through every position
	for x := range grid {
		for y := range grid[x] {
			current := grid[x][y]

			// Each element represents a check direction. When it is true, it means we can see further, and when it's false the view is blocked
			check := []bool{true,true,true,true}
			// Store a counter for the view distances in each direction
			views := []int{0,0,0,0}
			// Both of the above could have been split into 8 variables rather than 2 arrays for readability purposes but oh well

			// Here I go away from the current position by 1 in every direction with each iteration
			// Doing it all in one loop is another sacrifice to readability but I wanted to reduce number of loops
			// Since it would take unneccessary effort to calculate the stopping condition for this loop, I will just use a variable
			loop := true
			for i := 1; loop; i++ {
				// I instantly set loop to false because if none of the if statements below are true then we might as well stop here
				loop = false
				if x + i < len(grid) && check[0] {
					loop = true
					views[0]++
					if grid[x+i][y] >= current {
						check[0] = false
					}
				}
				if x - i >= 0 && check[1]{
					loop = true
					views[1]++
					if grid[x-i][y] >= current {
						check[1] = false
					}
				}
				if y + i < len(grid[x]) && check[2] {
					loop = true
					views[2]++
					if grid[x][y+i] >= current {
						check[2] = false
					}
				}
				if y - i >= 0 && check[3] {
					loop = true
					views[3]++
					if grid[x][y-i] >= current {
						check[3] = false
					}
				}
			}

			// Score starts at 1 because I am going to multiply all the view distances together
			score := 1

			for i := range views {
				score *= views[i]
			}

			// If this score is the best we've had so far, update bestScore
			if score > bestScore {
				bestScore = score
			}
		}
	}

	println(bestScore)
}