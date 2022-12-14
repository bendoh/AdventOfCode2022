package day10

import (
	"fmt"
	"strconv"
)

var X = 1
var clock = 0
var step = 0

const (
	width  = 40
	height = 6
	pixels = width * height
)

var crt [pixels]rune

func compute(instruction string) int {
	var val int
	var err error

	if instruction == "" {
		return 1
	}

	if instruction[0:4] == "addx" {
		if step < 1 {
			step++
			return 0
		} else {
			val, err = strconv.Atoi(instruction[5:])

			if err != nil {
				panic(err)
			}

			step = 0
			X += val
			return 1
		}
	}

	if instruction == "noop" {
		return 1
	}

	panic("Invalid instruction")
}

func printCrt() {
	print("\033[2J\033[H")
	for i := 0; i < pixels; i++ {
		if i > 0 && i%width == 0 {
			println()
		}
		print(string(crt[i]))
	}
	println()
	println()
}
func Day10(input []string) []string {
	instr := 0
	total := 0

	for i := 0; i < pixels; i++ {
		crt[i] = ' '
	}

	for instr < len(input) {
		if X-1 <= clock%width && clock%width <= X+1 {
			crt[clock] = '*'
		}
		clock++
		if (clock-20)%40 == 0 {
			total += clock * X
		}
		instr += compute(input[instr])
		printCrt()
	}
	return []string{fmt.Sprintf("Sum of signals: %d", total)}
}
