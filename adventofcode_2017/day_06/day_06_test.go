package day_06

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = "0 2 7 0"

func TestDistCyclesToRepeat(t *testing.T) {
	res, err := distCyclesToRepeat(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 5
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestLoopSize(t *testing.T) {
	res, err := loopSize(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 4
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
