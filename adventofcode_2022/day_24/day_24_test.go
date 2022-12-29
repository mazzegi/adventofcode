package day_24

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, 100)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 18
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, 100)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 54
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
