package day_24

import (
	"adventofcode_2017/testutil"
	"testing"
)

const inputTest = `
0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
`

func TestPart1MainFunc(t *testing.T) {
	res, err := strongestBridge(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 31
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := strongestLongestBridge(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 19
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
