package common

import "testing"

func TestDirection_String(t *testing.T) {
	if DirectionNorth.String() != "N" {
		t.Fail()
	}

	if DirectionEast.String() != "E" {
		t.Fail()
	}

	if DirectionSouth.String() != "S" {
		t.Fail()
	}

	if DirectionWest.String() != "W" {
		t.Fail()
	}

	if DirectionNorthEast.String() != "NE" {
		t.Fail()
	}

	if DirectionSouthEast.String() != "SE" {
		t.Fail()
	}

	if DirectionSouthWest.String() != "SW" {
		t.Fail()
	}

	if DirectionNorthWest.String() != "NW" {
		t.Fail()
	}
}

func TestDirection_Add(t *testing.T) {
	d := DirectionNorth
	d = d.Add(1)

	if d != DirectionEast {
		t.Fail()
	}

	d = d.Add(-1)
	if d != DirectionNorth {
		t.Fail()
	}

	d = d.Add(-1)
	if d != DirectionWest {
		t.Fail()
	}
}

func TestDirection_Negate(t *testing.T) {
	if DirectionNorth.Negate() != DirectionSouth {
		t.Fail()
	}

	if DirectionEast.Negate() != DirectionWest {
		t.Fail()
	}

	if DirectionSouth.Negate() != DirectionNorth {
		t.Fail()
	}

	if DirectionWest.Negate() != DirectionEast {
		t.Fail()
	}

	if DirectionNorthEast.Negate() != DirectionSouthWest {
		t.Fail()
	}

	if DirectionSouthEast.Negate() != DirectionNorthWest {
		t.Fail()
	}

	if DirectionSouthWest.Negate() != DirectionNorthEast {
		t.Fail()
	}

	if DirectionNorthWest.Negate() != DirectionSouthEast {
		t.Fail()
	}
}

func TestGridPos_Move(t *testing.T) {
	var p GridPos
	p = p.Move(DirectionEast, 1)

	if p != (GridPos{0, 1}) {
		t.Fail()
	}

	p = p.Move(DirectionSouth, 1)
	if p != (GridPos{1, 1}) {
		t.Fail()
	}

	p = p.Move(DirectionWest, 1)
	if p != (GridPos{1, 0}) {
		t.Fail()
	}

	p = p.Move(DirectionNorth, 1)
	if p != (GridPos{0, 0}) {
		t.Fail()
	}

	p = p.Move(DirectionSouthEast, 1)
	if p != (GridPos{1, 1}) {
		t.Fail()
	}

	p = p.Move(DirectionSouthWest, 1)
	if p != (GridPos{2, 0}) {
		t.Fail()
	}

	p = p.Move(DirectionNorthWest, 1)
	if p != (GridPos{1, -1}) {
		t.Fail()
	}

	p = p.Move(DirectionNorthEast, 1)
	if p != (GridPos{0, 0}) {
		t.Fail()
	}
}
