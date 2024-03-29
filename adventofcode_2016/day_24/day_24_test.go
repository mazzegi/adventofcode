package day_24

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2016/testutil"
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
	res, err := fewestStepsWithReturn(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 20
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
