package day_08

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
30373
25512
65332
33549
35390
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 21
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 8
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
