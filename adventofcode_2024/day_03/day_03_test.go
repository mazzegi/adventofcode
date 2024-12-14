package day_03

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const inputTest2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 161
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 48
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
