package main

import (
	"os"
	"strconv"
	"strings"
)

func drawPixels(cycle, sprite int) {
	// x reperesents the x coordinate of the pixel being drawn right now
	// Because cycle starts at 1 but the screen is 0 indexed, I subtract 1 before using modulo
	x := (cycle - 1) % 40

	// Sprite is three pixels wide so we need to check for all three of its pixels
	if x == sprite || x == sprite + 1 || x == sprite + 2 {
		// If there's an overlap we print a "bright" pixel
		print("#")
	}else{
		// Otherwise we print a dark pixel
		print(".")
	}

	if x == 0 {
		// If the current x is divisible by 40 then we must be at the end of a line so print a new line
		println()
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

	for l := range lines {
		parts := strings.Split(lines[l], " ")
		
		if parts[0] == "addx" {
			operand, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			cycle++
			// Pixels are drawn after every cycle increment
			drawPixels(cycle, X)

			cycle++
			drawPixels(cycle, X)		

			X += operand
		}else{
			cycle++
			drawPixels(cycle, X)		
		}
	}

	// The output is rendered on the screen progressively by drawPixels()
}