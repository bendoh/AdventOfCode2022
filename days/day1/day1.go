package day1

import (
	"AdventOfCode2022/util"
	"fmt"
)

func part1(numbers []int) string {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				return fmt.Sprintf("2020 = %d + %d, %d x %d = %d\n", numbers[i], numbers[j], numbers[i], numbers[j], numbers[i]*numbers[j])
			}
		}
	}
	return "<inconclusive!>"
}

func part2(numbers []int) string {
	return "<inconclusive!>"
}

func Day1(input []string) []string {
	numbers := util.ParseIntList(input)

	return []string{part1(numbers), part2(numbers)}
}
