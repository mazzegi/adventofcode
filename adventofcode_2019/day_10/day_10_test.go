package day_10

import (
	"testing"

	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest1 = `
.#..#
.....
#####
....#
...##
`

const inputTest2 = `
......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
`

const inputTest3 = `
.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`

func TestPart1MainFunc(t *testing.T) {
	t.Run("test_1", func(t *testing.T) {
		res, err := part1MainFunc(inputTest1)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 8
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

	t.Run("test_2", func(t *testing.T) {
		res, err := part1MainFunc(inputTest2)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 33
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

	t.Run("test_3", func(t *testing.T) {
		res, err := part1MainFunc(inputTest3)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 210
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})
}

const part2InputTest1 = `
.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##
`

func TestPart2MainFunc(t *testing.T) {
	t.Run("test_2_1", func(t *testing.T) {
		t.Skip()
		res, err := part2MainFunc(part2InputTest1, grid.Pt(8, 3))
		testutil.CheckUnexpectedError(t, err)
		var exp int = -42
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

	t.Run("test_2_1", func(t *testing.T) {

		res, err := part2MainFunc(inputTest3, grid.Pt(11, 13))
		testutil.CheckUnexpectedError(t, err)
		var exp int = 802
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

}
