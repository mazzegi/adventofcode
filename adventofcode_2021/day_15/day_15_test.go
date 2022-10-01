package day_15

import (
	"fmt"
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

const inputTest = `
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
`

func TestLowestTotalRisk(t *testing.T) {
	res, err := lowestTotalRisk(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 40
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestLowestTotalRiskDup5(t *testing.T) {
	res, err := lowestTotalRiskDup5(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 315
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestDup5(t *testing.T) {
	g, err := parseGraph(inputTest)
	testutil.CheckUnexpectedError(t, err)
	dg := dupPlus5(g)
	fmt.Printf("***\n%s\n***\n", dumpGraph(dg))
}
