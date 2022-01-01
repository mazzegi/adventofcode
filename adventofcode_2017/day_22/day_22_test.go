package day_22

import (
	"adventofcode_2017/testutil"
	"testing"
)

const inputTest = `
..#
#..
...
`

func TestPart1MainFunc(t *testing.T) {
	res, err := numInfections(inputTest, 10000)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 5587
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := numInfectionsExt(inputTest, 100)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 26
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFuncMany(t *testing.T) {
	res, err := numInfectionsExt(inputTest, 10000000)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 2511944
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
