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
	countZero := 0
	position := 50

	for rotation := range parse(input) {
		// effective change [-99..+99]
		if change := rotation % 100; change != 0 {
			position += change

			if position == 0 {
				// landing on 0
				countZero++
			} else if position < 0 {
				// landing below 0 (roll over to end)
				if position != change {
					// count only if previous position was not already zero
					countZero++
				}

				position += 100
			} else if position >= 100 {
				// landing over 99 (roll over to start)
				countZero++
				position -= 100
			}
		}

		// add full rotations
		if rotation < 0 {
			countZero += (-rotation) / 100
		} else {
			countZero += rotation / 100
		}
	}

	return countZero
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
