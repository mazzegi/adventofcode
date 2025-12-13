package day_05

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
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
	var exp int = 14
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
