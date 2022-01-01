package main

import (
	"adventofcode_2015/comb"
	"fmt"
	"sort"
)

var inputTest = []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}

var input = []int{1, 3, 5, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}

func main() {
	in := input
	sort.Ints(in)
	totalWeight := Group(in).Weight()
	if totalWeight%4 != 0 {
		panic("not solvable")
	}
	groupWeight := totalWeight / 4
	fmt.Printf("total-weight: %d, each group weight: %d\n", totalWeight, groupWeight)

	var minGroupSize int
	var sum int
	for i := len(in) - 1; i >= 0; i-- {
		sum += in[i]
		minGroupSize++
		if sum >= groupWeight {
			break
		}
	}
	maxGroupSize := len(in) - 3*minGroupSize
	fmt.Printf("min-group-size: %d\n", minGroupSize)
	fmt.Printf("max-group-size: %d\n", maxGroupSize)

	for gs := minGroupSize; gs <= len(in); gs++ {
		minQE := -1
		for part := range comb.Parts(in, gs) {
			if Group(part).Weight() != groupWeight {
				continue
			}

			rest := Group(in).Sub(part)
			g1, g2, g3, ok := rest.FindArrangementOf3(groupWeight, minGroupSize, maxGroupSize)
			if !ok {
				continue
			}
			qe := Group(part).QE()
			fmt.Printf("found arrangement: center %v (QE=%d), g1 %v, g2 %v, g3 %v\n", part, qe, g1, g2, g3)
			if minQE < 0 || qe < minQE {
				minQE = qe
			}
		}
		if minQE > -1 {
			fmt.Printf("min-qe: %d\n", minQE)
			break
		}
	}
}

type Group []int

func (g Group) Contains(v int) bool {
	for _, n := range g {
		if n == v {
			return true
		}
	}
	return false
}

func (g Group) Sub(og Group) Group {
	var sg Group
	for _, n := range g {
		if og.Contains(n) {
			continue
		}
		sg = append(sg, n)
	}
	return sg
}

func (g Group) Weight() int {
	var w int
	for _, v := range g {
		w += v
	}
	return w
}

func (g Group) QE() int {
	qe := 1
	for _, v := range g {
		qe *= v
	}
	return qe
}

func (g Group) Size() int {
	return len(g)
}

func (g Group) FindArrangementOf2(groupWeight int, minSize int, maxSize int) (g1, g2 Group, found bool) {
	for gs := minSize; gs <= maxSize; gs++ {
		for part := range comb.Parts(g, gs) {
			if Group(part).Weight() != groupWeight {
				continue
			}

			rest := g.Sub(part)
			if rest.Weight() == groupWeight {
				return part, rest, true
			}
		}
	}
	return nil, nil, false
}

func (g Group) FindArrangementOf3(groupWeight int, minSize int, maxSize int) (g1, g2, g3 Group, found bool) {
	for gs := minSize; gs <= maxSize; gs++ {
		for part := range comb.Parts(g, gs) {
			if Group(part).Weight() != groupWeight {
				continue
			}

			rest := g.Sub(part)
			g2, g3, ok := rest.FindArrangementOf2(groupWeight, minSize, maxSize)
			if !ok {
				continue
			}
			return part, g2, g3, true
		}
	}
	return nil, nil, nil, false
}
