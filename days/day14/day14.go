package day14

import (
	"fmt"
	"math"
	"strings"
)

func crd(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var minX, maxX, maxY int
var field map[string]rune
var sand [2]int
var startX = 500

func printMap(stopped int) {
	print("\033[2J\033[H")
	for y := 0; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if c, ok := field[crd(x, y)]; ok {
				print(string(c))
			} else if sand[0] == x && sand[1] == y {
				print("+")
			} else {
				print(" ")
			}
		}
		println()
	}
	fmt.Printf("%d stopped\n", stopped)
}
func Day14(input []string) []string {
	field = make(map[string]rune)
	minX = math.MaxInt

	for _, line := range input {
		parts := strings.Split(line, " -> ")

		for i := 1; i < len(parts); i++ {
			a, b := parts[i-1], parts[i]

			var sx, sy, ex, ey int
			fmt.Sscanf(a, "%d,%d", &sx, &sy)
			fmt.Sscanf(b, "%d,%d", &ex, &ey)

			minX = min(sx, min(ex, minX))
			maxX = max(sx, max(ex, maxX))
			maxY = max(sy, max(ey, maxY))

			if sx == ex {
				for y := min(sy, ey); y <= max(sy, ey); y++ {
					field[crd(sx, y)] = '#'
				}
			} else if sy == ey {
				for x := min(sx, ex); x <= max(sx, ex); x++ {
					field[crd(x, sy)] = '#'
				}
			} else {
				panic("Non-straight line asked for!")
			}
		}
	}

	part1 := computePart1()
	return []string{fmt.Sprintf("Part1 sand particles: %d\nPart2 sand particles with floor: %d\n", part1, computePart2(part1))}
}

func computePart1() int {
	sand = [2]int{startX, 1}
	stopped := 0

	printMap(stopped)

	for sand[1] <= maxY {
		var next = sand
		var ok bool
		next[1] += 1

		if _, ok = field[crd(next[0], next[1])]; !ok {
			sand = next
		} else if _, ok = field[crd(next[0]-1, next[1])]; !ok {
			sand = [2]int{next[0] - 1, next[1]}
		} else if _, ok = field[crd(next[0]+1, next[1])]; !ok {
			sand = [2]int{next[0] + 1, next[1]}
		} else {
			field[crd(sand[0], sand[1])] = 'o'
			stopped++
			sand = [2]int{startX, 1}
		}
		//		printMap(stopped)
	}

	printMap(stopped)

	return stopped
}

func computePart2(stopped int) int {
	sand = [2]int{startX, 0}
	maxY += 2

	printMap(stopped)

	for {
		var next = sand
		var ok bool
		next[1] += 1
		didStop := false
		for i := -1; i <= 1; i++ {
			field[crd(next[0]+i, maxY)] = '#'
		}
		if _, ok = field[crd(next[0], next[1])]; !ok {
			sand = next
		} else if _, ok = field[crd(next[0]-1, next[1])]; !ok {
			sand = [2]int{next[0] - 1, next[1]}
		} else if _, ok = field[crd(next[0]+1, next[1])]; !ok {
			sand = [2]int{next[0] + 1, next[1]}
		} else {
			field[crd(sand[0], sand[1])] = 'o'
			didStop = true
		}

		if sand[0] < minX {
			minX = sand[0]
		} else if sand[0] > maxX {
			maxX = sand[0]
		}
		//printMap(stopped)
		if didStop {
			stopped++
			if sand[0] == startX && sand[1] == 0 {
				break
			}
			sand = [2]int{startX, 0}
		}
	}

	printMap(stopped)

	return stopped
}
