package day_20

import (
	"adventofcode_2016/testutil"
	"testing"
)

const inputTest = `
5-8
0-2
4-7
`

func TestLowestNonBlocked(t *testing.T) {
	n, err := LowestNonBlocked(inputTest)
	testutil.CheckUnexpectedError(t, err)

	exp := 3
	if exp != n {
		t.Fatalf("want %d, have %d", exp, n)
	}
}

func TestAllowedCount(t *testing.T) {
	n, err := AllowedCount(inputTest, 9)
	testutil.CheckUnexpectedError(t, err)

	exp := 2
	if exp != n {
		t.Fatalf("want %d, have %d", exp, n)
	}
}
