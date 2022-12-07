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

	lines := strings.Split(string(data), "\n")

	// Store a stack holding the current working directory
	dirS := make([]string, 0)
	// Store a variable that determines wether the next input contains data about a directory
	lsMode := false

	sizes := make(map[string]int)

	for l := range lines {
		line := lines[l]

		// When we start a new command we're no longer reading size data from ls
		if line[0] == '$' {
			lsMode = false
		}

		if lsMode {
			parts := strings.Split(line, " ")

			var size int
			
			// If this is a directory, grab any existing size info about this directory and set size equal to that
			if parts[0] != "dir" {
			// Otherwise get the size of the current file
			// This needs to be converted from a string to an int
				sizeC, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				size = sizeC
			}

			// The current item is contributing to the size of all its parents so ALL of them need to have their size count increased
			for d := range dirS {

				// To cater for directories with the same names, I just make the keys the full directory path
				temp := ""
				for i := 0; i <= d; i++ {
					temp += dirS[i]
					if dirS[i] != "/" {
						temp += "/"
					}
				}

				sizes[temp] += size
			}
		}

		if line[2:4] == "cd" {
			// If we are going up a directory, pop the current directory off the top of the stack
			if line[5:] == ".." {
				dirS = dirS[:len(dirS) - 1]
			}else{
			// Otherwise just add the current directory to the stack
				dirS = append(dirS, line[5:])
			}
		}else if line[2:4] == "ls" {
			lsMode = true
		}
	}

	// This will store the minimum dir size
	min := -1
	
	for _, v := range sizes {
		// We want the current space used minus the directory we could delete to be under 40M
		if sizes["/"] - v <= 40000000 {
			if min == -1 || v < min {
				min = v
			}
		}
	}

	println(min)
}