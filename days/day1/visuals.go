package day1

import (
	"AdventOfCode2022/util"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"strings"
)

var numbers []int

var part1done bool
var part2done bool

var finalResults []string

func VisualSetup(_input []string) {
	numbers = util.ParseIntList(_input)
	part1indices = []int{0, 1}
	part1done, part2done = false, false
	finalResults = []string{}
}

func VisualStep(screen *ebiten.Image, timeElapsed int64) ([]string, bool) {
	var results []string

	if !part1done {
		results = append(results, part1step(screen, timeElapsed)...)

		if part1done {
			finalResults = append(finalResults, results...)
		}
	} else if part1done && !part2done {
		results = append(results, part2step(screen, timeElapsed)...)

		if part2done {
			finalResults = append(finalResults, results...)
		}
	} else {
		ebitenutil.DebugPrint(screen, strings.Join(finalResults, "\n"))
	}

	if part1done && part2done {
		return finalResults, true
	} else {
		return results, false
	}
}

var lastStep int64

func part1step(screen *ebiten.Image, timeElapsed int64) []string {
  result := "Step 1!"
  ebitenutil.DebugPrint(screen, result)
  part1done = true
	return []string{result}
}

func part2step(screen *ebiten.Image, timeElapsed int64) []string {
  result := "Step 2!"
  ebitenutil.DebugPrint(screen, result)
  part2done = true
	return []string{result}
}
