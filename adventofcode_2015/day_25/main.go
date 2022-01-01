package main

import "fmt"

//Enter the code at row 2978, column 3083.

func main() {
	// for row := 1; row <= 6; row++ {
	// 	fmt.Printf("%d: ", row)
	// 	for col := 1; col <= 6; col++ {
	// 		c := counterAt(row, col)
	// 		cd := code(c)
	// 		fmt.Printf("%d ", cd)
	// 	}
	// 	fmt.Printf("\n")
	// }
	cnt := counterAt(2978, 3083)
	cd := code(cnt)
	fmt.Printf("(%d) => %d\n", cnt, cd)
}

func code(n int) int {
	c := 20151125
	for i := 1; i < n; i++ {
		c = (c * 252533) % 33554393
	}
	return c
}

func counterAt(row int, col int) int {
	rf := rowFirstCounter(row)

	step := row + 1
	cn := rf
	for i := 1; i < col; i++ {
		cn = cn + step
		step++
	}
	return cn
}

func rowFirstCounter(row int) int {
	if row == 1 {
		return 1
	}
	if row == 2 {
		return 2
	}
	rnm := 1
	rn := 2
	for i := 0; i < row-2; i++ {
		prnm := rn
		rn = rn + (rn - rnm + 1)
		rnm = prnm
	}
	return rn
}
