package day_17

import (
	"adventofcode_2017/testutil"
	"testing"
)

const (
	inputStepTest         = 3
	inputInsertsTest      = 2017
	inputInsertsPart2test = 50000000
)

func TestPart1MainFunc(t *testing.T) {
	res, err := valAfterLastInsert(inputStepTest, inputInsertsTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 638
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := valAfter0(inputStepTest, inputInsertsPart2test)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 638
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
