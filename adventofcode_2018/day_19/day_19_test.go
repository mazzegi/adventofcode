package day_19

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2018/testutil"
)

const inputTestIP = 0

const inputTest = `
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, inputTestIP)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 6
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, inputTestIP)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
