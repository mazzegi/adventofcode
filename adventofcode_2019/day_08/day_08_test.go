package day_08

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = "123456789012"

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, 3, 2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 1
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, 3, 2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
