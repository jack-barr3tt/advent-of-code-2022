package main

import (
	"os"
	"strconv"
	"strings"
)

// Split into separate function since this same code block needs to be run twice in two different places
func addSum(sum *int, cycle, X int) {
	if (cycle - 20) % 40 == 0 {
		*sum += cycle * X
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	// X is the register
	X := 1
	// Cycle counter
	cycle := 1

	// Sum of signal strengths
	sum := 0

	for l := range lines {
		parts := strings.Split(lines[l], " ")
		
		if parts[0] == "addx" {
			operand, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			// We need to check if we need to add to the sum every time cycle is incremented
			cycle++
			addSum(&sum, cycle, X)
			cycle++

			// After the second cycle is complete the value of the register is actually updated
			X += operand
		}else{
		//  This case is used when we get a no op instruction, in which case we just increment the counter and move on
			cycle++
		}
		// Check if we need to add to the sum for the second time
		addSum(&sum, cycle, X)
	}

	// Print the sum after running through all instructions
	println(sum)
}