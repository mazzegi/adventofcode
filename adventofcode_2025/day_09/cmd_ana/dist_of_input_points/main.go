package main

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2025/day_09"
)

func main() {
	p1, p2, minD := day_09.MinDistOfInputPoints()
	fmt.Println(p1, " - ", p2, " = ", minD)
}
