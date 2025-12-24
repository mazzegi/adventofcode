package rat

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

func TestSimplify(t *testing.T) {
	tx := testutil.NewTx(t)

	tests := []struct {
		rat               Number
		wantSimplifiedRat Number
	}{
		{R(4, 2), R(2, 1)},
		{R(24, 32), R(3, 4)},

		{R(-2, -3), R(2, 3)},
		{R(-2, 3), R(-2, 3)},
		{R(2, -3), R(-2, 3)},

		{R(0, -34), R(0, 1)},
	}

	for _, test := range tests {
		simplifiedRat := Simplify(test.rat)
		tx.AssertEqual(test.wantSimplifiedRat, simplifiedRat)
	}
}

func TestAdd(t *testing.T) {
	tx := testutil.NewTx(t)

	tests := []struct {
		rat1, rat2 Number
		wantSum    Number
	}{
		{R(1, 2), R(2, 3), R(7, 6)},
		{R(12, 16), R(4, 8), R(5, 4)},
	}

	for _, test := range tests {
		sum := Simplify(Add(test.rat1, test.rat2))
		tx.AssertEqual(test.wantSum, sum)
	}
}

func TestSub(t *testing.T) {
	tx := testutil.NewTx(t)

	tests := []struct {
		rat1, rat2 Number
		wantSum    Number
	}{
		{R(1, 2), R(2, 3), R(-1, 6)},
		{R(12, 16), R(4, 8), R(1, 4)},
	}

	for _, test := range tests {
		sum := Simplify(Sub(test.rat1, test.rat2))
		tx.AssertEqual(test.wantSum, sum)
	}
}
