package day_11

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

func TestPart1MainFunc(t *testing.T) {
	tests := []struct {
		in  string
		exp int
	}{
		{
			in:  "ne,ne,ne",
			exp: 3,
		},
		{
			in:  "ne,ne,sw,sw",
			exp: 0,
		},
		{
			in:  "ne,ne,s,s",
			exp: 2,
		},
		{
			in:  "se,sw,se,sw,sw",
			exp: 3,
		},
	}

	for i, test := range tests {
		res, err := distance(test.in)
		testutil.CheckUnexpectedError(t, err)
		if test.exp != res {
			t.Fatalf("%d: want %d, have %d", i, test.exp, res)
		}
	}
}

func TestPart2MainFunc(t *testing.T) {
	// res, err := part2MainFunc(inputTest)
	// testutil.CheckUnexpectedError(t, err)
	// var exp int = -42
	// if exp != res {
	// 	t.Fatalf("want %d, have %d", exp, res)
	// }
}
