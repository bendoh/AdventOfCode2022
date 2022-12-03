package day1

import (
	"fmt"
	"sort"
	"strconv"
)

func Day1(input []string) []string {
	groups := make([][]int, 0)
	totals := make([]int, 0)

	group := make([]int, 0)
	total := 0

	for _, line := range input {
		if line == "" {
			totals = append(totals, total)
			total = 0
			groups = append(groups, group)
			continue
		}
		num, _ := strconv.Atoi(line)

		total += num
		group = append(group, num)
	}
	totals = append(totals, total)
	groups = append(groups, group)

	max := 0
	holder := 0
	for i, t := range totals {
		if t > max {
			max = t
			holder = i
		}
	}
	sort.Ints(totals)

	top3total := 0
	for i := 0; i < 3; i++ {
		top3total += totals[len(totals)-3+i]
	}
	return []string{
		fmt.Sprintf("Maximum calories: %d for holder %d", max, holder+1),
		fmt.Sprintf("Top 3 total: %d", top3total),
	}
}
