package day_03

import (
	"fmt"
	"testing"
)

func TestCarrySteps(t *testing.T) {
	tests := []struct {
		square int
		exp    int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res := CarrySteps(test.square)
			if res != test.exp {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}

func TestStressTest(t *testing.T) {
	tests := []struct {
		limit int
		exp   int
	}{
		{10, 11},
		{117, 122},
		{340, 351},
		{748, 806},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res := StressTest(test.limit)
			if res != test.exp {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}
