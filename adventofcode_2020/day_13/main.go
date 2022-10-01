package main

import (
	"fmt"
	"time"

	"github.com/mazzegi/adventofcode/adventofcode_2020/shuttle"
)

func main() {
	in, err := shuttle.ParseInput(inputTest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", in)

	res, err := in.FindEarliestDeparture()
	if err != nil {
		panic(err)
	}
	fmt.Printf("the earliest dep. is with bus %d after %d minutes (res = %d)\n", res.BusID, res.DepartAt, res.BusID*res.DepartAt)

	fmt.Printf("*** Part 2 ***\n\n")
	// bs, err := shuttle.ParseBusses(inputTest)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, b := range bs {
	// 	fmt.Printf("bus: %d, offset = %d\n", b.ID, b.Offset)
	// }
	// t, err := shuttle.FindEarliestSubsequent(bs)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("the earliest timestamp is %d\n", t)

	//runTest(inputTest)
	runTest(inputTest1)
	runTest(inputTest2)
	runTest(inputTest3)
	runTest(inputTest4)
	runTest(inputTest5)

	// fmt.Printf("\nspike ***\n")
	// shuttle.PrintSubsequentIntelli([]shuttle.Bus{
	// 	{
	// 		ID:     17,
	// 		Offset: 0,
	// 	},
	// 	{
	// 		ID:     13,
	// 		Offset: 2,
	// 	},
	// }, 5)

	// fmt.Printf("\nspike ***\n")
	// shuttle.PrintSubsequentIntelli([]shuttle.Bus{
	// 	{
	// 		ID:     17,
	// 		Offset: 0,
	// 	},
	// 	{
	// 		ID:     13,
	// 		Offset: 2,
	// 	},
	// 	{
	// 		ID:     19,
	// 		Offset: 3,
	// 	},
	// }, 5)

	runTest(input)
}

func runTest(rin shuttle.RawInput) {
	fmt.Printf("\n*** run test ***\n")
	bs, err := shuttle.ParseBusses(rin)
	if err != nil {
		panic(err)
	}
	t0 := time.Now()
	//t, err := shuttle.FindEarliestSubsequent(bs)
	t, err := shuttle.FindEarliestSubsequentIntelli(bs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%s]: => %d (%s)\n", rin.Shuttles, t, time.Since(t0))
}

var inputTest = shuttle.RawInput{
	EarliestTime: "939",
	Shuttles:     "7,13,x,x,59,x,31,19",
}

var inputTest1 = shuttle.RawInput{
	EarliestTime: "0",
	Shuttles:     "17,x,13,19",
}

var inputTest2 = shuttle.RawInput{
	EarliestTime: "0",
	Shuttles:     "67,7,59,61",
}

var inputTest3 = shuttle.RawInput{
	EarliestTime: "0",
	Shuttles:     "67,x,7,59,61",
}

var inputTest4 = shuttle.RawInput{
	EarliestTime: "0",
	Shuttles:     "67,7,x,59,61",
}

var inputTest5 = shuttle.RawInput{
	EarliestTime: "0",
	Shuttles:     "1789,37,47,1889",
}

var input = shuttle.RawInput{
	EarliestTime: "1000066",
	Shuttles:     "13,x,x,41,x,x,x,37,x,x,x,x,x,659,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,23,x,x,x,x,x,29,x,409,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17",
}
