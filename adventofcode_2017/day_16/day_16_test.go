package day_16

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = "s1,x3/4,pe/b"

func TestPart1MainFunc(t *testing.T) {
	res, err := danceOrder(inputTest, 5)
	testutil.CheckUnexpectedError(t, err)
	var exp string = "baedc"
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	loops := 10000
	res1, err := after(input, 16, loops, false)
	testutil.CheckUnexpectedError(t, err)
	log("after %d - without caching: %q:", loops, res1)

	res2, err := after(input, 16, loops, true)
	testutil.CheckUnexpectedError(t, err)
	log("after %d - with caching: %q:", loops, res2)
}
