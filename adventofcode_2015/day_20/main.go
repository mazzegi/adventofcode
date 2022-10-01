package main

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2015/numbers"
)

func main() {
	// houses := 9
	// for h := 1; h <= houses; h++ {
	// 	fmt.Printf("house %d got %d presents\n", h, presents(h))
	// }

	//#CAND: 3310049

	//target := 3310000
	//solve(target)
	//ps := presents(3310049)
	//fmt.Printf("target gets %d presents\n", ps)
	target := 33100000
	//first := math.Sqrt(2.0*float64(target)+0.25) - 0.5
	house := 0
	//house := int(first)
	var max int
	for {
		ps := presents(house)
		if ps >= target {
			fmt.Printf("house %d got %d presents\n", house, target)
			break
		}
		if ps > max {
			max = ps
			fmt.Printf("new max: house %d max = %d\n", house, max)
		}
		// if house%1000 == 0 {
		// 	fmt.Printf("checked house %d (=> %d, max = %d)\n", house, ps, max)
		// }
		house += 20
	}
}

func presents(house int) int {
	if house <= 0 {
		return 0
	}
	var ps int
	elveCount := house
	//for e := 1; e <= elveCount; e++ {
	for e := elveCount; e >= 1; e-- {
		if e*50 < house {
			continue
		}
		if house%e == 0 {
			//ps += 10 * e
			ps += 11 * e
		}
	}
	return ps
}

func solve(target int) {
	e := 2
	for {
		es := makeSlice(e)
		lcm := numbers.LeastCommonMutiple(es...)
		ps := presents(lcm)
		if ps == target {
			fmt.Printf("found house for target %d\n", lcm)
			return
		} else if ps > target {
			fmt.Printf("örks\n")
			return
		}
		e++
	}
}

func makeSlice(n int) []int {
	var ns []int
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
	}
	return ns
}
