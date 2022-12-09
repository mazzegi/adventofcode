package day_09

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

const inputTest2 = `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 13
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 36
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
