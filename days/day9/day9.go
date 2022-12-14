package day9

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func absDiff(a, b int) int {
	diff := a - b

	if diff < 0 {
		return -diff
	}

	return diff
}

func printField(knots [][]int) {

	minX := math.MaxInt
	maxX := 0
	minY := math.MaxInt
	maxY := 0

	for _, knot := range knots {
		if knot[0] < minX {
			minX = knot[0]
		}
		if knot[0] > maxX {
			maxX = knot[0]
		}
		if knot[1] < minY {
			minY = knot[1]
		}
		if knot[1] > maxY {
			maxY = knot[1]
		}
	}

	var leftX, rightX, topY, bottomY int
	for leftX = 0; leftX >= minX; leftX -= 5 {
	}
	for rightX = 0; rightX <= maxX; rightX += 5 {
	}
	for topY = 0; topY >= minY; topY -= 5 {
	}
	for bottomY = 0; bottomY <= maxY; bottomY += 5 {
	}

	height := bottomY - topY + 5
	width := rightX - leftX + 5

	field := make([][]byte, height)
	for i := 0; i < height; i++ {
		field[i] = bytes.Repeat([]byte{'.'}, width)
	}

	for i := len(knots) - 1; i >= 0; i-- {
		x, y := knots[i][0], knots[i][1]
		chr := i + '0'

		if i == 0 {
			chr = 'H'
		}

		field[y-topY][x-leftX] = byte(chr)
	}

	fmt.Print("\033[2J\033[H")
	for _, line := range field {
		fmt.Println(string(line))
	}
}

func Day9(input []string) []string {
	visited := make([]map[string]int, 2) // 0 => knot 1, 1 => knot 9
	visited[0] = make(map[string]int)
	visited[1] = make(map[string]int)
	knots := make([][]int, 10)
	for i := 0; i < len(knots); i++ {
		knots[i] = []int{0, 0}
	}

	printField(knots)
	for _, line := range input {
		if line == "" {
			continue
		}
		dir := line[0]
		dist, err := strconv.Atoi(line[2:])

		if err != nil {
			panic(err)
		}

		var step []int

		switch dir {
		case 'R':
			step = []int{1, 0}
		case 'L':
			step = []int{-1, 0}
		case 'U':
			step = []int{0, -1}
		case 'D':
			step = []int{0, 1}
		}

		for i := 0; i < dist; i++ {
			next := make([]int, 2)
			before := make([]int, 2)
			copy(before, knots[0])

			knots[0][0] += step[0]
			knots[0][1] += step[1]

			for j := 1; j < len(knots); j++ {
				copy(next, knots[j])
				if absDiff(knots[j-1][0], next[0]) > 1 || absDiff(knots[j-1][1], next[1]) > 1 {
					copy(knots[j], knots[j-1])
				}
				copy(before, next)
			}

			printField(knots)
		}

		knot1 := fmt.Sprintf("%dx%d", knots[1][0], knots[1][1])
		knot9 := fmt.Sprintf("%dx%d", knots[9][0], knots[9][1])

		visited[0][knot1]++
		visited[1][knot9]++
		printField(knots)
		fmt.Printf("After %s\n", line)
		fmt.Printf("n1=%d  n9=%d\n", len(visited[0]), len(visited[1]))
	}
	return []string{fmt.Sprintf("Node 1 has visited %d nodes, node 9 has visited %d nodes", len(visited[0]), len(visited[1]))}
}
