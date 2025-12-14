package day_06

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 4277556
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3263827
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
