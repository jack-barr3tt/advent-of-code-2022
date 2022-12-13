package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

// Recursively defined struct because we don't know how deeply nested the arrays will be
type Element struct {
	val      int
	children []Element
}

func ConvertElement(in string) Element {
	// If the current element is a number, make a new element with no children and this value
	thisNum, err := strconv.Atoi(in)
	if err == nil {
		return Element{thisNum, make([]Element, 0)}
	}

	children := make([]Element, 0)

	// If we've got this far then the first character of in will be [ so we start from index 1
	start := 1
	// I'm only using bracket count so I can ignore commas inside brackets
	bracketCount := 0
	for i := 1; i < len(in); i++ {
		if in[i] == '[' {
			bracketCount++
		}
		if in[i] == ']' {
			bracketCount--
		}
		if (in[i] == ',' && bracketCount == 0) || i == len(in)-1 {
			// Recursively get children
			children = append(children, ConvertElement(in[start:i]))
			start = i + 1
		}
	}

	return Element{-1, children}
}

// Returns >0 if in right order
// Returns <0 if in wrong order
// Returns 0 if undetermined
func orderCheck(left, right Element) int {
	// If both elements have a value
	if left.val >= 0 && right.val >= 0 {
		// This will return a positive, non-zero number if left is smaller, which is my right order case
		return right.val - left.val
	}

	// If only left has a value, put it inside a dummy element and then call the function using this
	if left.val >= 0 {
		return orderCheck(Element{-1, []Element{left}}, right)
	}
	// If only right has a value, put it inside a dummy element and then call the function using this
	if right.val >= 0 {
		return orderCheck(left, Element{-1, []Element{right}})
	}

	// Get the length of the longest children array
	maxLen := int(math.Max(float64(len(left.children)), float64(len(right.children))))

	for i := 0; i < maxLen; i++ {
		// If left runs out first, it's in the right order
		if i >= len(left.children) {
			return 1
		}
		// If right runs out first, it's in the wrong order
		if i >= len(right.children) {
			return -1
		}

		// Check the corresponding children from left and right
		check := orderCheck(left.children[i], right.children[i])

		// If a conclusion has been reached (i.e check is non zero) we should return that conclusion, otherwise keep looking
		if check != 0 {
			return check
		}
	}
	// If we got this far, we haven't determined if they are in the right order
	return 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	sum := 0

	// Useful data starts every 3 lines so just incrememnt the index by 3 each time
	for l := 0; l < len(lines); l += 3 {
		if orderCheck(ConvertElement(lines[l]), ConvertElement(lines[l+1])) > 0 {
			// A new pair is considered every 3 lines and indexing starts from 1
			sum += l/3 + 1
		}
	}

	println(sum)
}
