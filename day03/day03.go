package day03

import (
	"iter"
)

func SolvePart1(input <-chan string) uint64 {
	return solve(input, 2)
}

func SolvePart2(input <-chan string) uint64 {
	return solve(input, 12)
}

func solve(input <-chan string, batteries int) uint64 {
	sum := uint64(0)

	for bank := range parse(input) {
		num := uint64(0)

		for i := 0; i < batteries; i++ {
			num *= 10 // make room for another digit 1..9 (on first iteration, num is 0 so it doesnt make a difference)

			remainingPosToFill := batteries - i
			bestIdxForPos := 0

			for j, v := range bank {
				if (len(bank) - j) < remainingPosToFill {
					break
				}

				if v > bank[bestIdxForPos] {
					bestIdxForPos = j
				}
			}

			num += uint64(bank[bestIdxForPos])
			bank = bank[bestIdxForPos+1:]
		}

		sum += num
	}

	return sum
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
