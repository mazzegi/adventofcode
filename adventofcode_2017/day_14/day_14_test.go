package day_14

import (
	"adventofcode_2017/testutil"
	"encoding/hex"
	"testing"
)

const inputTest = "flqrgnkx"

func TestPart1MainFunc(t *testing.T) {
	res, err := usedSquares(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 8108
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := numRegions(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 1242
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestToBits(t *testing.T) {
	b := byte(1)
	bf := byteToBits(b)
	log("%d => %q", int(b), bf)

	b = byte(14)
	bf = byteToBits(b)
	log("%d => %q", int(b), bf)

	b = byte(15)
	bf = byteToBits(b)
	log("%d => %q", int(b), bf)

	//
	s := "a0c201"
	bs, _ := hex.DecodeString(s)
	bf = bytesToBits(bs)
	log("%q => %q", s, bf)
}

//101000001100001000000001
//101000001100001000000001
