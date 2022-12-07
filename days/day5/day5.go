package day5

import (
	"fmt"
)

func Day5(input []string) []string {
	var i int
	var line string

	var stacks [9][]rune
	var stacks2 [9][]rune

	for j := 0; j < len(stacks); j++ {
		stacks[j] = make([]rune, 0)
		stacks2[j] = make([]rune, 0)
	}

	for i, line = range input {
		if line[1] == '1' {
			break
		}

		for j := 0; j < len(stacks) && j*4+1 < len(line); j++ {
			if line[1+j*4] != ' ' {
				stacks[j] = append(stacks[j], rune(line[1+j*4]))
				stacks2[j] = append(stacks2[j], rune(line[1+j*4]))
			}
		}
	}

	i += 2

	for i, line = range input[i:] {
		var qty, src, dest int

		fmt.Sscanf(line, "move %d from %d to %d", &qty, &src, &dest)

		sourceStack := &stacks[src-1]
		destStack := &stacks[dest-1]

		for i := 0; i < qty; i++ {
			box := (*sourceStack)[0]
			*sourceStack = (*sourceStack)[1:]

			*destStack = append([]rune{box}, *destStack...)
		}

		sourceStack2 := &stacks2[src-1]
		destStack2 := &stacks2[dest-1]

		boxes := make([]rune, qty)
		copy(boxes, (*sourceStack2)[0:qty])

		*sourceStack2 = (*sourceStack2)[qty:]

		*destStack2 = append(boxes, *destStack2...)
	}

	res := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			res += string(stack[0])
		}
	}
	res2 := ""
	for _, stack := range stacks2 {
		if len(stack) > 0 {
			res2 += string(stack[0])
		}
	}
	return []string{"Top of each stack CrateMover 9000: " + res,
		"\nTop of each stack CrateMover 9001: " + res2}
}
