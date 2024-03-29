package day_12

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2016/testutil"
)

const inputTest = `
cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a
`

func TestRegAValue(t *testing.T) {
	n, err := RegAValue(inputTest, map[string]int{})
	testutil.CheckUnexpectedError(t, err)

	exp := 42
	if exp != n {
		t.Fatalf("want %d, have %d", exp, n)
	}
}
