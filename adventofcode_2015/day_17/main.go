package main

import (
	"adventofcode_2015/comb"
	"fmt"
)

var input = []int{43, 3, 4, 10, 21, 44, 4, 6, 47, 41, 34, 17, 17, 44, 36, 31, 46, 9, 27, 38}

var inputTest = []int{20, 15, 10, 5, 5}

func main() {
	test := false
	var total int
	var in []int
	if test {
		total = 25
		in = inputTest
	} else {
		total = 150
		in = input
	}

	var cnt int
	minSize := -1
	minCount := 0
	for size := 1; size < len(in); size++ {
		for part := range comb.Parts(in, size) {
			ps := sum(part)
			if ps == total {
				fmt.Printf("sum-of %v = %d\n", part, total)
				cnt++
				if minSize < 0 || len(part) < minSize {
					minSize = len(part)
					minCount = 1
				} else if len(part) == minSize {
					minCount++
				}
			}
		}
	}
	fmt.Printf("total: %d\n", cnt)
	fmt.Printf("min: %d, %d times\n", minSize, minCount)
}

func sum(ns []int) int {
	var s int
	for _, n := range ns {
		s += n
	}
	return s
}
