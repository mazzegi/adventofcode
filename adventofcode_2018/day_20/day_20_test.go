package day_20

import (
	"adventofcode_2018/testutil"
	"testing"
)

const inputTest0 = "^ENWWW(NEEE|SSE(EE|N))$"
const inputTest1 = "^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"
const inputTest2 = "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"

func TestPart1MainFunc(t *testing.T) {
	var res, exp int
	var err error

	res, err = part1MainFunc(inputTest0)
	testutil.CheckUnexpectedError(t, err)
	exp = 31
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}

	// res, err = part1MainFunc(inputTest1)
	// testutil.CheckUnexpectedError(t, err)
	// exp = 23
	// if exp != res {
	// 	t.Fatalf("want %d, have %d", exp, res)
	// }

	// res, err = part1MainFunc(inputTest2)
	// testutil.CheckUnexpectedError(t, err)
	// exp = 31
	// if exp != res {
	// 	t.Fatalf("want %d, have %d", exp, res)
	// }
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest1)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
