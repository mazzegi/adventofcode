package day_25

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

const inputTest = `
v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>
`

func TestPart1MainFunc(t *testing.T) {
	res, err := stepsToNoMoves(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 58
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
