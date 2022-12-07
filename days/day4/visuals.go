package day4

import (
	"AdventOfCode2022/config"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"strings"
)

var isDone bool

var counts []map[int]int
var parsed [][][]int

var min, max int

func setMax(a, b int) {
	if a > max {
		max = a
	}
	if b > max {
		max = b
	}
}
func setMin(a, b int) {
	if a < min {
		min = a
	}
	if b < min {
		min = b
	}
}

var blocks []*ebiten.Image
var bottomSet [][]int

const (
	BlocksTall  = 50
	BlockHeight = config.GameHeight / BlocksTall
	Speed       = 15
	NumStacked  = 10 // 10 blocks at the bottom
)

var blockWidth int

func VisualSetup(_input []string) {
	counts = make([]map[int]int, 0)
	parsed = make([][][]int, len(_input))

	for i, line := range _input {
		pairs := strings.Split(line, ",")

		parsed[i] = make([][]int, 2)
		parsed[i][0] = make([]int, 2)
		parsed[i][1] = make([]int, 2)

		fmt.Sscanf(pairs[0], "%d-%d", &parsed[i][0][0], &parsed[i][0][1])
		fmt.Sscanf(pairs[1], "%d-%d", &parsed[i][1][0], &parsed[i][1][1])

		setMax(parsed[i][0][1], parsed[i][1][1])
		setMin(parsed[i][0][0], parsed[i][1][0])
	}

	blockWidth = config.GameWidth / (max - min)

	blocks = make([]*ebiten.Image, 3)
	blocks[0] = ebiten.NewImage(blockWidth, BlockHeight)
	blocks[1] = ebiten.NewImage(blockWidth, BlockHeight)
	blocks[2] = ebiten.NewImage(blockWidth, BlockHeight)

	blocks[0].Fill(color.RGBA{255, 0, 255, 255})
	blocks[1].Fill(color.RGBA{0, 255, 0, 255})
	blocks[2].Fill(color.RGBA{255, 255, 255, 255})
	bottomSet = make([][]int, 10)

	for i := 0; i < 10; i++ {
		bottomSet[i] = make([]int, max-min+1)
	}
}

var nOverlaps, nAny int
var index int

func drawLine(screen *ebiten.Image, yPos int, values []int) int {
	drawn := 0
	trans := ebiten.GeoM{}
	trans.Translate(0, float64(config.GameHeight-yPos*BlockHeight))

	for _, v := range values {
		if v > 0 {
			drawn++
			trans.Translate(float64(blockWidth), 0)
			screen.DrawImage(blocks[v-1], &ebiten.DrawImageOptions{GeoM: trans})
		}
	}

	return drawn
}

func renderLine(ranges [][]int) []int {
	res := make([]int, max-min+1)

	for i := ranges[0][0]; i <= ranges[0][1]; i++ {
		res[i]++
	}
	for i := ranges[1][0]; i <= ranges[1][1]; i++ {
		res[i] += 2
	}

	return res
}

var lastPos int

func VisualStep(screen *ebiten.Image, timeElapsed int64) ([]string, bool) {
	if index > len(parsed) {
		return []string{fmt.Sprintf("Out of %d there are %d fully overlapping, %d with any overlap", len(parsed), nOverlaps, nAny)}, true
	}

	seconds := float64(timeElapsed) / 1e9
	pos := int(Speed*seconds) - lastPos

	// Draw and count the bottom set
	foundStacked := 0
	for i := 0; i < NumStacked; i++ {
		drawn := drawLine(screen, BlocksTall-i, bottomSet[NumStacked-i-1])

		if drawn > 0 {
			foundStacked++
		}
	}

	if pos < BlocksTall-foundStacked {
		drawLine(screen, BlocksTall-pos, renderLine(parsed[index]))
	} else {
		line := renderLine(parsed[index])
		count := 0

		for _, c := range line {
			if c == 2 {
				nAny++
				count++
			}
		}

		if count > 0 {
			nAny++
		}
		if count == (parsed[index][0][1]-parsed[index][0][0]+1) ||
			count == (parsed[index][1][1]-parsed[index][1][0]+1) {
			nOverlaps++
		}

		lastPos = pos
		pos = 0
		index++

		if foundStacked == NumStacked {
			for i := NumStacked - 1; i > 0; i-- {
				copy(bottomSet[i], bottomSet[i-1])
			}
		} else {
			copy(bottomSet[NumStacked-foundStacked-1], line)
		}
	}

	return []string{}, false
}
