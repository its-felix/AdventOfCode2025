package day01

import (
	"iter"
	"strconv"
)

func SolvePart1(input <-chan string) int {
	countZero := 0
	position := 50

	for rotation := range parse(input) {
		position += rotation
		position %= 100

		if position == 0 {
			countZero++
		}
	}

	return countZero
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func parse(input <-chan string) iter.Seq[int] {
	return func(yield func(int) bool) {
		for line := range input {
			if line == "" {
				continue
			}

			isLeft := line[0] == 'L'
			times, err := strconv.Atoi(line[1:])
			if err != nil {
				panic(err)
			}

			if isLeft {
				times = -times
			}

			if !yield(times) {
				break
			}
		}
	}
}
