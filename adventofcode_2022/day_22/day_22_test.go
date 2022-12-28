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

const strConnectionsTets = `
	# Side 1
	up:    8,0  ... 11,0  => down: 3,4 ...0,4
	left:  8,0  ... 8,3   => down: 4,4 ...7,4
	right: 11,0 ... 11,3  => left: 15,11 ...15,8

	# Side 4
	right: 11,4 ... 11,7  => down: 15,8 ... 12,8

	# Side 2
	left:  0,4  ... 0,7   => up  : 15,11 ... 12,11
	down:  0,7  ... 3,7   => up  : 11,11 ... 8,11

	# Side 3
	down:  4,7  ... 7,7   => right: 8,11 ...11,11
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, inputTestPath)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 6032
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func connectTest(cube *Cube) {
	cube.ConnectRange(Up, P(8, 0), P(11, 0), Down, P(3, 4), P(0, 4))
	cube.ConnectRange(Left, P(8, 0), P(8, 3), Down, P(4, 4), P(7, 4))
	cube.ConnectRange(Right, P(11, 0), P(11, 3), Left, P(15, 11), P(15, 8))
	cube.ConnectRange(Right, P(11, 4), P(11, 7), Down, P(15, 8), P(12, 8))
	cube.ConnectRange(Left, P(0, 4), P(0, 7), Up, P(15, 11), P(12, 11))
	cube.ConnectRange(Down, P(0, 7), P(3, 7), Up, P(11, 11), P(8, 11))
	cube.ConnectRange(Down, P(4, 7), P(7, 7), Right, P(8, 11), P(11, 11))
}

func TestPart2MainFunc(t *testing.T) {
	cube := mustParseCube(inputTest)
	connectTest(cube)

	res, err := part2MainFunc(cube, inputTestPath)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 5031
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
