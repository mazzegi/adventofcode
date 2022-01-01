package main

import (
	"fmt"
	"time"
)

var inputTest0 = []int{0, 3, 6}
var inputTest1 = []int{1, 3, 2}
var inputTest2 = []int{2, 1, 3}
var inputTest3 = []int{1, 2, 3}
var inputTest4 = []int{2, 3, 1}
var inputTest5 = []int{3, 2, 1}
var inputTest6 = []int{3, 1, 2}

var input = []int{2, 0, 1, 9, 5, 19}

func main() {
	//turns := 2020
	turns := 30000000
	tests := [][]int{
		inputTest0, inputTest1, inputTest2, inputTest3, inputTest4, inputTest5, inputTest6, input,
	}
	for _, in := range tests {
		runTest(in, turns)
	}

	// in := inputTest6
	// //in := input
	// fmt.Printf("input: %v\n", in)
	// //turns := 2020 - len(in)
	// turns := 30000000 - len(in)

}

func runTest(in []int, turns int) {
	sin := fmt.Sprintf("%v", in)
	turns = turns - len(in)

	lookup := map[int]int{}
	for i := 0; i < len(in)-1; i++ {
		lookup[in[i]] = i
	}
	lastSpoken := in[len(in)-1]

	findTurn := func(num int) (int, bool) {
		pos, ok := lookup[num]
		return pos, ok
	}

	t0 := time.Now()
	for i := 0; i < turns; i++ {
		pos, found := findTurn(lastSpoken)
		//but now save last spoken
		lookup[lastSpoken] = len(in) - 1

		if !found {
			in = append(in, 0)
		} else {
			in = append(in, len(in)-pos-1)
		}
		lastSpoken = in[len(in)-1]
		// if i%100000 == 0 {
		// 	p := 100.0 * float64(i) / float64(turns)
		// 	fmt.Printf("reached %.1f %%\n", p)
		// }

	}
	fmt.Printf("in: %s => last number spoken: %d (%s)\n", sin, in[len(in)-1], time.Since(t0))
}
