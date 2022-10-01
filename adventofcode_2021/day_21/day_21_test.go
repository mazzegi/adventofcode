package day_21

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

func TestPart1MainFunc(t *testing.T) {
	res, err := practiceCoefficient(4, 8)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 739785
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := winUniverseCount(4, 8)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 444356092776315
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
