package day_01

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

const inputTest2 = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 142
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 281
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
