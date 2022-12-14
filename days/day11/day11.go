package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operation func(int) int
	test      func(int) bool
	targets   [2]int
}

func Day11(input []string) []string {
	monkeys := make([]*monkey, 0)
	curMonkey, monkeyNum := 0, 0

	for _, line := range input {
		nParsed, err := fmt.Sscanf(line, "Monkey %d:", &monkeyNum)

		if err != nil {
			panic(err)
		}

		if nParsed > 0 {
			if len(monkeys) != monkeyNum {
				panic("Weird monkeys")
			}
			monkeys = append(monkeys, &monkey{items: []int{}, targets: [2]int{}})
			curMonkey = len(monkeys) - 1

			continue
		}

		starters := ""
		nParsed, err = fmt.Sscanf(line, "  Starting items: %s", &starters)

		if err != nil {
			panic(err)
		}

		for _, worryString := range strings.Split(starters, ", ") {
			var worryVal int
			worryVal, err = strconv.Atoi(worryString)

			if err != nil {
				panic(err)
			}

			monkeys[curMonkey].items = append(monkeys[curMonkey].items, worryVal)
		}

		if len(starters) > 0 {
			continue
		}

		operation := ""
		nParsed, err = fmt.Sscanf(line, "  Operation: new = %s", &operation)

		if err != nil {
			panic(err)
		}

		if nParsed > 0 {
			monkeys[curMonkey].operation = (func(op string) func(int) int {
				return func(old int) int {
					terms := make([]int, 0, 2)
					operation = ""

					for _, term := range strings.Split(" ", op) {
						val := 0

						if term == "old" {
							val = old
						} else if term == "*" || term == "+" {
							operation = term
							continue
						} else {
							val, err = strconv.Atoi(term)

							if err != nil {
								panic(err)
							}
						}
						terms = append(terms, val)
					}

					if operation == "+" {
						return terms[0] + terms[1]
					} else if operation == "*" {
						return terms[0] * terms[1]
					} else {
						panic("Unknown operation")
					}
				}
			})(operation)

			continue
		}

		divisor := 0
		nParsed, err = fmt.Sscanf(line, "  Test: divisible by %d", &divisor)

		if err != nil {
			panic(err)
		}

		if nParsed > 0 {
			monkeys[curMonkey].test = (func(d int) func(int) bool {
				return func(val int) bool {
					return val%d == 0
				}
			})(divisor)
			continue
		}

		var targetTrue, targetFalse int
		nParsed, err = fmt.Sscanf(line, "    If true: throw to monkey %d", &targetTrue)
		if nParsed > 0 {
			monkeys[curMonkey].targets[0] = targetTrue
			continue
		}
		nParsed, err = fmt.Sscanf(line, "    If false: throw to monkey %d", &targetFalse)
		if nParsed > 0 {
			monkeys[curMonkey].targets[1] = targetFalse
			continue
		}

	}

	// Play 20 rounds
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]
			for _, item := range m.items {
				nextWorry := m.operation(item)
				var target int

				if m.test(nextWorry) {
					target = m.targets[0]
				} else {
					target = m.targets[1]
				}

				monkeys[target].items = append(monkeys[target].items)
			}
		}

	}
	return []string{"Welcome to AoC 2022 in GO!"}
}
