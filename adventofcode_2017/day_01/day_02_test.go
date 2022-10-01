package day_01

import (
	"fmt"
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		in  string
		exp int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res, err := Solve(test.in)
			testutil.CheckUnexpectedError(t, err)
			if res != test.exp {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}

func TestSolveHalfWayRound(t *testing.T) {
	tests := []struct {
		in  string
		exp int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res, err := SolveHalfWayRound(test.in)
			testutil.CheckUnexpectedError(t, err)
			if res != test.exp {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}
