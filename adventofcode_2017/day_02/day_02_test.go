package day_02

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = `
5 1 9 5
7 5 3
2 4 6 8
`

const inputTestED = `
5 9 2 8
9 4 7 3
3 8 6 5
`

func TestChecksum(t *testing.T) {
	res, err := Checksum(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 18
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestChecksumED(t *testing.T) {
	res, err := ChecksumEvenlyDiv(inputTestED)
	testutil.CheckUnexpectedError(t, err)
	exp := 9
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
