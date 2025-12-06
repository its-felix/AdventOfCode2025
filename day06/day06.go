package day06

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

const (
	operationMul = rune('*')
	operationAdd = rune('+')
)

type problem struct {
	numbers   []int
	operation rune
}

func SolvePart1(input <-chan string) int {
	sum := 0
	for _, p := range parse(input) {
		opResult := p.numbers[0]

		switch p.operation {
		case operationMul:
			for _, n := range p.numbers[1:] {
				opResult *= n
			}

		case operationAdd:
			for _, n := range p.numbers[1:] {
				opResult += n
			}
		}

		sum += opResult
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	grid := make([][]rune, 0)
	maxColumns := 0

	for line := range input {
		if line == "" {
			continue
		}

		columns := []rune(line)
		grid = append(grid, columns)
		maxColumns = max(maxColumns, len(columns))
	}

	const operationUnknown = rune('#')

	sum := 0
	operation := operationUnknown
	currentNum := 0
	result := 0

	for col := 0; col < maxColumns; col++ {
		isInitial := false
		isColumnDelimiter := true

		for row := 0; row < len(grid); row++ {
			if col >= len(grid[row]) || grid[row][col] == ' ' {
				continue
			}

			isColumnDelimiter = false

			if grid[row][col] >= '0' && grid[row][col] <= '9' {
				currentNum *= 10
				currentNum += int(grid[row][col] - '0')
			} else if grid[row][col] == operationMul || grid[row][col] == operationAdd {
				operation = grid[row][col]
				result = currentNum
				isInitial = true
			} else {
				panic("invalid input")
			}
		}

		if isColumnDelimiter {
			sum += result
			operation = operationUnknown
			result = 0
		} else if !isInitial {
			switch operation {
			case operationMul:
				result *= currentNum

			case operationAdd:
				result += currentNum
			}
		}

		currentNum = 0
	}

	return sum + result
}

func parse(input <-chan string) []problem {
	minLength, maxLength := math.MaxInt, 0
	numbers := make([][]int, 0)
	problems := make([]problem, 0)

	for line := range input {
		if line == "" {
			continue
		}

		numberOrSymbols := strings.FieldsFunc(line, unicode.IsSpace)
		minLength, maxLength = min(minLength, len(numberOrSymbols)), max(maxLength, len(numberOrSymbols))

		if minLength != maxLength {
			panic("invalid input")
		}

		switch rune(numberOrSymbols[0][0]) {
		case operationMul, operationAdd:
			for i, symbol := range numberOrSymbols {
				problems = append(problems, problem{
					numbers:   numbers[i],
					operation: rune(symbol[0]),
				})
			}

		default:
			for i, number := range numberOrSymbols {
				for i >= len(numbers) {
					numbers = append(numbers, make([]int, 0, len(numberOrSymbols)))
				}

				num, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}

				numbers[i] = append(numbers[i], num)
			}
		}
	}

	return problems
}
