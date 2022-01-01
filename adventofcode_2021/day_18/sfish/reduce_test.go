package sfish

import (
	"adventofcode_2021/testutil"
	"fmt"
	"testing"
)

func TestExplode(t *testing.T) {
	tests := []struct {
		in   string
		out  string
		skip bool
	}{
		{
			in:   "[[[[[9,8],1],2],3],4]",
			out:  "[[[[0,9],2],3],4]",
			skip: true,
		},
		{
			in:   "[7,[6,[5,[4,[3,2]]]]]",
			out:  "[7,[6,[5,[7,0]]]]",
			skip: true,
		},
		{
			in:   "[[6,[5,[4,[3,2]]]],1]",
			out:  "[[6,[5,[7,0]]],3]",
			skip: true,
		},
		{
			in:   "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			out:  "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			skip: false,
		},
		{
			in:   "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			out:  "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
			skip: true,
		},
	}

	doSkip := true
	for i, test := range tests {
		if doSkip && test.skip {
			t.Logf("skip %d", i+1)
			continue
		}
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			p, err := Parse(test.in)
			testutil.CheckUnexpectedError(t, err)
			explode(p, 0)
			if p.String() != test.out {
				t.Fatalf("want %q, have %q", test.out, p.String())
			}
		})
	}
}
