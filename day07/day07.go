package day07

import (
	"github.com/its-felix/AdventOfCode2025/common"
)

const (
	GridItemSpace    = GridItem(0)
	GridItemSplitter = GridItem(1)
)

type GridItem uint8

func SolvePart1(input <-chan string) int {
	beamStartPos, grid := parse(input)
	beamColumns := make([]bool, len(grid[0]))
	beamColumns[beamStartPos.Col()] = true

	splits := 0

	for row := 0; row < len(grid); row++ {
		nextRowBeamColumns := make([]bool, len(grid[row]))

		for col := 0; col < len(grid[row]); col++ {
			if beamColumns[col] {
				switch grid[row][col] {
				case GridItemSpace:
					nextRowBeamColumns[col] = true

				case GridItemSplitter:
					splits++

					if col-1 >= 0 {
						nextRowBeamColumns[col-1] = true
					}

					if col+1 < len(grid[row]) {
						nextRowBeamColumns[col+1] = true
					}
				}
			}
		}

		beamColumns = nextRowBeamColumns
	}

	return splits
}

func SolvePart2(input <-chan string) int {
	beamStartPos, grid := parse(input)
	beamColumnTimelines := make([]int, len(grid[0]))
	beamColumnTimelines[beamStartPos.Col()] = 1
	timelines := 1

	for row := 0; row < len(grid); row++ {
		nextRowBeamColumnTimelines := make([]int, len(grid[row]))

		for col := 0; col < len(grid[row]); col++ {
			if beamColumnTimelines[col] > 0 {
				switch grid[row][col] {
				case GridItemSpace:
					nextRowBeamColumnTimelines[col] += beamColumnTimelines[col]

				case GridItemSplitter:
					timelines += beamColumnTimelines[col]

					if col-1 >= 0 {
						nextRowBeamColumnTimelines[col-1] += beamColumnTimelines[col]
					}

					if col+1 < len(grid[row]) {
						nextRowBeamColumnTimelines[col+1] += beamColumnTimelines[col]
					}
				}
			}
		}

		beamColumnTimelines = nextRowBeamColumnTimelines
	}

	return timelines
}

func parse(input <-chan string) (common.GridPos, common.Grid[GridItem]) {
	var beamStartPos common.GridPos
	grid := make(common.Grid[GridItem], 0)

	for line := range input {
		runes := []rune(line)
		row := make([]GridItem, len(runes))
		for i, c := range runes {
			switch c {
			case '.':
				row[i] = GridItemSpace

			case '^':
				row[i] = GridItemSplitter

			case 'S':
				row[i] = GridItemSpace
				beamStartPos = common.GridPos{len(grid), i}

			default:
				panic("invalid input")
			}
		}

		grid = append(grid, row)
	}

	return beamStartPos, grid
}
