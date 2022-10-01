package day_07

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

const inputTest = `
16,1,2,0,4,2,7,1,2,14
`

func TestLeastFuelLinearCost(t *testing.T) {
	res, err := LeastFuel(inputTest, linearCost)
	testutil.CheckUnexpectedError(t, err)
	exp := 37
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestLeastFuelProgressiveCost(t *testing.T) {
	res, err := LeastFuel(inputTest, progressiveCost)
	testutil.CheckUnexpectedError(t, err)
	exp := 168
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
