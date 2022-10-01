package day_11

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

const inputTest = `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

func TestTotalFlashes(t *testing.T) {
	res, err := totalFlashes(inputTest, 100)
	testutil.CheckUnexpectedError(t, err)
	exp := 1656
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestTotalFlashes10(t *testing.T) {
	res, err := totalFlashes(inputTest, 10)
	testutil.CheckUnexpectedError(t, err)
	exp := 204
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestStepsToAllFlash(t *testing.T) {
	res, err := stepsToAllFlash(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 195
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
