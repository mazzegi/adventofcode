package day_20

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = `
p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>
`

func TestPart1MainFunc(t *testing.T) {
	res, err := closestToOrigin(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 0
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := leftAfterCollisions(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
