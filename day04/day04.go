package day04

import "math"

const (
	roll  = rune('@')
	space = rune('.')
)

func SolvePart1(input <-chan string) int {
	grid := parse(input)
	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == roll && countAdjacent(grid, row, col, roll) < 4 {
				count++
			}
		}
	}

	return count
}

func SolvePart2(input <-chan string) int {
	grid := parse(input)
	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			count += processPosition(grid, row, col)
		}
	}

	return count
}

func processPosition(grid [][]rune, row, col int) int {
	removed := 0

	if grid[row][col] == roll {
		adjacentRolls := findAdjacent(grid, row, col, roll)
		if len(adjacentRolls) < 4 {
			grid[row][col] = space
			removed++

			for _, pos := range adjacentRolls {
				removed += processPosition(grid, pos[0], pos[1])
			}
		}
	}

	return removed
}

func countAdjacent(grid [][]rune, currRow, currCol int, match rune) int {
	return len(findAdjacent(grid, currRow, currCol, match))
}

func findAdjacent(grid [][]rune, currRow, currCol int, match rune) [][2]int {
	positions := make([][2]int, 0)

	for row := max(currRow-1, 0); row <= min(currRow+1, len(grid)-1); row++ {
		for col := max(currCol-1, 0); col <= min(currCol+1, len(grid[row])-1); col++ {
			if (row != currRow || col != currCol) && grid[row][col] == match {
				positions = append(positions, [2]int{row, col})
			}
		}
	}

	return positions
}

func parse(input <-chan string) [][]rune {
	grid := make([][]rune, 0)
	minWidth, maxWidth := math.MaxInt, 0

	for line := range input {
		row := make([]rune, len(line))
		for i, c := range line {
			if c != space && c != roll {
				panic("invalid input")
			}

			row[i] = c
		}

		minWidth = min(minWidth, len(row))
		maxWidth = max(maxWidth, len(row))

		if minWidth != maxWidth {
			panic("invalid input")
		}

		grid = append(grid, row)
	}

	return grid
}
