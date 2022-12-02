package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if(err != nil) {
		panic(err);
	}

	lines := strings.Split(string(data), "\n")

	max := [3]int{0,0,0}
	temp := 0

	for i := range lines {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			// Find the smallest max value
			min := 0
			for j := range max {
				if(max[j] < max[min]) {
					min = j
				}
			}

			// If this value is bigger, update it
			if temp > max[min] {
				max[min] = temp
			}

			// Reset temp counter
			temp = 0
			continue
		}
		temp += num
	}
	
	println(max[0] + max[1] + max[2])
}