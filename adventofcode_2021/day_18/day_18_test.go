package day_18

import (
	"fmt"
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/day_18/sfish"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

const inputTest = `
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`

func TestParse(t *testing.T) {
	lines := readutil.ReadLines(inputTest)
	for _, line := range lines {
		p, err := sfish.Parse(line)
		testutil.CheckUnexpectedError(t, err)
		if line != p.String() {
			t.Fatalf("expect %q, got %q", line, p.String())
		}
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		in   string
		exp  int
		skip bool
	}{
		{
			in:   "[[1,2],[[3,4],5]]",
			exp:  143,
			skip: true,
		},
		{
			in:   "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			exp:  1384,
			skip: true,
		},
		{
			in:   "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			exp:  445,
			skip: false,
		},
		{
			in:   "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			exp:  791,
			skip: true,
		},
		{
			in:   "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			exp:  1137,
			skip: true,
		},
		{
			in:   "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			exp:  3488,
			skip: true,
		},
	}

	doSkip := false
	for i, test := range tests {
		if doSkip && test.skip {
			t.Logf("skip %d", i+1)
			continue
		}
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			p, err := sfish.Parse(test.in)
			testutil.CheckUnexpectedError(t, err)
			res := p.Magnitude()
			if res != test.exp {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}

func TestMagnitudeOfFinalSum(t *testing.T) {
	res, err := magnitudeOfFinalSum(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 4140
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestLargestSumMagnitude(t *testing.T) {
	res, err := largestSumMagnitude(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3993
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
