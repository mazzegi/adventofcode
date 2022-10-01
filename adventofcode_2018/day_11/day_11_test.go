package day_11

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2018/testutil"
)

func TestPowerLevel(t *testing.T) {
	var res int

	res = powerLevel(point{3, 5}, 8)
	testutil.Assert(t, res == 4, 4, res)

	res = powerLevel(point{122, 79}, 57)
	testutil.Assert(t, res == -5, -5, res)

	res = powerLevel(point{217, 196}, 39)
	testutil.Assert(t, res == 0, 0, res)

	res = powerLevel(point{101, 153}, 71)
	testutil.Assert(t, res == 4, 4, res)
}

func TestPart1MainFunc(t *testing.T) {
	res, err := coordOfLargestPowerLevelSquare(18)
	testutil.CheckUnexpectedError(t, err)
	exp := point{33, 45}
	if exp != res {
		t.Fatalf("want %s, have %s", exp, res)
	}

	res, err = coordOfLargestPowerLevelSquare(42)
	testutil.CheckUnexpectedError(t, err)
	exp = point{21, 61}
	if exp != res {
		t.Fatalf("want %s, have %s", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := coordOfLargestPowerLevelFlexSquare(18)
	testutil.CheckUnexpectedError(t, err)
	testutil.Assert(t, res == "90,269,16", "90,269,16", res)

	res, err = coordOfLargestPowerLevelFlexSquare(42)
	testutil.CheckUnexpectedError(t, err)
	testutil.Assert(t, res == "232,251,12", "232,251,12", res)

}
