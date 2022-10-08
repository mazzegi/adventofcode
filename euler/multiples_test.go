package euler

import (
	"fmt"
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	tests := []struct {
		ns []int
		sm int
	}{
		{[]int{9, 12}, 36},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2520},
		{[]int{18, 28, 44}, 2772},
		{[]int{2028, 4702, 5898}, 4686774924},
		{[]int{23326, 231614, 268296}, 362375881472136},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := SmallestMultipleOf(test.ns...)
			if test.sm != res {
				t.Fatalf("want %v, have %v", test.sm, res)
			}
		})
	}
}
