package day_05

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

const inputTest = `
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

func TestOverlappingCountHV(t *testing.T) {
	cnt, err := OverlappingCount(inputTest, LinePointsHV)
	testutil.CheckUnexpectedError(t, err)

	exp := 5
	if exp != cnt {
		t.Fatalf("want %d, have %d", exp, cnt)
	}
}

func TestOverlappingCountHVD(t *testing.T) {
	cnt, err := OverlappingCount(inputTest, LinePointsHVD)
	testutil.CheckUnexpectedError(t, err)

	exp := 12
	if exp != cnt {
		t.Fatalf("want %d, have %d", exp, cnt)
	}
}
