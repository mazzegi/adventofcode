package day_23

import (
	"adventofcode_2017/testutil"
	"math"
	"testing"
)

const inputTest = ""

func TestPart1MainFunc(t *testing.T) {
	res, err := numMulInvoked(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := fixedRegH(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

const eps = 0.0000001

func isMultiple(a int, of int) (int, bool) {
	m := float64(a) / float64(of)
	if math.Abs(m-math.Floor(m)) < eps {
		return int(math.Round(m)), true
	}
	return 0, false
}

func TestDo(t *testing.T) {
	h := 0
	for b := 109300; b <= 126300; b += 17 {
		log("b = %d", b)

		addH := false
		for d := 2; d < b; d++ {
			e, ok := isMultiple(b, d)
			if !ok {
				continue
			}
			if e >= 2 {
				addH = true
			}

			// for e := 2; e < b; e++ {
			// 	if e*d == b {
			// 		h--
			// 	}
			// }
		}
		if addH {
			h++
		}
	}
	log("h = %d", h)
}
