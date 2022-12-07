package day6

import "fmt"

func isMarker(chars string) bool {
	counts := make(map[rune]int)

	for _, char := range chars {
		if _, ok := counts[char]; !ok {
			counts[char] = 0
		}
		counts[char]++

		if counts[char] > 1 {
			return false
		}
	}

	return true

}
func Day6(input []string) []string {
	line := input[0]
	var part1, part2 int

	for i := 4; i < len(line); i++ {
		if isMarker(line[i-4:i]) && part1 == 0 {
			part1 = i
		}
		if i >= 14 && isMarker(line[i-14:i]) && part2 == 0 {
			part2 = i
		}
	}

	return []string{fmt.Sprintf("%d", part1), fmt.Sprintf("%d", part2)}
}
