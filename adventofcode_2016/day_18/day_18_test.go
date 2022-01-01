package day_18

import (
	"adventofcode_2016/testutil"
	"testing"
)

const inputTest = ".^^.^.^^^^"

func TestPart1MainFunc(t *testing.T) {
	res, err := safeTiles(inputTest, 10)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 38
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
