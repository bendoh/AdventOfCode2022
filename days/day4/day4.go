package day4

import (
	"fmt"
	"strings"
)

func Day4(input []string) []string {
	overlaps := 0
	any := 0
	for _, line := range input {
		counts := make(map[int]int)

		pairs := strings.Split(line, ",")

		range1 := make([]int, 2)
		range2 := make([]int, 2)

		fmt.Sscanf(pairs[0], "%d-%d", &range1[0], &range1[1])
		fmt.Sscanf(pairs[1], "%d-%d", &range2[0], &range2[1])

		for i := range1[0]; i <= range1[1]; i++ {
			if _, ok := counts[i]; !ok {
				counts[i] = 0
			}
			counts[i]++
		}
		for i := range2[0]; i <= range2[1]; i++ {
			if _, ok := counts[i]; !ok {
				counts[i] = 0
			}
			counts[i]++
		}

		overlap := 0
		for _, c := range counts {
			if c == 2 {
				overlap++
			}
		}

		if overlap > 0 {
			any++
		}
		if overlap == (range1[1]-range1[0]+1) || overlap == (range2[1]-range2[0]+1) {
			overlaps++
		}
	}
	return []string{fmt.Sprintf("Out of %d there are %d fully overlapping, %d with any overlap", len(input), overlaps, any)}
}
