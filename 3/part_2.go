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

	// Using this style instead of range so I can control the incrememnt amount
	for l := 0; l < len(lines); l+=3 {
		// This will get a sub array containing 3 backpack strings
		triple := lines[l:l+3]
		
		counts := make(map[byte]int)

		var common byte

		for p := range triple {
			// Store a map of the letters seen for each backpack so we don't add them to counts twice
			used := make(map[byte]int)
			for i := range triple[p] {
				// If this letter is uncounted we should deal with it
				if used[triple[p][i]] < 1 {
					// Flag it as used
					used[triple[p][i]]++
					// If this is the third item in the triple and its been counted twice already, this is the common letter
					if p == 2 && counts[triple[p][i]] == 2 {
						common = triple[p][i]
						break;
					}
					// If we got this far we just need to increment the count for this letter
					counts[triple[p][i]]++
				}
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
