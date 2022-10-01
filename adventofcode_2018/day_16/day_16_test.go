package day_16

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2018/testutil"
)

const inputTest = `
Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 1
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, "foo")
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
