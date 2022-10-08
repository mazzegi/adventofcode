package euler

import (
	"fmt"
	"math/big"
)

func IsPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(1)
}

func NextPrime(n int) int {
	for {
		n++
		if IsPrime(n) {
			return n
		}
		if n < 0 {
			//overflow
			panic("overflow in next prime")
		}
	}
}

func PrimeFactors(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	var pfs []int
	rem := n
outer:
	for {
		if IsPrime(rem) {
			pfs = append(pfs, rem)
			break
		}

		for f := 2; f < rem; f = NextPrime(f) {
			if rem%f == 0 {
				pfs = append(pfs, f)
				rem /= f
				continue outer
			}
		}
		panic(fmt.Sprintf("found no prime fac of %d though its not a prime", rem))
	}

	return pfs
}
