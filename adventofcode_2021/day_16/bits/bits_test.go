package bits

import (
	"fmt"
	"testing"
)

func TestBitsToNumber(t *testing.T) {
	tests := []struct {
		in  []bool
		exp int
	}{
		{
			in:  []bool{true, false, false},
			exp: 4,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res := bitsToNumber(test.in)
			if res != test.exp {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}
