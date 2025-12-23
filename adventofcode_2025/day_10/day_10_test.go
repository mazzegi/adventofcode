package day_10

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`

const inputTest2 = `
[#.####.##.] (2,9) (0,1,3,4,5,6,8,9) (7,9) (0,4,9) (4,5,6) (1,2,4,7,8,9) (2,7,9) (1,2,3,4,5) (0,1,3,4,5,7,9) (0,7) (1,3,4,5,7) (1,3) {27,64,209,46,62,34,10,220,27,233}
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 7
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart1MainFuncInput2(t *testing.T) {
	res, err := part1MainFunc(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 7
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 33
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
