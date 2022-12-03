package days

import (
	"AdventOfCode2022/days/day0"
	"AdventOfCode2022/days/day1"
	"AdventOfCode2022/days/day2"
	"AdventOfCode2022/days/day3"
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
