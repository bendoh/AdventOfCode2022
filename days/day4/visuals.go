package day4

import (
	"AdventOfCode2022/config"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	StartingBlocksTall  = 50
	StartingBlockHeight = config.GameHeight / StartingBlocksTall
	Speed               = 4000
	NumStacked          = 10 // 10 blocks at the bottom
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
	blocks[0] = ebiten.NewImage(blockWidth, StartingBlockHeight)
	blocks[1] = ebiten.NewImage(blockWidth, StartingBlockHeight)
	blocks[2] = ebiten.NewImage(blockWidth, StartingBlockHeight)

	blocks[0].Fill(color.RGBA{128, 128, 0, 255})
	blocks[1].Fill(color.RGBA{0, 128, 128, 255})
	blocks[2].Fill(color.RGBA{128, 255, 128, 255})
	bottomSet = make([][]int, 10)

	for i := 0; i < 10; i++ {
		bottomSet[i] = make([]int, max-min+1)
	}
}

var nOverlaps, nAny int
var index int

func drawLine(screen *ebiten.Image, yPos int, heightScale float64, values []int) int {
	drawn := 0
	trans := ebiten.GeoM{}

	height := heightScale * StartingBlockHeight

	trans.Scale(0, height)
	trans.Translate(0, float64(yPos)-height)

	for _, v := range values {
		trans.Translate(float64(blockWidth), 0)
		if v > 0 {
			drawn++
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

var lastDrop int64

func VisualStep(screen *ebiten.Image, timeElapsed int64) ([]string, bool) {
	if index >= len(parsed) {
		return []string{fmt.Sprintf("Out of %d there are %d fully overlapping, %d with any overlap", len(parsed), nOverlaps, nAny)}, true
	}

	seconds := float64(timeElapsed-lastDrop) / 1e9
	pos := int(Speed * seconds)

	// Draw and count the bottom set
	numDropped := index

	if numDropped > NumStacked {
		numDropped = NumStacked
	}

	for i := 0; i < NumStacked; i++ {
		height := numDropped
		yPos := config.GameHeight - height - i
		drawLine(screen, yPos, float64(height), bottomSet[NumStacked-i-1])
	}

	if pos < index {
		drawLine(screen, config.GameHeight-pos*StartingBlockHeight, StartingBlockHeight, renderLine(parsed[index]))
	} else {
		line := renderLine(parsed[index])
		count := 0

		for _, c := range line {
			if c == 2 {
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

		lastDrop = timeElapsed
		index++

		if index >= NumStacked {
			for i := NumStacked - 1; i > 0; i-- {
				copy(bottomSet[i], bottomSet[i-1])
			}
			copy(bottomSet[0], line)
		} else {
			copy(bottomSet[NumStacked-numDropped-1], line)
		}
	}

	res := fmt.Sprintf("Index: %d, Overlaps: %d  Any: %d", index, nOverlaps, nAny)

	ebitenutil.DebugPrint(screen, res)
	return []string{res}, false
}
