package day8

import (
	"fmt"
	"strconv"
)

var field [][]int
var perimeterCount int

func Day8(input []string) []string {
	field = make([][]int, 0)

	for i, line := range input {
		if line == "" {
			continue
		}

		field = append(field, make([]int, len(line)))

		for j, val := range line {
			intVal, err := strconv.Atoi(string(val))

			if err != nil {
				panic(err)
			}

			field[i][j] = intVal
		}
	}

	visibleCount := 0
	maxScore := 0

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			isVisible, score := getVisibility(i, j)

			fmt.Printf("[%d, %d] (%d) ", i, j, field[i][j])
			if isVisible {
				visibleCount++
				fmt.Printf("(VISIBLE)")
			}
			fmt.Printf(" Score: %d\n", score)

			if score > maxScore {
				maxScore = score
			}
		}
	}

	return []string{fmt.Sprintf("Visible trees: %d (%d perimeter); Max score %d",
		visibleCount, perimeterCount, maxScore)}
}

func getVisibility(i, j int) (bool, int) {
	val := field[i][j]
	linesBlocked := make([]bool, 4)
	views := make([]int, 4)

	if i == 0 || i == len(field)-1 || j == 0 || j == len(field[i])-1 {
		perimeterCount++
		return true, 0
	}

	for v := i - 1; v >= 0; v-- {
		views[0]++
		if field[v][j] >= val {
			linesBlocked[0] = true
			break
		}
	}
	for v := i + 1; v < len(field); v++ {
		views[1]++
		if field[v][j] >= val {
			linesBlocked[1] = true
			break
		}
	}
	for u := j - 1; u >= 0; u-- {
		views[2]++
		if field[i][u] >= val {
			linesBlocked[2] = true
			break
		}
	}
	for u := j + 1; u < len(field[i]); u++ {
		views[3]++
		if field[i][u] >= val {
			linesBlocked[3] = true
			break
		}
	}

	visible := false
	score := 1
	for i := 0; i < 4; i++ {
		if !linesBlocked[i] {
			visible = true
		}

		score *= views[i]
	}
	return visible, score
}
