package day_24

import (
	"adventofcode_2021/testutil"
	"testing"
)

const inputTest = `
inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2
`

func TestPart1MainFunc(t *testing.T) {
	res, err := largestModelNumber2(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestAlu(t *testing.T) {
	iss, err := parseInstructions(inputTest)
	testutil.CheckUnexpectedError(t, err)

	a := &alu{iss: iss}
	a.run([]int{14}, 0)
	log("%d, %d, %d, %d", a.w, a.x, a.y, a.z)
}

func TestProbe(t *testing.T) {
	err := probeSubs(input)
	testutil.CheckUnexpectedError(t, err)
}

func TestCompare(t *testing.T) {
	err := compareSubs(input)
	testutil.CheckUnexpectedError(t, err)
}

func TestTestExtra(t *testing.T) {
	err := testExtra(input)
	testutil.CheckUnexpectedError(t, err)
}

func TestTestSolveExtra(t *testing.T) {
	err := testSolveExtra(input)
	testutil.CheckUnexpectedError(t, err)
}

func TestReverseSolve(t *testing.T) {
	err := reverseSolve(input)
	testutil.CheckUnexpectedError(t, err)
}
