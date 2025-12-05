package day05

import (
	"strconv"
	"strings"
)

type Range [2]int

func (r Range) Contains(n int) bool {
	return n >= r[0] && n <= r[1]
}

type Ranges []Range

func (rs Ranges) Contains(n int) bool {
	for _, r := range rs {
		if r.Contains(n) {
			return true
		}
	}

	return false
}

func SolvePart1(input <-chan string) int {
	ranges, nums := parse(input)
	count := 0

	for _, n := range nums {
		if ranges.Contains(n) {
			count++
		}
	}

	return count
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func parse(input <-chan string) (Ranges, []int) {
	ranges := make(Ranges, 0)
	nums := make([]int, 0)

	for line := range input {
		// split between ranges and numbers
		if line == "" {
			break
		}

		startRaw, endRaw, ok := strings.Cut(line, "-")
		if !ok {
			panic("invalid range")
		}

		start, err := strconv.Atoi(startRaw)
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(endRaw)
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, Range{start, end})
	}

	for line := range input {
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	return ranges, nums
}
