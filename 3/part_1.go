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

	lines := strings.Split(string(data), "\n")

	sum := 0

	for l := range lines {
		line := lines[l]
		// Split each line into equal sized halves
		cLen := len(line) / 2
		part1 := line[0:cLen]
		part2 := line[cLen:]

		counts := make(map[byte]int)

		var common byte

		// Count the letters in part 1
		for i := range part1 {
			counts[part1[i]] += 1
		}
		// If a letter was counted in part 1 (i.e. it's count is greater than 1) and it appears in part 2, it is the common letter
		for i := range part2 {
			if counts[part2[i]] > 0 {
				common = part2[i]
				break
			}
		}

		// Add the byte value to the sum
		sum += int(common)

		// Byte value != actual priority
		// Subtract 96 for lowercase letters and 38 for uppercase letters to get correct priority
		if common >= 97 {
			sum -= 96
		} else {
			sum -= 38
		}
	}

	println(sum)
}
