package day_09

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

const inputTest = `
2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestSumOfRiskLevels(t *testing.T) {
	res, err := sumOfRiskLevels(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 15
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestProductOf3LargestBasinSizes(t *testing.T) {
	res, err := productOf3LargestBasinSizes(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 1134
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
