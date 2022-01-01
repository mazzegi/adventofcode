package comb

import (
	"fmt"
	"math/big"
)

func fac(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}

func binCoeff(n, k int) int {
	f := float64(fac(n)) / (float64(fac(k)) * float64(fac(n-k)))
	return int(f)
}

func sumBinCoeff(n int, startK int) int {
	sum := 0
	for k := startK; k <= n; k++ {
		bc := binCoeff(n, k)
		sum += bc
	}
	return sum
}

func cruxOne(n, num int, size int) int {
	fac := binCoeff(num, n)
	sign := 1
	sum := 0
	bs := (num - n + 1) * 3
	for k := 0; k <= num-n; k++ {
		c := binCoeff(num-n, k)
		p := sumBinCoeff(bs, 0)
		sum += sign * c * p
		sign *= -1
		bs -= 3
	}
	return fac * sum
}

func crux(num int, size int) int {
	fmt.Printf("crux: num=%d, size=%d\n", num, size)
	sum := 0
	for n := 1; n <= num; n++ {
		co := cruxOne(n, num, size)
		fmt.Printf("cx(%d) => %d\n", n, co)
		sum += co
	}
	return sum
}

//
func Crux3One(n, num, size *big.Int) *big.Int {
	fac := BinCoeff(num, n)
	sign := int64(1)
	sum := Zero()

	ndiff := big.NewInt(0).Sub(num, n)
	bs := big.NewInt(0).Add(ndiff, One())
	bs.Mul(bs, big.NewInt(3))

	for k := Zero(); k.Cmp(ndiff) <= 0; k.Add(k, One()) {
		c := BinCoeff(ndiff, k)
		c.Mul(c, big.NewInt(sign))

		p := SumBinCoeff(bs, Zero())

		sum.Add(sum, p.Mul(p, c))

		sign *= -1
		bs.Sub(bs, big.NewInt(3))
	}

	return sum.Mul(sum, fac)
}

func Crux3(num, size *big.Int) *big.Int {
	sum := Zero()
	for n := One(); n.Cmp(num) <= 0; n.Add(n, One()) {
		co := Crux3One(n, num, size)
		fmt.Printf("bcx(%s) => %s\n", n.String(), co.String())
		sum.Add(sum, co)
	}
	return sum
}
