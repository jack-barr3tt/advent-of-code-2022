package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Split the input into lines
	lines := strings.Split(string(data), "\n")

	count := 0

	for l := range lines {
		// Split the line into the two elves
		elves := strings.Split(lines[l], ",")
		// Split each elf into the start and end of their ranges
		e1 := strings.Split(elves[0], "-")
		e2 := strings.Split(elves[1], "-")

		// Convert all the intervals to integers to make them easier to work with
		e1s, err := strconv.Atoi(e1[0])
		e1e, err := strconv.Atoi(e1[1])
		e2s, err := strconv.Atoi(e2[0])
		e2e, err := strconv.Atoi(e2[1])
		if err != nil {
			panic(err)
		}

		// Check if elf1's task is fully contained within elf2's task
		if e1s >= e2s && e1e <= e2e {
			count++
		// Check if elf2's task is fully contained within elf1's task
		} else if e2s >= e1s && e2e <= e1e {
			count++
		}
	}

	println(count)
}