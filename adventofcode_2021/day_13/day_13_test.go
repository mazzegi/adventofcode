package day_13

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

const inputDotsTest = `
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0
`

const inputFoldsTest = `
fold along y = 7
fold along x = 5
`

func TestDotsAfterFirstFold(t *testing.T) {
	res, err := dotsAfterFirstFold(inputDotsTest, inputFoldsTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 17
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	_, err := displayCode(inputDotsTest, inputFoldsTest)
	testutil.CheckUnexpectedError(t, err)

}
