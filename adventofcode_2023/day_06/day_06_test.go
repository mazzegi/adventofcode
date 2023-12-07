package day_06

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

var inputTest = []RaceResult{
	{Time: 7, Distance: 9},
	{Time: 15, Distance: 40},
	{Time: 30, Distance: 200},
}

var inputTest2 = RaceResult{
	Time:     71530,
	Distance: 940200,
}

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 288
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 71503
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
