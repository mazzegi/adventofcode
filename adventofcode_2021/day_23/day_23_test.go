package day_23

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

func TestPart1MainFunc(t *testing.T) {
	res, err := leastEnergy(setupTest())
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(setupTest())
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPossibleMoves(t *testing.T) {
	b := setupTest()
	mvs := b.possibleMoves()
	for _, mv := range mvs {
		log("%s", mv.String())
	}
}
