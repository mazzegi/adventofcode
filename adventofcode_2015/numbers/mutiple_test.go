package numbers

import (
	"fmt"
	"testing"
)

func TestLeastCommonMultipleOfPair(t *testing.T) {
	var n1 int
	var n2 int

	n1, n2 = 12345, 21324
	fmt.Printf("%d, %d => %d\n", n1, n2, LeastCommonMutipleOfPair(n1, n2))
}

func TestLeastCommonMultiple(t *testing.T) {
	var n int
	n = 27
	var ns []int
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
	}

	fmt.Printf("%v => %d\n", ns, LeastCommonMutiple(ns...))
}

func TestPrimeFactors(t *testing.T) {
	n := 86792
	fmt.Printf("%d: %v\n", n, PrimeFactors(n))
}

func TestPrime(t *testing.T) {
	//   3309997
	n := 3310000
	for {
		if IsPrime(n) {
			fmt.Printf("%d is prime\n", n)
			break
		}
		n--
	}
}
