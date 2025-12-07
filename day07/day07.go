package day07

import "github.com/its-felix/AdventOfCode2025/common"

const (
	GridItemEmpty    = GridItem(0)
	GridItemSplitter = GridItem(1)
)

type GridItem uint8

func SolvePart1(input <-chan string) int {
	beamStartPos, grid := parse(input)
	beamColumns := make([]bool, len(grid[0]))
	beamColumns[beamStartPos.Col()] = true

	splits := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if beamColumns[col] && grid[row][col] == GridItemSplitter {
				splits++
				beamColumns[col] = false

				if col-1 >= 0 {
					beamColumns[col-1] = true
				}

				if col+1 < len(grid[row]) {
					beamColumns[col+1] = true
				}
			}
		}
	}

	return splits
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
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
				row[i] = GridItemEmpty

			case '^':
				row[i] = GridItemSplitter

			case 'S':
				row[i] = GridItemEmpty
				beamStartPos = common.GridPos{len(grid), i}

			default:
				panic("invalid input")
			}
		}

		grid = append(grid, row)
	}

	return beamStartPos, grid
}
