package day_09

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = `

`

func TestGroupScore(t *testing.T) {
	tests := []struct {
		in    string
		score int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for i, test := range tests {
		score, _, err := groupScore(test.in)
		testutil.CheckUnexpectedError(t, err)
		if test.score != score {
			t.Fatalf("%d: want %d, have %d", i, test.score, score)
		}
	}
}

func TestGarbage(t *testing.T) {
	tests := []struct {
		in      string
		garbage int
	}{
		{"{<>}", 0},
		{"{<random characters>}", 17},
		{"{<<<<>}", 3},
		{"{<{!>}>}", 2},
		{"{<!!>}", 0},
		{"{<!!!>>}", 0},
		{"{<{odi!a,<{i<a>}", 10},
	}

	for i, test := range tests {
		_, garbage, err := groupScore(test.in)
		testutil.CheckUnexpectedError(t, err)
		if test.garbage != garbage {
			t.Fatalf("%d: want %d, have %d", i, test.garbage, garbage)
		}
	}
}
