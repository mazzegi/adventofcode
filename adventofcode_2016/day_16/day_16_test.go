package day_16

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2016/testutil"
)

func TestPart1MainFunc(t *testing.T) {
	res, err := checksum([]byte("10000"), 20)
	testutil.CheckUnexpectedError(t, err)
	var exp string = "01100"
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}
}
