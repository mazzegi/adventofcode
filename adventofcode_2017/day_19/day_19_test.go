package day_19

import (
	"adventofcode_2017/testutil"
	"testing"
)

const inputTest = `
     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`

func TestPart1MainFunc(t *testing.T) {
	res, err := letters(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp string = "ABCDEF"
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
