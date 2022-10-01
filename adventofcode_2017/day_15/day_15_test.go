package day_15

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

func TestPart1MainFunc(t *testing.T) {
	res, err := judgeCount(testGenValues, Mio40)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 588
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := judgeCountMult(testGenValues, Mio5)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 309
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
