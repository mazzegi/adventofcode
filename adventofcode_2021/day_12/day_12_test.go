package day_12

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

/*
    start
    /   \
c--A-----b--d
    \   /
     end
*/

const inputTest1 = `
start-A
start-b
A-b
A-c
b-d
A-end
b-end
`

const inputTest2 = `
dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`

const inputTest3 = `
fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`

func TestPathCount(t *testing.T) {
	res, err := pathCount(inputTest1)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 10
	if exp != res {
		t.Fatalf("test1: want %d, have %d", exp, res)
	}

	res, err = pathCount(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	exp = 19
	if exp != res {
		t.Fatalf("test2: want %d, have %d", exp, res)
	}

	res, err = pathCount(inputTest3)
	testutil.CheckUnexpectedError(t, err)
	exp = 226
	if exp != res {
		t.Fatalf("test3: want %d, have %d", exp, res)
	}
}

func TestPathCountV2(t *testing.T) {
	res, err := pathCountV2(inputTest1)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 36
	if exp != res {
		t.Fatalf("test1: want %d, have %d", exp, res)
	}

	res, err = pathCountV2(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	exp = 103
	if exp != res {
		t.Fatalf("test2: want %d, have %d", exp, res)
	}

	res, err = pathCountV2(inputTest3)
	testutil.CheckUnexpectedError(t, err)
	exp = 3509
	if exp != res {
		t.Fatalf("test3: want %d, have %d", exp, res)
	}
}
