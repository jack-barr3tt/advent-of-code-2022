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

	intervals := make([][2]int, 0)
	minX, maxX := 9999999, 0

	// Calculate the intervals for each sensor and also the min and max X values at the same time because speed
	for i := range sensors {
		dx := distance(sensors[i], beacons[i]) - int(math.Abs(float64(sensors[i][1]-Y)))

		if dx > 0 {
			sL := sensors[i][0] - dx
			sU := sensors[i][0] + dx
			intervals = append(intervals, [2]int{sL, sU})

			if sL < minX {
				minX = sL
			}
			if sU > maxX {
				maxX = sU
			}
		}
	}

	result := 0

	// go along the line between the minimum and maximum X values, and check if each point is in range of any of the sensors
	for i := minX; i <= maxX; i++ {
		if onLine[i] > 0 {
			continue
		}

		for j := range intervals {
			if intervals[j][0] <= i && i <= intervals[j][1] {
				result++
				break
			}
		}
	}

	println(result)
}
