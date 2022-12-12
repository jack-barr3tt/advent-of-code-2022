package main

import (
	"os"
	"strings"
)

type Node struct {
	height  byte
	working []int
	final   int
}

func getGrid(filename string) ([][]*Node, int, int, int, int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	grid := make([][]*Node, 0)

	var Sx, Sy, Ex, Ey int

	for x := range lines {
		temp := make([]*Node, 0)
		for y := range lines[x] {
			// To make our lives easier in the pathfinding, we switch the S and E letters out for their height equivalents
			// but store their locations in the appropriate variables
			letter := lines[x][y]
			if lines[x][y] == 'S' {
				Sx = x
				Sy = y
				letter = 'a'
			}
			if lines[x][y] == 'E' {
				Ex = x
				Ey = y
				letter = 'z'
			}
			temp = append(temp, &Node{letter, make([]int, 0), -1})
		}
		grid = append(grid, temp)
	}

	return grid, Sx, Sy, Ex, Ey
}

func shortestPathLength(grid [][]*Node, Sx, Sy, Ex, Ey int) int {
	// Essentially this is dijkstra's shortest path algorithm

	// The final value for the starting position should be set to 0 as the base case
	grid[Sx][Sy].final = 0

	// Start from the starting position
	x := Sx
	y := Sy

	// The four possible movements from any node, represented as vectors
	moves := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for x != Ex || y != Ey {
		for m := range moves {
			// Calculate the new x and y after the movement
			dx := x + moves[m][0]
			dy := y + moves[m][1]

			// The following cases mean this movement is invalid:
			// The new coordinates lie outside of the grid
			// The next node already has a final value
			// The next node is too high for our character to move to
			if dx < 0 || dx >= len(grid) || dy < 0 || dy >= len(grid[dx]) || grid[dx][dy].final != -1 || grid[dx][dy].height > grid[x][y].height+1 {
				continue
			}

			// If the next node has no working values, or the path to it is smaller than the smallest working value, we should add the path to the list of working values
			if len(grid[dx][dy].working) == 0 || grid[x][y].final+1 < grid[dx][dy].working[len(grid[dx][dy].working)-1] {
				grid[dx][dy].working = append(grid[dx][dy].working, grid[x][y].final+1)
			}
		}

		nextX := 0
		nextY := 0
		nextSmallest := grid[nextX][nextY]

		for cx := range grid {
			for cy := range grid[cx] {
				// If this node has no working values or its already finalised there's no point considering it
				if len(grid[cx][cy].working) == 0 || grid[cx][cy].final != -1 {
					continue
				}

				// If our next node has no working values or it's already finalised, we need to switch it out
				// However the main case for switching is if the node we're checking right now has a smallest minimum working value than the node in our nextSmallest variable
				if len(nextSmallest.working) == 0 || nextSmallest.final != -1 || grid[cx][cy].working[len(grid[cx][cy].working)-1] < nextSmallest.working[len(nextSmallest.working)-1] {
					nextSmallest = grid[cx][cy]
					nextX = cx
					nextY = cy
				}
			}
		}

		// If we get to this siuation we're lost
		if len(nextSmallest.working) == 0 {
			break
		}

		// Once we've found our best case, we set its final value...
		grid[nextX][nextY].final = nextSmallest.working[len(nextSmallest.working)-1]

		// ...and update our x and y values
		x = nextX
		y = nextY
	}

	return grid[Ex][Ey].final
}
func main() {
	filename := "input.txt"

	// We don't care about the starting values for this challenge because we will choose our own
	grid, _, _, Ex, Ey := getGrid(filename)

	starts := make([][2]int, 0)

	// Grab all the positions with a height of 'a'
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y].height == 'a' {
				starts = append(starts, [2]int{x, y})
			}
		}
	}

	best := -1

	for s := range starts {
		x := starts[s][0]
		y := starts[s][1]

		// We need to get the grid again every time so we're always starting afresh
		grid, _, _, Ex, Ey = getGrid(filename)

		// Get the shortest path for this starting point
		len := shortestPathLength(grid, x, y, Ex, Ey)

		// Sometimes the function will return -1, indicating no path found
		// Otherwise we just want to keep updating best with the shortest path
		if len > 0 && (best < 0 || len < best) {
			best = len
		}
	}

	// Output the best path length
	println(best)
}
