package day_05

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

/*
[D]
[N] [C]
[Z] [M] [P]
 1   2   3
*/

var inputTestStacks = []Stack{
	{'N', 'Z'},
	{'D', 'C', 'M'},
	{'P'},
}

const inputTest = `
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, inputTestStacks)
	testutil.CheckUnexpectedError(t, err)
	var exp string = "CMZ"
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, inputTestStacks)
	testutil.CheckUnexpectedError(t, err)
	var exp string = "MCD"
	if exp != res {
		t.Fatalf("want %s, have %s", exp, res)
	}
}
