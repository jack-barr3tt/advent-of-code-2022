package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

// Store the minimum and maximum x and y positions of rocks so we only store the relevant part of the cave
// These are global variables because basically everything needs them anyway
var minX, minY, maxX, maxY int

// Parse the rocks from input into a 2D array of point pairs
func getRocks(lines []string) [][][2]int {
	rocks := make([][][2]int, 0)
	// Each line represents a rock
	for l := range lines {
		temp := make([][2]int, 0)
		points := strings.Split(lines[l], " -> ")
		// Each point represents a vertex of the rock
		for p := range points {
			parts := strings.Split(points[p], ",")
			x, err := strconv.Atoi(parts[0])
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			temp = append(temp, [2]int{x, y})
		}
		rocks = append(rocks, temp)
	}
	return rocks
}

// Set the values of the min and max variables
func setMinMax(rocks [][][2]int) {
	for r := range rocks {
		for p := range rocks[r] {
			x := rocks[r][p][0]
			y := rocks[r][p][1]

			if x < minX || minX == 0 {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY || minY == 0 {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	// Extend the width of the grid to accomodate the maximum possible width of a sand pile	
	if 500 + (2 * maxY - 1) > maxX {
		maxX = 500 + (2 * maxY - 1)
	}
	if 500 - (2 * maxY - 1) < minX {
		minX = 500 - (2 * maxY - 1)
	}
	// Extend the height of the grid to accomodate for the floor
 	maxY += 2
}

// Just used for converting to non-variadic (thanks Go)
func convertCoordsArr(coords [2]int) (int, int) {
	return convertCoords(coords[0], coords[1])
}

// Basically just subtracts the mininum X value from X
func convertCoords(x, y int) (int, int) {
	return x - minX, y
}

// Returns the amount to move in the x and y direction according to this rock pair
func calculateMoveDir(rockpair [][2]int) [2]int {
	if len(rockpair) != 2 {
		panic("Not a pair")
	}

	// Magnitudes of the required movements
	xMag := rockpair[1][0] - rockpair[0][0]
	yMag := rockpair[1][1] - rockpair[0][1]
	moveDir := [2]int{0,0}

	// These if statements are to avoid divide by 0 errors
	if xMag != 0 {
		// This makes it so that the move direction is only every 0, 1 or -1
		moveDir[0] = xMag / int(math.Abs(float64(xMag)))
	}
	if yMag != 0 {
		moveDir[1] = yMag / int(math.Abs(float64(yMag)))
	}

	return moveDir
}

// Generates the grid and populates it with rocks
func generateGrid(rocks [][][2]int) [][]byte {
	// This makes a grid of the required size where every position is empty
	grid := make([][]byte, 0)
	for y := 0; y <= maxY; y++ {
		temp := make([]byte, 0)
		for x := 0; x <= maxX-minX; x++ {
			// The bottom row is also floor (made of rock)
			if y == maxY {
				temp = append(temp, '#')
			}else{
				temp = append(temp, '.')
			}
		}
		grid = append(grid, temp)
	}

	for r := range rocks {
		for p := 1; p < len(rocks[r]); p++ {
			moveDir := calculateMoveDir(rocks[r][p-1:p+1])

			x, y := convertCoordsArr(rocks[r][p-1])
			tarX, tarY := convertCoordsArr(rocks[r][p])

			// Keep placing down # until we reach the target point
			for x != tarX || y != tarY {
				grid[y][x] = '#'
				x += moveDir[0]
				y += moveDir[1]
			}
			grid[y][x] = '#'
		}
	}


	return grid
}

func dropSand(gridRef *[][]byte) bool {
	// The source of the sand is at (500,0)
	x, y := convertCoords(500,0)

	// Just for pointer funsies
	grid := *gridRef

	// If the source point is already filled the sand can't settle so return false
	if grid[y][x] != '.' {
		return false
	}

	for y < maxY {
		// If there is something under us
		if grid[y+1][x] != '.' {
			// If we can move left, go left
			if x - 1 < 0 || grid[y+1][x-1] == '.' {
				x--
			// Otherwise if can move right, go right
			}else if x + 1 >= len(grid[0]) || grid[y+1][x+1] == '.' {
				x++
			}else{
			// If there's nowhere else to go, just sit right here
				grid[y][x] = 'O'
				return true
			}
		}
		// If we've gone out of bounds, break out (meaning we return the sand doesn't settle)
		if x < 0 || x >= len(grid[0]) {
			break
		}

		y++
	}

	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	rocks := getRocks(lines)
	setMinMax(rocks)
	grid := generateGrid(rocks)

	// Count up the number of times we can drop sand
	sandCount := 0
	for dropSand(&grid) {
		sandCount++
	}

	println(sandCount)
}
