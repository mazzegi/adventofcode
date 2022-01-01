package day_13

import (
	"adventofcode_2017/testutil"
	"testing"
)

const inputTest = `
0: 3
1: 2
4: 4
6: 4
`

func TestPart1MainFunc(t *testing.T) {
	res, err := tripSeverity(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 24
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := minDelay(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 10
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
