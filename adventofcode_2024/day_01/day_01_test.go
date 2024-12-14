package day_01

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 11
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 31
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
