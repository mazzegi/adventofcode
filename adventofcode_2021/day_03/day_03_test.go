package day_03

import (
	"adventofcode_2021/testutil"
	"testing"
)

const inputTest = `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

func TestCalcRate(t *testing.T) {
	r, err := CalcRate(inputTest)
	testutil.CheckUnexpectedError(t, err)

	exp := Rate{
		gamma:   22,
		epsilon: 9,
	}
	if exp != r {
		t.Fatalf("want %v, have %v", exp, r)
	}
}

func TestCalcLifeSupportRating(t *testing.T) {
	r, err := CalcLifeSupportRating(inputTest)
	testutil.CheckUnexpectedError(t, err)

	exp := LifeSupportRating{
		oxyGen:   23,
		co2Scrub: 10,
	}
	if exp != r {
		t.Fatalf("want %v, have %v", exp, r)
	}
}
