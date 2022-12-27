package day_22

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.
`

const inputTestPath = "10R5L5R10L4R5L5"

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, inputTestPath)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, inputTestPath)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
