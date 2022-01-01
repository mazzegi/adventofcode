package day_03

import (
	"fmt"
	"testing"
)

func TestIsPossibleTriangle(t *testing.T) {
	tests := []struct {
		in       [3]int
		possible bool
	}{
		{
			in:       [3]int{1, 2, 3},
			possible: false,
		},
		{
			in:       [3]int{5, 10, 25},
			possible: false,
		},
		{
			in:       [3]int{5, 10, 12},
			possible: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			isp := IsPossibleTriangle(test.in)
			if isp != test.possible {
				t.Fatalf("expect %t, got %t", test.possible, isp)
			}
		})
	}
}
