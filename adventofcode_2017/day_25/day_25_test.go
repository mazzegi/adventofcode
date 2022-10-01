package day_25

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

/*
Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.
*/

var inputBlueprintTest = blueprint{
	startStateID: "A",
	steps:        6,
	states: map[string]state{
		"A": {
			id: "A",
			if0: instruction{
				write:         1,
				move:          moveRight,
				continueState: "B",
			},
			if1: instruction{
				write:         0,
				move:          moveLeft,
				continueState: "B",
			},
		},
		"B": {
			id: "B",
			if0: instruction{
				write:         1,
				move:          moveLeft,
				continueState: "A",
			},
			if1: instruction{
				write:         1,
				move:          moveRight,
				continueState: "A",
			},
		},
	},
}

func TestPart1MainFunc(t *testing.T) {
	res, err := checksum(inputBlueprintTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

// func TestPart2MainFunc(t *testing.T) {
// 	res, err := part2MainFunc(inputTest)
// 	testutil.CheckUnexpectedError(t, err)
// 	var exp int = -42
// 	if exp != res {
// 		t.Fatalf("want %d, have %d", exp, res)
// 	}
// }
