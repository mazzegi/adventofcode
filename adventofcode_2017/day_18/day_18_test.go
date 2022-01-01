package day_18

import (
	"adventofcode_2017/testutil"
	"testing"
)

const inputTest = `
set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2
`

const inputTest2 = `
snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d
`

func TestPart1MainFunc(t *testing.T) {
	res, err := recFreq(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 4
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := sendValueCountOf1(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 3
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
