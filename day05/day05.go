package day05

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Range [2]int

func (r Range) Contains(n int) bool {
	return n >= r[0] && n <= r[1]
}

func (r Range) Length() int {
	return r[1] - r[0] + 1
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

func (rs Ranges) Sort() {
	slices.SortFunc(rs, func(a, b Range) int {
		return cmp.Or(
			a[0]-b[0],
			a[1]-b[1],
		)
	})
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
	ranges, _ := parse(input)
	if len(ranges) < 1 {
		return 0
	}

	ranges.Sort()

	count := ranges[0].Length()
	end := ranges[0][1]

	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		if r[1] <= end {
			continue
		}

		if r[0] > end {
			// starts after previous range
			count += r.Length()
			end = r[1]
		} else {
			// starts within previous range
			// but at or after start of previous range (input is sorted)
			count += r[1] - end
			end = r[1]
		}
	}

	return count
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
