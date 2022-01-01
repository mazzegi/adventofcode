package main

import (
	"adventofcode_2015/numbers"
	"fmt"
	"math/big"
)

func main() {
	max := 3310000
	n := 1
	sum := 1
	at := big.NewInt(1)
	for sum <= max {
		if numbers.IsPrime(n) {
			fmt.Printf("%d (sum=%d) (at=%d)\n", n, sum, at)
			sum += n
			at.Mul(at, big.NewInt(int64(n)))
		}
		n++
	}
	// target := 3310000
	// pfs := numbers.PrimeFactors(target)
	// fmt.Printf("%v\n", pfs)

	// 331 * 10000
}
