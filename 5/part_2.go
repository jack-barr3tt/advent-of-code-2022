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

	// Go doesn't have stacks and that so we just use slices
	lines := strings.Split(string(data), "\n")

	stacks := make([][]byte, 0)

	l := 0
	var s int

	// Parse crate stack data into actual stacks
	for lines[l][1] > 57 || lines[l][1] < 48 {
		for i := range lines[l] {
			c := lines[l][i]
			if c < 65 || c > 90 {
				continue
			}
			// The stack number can be found based on how far across the line we've got
			s = (i - 1) / 4
			// If we don't have enough stacks, make more until we do
			for s >= len(stacks) {
				stacks = append(stacks, make([]byte, 0))
			}
			// Add on the current data to the appropriate stack
			stacks[s] = append(stacks[s], c)
		}
		l++
	}
	// Skip forward 2 because the that's where the next useful information is
	l += 2

	// l currently points to the first instruction so we can just loop to the end from here
	for l < len(lines) {
		// instruction elements are separated by spaces
		parts := strings.Split(lines[l], " ")

		// get the count, source and destination
		count, err := strconv.Atoi(parts[1])
		src, err := strconv.Atoi(parts[3])
		dest, err := strconv.Atoi(parts[5])
		if err != nil {
			panic(err)
		}
		// Zero indexing is a thing so we decrement the parts so they refer to the right place is the slice
		src--
		dest--

		// Grab the stuff being moved
		tmp := stacks[src][0:count]
		// Remove it from the source
		stacks[src] = stacks[src][count:]

		// Save a copy of the destination as of now
		copy := stacks[dest]

		// Wipe the destination
		stacks[dest] = make([]byte, 0)

		// Add the stuff being moved to the destination - it can be in order this time because the crane can grab multiple things at once
		for i := range tmp {
			stacks[dest] = append(stacks[dest], tmp[i])
		}
		// Then add back the original contents as well
		for i := range copy {
			stacks[dest] = append(stacks[dest], copy[i])
		}

		l++
	}

	// Print the first thing from each stack because that's the answer
	for sn := range stacks {
		print(string(stacks[sn][0]))
	}
	println()
}