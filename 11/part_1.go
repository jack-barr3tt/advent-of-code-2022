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
				newMonkey.operand = 2
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

func main() {
	monkeys := parseFile("input.txt")

	rounds := 20

	for r := 0; r < rounds; r++ {
		for m := range monkeys {
			monkey := monkeys[m]

			temp := make([]int, 0)

			for i := range monkey.items {
				monkeys[m].inspections++
				newNum := monkey.items[i]
				if monkey.operator == '*' {
					newNum *= monkey.operand
				} else if monkey.operator == '+' {
					newNum += monkey.operand
				} else if monkey.operator == '^' {
					newNum = int(math.Pow(float64(newNum), 2))
				} else {
					panic("Invalid operand")
				}
				temp = append(temp, newNum/3)
			}

			for i := range temp {
				if temp[i]%monkey.test == 0 {
					monkeys[monkey.moveTrue].items = append(monkeys[monkey.moveTrue].items, temp[i])
				} else {
					monkeys[monkey.moveFalse].items = append(monkeys[monkey.moveFalse].items, temp[i])
				}
			}

			monkeys[m].items = make([]int, 0)
		}
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	println(monkeys[0].inspections * monkeys[1].inspections)
}
