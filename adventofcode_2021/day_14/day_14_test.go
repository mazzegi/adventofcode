package day_14

import (
	"adventofcode_2021/testutil"
	"testing"
)

const inputTestTemplate = "NNCB"

const inputTestRules = `
CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

func TestDiffMostLeastCommon(t *testing.T) {
	res, err := diffMostLeastCommon(inputTestTemplate, inputTestRules, 10)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 1588
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestDiffMostLeastCommonSpliced1(t *testing.T) {
	res, err := diffMostLeastCommonSpliced(inputTestTemplate, inputTestRules, 1)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 0
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestDiffMostLeastCommonSpliced2(t *testing.T) {
	res, err := diffMostLeastCommonSpliced(inputTestTemplate, inputTestRules, 2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 5
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestDiffMostLeastCommonSpliced3(t *testing.T) {
	res, err := diffMostLeastCommonSpliced(inputTestTemplate, inputTestRules, 3)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 7
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestDiffMostLeastCommonSpliced10(t *testing.T) {
	res, err := diffMostLeastCommonSpliced(inputTestTemplate, inputTestRules, 10)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 1588
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestDiffMostLeastCommonSpliced40(t *testing.T) {
	res, err := diffMostLeastCommonSpliced(inputTestTemplate, inputTestRules, 40)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 2188189693529
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestDiffMostLeastCommonSplicedHashed10(t *testing.T) {
	res, err := diffMostLeastCommonSplicedHashed(inputTestTemplate, inputTestRules, 10)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 1588
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
