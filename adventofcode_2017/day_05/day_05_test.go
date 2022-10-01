package day_05

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = `
0
3
0
1
-3
`

func TestStepsToExit(t *testing.T) {
	res, err := StepsToExit(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 5
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestStepsToExitDec(t *testing.T) {
	res, err := StepsToExitDec(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 10
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
