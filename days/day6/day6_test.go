package day6

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay6(t *testing.T) {
	cases := []struct {
		input  string
		result int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, c := range cases {
		assert.Equal(t, []string{fmt.Sprintf("%d", c.result)}, Day6([]string{c.input}))
	}
}
