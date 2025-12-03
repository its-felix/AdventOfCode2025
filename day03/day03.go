package day03

import (
	"iter"
)

func SolvePart1(input <-chan string) int {
	sum := 0

	for bank := range parse(input) {
		high := [2]uint8{0, 0}

		for i := 0; i < len(bank); i++ {
			if bank[i] > high[0] && (i+1) < len(bank) {
				high[0] = bank[i]
				high[1] = bank[i+1]
			} else if bank[i] > high[1] {
				high[1] = bank[i]
			}
		}

		num := (high[0] * 10) + high[1]
		sum += int(num)
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func parse(input <-chan string) iter.Seq[[]uint8] {
	return func(yield func([]uint8) bool) {
		for line := range input {
			bank := make([]uint8, 0, len(line))
			for _, num := range line {
				if num < '0' || num > '9' {
					panic("invalid input")
				}

				bank = append(bank, uint8(num-'0'))
			}

			if !yield(bank) {
				break
			}
		}
	}
}
