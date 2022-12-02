package main

import (
	"os"
	"strings"
)

func checkWin(a, b byte) bool {
	if a == 65 && b == 67 {return true}
	if a == 66 && b == 65 {return true}
	if a == 67 && b == 66 {return true}
	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	score := 0

	for line := range lines {
		if(lines[line] == "") {continue}
		parts := strings.Split(lines[line], " ")
		plays := [2]byte{ parts[0][0], parts[1][0] - 23 }

		// Part of your score is based on whether you play rock, paper or scissors
		typeScore := int(plays[1]) - 64

		// The other part of your score is based on the result of the game
		var resultScore int
		if plays[0] == plays[1] {
			resultScore = 3
		}else{
			if checkWin(plays[1], plays[0]) {
				resultScore = 6
			}else{
				resultScore = 0
			}
		}

		score += typeScore + resultScore
	}

	println(score)
}