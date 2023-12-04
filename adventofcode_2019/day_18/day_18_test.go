package day_18

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest1 = `
#########
#b.A.@.a#
#########
`

const inputTest2 = `
########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################
`

const inputTest3 = `
########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################
`

const inputTest4 = `
#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################
`

const inputTest5 = `
########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################
`

func TestPart1MainFunc(t *testing.T) {
	{
		res, err := part1MainFunc(inputTest1)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 8
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	}
	{
		res, err := part1MainFunc(inputTest2)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 86
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	}
	{
		res, err := part1MainFunc(inputTest3)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 132
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	}
	{
		res, err := part1MainFunc(inputTest4)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 136
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	}
	{
		res, err := part1MainFunc(inputTest5)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 81
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest1)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
