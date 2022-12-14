package days

import (
	"AdventOfCode2022/days/day0"
	"AdventOfCode2022/days/day1"
	"AdventOfCode2022/days/day11"
	"AdventOfCode2022/days/day12"
	"AdventOfCode2022/days/day13"
	"AdventOfCode2022/days/day2"
	"AdventOfCode2022/days/day3"
	"AdventOfCode2022/days/day4"
	"AdventOfCode2022/days/day5"
	"AdventOfCode2022/days/day6"
	"AdventOfCode2022/days/day7"
	"AdventOfCode2022/days/day8"
	"AdventOfCode2022/days/day9"
	"bufio"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"io"
	"os"
)

type Day struct {
	Number      int
	CLI         func([]string) []string
	VisualSetup func([]string)
	VisualStep  func(*ebiten.Image, int64) ([]string, bool)
}

var dayFunctions = []Day{
	{0, day0.Day0, nil, nil},
	{1, day1.Day1, day1.VisualSetup, day1.VisualStep},
	{2, day2.Day2, nil, nil},
	{3, day3.Day3, nil, nil},
	{4, day4.Day4, day4.VisualSetup, day4.VisualStep},
	{5, day5.Day5, nil, nil},
	{6, day6.Day6, nil, nil},
	{7, day7.Day7, nil, nil},
	{8, day8.Day8, nil, nil},
	{9, day9.Day9, nil, nil},
	{10, nil, nil, nil},
	{11, day11.Day11, nil, nil},
	{12, day12.Day12, nil, nil},
	{13, day13.Day13, nil, nil},
}

func Get(day int) Day {
	return dayFunctions[day]
}

func NumberDays() int {
	return len(dayFunctions)
}

func GetInput(inputFile string) []string {
	input, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(input)
	var lines []string

	for {
		line, err := reader.ReadString('\n')

		if len(line) > 0 && line[len(line)-1] == '\n' {
			lines = append(lines, line[:len(line)-1])
		} else {
			lines = append(lines, line)
		}

		if err == io.EOF {
			break
		}

	}

	return lines
}

func GetInputLines(day int) []string {
	lines := []string{}

	inputFilename := fmt.Sprintf("days/day%d/input", day)

	if _, err := os.Stat(inputFilename); err == nil {
		lines = GetInput(inputFilename)
	}

	return lines
}
