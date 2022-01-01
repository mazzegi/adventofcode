package day_13

import (
	"adventofcode_2016/testutil"
	"testing"
)

func TestPart1MainFunc(t *testing.T) {
	res, err := minSteps(10, p(7, 4))
	testutil.CheckUnexpectedError(t, err)
	var exp int = 11
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := locsIn50(10)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestMakeMaze(t *testing.T) {
	m := makeMaze(10)
	log("\n%s", m.dump(9, 6))
}
