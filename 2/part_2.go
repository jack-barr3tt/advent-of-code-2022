package main

import (
	"os"
	"strings"
)

func getPlay(played, outcome byte) byte {
	if outcome == 89 {
		// Need to draw so play same as other player
		return played
	}
	if outcome == 90 {
		// Need to win
		if played == 65 {return 66}
		if played == 66 {return 67}
		if played == 67 {return 65}
	} else {
		// Need to lose
		if played == 65 {return 67}
		if played == 66 {return 65}
		if played == 67 {return 66}
	}
	return 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	score := 0

	for line := range lines {
		if lines[line] == "" {
			break
		}
		parts := strings.Split(lines[line], " ")
		plays := [2]byte{parts[0][0], parts[1][0]}

		// Part of the score is based on what you play
		playScore := getPlay(plays[0], plays[1]) - 64

		// The other part of the score is based on the outcome of the game
		resultScore := 3 * (plays[1] - 88)

		score += int(playScore) + int(resultScore)
	}

	println(score)
}
