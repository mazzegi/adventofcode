package day_01

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 24000
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 45000
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
