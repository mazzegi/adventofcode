package bigx

import (
	"fmt"
	"math/big"
)

func IsDivisibleby(x *big.Int, n int) bool {
	switch n {
	case 2:
		return isDivBy2(x)
	case 3:
		return isDivBy3(x)
	case 5:
		return isDivBy5(x)
	default:
		panic(fmt.Errorf("div by %d not implemented", n))
	}
}

func digits(x *big.Int) []int {
	bs := []byte(x.String())
	ds := make([]int, len(bs))
	for i, b := range bs {
		ds[i] = int(b) - 48
	}
	return ds
}

func lastDigit(x *big.Int) int {
	bs := []byte(x.String())
	return int(bs[len(bs)-1]) - 48
}

func isDivBy2(x *big.Int) bool {
	return lastDigit(x)%2 == 0
}

func isDivBy3(x *big.Int) bool {
	var sum int
	for _, d := range digits(x) {
		sum += d
	}
	return sum%3 == 0
}

func isDivBy5(x *big.Int) bool {
	return lastDigit(x)%5 == 0
}
