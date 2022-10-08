package euler

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		n    int
		facs []int
	}{
		{13195, []int{5, 7, 13, 29}},
		{1, []int{1}},
		{2, []int{2}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := PrimeFactors(test.n)
			if !reflect.DeepEqual(test.facs, res) {
				t.Fatalf("want %v, have %v", test.facs, res)
			}
		})
	}
}
