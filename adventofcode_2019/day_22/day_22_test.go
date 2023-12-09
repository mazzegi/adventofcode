package day_22

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest1 = `
deal with increment 7
deal into new stack
deal into new stack
`

var testResult1 = []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}

const inputTest2 = `
cut 6
deal with increment 7
deal into new stack
`

var testResult2 = []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}

const inputTest3 = `
deal with increment 7
deal with increment 9
cut -2
`

var testResult3 = []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}

const inputTest4 = `
deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1
`

var testResult4 = []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}

func TestPart1MainFunc(t *testing.T) {
	{
		res := shuffle(inputTest1, 10)
		testutil.Assert(t, testResult1, res)
	}
	{
		res := shuffle(inputTest2, 10)
		testutil.Assert(t, testResult2, res)
	}
	{
		res := shuffle(inputTest3, 10)
		testutil.Assert(t, testResult3, res)
	}
	{
		res := shuffle(inputTest4, 10)
		testutil.Assert(t, testResult4, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest1, 10)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
