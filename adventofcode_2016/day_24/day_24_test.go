package day_24

import (
	"adventofcode_2016/testutil"
	"testing"
)

const inputTest = `
###########
#0.1.....2#
#.#######.#
#4.......3#
###########
`

func TestPart1MainFunc(t *testing.T) {
	res, err := fewestSteps(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 14
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
