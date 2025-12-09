package day09

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1(input <-chan string) int {
	positions := parse(input)
	slices.SortFunc(positions, func(a, b [2]int) int {
		return cmp.Or(
			a[0]-b[0],
			a[1]-b[1],
		)
	})

	largest := 0
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			width := max(positions[i][0], positions[j][0]) - min(positions[i][0], positions[j][0]) + 1
			height := max(positions[i][1], positions[j][1]) - min(positions[i][1], positions[j][1]) + 1
			largest = max(largest, width*height)
		}
	}

	return largest
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func parse(input <-chan string) [][2]int {
	rows := make([][2]int, 0)

	for line := range input {
		if line == "" {
		}

		left, right, ok := strings.Cut(line, ",")
		if !ok {
			panic("invalid input")
		}

		var position [2]int
		var err error

		position[0], err = strconv.Atoi(left)
		if err != nil {
			panic(err)
		}

		position[1], err = strconv.Atoi(right)
		if err != nil {
			panic(err)
		}

		rows = append(rows, position)
	}

	return rows
}
