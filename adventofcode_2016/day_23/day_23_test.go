package day_23

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2016/testutil"
)

const inputTest = `
cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a
`

func TestRegAValue(t *testing.T) {
	n, err := RegAValue(inputTest, map[string]int{})
	testutil.CheckUnexpectedError(t, err)

	exp := 3
	if exp != n {
		t.Fatalf("want %d, have %d", exp, n)
	}
}

func TestCalc(t *testing.T) {
	res := calc(12)
	log("12 => %d", res)
}
