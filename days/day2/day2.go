package day2

import (
	"fmt"
	"strings"
)

func Day2(input []string) []string {
	guide := make([][]string, len(input))
	for i, line := range input {
		guide[i] = strings.Split(line, " ")
	}

	return []string{part1(guide), part2(guide)}
}

func part1(guide [][]string) string {
	score := 0
	for _, pair := range guide {
		if len(pair) < 2 {
			continue
		}
		a := pair[0]
		b := pair[1]
		shapescore := 0
		winscore := 3

		switch a {
		case "A":
			switch b {
			case "X": // rock X draws rock A
				shapescore = 1
				winscore = 3
			case "Y": // paper Y beats rock A
				shapescore = 2
				winscore = 6
			case "Z": // scissors Z loses to rock A
				shapescore = 3
				winscore = 0
			}
		case "B": // paper
			switch b {
			case "X": // rock X loses to paper B
				shapescore = 1
				winscore = 0
			case "Y": // paper Y draws to paper B
				shapescore = 2
				winscore = 3
			case "Z": // scissors Z beats paper B
				shapescore = 3
				winscore = 6
			}
		case "C": // scissors
			switch b {
			case "X": // scissors beats paper B
				shapescore = 1
				winscore = 6
			case "Y": // paper Y draws to paper B
				shapescore = 2
				winscore = 0
			case "Z": // draw
				shapescore = 3
				winscore = 3
			}
		}

		score += shapescore + winscore
	}
	return fmt.Sprintf("Part 1: %d", score)
}

func part2(guide [][]string) string {
	score := 0
	for _, pair := range guide {
		if len(pair) < 2 {
			continue
		}
		a := pair[0]
		b := pair[1]
		shapescore := 0
		winscore := 3

		switch a {
		case "A":
			switch b {
			case "X": // needs to lose to Rock: scissors
				shapescore = 3
				winscore = 0
			case "Y": // needs to draw to Rock: rock
				shapescore = 1
				winscore = 3
			case "Z": // needs to win to Rock: paper
				shapescore = 2
				winscore = 6
			}
		case "B": // paper
			switch b {
			case "X": // needs to lose to Paper: rock
				shapescore = 1
				winscore = 0
			case "Y": // needs to draw to Paper: paper
				shapescore = 2
				winscore = 3
			case "Z": // needs to win against Paper: scissors
				shapescore = 3
				winscore = 6
			}
		case "C": // scissors
			switch b {
			case "X": // needs to lose to scissors: paper
				shapescore = 2
				winscore = 0
			case "Y": // needs to draw scissors: scissors
				shapescore = 3
				winscore = 3
			case "Z": // needs to win against scissors: rock
				shapescore = 1
				winscore = 6
			}
		}

		score += shapescore + winscore
	}
	return fmt.Sprintf("Part2: %d", score)
}
