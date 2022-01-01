package comb

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFac(t *testing.T) {
	for n := int64(0); n < 20; n++ {
		r := Fac(big.NewInt(n))
		exp := big.NewInt(int64(fac(int(n))))
		if exp.Cmp(r) != 0 {
			t.Fatalf("%d: want %s have %s", n, exp.String(), r.String())
		}
	}
}

func TestBinCoeff(t *testing.T) {
	for n := int64(0); n < 20; n++ {
		for k := int64(0); k <= n; k++ {
			r := BinCoeff(big.NewInt(n), big.NewInt(k))
			exp := big.NewInt(int64(binCoeff(int(n), int(k))))
			if exp.Cmp(r) != 0 {
				t.Fatalf("(%d, %d): want %s have %s", n, k, exp.String(), r.String())
			}
			//fmt.Printf("(%d, %d) => %s (%s)\n", n, k, r.String(), exp.String())
		}
	}
}

func TestSumBinCoeff(t *testing.T) {
	for n := int64(0); n < 20; n++ {
		r := SumBinCoeff(big.NewInt(n), Zero())
		exp := big.NewInt(int64(sumBinCoeff(int(n), 0)))
		if exp.Cmp(r) != 0 {
			t.Fatalf("%d: want %s have %s", n, exp.String(), r.String())
		}
		//fmt.Printf("(%d) => %s (%s)\n", n, r.String(), exp.String())
	}
}

func TestCrux3(t *testing.T) {
	// size := int64(15)
	// num := int64(4)
	size := int64(43)
	num := int64(11)

	bsize := big.NewInt(size)
	bnum := big.NewInt(num)

	notValid := Crux3(bnum, bsize)

	//exp := crux(int(num), int(size))
	fmt.Printf("(%d, %d) => %s", num, size, notValid.String())

	tot := SumBinCoeff(bsize, Zero())
	rem := tot.Sub(tot, notValid)
	t.Logf("(tot: %s, notvalid: %s) rem => %s", tot.String(), notValid.String(), rem.String())

	// for n := int64(0); n < 20; n++ {
	// 	r := SumBinCoeff(big.NewInt(n), Zero())
	// 	exp := big.NewInt(int64(sumBinCoeff(int(n), 0)))
	// 	if exp.Cmp(r) != 0 {
	// 		t.Fatalf("%d: want %s have %s", n, exp.String(), r.String())
	// 	}
	// 	//fmt.Printf("(%d) => %s (%s)\n", n, r.String(), exp.String())
	// }
}
