package day_17

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

var targetAreaTest = area{
	xmin: 20,
	xmax: 30,
	ymin: -10,
	ymax: -5,
}

func TestPart1MainFunc(t *testing.T) {
	res, err := highestYPos(targetAreaTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 45
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(targetAreaTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
