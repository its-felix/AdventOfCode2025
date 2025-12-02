package day02

import (
	"iter"
	"math"
	"strconv"
	"strings"
)

func SolvePart1(input <-chan string) int {
	sum := 0

	for start, end := range parse(input) {
		for invalidId := range invalidIds(start) {
			if invalidId > end {
				break
			}

			sum += invalidId
		}
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func invalidIds(start int) iter.Seq[int] {
	return func(yield func(int) bool) {
		curr := start - 1
		for {
			curr = nextInvalidId(curr)
			if !yield(curr) {
				break
			}
		}
	}
}

func nextInvalidId(start int) int {
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
		return nextInvalidId(start + 1)
	} else {
		num = firstHalf + 1
	}

	return num*int(math.Pow10(len(str)/2)) + num
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
