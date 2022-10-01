package day_08

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = `
b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
`

func TestLargestRegisterValue(t *testing.T) {
	res, _, err := largestRegisterValue(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 1
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestHighestEver(t *testing.T) {
	_, he, err := largestRegisterValue(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 10
	if exp != he {
		t.Fatalf("want %d, have %d", exp, he)
	}
}
