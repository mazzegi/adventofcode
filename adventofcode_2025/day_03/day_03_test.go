package day_03

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
987654321111111
811111111111119
234234234234278
818181911112111
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 357
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3121910778619
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
