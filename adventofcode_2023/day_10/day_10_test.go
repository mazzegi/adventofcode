package day_10

import (
	"fmt"
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest1 = `
.....
.S-7.
.|.|.
.L-J.
.....
`

const inputTest2 = `
..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`

func TestPart1MainFunc(t *testing.T) {
	{
		res, err := part1MainFunc(inputTest1)
		testutil.CheckUnexpectedError(t, err)
		testutil.Assert(t, 4, res)
	}
	{
		res, err := part1MainFunc(inputTest2)
		testutil.CheckUnexpectedError(t, err)
		testutil.Assert(t, 8, res)
	}
}

const imputPart2Test1 = `
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`

const imputPart2Test2 = `
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`

const imputPart2Test3 = `
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`

func TestPart2MainFunc(t *testing.T) {
	{
		res, err := part2MainFunc(imputPart2Test1)
		testutil.CheckUnexpectedError(t, err)
		testutil.Assert(t, 4, res)
	}
	{
		res, err := part2MainFunc(imputPart2Test2)
		testutil.CheckUnexpectedError(t, err)
		testutil.Assert(t, 8, res)
	}
	{
		res, err := part2MainFunc(imputPart2Test3)
		testutil.CheckUnexpectedError(t, err)
		testutil.Assert(t, 10, res)
	}
}

func TestInsideCount(t *testing.T) {
	tests := []struct {
		in    string
		count int
	}{
		{".F-------7.", 0},
		{".||.....||.", 0},
		{".|..|.|..|.", 4},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_#%02d", i), func(t *testing.T) {
			res := insideCount(test.in)
			testutil.Assert(t, test.count, res)
		})
	}
}
