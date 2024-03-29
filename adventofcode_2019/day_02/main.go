package main

import (
	"fmt"
	"os"

	"github.com/mazzegi/adventofcode/adventofcode_2019/intcode"
)

func main() {
	input[1] = 12
	input[2] = 2

	out, err := intcode.Exec(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("out:", out)

	//
	target := 19690720
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			input[1] = noun
			input[2] = verb
			out, err = intcode.Exec(input)
			if err != nil {
				panic(err)
			}
			if out[0] == target {
				fmt.Printf("match: noun=%d verb=%d -> %d\n", noun, verb, 100*noun+verb)
				os.Exit(0)
			}
		}
	}
}

var input = []int{
	1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 1, 23, 13, 27, 2, 6, 27, 31, 1, 5, 31, 35, 2, 10, 35, 39, 1, 6, 39,
	43, 1, 13, 43, 47, 2, 47, 6, 51, 1, 51, 5, 55, 1, 55, 6, 59, 2, 59, 10, 63, 1, 63, 6, 67, 2, 67, 10, 71, 1, 71, 9, 75, 2, 75, 10, 79, 1, 79, 5,
	83, 2, 10, 83, 87, 1, 87, 6, 91, 2, 9, 91, 95, 1, 95, 5, 99, 1, 5, 99, 103, 1, 103, 10, 107, 1, 9, 107, 111, 1, 6, 111, 115, 1, 115, 5, 119, 1,
	10, 119, 123, 2, 6, 123, 127, 2, 127, 6, 131, 1, 131, 2, 135, 1, 10, 135, 0, 99, 2, 0, 14, 0,
}
