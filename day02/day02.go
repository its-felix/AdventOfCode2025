package day02

import (
	"iter"
	"math"
	"strconv"
	"strings"
)

func SolvePart1(input <-chan string) int {
	return solve(input, nextInvalidIdPart1)
}

func SolvePart2(input <-chan string) int {
	return solve(input, nextInvalidIdPart2)
}

func solve(input <-chan string, nextFn func(int) int) int {
	sum := 0

	for start, end := range parse(input) {
		for invalidId := range iterGenerator(start-1, nextFn) {
			if invalidId > end {
				break
			}

			sum += invalidId
		}
	}

	return sum
}

func iterGenerator(start int, nextFn func(int) int) iter.Seq[int] {
	return func(yield func(int) bool) {
		curr := start
		for {
			curr = nextFn(curr)
			if !yield(curr) {
				break
			}
		}
	}
}

func nextInvalidIdPart1(start int) int {
	if start < 11 {
		return 11
	}

	str := strconv.Itoa(start)
	if len(str)%2 != 0 {
		// any 3-digit number -> 1010, any 5-digit number -> 100100
		num := int(math.Pow10(((len(str) + 1) / 2) - 1))
		return (num * num * 10) + num
	}

	firstHalf, _ := strconv.Atoi(str[:len(str)/2])
	secondHalf, _ := strconv.Atoi(str[len(str)/2:])

	var num int
	if firstHalf > secondHalf {
		num = firstHalf
	} else if (firstHalf+1)%10 == 0 {
		return nextInvalidIdPart1(start + 1)
	} else {
		num = firstHalf + 1
	}

	return num*int(math.Pow10(len(str)/2)) + num
}

func nextInvalidIdPart2(start int) int {
	if start < 11 {
		return 11
	}

	for i := start + 1; ; i++ {
		if isInvalidPart2(i) {
			return i
		}
	}
}

func isInvalidPart2(num int) bool {
	str := strconv.Itoa(num)
	for i := 1; i < len(str); i++ {
		if isInvalidPart2Step(str, i) {
			return true
		}
	}

	return false
}

func isInvalidPart2Step(v string, step int) bool {
	if step >= len(v) || len(v)%step != 0 {
		return false
	}

	part := v[:step]
	for i := step; i+step <= len(v); i += step {
		if v[i:i+step] != part {
			return false
		}
	}

	return true
}

func parse(input <-chan string) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for line := range input {
			for rawRange := range strings.SplitSeq(line, ",") {
				if rawRange == "" {
					continue
				}

				rawStart, rawEnd, ok := strings.Cut(rawRange, "-")
				if !ok {
					panic("invalid range")
				}

				start, err := strconv.Atoi(rawStart)
				if err != nil {
					panic(err)
				}

				end, err := strconv.Atoi(rawEnd)
				if err != nil {
					panic(err)
				}

				if !yield(start, end) {
					return
				}
			}
		}
	}
}
