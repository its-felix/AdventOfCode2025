package day10

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	TargetState uint64
	Buttons     []uint64
	Joltages    []int
}

func SolvePart1(input <-chan string) int {
	machines := parse(input)
	sum := 0

	for _, m := range machines {
		presses, success := solvePart1(0, m.Buttons, m.TargetState, math.MaxInt)
		if !success {
			panic("not solvable")
		}

		sum += presses
	}

	return sum
}

func solvePart1(state uint64, buttons []uint64, targetState uint64, maxPresses int) (int, bool) {
	if state == targetState {
		return 0, true
	} else if maxPresses < 1 {
		return 0, false
	}

	success := false

	for i := 0; i < len(buttons); i++ {
		presses, solved := solvePart1(state^buttons[i], buttons[i+1:], targetState, maxPresses-1)
		if solved {
			success = true
			presses++

			if presses < maxPresses {
				maxPresses = presses
			}
		}
	}

	return maxPresses, success
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func parse(input <-chan string) []Machine {
	machines := make([]Machine, 0)
	for line := range input {
		if line == "" {
			continue
		}

		m, err := parseMachine(line)
		if err != nil {
			panic(err)
		}

		machines = append(machines, m)
	}

	return machines
}

func parseMachine(line string) (Machine, error) {
	var m Machine

	if line[0] != '[' {
		return m, fmt.Errorf("invalid line: %q", line)
	}

	start := strings.IndexRune(line, '[')
	end := strings.IndexRune(line, ']')
	if start == -1 || end == -1 {
		return m, fmt.Errorf("invalid line: %q", line)
	}

	var err error
	m.TargetState, err = parseIndicators(line[start+1 : end])
	if err != nil {
		return m, err
	}

	line, m.Buttons, err = parseButtons(line[end+1:])
	if err != nil {
		return m, err
	}

	return m, nil
}

func parseIndicators(indicators string) (uint64, error) {
	result := uint64(0)

	runes := []rune(indicators)
	for i := len(runes) - 1; i >= 0; i-- {
		result <<= 1

		switch runes[i] {
		case '#':
			result |= 1

		case '.':
			// no-op

		default:
			return result, fmt.Errorf("invalid indicator: %q", indicators)
		}
	}

	return result, nil
}

func parseButtons(line string) (string, []uint64, error) {
	buttons := make([]uint64, 0)

	start := strings.IndexRune(line, '(')
	for start != -1 {
		line = line[start+1:]
		end := strings.IndexRune(line, ')')
		if end == -1 {
			return line, buttons, fmt.Errorf("invalid line, missing closing parenthesis: %q", line)
		}

		button, err := parseButton(line[:end])
		if err != nil {
			return line, buttons, err
		}

		buttons = append(buttons, button)

		line = line[end+1:]
		start = strings.IndexRune(line, '(')
	}

	return line, buttons, nil
}

func parseButton(button string) (uint64, error) {
	result := uint64(0)

	for _, v := range strings.Split(button, ",") {
		idx, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}

		if idx >= 64 {
			return 0, fmt.Errorf("invalid button: %q", button)
		}

		result |= 1 << uint64(idx)
	}

	return result, nil
}
