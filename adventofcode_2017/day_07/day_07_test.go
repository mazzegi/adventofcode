package day_07

import (
	"adventofcode_2017/testutil"
	"testing"
)

/*
                gyxo
              /
         ugml - ebii
       /      \
      |         jptl
      |
      |         pbga
     /        /
tknk --- padx - havc
     \        \
      |         qoyq
      |
      |         ktlj
       \      /
         fwft - cntj
              \
                xhth
*/

const inputTest = `
pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
`

func TestNameOfBottomProgram(t *testing.T) {
	res, err := nameOfBottomProgram(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp string = "tknk"
	if exp != res {
		t.Fatalf("want %s, have %s", exp, res)
	}
}

func TestCorrectionWeight(t *testing.T) {
	res, err := correctionWeight(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 60
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
