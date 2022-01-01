package day_21

import (
	"adventofcode_2016/testutil"
	"testing"
)

const inputTest = `
swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d
`

func TestPart1MainFunc(t *testing.T) {
	res, err := scrambled(inputTest, "abcde")
	testutil.CheckUnexpectedError(t, err)
	var exp string = "decab"
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}
}

// func TestPart2MainFunc(t *testing.T) {
// 	res, err := unscramble(inputTest)
// 	testutil.CheckUnexpectedError(t, err)
// 	var exp int = -42
// 	if exp != res {
// 		t.Fatalf("want %d, have %d", exp, res)
// 	}
// }

func TestReverse(t *testing.T) {
	inp := "dcafegbh"
	exp := "dbgefach"
	res := reversePositions{x: 1, y: 6}.apply([]byte(inp))
	if exp != string(res) {
		t.Fatalf("want %q, have %q", exp, res)
	}
}
