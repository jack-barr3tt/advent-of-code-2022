package main

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int
	operator    byte
	operand     int
	test        int
	moveTrue    int
	moveFalse   int
	inspections int
}

func parseFile(file string) []Monkey {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	monkeys := make([]Monkey, 0)

	var newMonkey Monkey

	for l := range lines {
		if len(lines[l]) == 0 {
			continue
		}

		line := lines[l]

		if line[0] == 'M' {
			if l != 0 {
				monkeys = append(monkeys, newMonkey)
			}

			// A new monkey is being defined
			newMonkey = Monkey{}
		} else if line[2] == 'S' {
			// The starting items for a monkey are being defined
			itemsStr := strings.Split(line[18:], ", ")
			itemsInt := make([]int, 0)

			for i := range itemsStr {
				num, err := strconv.Atoi(itemsStr[i])
				if err != nil {
					panic(err)
				}

				itemsInt = append(itemsInt, num)
			}

			newMonkey.items = itemsInt
		} else if line[2] == 'O' {
			// The operation for a monkey is being defined

			if line[25] == 'o' {
				newMonkey.operator = '^'
			} else {
				newMonkey.operator = line[23]

				num, err := strconv.Atoi(line[25:])
				if err != nil {
					panic(err)
				}

				newMonkey.operand = num
			}

		} else if line[2] == 'T' {
			// The test for a monkey is being defined

			num, err := strconv.Atoi(line[21:])
			if err != nil {
				panic(err)
			}

			newMonkey.test = num
		} else if line[7] == 't' {
			// The true case for a monkey test is being defined

			num, err := strconv.Atoi(line[29:])
			if err != nil {
				panic(err)
			}

			newMonkey.moveTrue = num

		} else if line[7] == 'f' {
			// The false case for a monkey test is being defined

			num, err := strconv.Atoi(line[30:])
			if err != nil {
				panic(err)
			}

			newMonkey.moveFalse = num

		} else {
			panic("We shouldn't be here")
		}
	}
	monkeys = append(monkeys, newMonkey)

	return monkeys
}

// Use the euclidean algorithm for finding LCM which uses the greatest common factor
func greatestCommonFactor(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

// This is the other half of the euclidean algorithm
func lowestCommonMultiple(a, b int) int {
	return int(math.Abs(float64(a*b))) / greatestCommonFactor(a, b)
}

// Find the LCM of more than 2 numbers by finding LCM of the first two and then repeatedly finding the LCM of that number and the next in the array
func arrayLCM(ints []int) int {
	lcm := lowestCommonMultiple(ints[0], ints[1])
	for i := 2; i < len(ints); i++ {
		lcm = lowestCommonMultiple(lcm, ints[i])
	}
	return lcm
}

func main() {
	monkeys := parseFile("input.txt")

	rounds := 10000

	divs := make([]int, 0)
	for m := range monkeys {
		divs = append(divs, monkeys[m].test)
	}

	// This will be useful later
	lcm := arrayLCM(divs)

	for r := 0; r < rounds; r++ {
		for m := range monkeys {
			monkey := monkeys[m]

			for i := range monkey.items {
				monkeys[m].inspections++
				if monkey.operator == '*' {
					monkey.items[i] *= monkey.operand
				} else if monkey.operator == '+' {
					monkey.items[i] += monkey.operand
				} else if monkey.operator == '^' {
					monkey.items[i] *= monkey.items[i]
				} else {
					panic("Invalid operand")
				}
				// By taking the modulo of the number and the LCM of all divisors, we are reducing our worry level (as required by the brief)
				// but without comprimising our ability to check divisibility
				monkey.items[i] = monkey.items[i] % lcm

				if monkey.items[i]%monkey.test == 0 {
					monkeys[monkey.moveTrue].items = append(monkeys[monkey.moveTrue].items, monkey.items[i])
				} else {
					monkeys[monkey.moveFalse].items = append(monkeys[monkey.moveFalse].items, monkey.items[i])
				}
			}

			monkeys[m].items = monkeys[m].items[:0]
		}
	}

	// I could do something smart to get the top 2 inspection counts but since there's not many monkeys I'll just sort them
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	println(monkeys[0].inspections * monkeys[1].inspections)
}
