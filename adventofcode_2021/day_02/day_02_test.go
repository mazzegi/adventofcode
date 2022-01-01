package day_02

import (
	"adventofcode_2021/testutil"
	"testing"
)

const testInput = `
forward 5
down 5
forward 8
up 3
down 8
forward 2
`

func TestProcessInput(t *testing.T) {
	pos, err := ProcessInput(testInput, Position2D{}, AppliedCommand)
	testutil.CheckUnexpectedError(t, err)

	exp := Position2D{
		X:     15,
		Depth: 10,
	}
	if pos != exp {
		t.Fatalf("want %v, have %v", exp, pos)
	}
}

func TestProcessInputAim(t *testing.T) {
	pos, err := ProcessInput(testInput, Position2D{}, AppliedCommandAim)
	testutil.CheckUnexpectedError(t, err)

	exp := Position2D{
		X:     15,
		Depth: 60,
		Aim:   10,
	}
	if pos != exp {
		t.Fatalf("want %v, have %v", exp, pos)
	}
}
