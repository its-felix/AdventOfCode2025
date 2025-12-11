package day10

import (
	"context"
	"fmt"
	"iter"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type Machine[B any] struct {
	TargetState uint64
	Buttons     []B
	Joltages    []int
}

func SolvePart1(input <-chan string) int {
	machines := parse(input, parseButtonPart1)
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

func SolvePart2(input <-chan string) uint64 {
	machines := parse(input, parseButtonPart2)

	var sum atomic.Uint64
	g, ctx := errgroup.WithContext(context.Background())

	for i, m := range machines {
		g.Go(func() error {
			presses, solvable, err := solvePart2(ctx, make([]int, len(m.Joltages)), m.Buttons, m.Joltages, math.MaxInt)
			if err != nil {
				return err
			}

			if !solvable {
				return fmt.Errorf("machine %d: not solvable", i+1)
			}

			fmt.Printf("machine %d: %d presses\n", i+1, presses)

			sum.Add(uint64(presses))

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}

	return sum.Load()
}

func solvePart2(ctx context.Context, state []int, buttons [][]int, targetState []int, maxPresses int) (int, bool, error) {
	select {
	case <-ctx.Done():
		return 0, false, ctx.Err()

	default:
	}

	solved, possible := isSolvedPart2(state, targetState)
	if solved {
		return 0, true, nil
	} else if !possible {
		return 0, false, nil
	} else if maxPresses < 1 {
		return 0, false, nil
	}

	for i := 0; i < len(buttons); i++ {
		for presses, state := range applyButtonManyTimes(state, buttons[i], targetState) {
			if presses > maxPresses {
				break
			}

			additionalPresses, wasSolved, err := solvePart2(ctx, state, buttons[i+1:], targetState, maxPresses-presses)
			if err != nil {
				return 0, false, err
			}

			if wasSolved {
				solved = true
				presses += additionalPresses

				if presses < maxPresses {
					maxPresses = presses
				}
			}
		}
	}

	return maxPresses, solved, nil
}

func applyButtonManyTimes(state, button, targetState []int) iter.Seq2[int, []int] {
	return func(yield func(int, []int) bool) {
		presses := 0
		for {
			state = applyButtonPart2(state, button)
			presses++
			if _, possible := isSolvedPart2(state, targetState); !possible {
				break
			}

			if !yield(presses, state) {
				break
			}
		}
	}
}

func applyButtonPart2(state, button []int) []int {
	result := slices.Clone(state)
	for _, idx := range button {
		result[idx]++
	}

	return result
}

func isSolvedPart2(state, targetState []int) (bool, bool) {
	solved := true

	for i := 0; i < len(state); i++ {
		if state[i] > targetState[i] {
			return false, false
		} else if state[i] < targetState[i] {
			solved = false
		}
	}

	return solved, true
}

func parse[B any](input <-chan string, parseButton func(string) (B, error)) []Machine[B] {
	machines := make([]Machine[B], 0)
	for line := range input {
		if line == "" {
			continue
		}

		m, err := parseMachine(line, parseButton)
		if err != nil {
			panic(err)
		}

		machines = append(machines, m)
	}

	return machines
}

func parseMachine[B any](line string, parseButton func(string) (B, error)) (Machine[B], error) {
	var m Machine[B]

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

	line, m.Buttons, err = parseButtons(line[end+1:], parseButton)
	if err != nil {
		return m, err
	}

	m.Joltages, err = parseJoltages(line)
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

func parseButtons[B any](line string, parseButton func(string) (B, error)) (string, []B, error) {
	buttons := make([]B, 0)

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

func parseJoltages(joltages string) ([]int, error) {
	start := strings.IndexRune(joltages, '{')
	end := strings.IndexRune(joltages, '}')
	if start == -1 || end == -1 {
		return nil, fmt.Errorf("invalid line, missing joltages: %q", joltages)
	}

	result := make([]int, 0)

	for _, v := range strings.Split(joltages[start+1:end], ",") {
		idx, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		result = append(result, idx)
	}

	return result, nil
}

func parseButtonPart1(button string) (uint64, error) {
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

func parseButtonPart2(button string) ([]int, error) {
	result := make([]int, 0)

	for _, v := range strings.Split(button, ",") {
		idx, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		result = append(result, idx)
	}

	return result, nil
}
