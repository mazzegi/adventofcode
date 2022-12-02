package day_02

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
A Y
B X
C Z
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 15
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 12
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
