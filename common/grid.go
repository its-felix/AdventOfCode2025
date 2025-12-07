package common

import "math/bits"

const (
	DirectionNorth = Direction(0b00000010)
	DirectionEast  = Direction(0b00001000)
	DirectionSouth = Direction(0b00100000)
	DirectionWest  = Direction(0b10000000)

	DirectionNorthEast = DirectionNorth | DirectionEast
	DirectionSouthEast = DirectionSouth | DirectionEast
	DirectionNorthWest = DirectionNorth | DirectionWest
	DirectionSouthWest = DirectionSouth | DirectionWest
)

type Direction uint8

func (d Direction) String() string {
	s := ""
	if d&DirectionNorth != 0 {
		s += "N"
	}

	if d&DirectionSouth != 0 {
		s += "S"
	}

	if d&DirectionEast != 0 {
		s += "E"
	}

	if d&DirectionWest != 0 {
		s += "W"
	}

	return s
}

func (d Direction) Add(v int) Direction {
	return Direction(bits.RotateLeft8(uint8(d), v*2))
}

func (d Direction) Next() Direction {
	return d.Add(1)
}

func (d Direction) Negate() Direction {
	return d.Add(2)
}

type Grid[T any] [][]T

func (g Grid[T]) Contains(p GridPos) bool {
	if p.Row() < 0 || p.Row() >= len(g) {
		return false
	}

	if p.Col() < 0 || p.Col() >= len(g[p.Row()]) {
		return false
	}

	return true
}

func (g Grid[T]) Move(p GridPos, d Direction, steps int) (GridPos, bool) {
	p = p.Move(d, steps)
	return p, g.Contains(p)
}

func (g Grid[T]) MoveRollover(p GridPos, d Direction, steps int) GridPos {
	p = p.Move(d, steps)

	if p.Row() < 0 {
		p[0] = len(g) + (p.Row() % len(g))
	} else if p.Row() >= len(g) {
		p[0] = p.Row() % len(g)
	}

	if p.Col() < 0 {
		p[1] = len(g[p.Row()]) + (p.Col() % len(g[p.Row()]))
	} else if p.Col() >= len(g[p.Row()]) {
		p[1] = p.Col() % len(g[p.Row()])
	}

	return p
}

type GridPos [2]int

func (p GridPos) Row() int {
	return p[0]
}

func (p GridPos) Col() int {
	return p[1]
}

func (p GridPos) Move(d Direction, steps int) GridPos {
	newPos := p
	if d&DirectionNorth != 0 {
		newPos[0] -= steps
	}

	if d&DirectionSouth != 0 {
		newPos[0] += steps
	}

	if d&DirectionEast != 0 {
		newPos[1] += steps
	}

	if d&DirectionWest != 0 {
		newPos[1] -= steps
	}

	return newPos
}
