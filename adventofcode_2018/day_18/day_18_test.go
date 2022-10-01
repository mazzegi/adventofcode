package day_18

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2018/testutil"
)

const inputTest = `
.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, 10)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 1147
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
