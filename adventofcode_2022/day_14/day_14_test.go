package day_14

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 24
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 93
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
