package common

const (
	DirectionNorth = Direction(0)
	DirectionEast  = Direction(1)
	DirectionSouth = Direction(2)
	DirectionWest  = Direction(3)
)

type Direction uint8

func (d Direction) String() string {
	switch d {
	case DirectionNorth:
		return "N"

	case DirectionEast:
		return "E"

	case DirectionSouth:
		return "S"

	case DirectionWest:
		return "W"
	}

	panic("invalid direction")
}

func (d Direction) Add(v int) Direction {
	return (d + Direction(v)) % 4
}

func (d Direction) Next() Direction {
	return d.Add(1)
}

func (d Direction) Negate() Direction {
	return d.Add(2)
}

type IntercardinalDirection [2]Direction

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
	switch d {
	case DirectionNorth:
		newPos[0] -= steps

	case DirectionEast:
		newPos[1] += steps

	case DirectionSouth:
		newPos[0] += steps

	case DirectionWest:
		newPos[1] -= steps
	}

	return newPos
}
