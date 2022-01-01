package day_19

import (
	"adventofcode_2016/testutil"
	"testing"
)

func TestPart1MainFunc(t *testing.T) {
	res, err := winningElf(5)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := winningElfSteelAcross(5)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 2
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
