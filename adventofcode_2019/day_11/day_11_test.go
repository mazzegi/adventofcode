package day_11

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

var inputTest = []int{}

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, 0)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
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
