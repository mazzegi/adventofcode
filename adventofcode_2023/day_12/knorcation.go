package day_12

import "github.com/mazzegi/adventofcode/slices"

func IntKnorcations(oneCount int, size int) [][]int {
	var kns [][]int
	if oneCount == 0 {
		allZeros := slices.Repeat(0, size)
		kns = append(kns, allZeros)
		return kns
	}
	if oneCount == size {
		allOnes := slices.Repeat(1, oneCount)
		kns = append(kns, allOnes)
		return kns
	}
	if oneCount > size {
		panic("you shall not pass!")
	}

	rkns1 := IntKnorcations(oneCount-1, size-1)
	for _, rkn1 := range rkns1 {
		vs := append([]int{1}, rkn1...)
		kns = append(kns, vs)
	}
	rkns0 := IntKnorcations(oneCount, size-1)
	for _, rkn0 := range rkns0 {
		vs := append([]int{0}, rkn0...)
		kns = append(kns, vs)
	}

	return kns
}

func RuneKnorcations(oneCount int, size int, oneRune, zeroRune rune) [][]rune {
	var kns [][]rune
	if oneCount == 0 {
		allZeros := slices.Repeat(zeroRune, size)
		kns = append(kns, allZeros)
		return kns
	}
	if oneCount == size {
		allOnes := slices.Repeat(oneRune, oneCount)
		kns = append(kns, allOnes)
		return kns
	}
	if oneCount > size {
		panic("you shall not pass!")
	}

	rkns1 := RuneKnorcations(oneCount-1, size-1, oneRune, zeroRune)
	for _, rkn1 := range rkns1 {
		vs := append([]rune{oneRune}, rkn1...)
		kns = append(kns, vs)
	}
	rkns0 := RuneKnorcations(oneCount, size-1, oneRune, zeroRune)
	for _, rkn0 := range rkns0 {
		vs := append([]rune{zeroRune}, rkn0...)
		kns = append(kns, vs)
	}

	return kns
}
