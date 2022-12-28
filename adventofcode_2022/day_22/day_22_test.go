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

func connectTest(cb *CubeBoard) {
	cb.ConnectRange(P(8, 0), P(11, 0), Up, P(3, 4), P(0, 4), Down)
	cb.ConnectRange(P(8, 0), P(8, 3), Left, P(4, 4), P(7, 4), Down)
	cb.ConnectRange(P(11, 0), P(11, 3), Right, P(15, 11), P(15, 8), Left)

	// cb.Connect(P(8, 0), Up, P(3, 4), Down)
	// cb.Connect(P(9, 0), Up, P(2, 4), Down)
	// cb.Connect(P(10, 0), Up, P(1, 4), Down)
	// cb.Connect(P(11, 0), Up, P(0, 4), Down)

	// cb.Connect(P(8, 0), Left, P(4, 4), Down)
	// cb.Connect(P(8, 1), Left, P(5, 4), Down)
	// cb.Connect(P(8, 2), Left, P(6, 4), Down)
	// cb.Connect(P(8, 3), Left, P(7, 4), Down)

	// cb.Connect(P(11, 0), Right, P(15, 11), Left)
	// cb.Connect(P(11, 1), Right, P(15, 10), Left)
	// cb.Connect(P(11, 2), Right, P(15, 9), Left)
	// cb.Connect(P(11, 3), Right, P(15, 8), Left)
	//
}

const inputTestPath = "10R5L5R10L4R5L5"

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, inputTestPath)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 6032
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, inputTestPath, connectTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 5031
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
