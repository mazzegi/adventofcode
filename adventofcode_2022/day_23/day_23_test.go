package day_23

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest1 = `
.....
..##.
..#..
.....
..##.
.....
`

const inputTest2 = `
.......#......
.....###.#....
...#...#.#....
....#...##....
...#.###......
...##.#.##....
....#..#......
`

func TestPart1MainFunc(t *testing.T) {
	// res, err := part1MainFunc(inputTest1)
	// testutil.CheckUnexpectedError(t, err)
	// var exp int = 25
	// if exp != res {
	// 	t.Fatalf("want %d, have %d", exp, res)
	// }

	res, err := part1MainFunc(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	exp := 110
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
