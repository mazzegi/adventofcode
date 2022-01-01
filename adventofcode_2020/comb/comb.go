package comb

import "math/big"

func Zero() *big.Int {
	return big.NewInt(0)
}

func One() *big.Int {
	return big.NewInt(1)
}

func Fac(n *big.Int) *big.Int {
	cmp := n.Cmp(One())
	if cmp <= 0 {
		return One()
	}
	return big.NewInt(0).Mul(n, Fac(big.NewInt(0).Sub(n, One())))
}

func BinCoeff(n, k *big.Int) *big.Int {
	return big.NewInt(0).Div(Fac(n), big.NewInt(0).Mul(Fac(k), Fac(big.NewInt(0).Sub(n, k))))
}

/*
func sumBinCoeff(n int, startK int) int {
	sum := 0
	for k := startK; k <= n; k++ {
		bc := binCoeff(n, k)
		sum += bc
	}
	return sum
}
*/

func SumBinCoeff(n, startK *big.Int) *big.Int {
	sum := Zero()
	k := big.NewInt(0).Add(Zero(), startK)
	for ; k.Cmp(n) <= 0; k.Add(k, One()) {
		sum.Add(sum, BinCoeff(n, k))
	}
	return sum
}
