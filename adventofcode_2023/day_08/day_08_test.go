package day_08

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest1Seq = "RL"
const inputTest1Nodes = `
AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

const inputTest2Seq = "LLR"
const inputTest2Nodes = `
AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

const inputTest3Seq = "LR"
const inputTest3Nodes = `
11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

func TestPart1MainFunc(t *testing.T) {
	{
		res, err := part1MainFunc(inputTest1Seq, inputTest1Nodes)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 2
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	}
	{
		res, err := part1MainFunc(inputTest2Seq, inputTest2Nodes)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 6
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest3Seq, inputTest3Nodes)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 6
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
