package numbers

import (
	"math/big"
	"sort"
)

func IsPrime(n int) bool {
	if n <= 3 {
		return true
	}
	return big.NewInt(int64(n)).ProbablyPrime(n)
}

func PrimeFactors(n int) []int {
	if IsPrime(n) {
		return []int{n}
	}
	for p := 2; p < n; p++ {
		if IsPrime(p) && n%p == 0 {
			n /= p
			subps := PrimeFactors(n)
			return append([]int{p}, subps...)
		}
	}
	panic("found no further prime")
}

func LeastCommonMutiple(ns ...int) int {
	if len(ns) == 0 {
		return 0
	}
	if len(ns) == 1 {
		return ns[0]
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i] > ns[j]
	})
	return LeastCommonMutipleOfPair(ns[0], LeastCommonMutiple(ns[1:]...))
}

func histogram(ns []int) map[int]int {
	h := map[int]int{}
	for _, n := range ns {
		h[n]++
	}
	return h
}

func LeastCommonMutipleOfPair(n1, n2 int) int {
	h1 := histogram(PrimeFactors(n1))
	h2 := histogram(PrimeFactors(n2))

	for p, c2 := range h2 {
		if c1, ok := h1[p]; ok {
			if c2 > c1 {
				h1[p] = c2
			}
		} else {
			h1[p] = c2
		}
	}

	m := 1
	for p, c := range h1 {
		m *= powInt(p, c)
	}
	return m
}

func powInt(n, ex int) int {
	if ex == 0 {
		return 1
	}
	r := 1
	for i := 0; i < ex; i++ {
		r *= n
	}
	return r
}
