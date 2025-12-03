package day03

import (
	"iter"
)

func SolvePart1(input <-chan string) uint64 {
	return solve[[2]uint8](input)
}

func SolvePart2(input <-chan string) uint64 {
	return solve[[12]uint8](input)
}

func solve[T interface{ ~[2]uint8 | ~[12]uint8 }](input <-chan string) uint64 {
	sum := uint64(0)

	for bank := range parse(input) {
		var high T

		for i := 0; i < len(high); i++ {
			remainingPosToFill := len(high) - i
			bestIdxForPos := 0

			for j, v := range bank {
				if (len(bank) - j) < remainingPosToFill {
					break
				}

				if v > bank[bestIdxForPos] {
					bestIdxForPos = j
				}
			}

			high[i] = bank[bestIdxForPos]
			bank = bank[bestIdxForPos+1:]
		}

		num := uint64(0)
		for i := range len(high) {
			num *= 10
			num += uint64(high[i])
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
