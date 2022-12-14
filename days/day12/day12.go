package day12

import "fmt"

type point [2]int

var field [][]int

var loc, target point

/* BFS */
/*
func search(start point) int {
	paths := [][]point{{start}}
	for {

		steps := [4]point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

		nextPaths := make([][]point, 0)

		for _, thisPath := range paths {
			cur := thisPath[len(thisPath)-1]
			pos := field[cur[0]][cur[1]]
			height := pos.val
			for k, s := range steps {
				if pos.history[k] {
					continue
				}
				pos.history[k] = true
				var nextPoint point
				nextPoint[0] = cur[0] + s[0]
				nextPoint[1] = cur[1] + s[1]

				if nextPoint[0] < 0 ||
					nextPoint[0] >= len(field) ||
					nextPoint[1] < 0 ||
					nextPoint[1] >= len(field[0]) {
					continue
				}

				nextStep := field[nextPoint[0]][nextPoint[1]].val

				if nextStep == 27 {
					return len(thisPath)
				}
				if nextStep <= height+1 {
					seen := false
					for i := 0; i < len(thisPath); i++ {
						if thisPath[i][0] == nextPoint[0] && thisPath[i][1] == nextPoint[1] {
							seen = true
						}
					}

					if seen {
						continue
					}

					nextPaths = append(nextPaths, append(thisPath, nextPoint))
				}
			}
		}

		paths = nextPaths
	}

	return 0
}
*/

var width, height int

/* DFS */
func crd(s point) int {
	return s[0] + s[1]*width
}

type qi struct {
	p point
	d int
}

func search(start point) int {
	steps := [4]point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	cost := make(map[int]int)
	for l, line := range field {
		for j, _ := range line {
			cost[crd(point{l, j})] = 9999
		}
	}
	cost[crd(start)] = 0
	q := []qi{{start, 0}}

	for len(q) > 0 {
		nq := make([]qi, 0)

		for _, p := range q {
			h := field[p.p[0]][p.p[1]]

			for _, s := range steps {
				var np point
				np[0] = p.p[0] + s[0]
				np[1] = p.p[1] + s[1]
				if np[0] < 0 || np[0] >= len(field) ||
					np[1] < 0 || np[1] >= len(field[0]) {
					continue
				}

				coord := crd(np)

				if cost[coord] <= p.d+1 {
					continue
				}

				cost[coord] = p.d + 1
				nextStep := field[np[0]][np[1]]

				if nextStep == 27 {
					return p.d
				}

				if nextStep <= h+1 {
					nq = append(nq, qi{np, p.d + 1})
				}
			}
		}

		q = nq
	}

	return 0
}

func Day12(input []string) []string {
	field = make([][]int, len(input))
	for i, line := range input {
		field[i] = make([]int, len(line))

		for j, c := range line {
			if c == 'S' {
				loc[0] = i
				loc[1] = j
				field[i][j] = -1
			} else if c == 'E' {
				target[0] = i
				target[1] = j
				field[i][j] = 27
			} else {
				field[i][j] = int(c - 'a')
			}
		}
	}
	height = len(field)
	width = len(field[0])

	nSteps := search(loc)
	return []string{fmt.Sprintf("Number of steps: %d", nSteps)}
}
