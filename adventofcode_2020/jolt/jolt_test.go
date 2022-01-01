package jolt

import "testing"

func TestCombinations(t *testing.T) {
	min := 1
	max := 11
	cnt := 2
	bl := map[int]bool{
		1:  true,
		2:  true,
		5:  true,
		6:  true,
		8:  true,
		9:  true,
		10: true,
		11: true,
	}
	//1 2 5 6 8 9 10 11
	res := combinations(cnt, min, max, bl)
	for comb := range res {
		t.Logf("comb: %v", comb)
	}
	t.Logf("comb: done")
}

func TestUniqueCombinations(t *testing.T) {

	vals := []int{3, 4, 5, 6, 7}

	res := uniqueCombis(3, vals)
	for comb := range res {
		t.Logf("comb: %v", comb)
	}
	t.Logf("comb: done")
}

func TestCrux(t *testing.T) {
	size := 15
	num := 4
	// size := 43
	// num := 11
	// notValid := cruxOne(1, num, size)
	// t.Logf("=> %d", notValid)

	notValid := crux(num, size)
	t.Logf("=> %d", notValid)

	//sub3 := sumBinCoeff(size-3, 0)
	tot := sumBinCoeff(size, 0)
	rem := tot - notValid
	t.Logf("rem => %d", rem)
}
