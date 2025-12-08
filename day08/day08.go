package day08

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

type _3DPosition [3]int

func (p _3DPosition) Distance(other _3DPosition) float64 {
	distanceA := float64(max(p[0], other[0]) - min(p[0], other[0]))
	distanceB := float64(max(p[1], other[1]) - min(p[1], other[1]))
	distanceC := float64(max(p[2], other[2]) - min(p[2], other[2]))

	return math.Sqrt(distanceA*distanceA + distanceB*distanceB + distanceC*distanceC)
}

type Circuit []_3DPosition

func SolvePart1(input <-chan string, connect int) int {
	circuits := parse(input)
	connected := make(map[[2]_3DPosition]struct{})

	for len(connected) < connect {
		var ok bool
		if circuits, ok = connectClosest(connected, circuits); !ok {
			panic("no more connections possible")
		}
	}

	slices.SortFunc(circuits, func(a, b Circuit) int {
		return len(b) - len(a)
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func connectClosest(connected map[[2]_3DPosition]struct{}, circuits []Circuit) ([]Circuit, bool) {
	minDistance := math.MaxFloat64
	leftCircuitIdx, leftPositionIdx := -1, -1
	rightCircuitIdx, rightPositionIdx := -1, -1

	for circuitI := 0; circuitI < len(circuits); circuitI++ {
		for positionI := 0; positionI < len(circuits[circuitI]); positionI++ {
			for circuitJ := circuitI; circuitJ < len(circuits); circuitJ++ {
				for positionJ := 0; positionJ < len(circuits[circuitJ]); positionJ++ {
					if circuitI == circuitJ && positionI == positionJ {
						// same entry, skip
						continue
					} else if _, ok := connected[key(circuits[circuitI][positionI], circuits[circuitJ][positionJ])]; ok {
						// already connected, skip
						continue
					}

					distance := circuits[circuitI][positionI].Distance(circuits[circuitJ][positionJ])
					if distance < minDistance {
						minDistance = distance
						leftCircuitIdx, leftPositionIdx = circuitI, positionI
						rightCircuitIdx, rightPositionIdx = circuitJ, positionJ
					}
				}
			}
		}
	}

	if leftCircuitIdx == -1 || leftPositionIdx == -1 || rightCircuitIdx == -1 || rightPositionIdx == -1 {
		return circuits, false
	}

	return connect(connected, circuits, leftCircuitIdx, leftPositionIdx, rightCircuitIdx, rightPositionIdx), true
}

func connect(connected map[[2]_3DPosition]struct{}, circuits []Circuit, leftCircuitIdx, leftPositionIdx, rightCircuitIdx, rightPositionIdx int) []Circuit {
	connected[key(circuits[leftCircuitIdx][leftPositionIdx], circuits[rightCircuitIdx][rightPositionIdx])] = struct{}{}

	if leftCircuitIdx == rightCircuitIdx {
		return circuits
	}

	newCircuits := make([]Circuit, 0, len(circuits))
	for i, circuit := range circuits {
		if i == leftCircuitIdx {
			// combine left and right
			circuit = append(circuit, circuits[rightCircuitIdx]...)
			newCircuits = append(newCircuits, circuit)
		} else if i != rightCircuitIdx {
			// skip right, already added by left
			newCircuits = append(newCircuits, circuit)
		}
	}

	return newCircuits
}

func key(a, b _3DPosition) [2]_3DPosition {
	if cmp.Or(a[0]-b[0], a[1]-b[1], a[2]-b[2]) >= 0 {
		return [2]_3DPosition{a, b}
	} else {
		return [2]_3DPosition{b, a}
	}
}

func parse(input <-chan string) []Circuit {
	result := make([]Circuit, 0)

	for line := range input {
		if line == "" {
			continue
		}

		var position _3DPosition
		parts := strings.SplitN(line, ",", 3)
		if len(parts) != 3 {
			panic("invalid input")
		}

		for i := 0; i < 3; i++ {
			var err error
			position[i], err = strconv.Atoi(parts[i])
			if err != nil {
				panic(err)
			}
		}

		result = append(result, Circuit{position})
	}

	return result
}
