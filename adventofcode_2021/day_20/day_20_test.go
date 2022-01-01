package day_20

import (
	"adventofcode_2021/testutil"
	"testing"
)

const inputAlgoTest = `
..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##
#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###
.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#.
.#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#.....
.#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#..
...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.....
..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#
`

const inputImgTest = `
#..#.
#....
##..#
..#..
..###
`

func TestPart1MainFunc(t *testing.T) {
	res, err := numLitPixels(inputAlgoTest, inputImgTest, 2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 35
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := numLitPixels(inputAlgoTest, inputImgTest, 50)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3351
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestBitsToNumber(t *testing.T) {
	//000100010
	in := []bool{false, false, false, true, false, false, false, true, false}
	num := bitsToNumber(in)
	if num != 34 {
		t.Fatalf("want %d, have %d", 34, num)
	}
}
