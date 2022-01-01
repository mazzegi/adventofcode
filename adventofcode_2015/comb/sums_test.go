package comb

import (
	"fmt"
	"testing"
)

func TestFixedSumInts(t *testing.T) {
	var sum int
	var len int

	sum = 10
	len = 3
	for ns := range FixedSumInts(sum, len) {
		fmt.Printf("%v\n", ns)
	}
}
