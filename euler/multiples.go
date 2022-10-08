package euler

import "sort"

func SmallestMultipleOf(ns ...int) int {
	sort.Ints(ns)
	var smpfs []int
	for _, n := range ns {
		pfs := PrimeFactors(n)
		rem := SliceRemain(pfs, smpfs)
		smpfs = append(smpfs, rem...)
	}
	return SliceProduct(smpfs)
}
