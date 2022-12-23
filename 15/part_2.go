package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func distance(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0]))) + int(math.Abs(float64(a[1]-b[1])))
}

func main() {
	// Row we are checking
	Y := 2000000
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	// Array of sensors stored as [x, y] pairs
	sensors := make([][2]int, 0)
	// Array of beacons stored as [x, y] pairs
	beacons := make([][2]int, 0)
	// Map of all x coordinates that have a beacon at Y. I use a map so I can access elements using indexes rather than having to loop through an array
	onLine := make(map[int]int)
	
	for l := range lines {
		parts := strings.Split(lines[l], " ")

		sx, err := strconv.Atoi(parts[2][2 : len(parts[2])-1])
		sy, err := strconv.Atoi(parts[3][2 : len(parts[3])-1])
		bx, err := strconv.Atoi(parts[8][2 : len(parts[8])-1])
		by, err := strconv.Atoi(parts[9][2:])
		if err != nil {
			panic(err)
		}

		sensors = append(sensors, [2]int{sx, sy})
		beacons = append(beacons, [2]int{bx, by})

		// If we find a beacon at Y, add the X value to the map by incrementing the value at that point
		if by == Y {
			onLine[bx]++
		}
	}

	// Lists of the positive and negative lines that form the boundaries of the ranges of the sensors
	// I am representing lines in the form x - y = c where the two arrays below contain the values of c
	positiveLines := make([]int, 0)
	negativeLines := make([]int, 0)

	for i := range sensors {
		dist := distance(sensors[i], beacons[i])
		// Add the two positive lines which are above and below the sensor by a distance of dist
		positiveLines = append(positiveLines, sensors[i][0] + sensors[i][1] - dist)
		positiveLines = append(positiveLines, sensors[i][0] + sensors[i][1] + dist)
		// Add the two negative lines which are above and below the sensor by a distance of dist
		negativeLines = append(negativeLines, sensors[i][0] - sensors[i][1] - dist)
		negativeLines = append(negativeLines, sensors[i][0] - sensors[i][1] + dist)
	}

	var positive, negative int

	// compare every positive line to every subsequent positive line and every negative line to every subsequent negative line
	for i := range positiveLines {
		for j := i + 1; j < len(positiveLines); j++ {
			// The lines will be a distance of 2 apart if they are what we are looking for
			if math.Abs(float64(positiveLines[i] - positiveLines[j])) == 2 {
				positive = int(math.Min(float64(positiveLines[i]), float64(positiveLines[j]))) + 1
			}

			if math.Abs(float64(negativeLines[i] - negativeLines[j])) == 2 {
				negative = int(math.Min(float64(negativeLines[i]), float64(negativeLines[j]))) + 1
			}
		}
	}

	// This just comes from the puzzle brief
	tuningFrequency := (positive + negative) / 2 * (Y * 2) + (positive - negative) / 2

	println(tuningFrequency)
}
