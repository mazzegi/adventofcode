package comb

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	var in []int
	var size int

	in = []int{1, 2, 3, 4}
	size = 3
	for part := range Parts(in, size) {
		fmt.Printf("%v\n", part)
	}
}
