package day_11

import (
	"testing"

	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

const inputTestExpanded = `
....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 374
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	{
		res, err := part2MainFunc(inputTest, 10)
		testutil.CheckUnexpectedError(t, err)
		testutil.Assert(t, 1030, res)
	}
	{
		res, err := part2MainFunc(inputTest, 100)
		testutil.CheckUnexpectedError(t, err)
		testutil.Assert(t, 8410, res)
	}

}

func TestExpand(t *testing.T) {
	erows := expanded(readutil.ReadLines(inputTest))
	testutil.Assert(t, readutil.ReadLines(inputTestExpanded), erows)
}
