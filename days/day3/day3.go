package day3

import (
	"fmt"
	"strings"
)

func makeScore(chr string) int {
	if chr[0] >= 'A' && chr[0] <= 'Z' {
		return int((chr[0] - 'A') + 27)
	} else if chr[0] >= 'a' && chr[0] <= 'z' {
		return int((chr[0] - 'a') + 1)
	}

	return 0
}

func Day3(input []string) []string {
	rucksacks := make([][]string, len(input))

	score := 0
	score2 := 0

	for i, line := range input {
		rucksacks[i] = make([]string, 2)

		rucksacks[i][0] = line[0 : len(line)/2]
		rucksacks[i][1] = line[len(line)/2:]

		found := ""
		for j := 0; j < len(rucksacks[i][0]); j++ {
			for k := 0; k < len(rucksacks[i][1]); k++ {
				if rucksacks[i][0][j] == rucksacks[i][1][k] {
					found += string(rucksacks[i][0][j])
				}
			}
		}
		if len(found) == 0 {
			panic("found none matching")
		}

		first := found[0]
		for k := 0; k < len(found); k++ {
			if found[k] != first {
				panic("Wasn't all the same!")
			}
		}

		if i%3 == 2 {
			for j := 0; j < len(input[i-2]); j++ {
				chr := string(input[i-2][j])

				if strings.Contains(input[i-1], chr) && strings.Contains(input[i], chr) {
					score2 += makeScore(chr)
					break
				}
			}
		}

		score += makeScore(found)
	}

	return []string{fmt.Sprintf("Score 1: %d  Score 2: %d", score, score2)}
}
