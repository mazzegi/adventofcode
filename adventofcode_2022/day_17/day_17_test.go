package day_17

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

//                 >>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, 2022)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3068
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, 1000000000000)

	testutil.CheckUnexpectedError(t, err)
	var exp int = 1514285714288
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
