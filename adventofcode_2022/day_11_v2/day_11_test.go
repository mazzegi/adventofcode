package day_11_v2

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

var inputTest = []InputMonkey{
	{
		Items:          []int{79, 98},
		Operation:      MultBy(19),
		TestDivBy:      23,
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 3,
	},
	{
		Items:          []int{54, 65, 75, 74},
		Operation:      Add(6),
		TestDivBy:      19,
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 0,
	},
	{
		Items:          []int{79, 60, 97},
		Operation:      Square,
		TestDivBy:      13,
		ThrowToIfTrue:  1,
		ThrowToIfFalse: 3,
	},
	{
		Items:          []int{74},
		Operation:      Add(3),
		TestDivBy:      17,
		ThrowToIfTrue:  0,
		ThrowToIfFalse: 1,
	},
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 2713310158
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
