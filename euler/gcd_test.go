package euler

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

func TestGCD(t *testing.T) {
	tx := testutil.NewTx(t)

	tests := []struct {
		a       int
		b       int
		wantGCD int
	}{
		{12, 8, 4},
		{0, 3, 3},
		{144, 88, 8},
		{567, 81, 81},
		{3, 9, 3},
		{517, 1023, 11},
		{37, 53, 1},
	}

	for _, test := range tests {
		gcd := GCD(test.a, test.b)
		tx.AssertEqual(test.wantGCD, gcd)
	}
}
