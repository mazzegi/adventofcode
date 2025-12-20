package day_09

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 50
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFuncV2(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 24
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
