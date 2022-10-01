package day_21

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = `
../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#
`

func TestPart1MainFunc(t *testing.T) {
	res, err := onAfter(inputTest, 2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 12
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
