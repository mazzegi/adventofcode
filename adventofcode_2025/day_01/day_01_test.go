package day_01

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 6
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
