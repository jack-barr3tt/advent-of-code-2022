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

	max := 0
	temp := 0

	for i := range lines {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			// If this value is bigger, update it
			if temp > max {
				max = temp
			}

			// Reset temp counter
			temp = 0
			continue
		}
		temp += num
	}
	
	println(max)
}