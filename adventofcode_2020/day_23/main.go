package main

import (
	"adventofcode_2020/ring2"
	"fmt"
	"time"
)

func main() {
	in := input
	// doSpike := false
	// if doSpike {
	// 	spike(in)
	// 	os.Exit(0)
	// }

	for n := 10; n <= 1000000; n++ {
		in = append(in, n)
	}
	buf := ring2.NewBuffer(in)

	fmt.Printf("start\n")
	moves := 10000000
	t0 := time.Now()
	for i := 0; i < moves; i++ {
		step := i + 2
		if step%100 == 0 {
			dur := time.Since((t0))
			t0 = time.Now()
			fmt.Printf("%d / %d (%s)\n", step, moves, dur.Round(time.Millisecond))
		}
		buf.Step()
	}

	nto, prod := buf.TwoNextToOne()
	fmt.Printf("next to one: %s => %d\n", nto, prod)
	fmt.Printf("done\n")
}

// func spike(in []int) {
// 	fillTo := 14
// 	for n := 10; n <= fillTo; n++ {
// 		in = append(in, n)
// 	}
// 	buf := ring.NewBuffer(in)

// 	h := buf.NextToOne()
// 	visited := map[string]int{}
// 	visited[h] = 0

// 	moves := 100000
// 	fmt.Printf("%d: %s\n", 1, buf.NextToOne())
// 	for i := 0; i < moves; i++ {
// 		step := i + 2
// 		buf.Step()
// 		fmt.Printf("%d: %s\n", step, buf.NextToOne())

// 		h = buf.NextToOne()
// 		if vstep, ok := visited[h]; ok {
// 			fmt.Printf("step %d: already visited %q in step %d\n", step, h, vstep)
// 			break
// 		} else {
// 			visited[h] = step
// 		}
// 	}

// 	fmt.Printf("next to one: %s\n", buf.NextToOne())
// 	fmt.Printf("done\n")
// }

var inputTest = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

var input = []int{8, 7, 2, 4, 9, 5, 1, 3, 6}
