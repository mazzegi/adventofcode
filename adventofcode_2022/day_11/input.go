package day_11

import "math/big"

var input = []*Monkey{
	{
		ID:             0,
		Items:          []int{72, 97},
		Operation:      func(n int) int { return n * 13 },
		Test:           func(n int) bool { return n%19 == 0 },
		ThrowToIfTrue:  5,
		ThrowToIfFalse: 6,
		Activity:       0,
	},
	{
		ID:             1,
		Items:          []int{55, 70, 90, 74, 95},
		Operation:      func(n int) int { return n * n },
		Test:           func(n int) bool { return n%7 == 0 },
		ThrowToIfTrue:  5,
		ThrowToIfFalse: 0,
		Activity:       0,
	},
	{
		ID:             2,
		Items:          []int{74, 97, 66, 57},
		Operation:      func(n int) int { return n + 6 },
		Test:           func(n int) bool { return n%17 == 0 },
		ThrowToIfTrue:  1,
		ThrowToIfFalse: 0,
		Activity:       0,
	},
	{
		ID:             3,
		Items:          []int{86, 54, 53},
		Operation:      func(n int) int { return n + 2 },
		Test:           func(n int) bool { return n%13 == 0 },
		ThrowToIfTrue:  1,
		ThrowToIfFalse: 2,
		Activity:       0,
	},
	{
		ID:             4,
		Items:          []int{50, 65, 78, 50, 62, 99},
		Operation:      func(n int) int { return n + 3 },
		Test:           func(n int) bool { return n%11 == 0 },
		ThrowToIfTrue:  3,
		ThrowToIfFalse: 7,
		Activity:       0,
	},
	{
		ID:             5,
		Items:          []int{90},
		Operation:      func(n int) int { return n + 4 },
		Test:           func(n int) bool { return n%2 == 0 },
		ThrowToIfTrue:  4,
		ThrowToIfFalse: 6,
		Activity:       0,
	},
	{
		ID:             6,
		Items:          []int{88, 92, 63, 94, 96, 82, 53, 53},
		Operation:      func(n int) int { return n + 8 },
		Test:           func(n int) bool { return n%5 == 0 },
		ThrowToIfTrue:  4,
		ThrowToIfFalse: 7,
		Activity:       0,
	},
	{
		ID:             7,
		Items:          []int{70, 60, 71, 69, 77, 70, 98},
		Operation:      func(n int) int { return n * 7 },
		Test:           func(n int) bool { return n%3 == 0 },
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 3,
		Activity:       0,
	},
}

// func bigMulN(x *big.Int, n int) *big.Int {
// 	y := big.NewInt(0).Set(x)
// 	if n <= 0 {
// 		return y
// 	}
// 	for i := 0; i < n-1; i++ {
// 		y.Add(y, x)
// 	}
// 	return y
// }

// func bigSqrN(x *big.Int) *big.Int {
// 	return big.NewInt(0).Exp(x, big.NewInt(2), nil)
// }

// func bigClone(x *big.Int) *big.Int {
// 	cx := big.NewInt(0)
// 	cx.Set(x)
// 	return cx
// }

var inputBig = []*BigMonkey{
	{
		ID:             0,
		Items:          []*big.Int{big.NewInt(79), big.NewInt(98)},
		Operation:      func(n *big.Int) *big.Int { return big.NewInt(0).Mul(n, big.NewInt(19)) },
		Test:           func(n *big.Int) bool { return big.NewInt(0).Mod(n, big.NewInt(23)).Int64() == 0 },
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 3,
		Activity:       0,
		cache:          newCache(),
	},
	{
		ID:             1,
		Items:          []*big.Int{big.NewInt(54), big.NewInt(65), big.NewInt(75), big.NewInt(74)},
		Operation:      func(n *big.Int) *big.Int { return big.NewInt(0).Add(n, big.NewInt(6)) },
		Test:           func(n *big.Int) bool { return big.NewInt(0).Mod(n, big.NewInt(19)).Int64() == 0 },
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 0,
		Activity:       0,
		cache:          newCache(),
	},
	{
		ID:             2,
		Items:          []*big.Int{big.NewInt(79), big.NewInt(60), big.NewInt(97)},
		Operation:      func(n *big.Int) *big.Int { return big.NewInt(0).Mul(n, n) },
		Test:           func(n *big.Int) bool { return big.NewInt(0).Mod(n, big.NewInt(13)).Int64() == 0 },
		ThrowToIfTrue:  1,
		ThrowToIfFalse: 3,
		Activity:       0,
		cache:          newCache(),
	},
	{
		ID:             3,
		Items:          []*big.Int{big.NewInt(74)},
		Operation:      func(n *big.Int) *big.Int { return big.NewInt(0).Add(n, big.NewInt(3)) },
		Test:           func(n *big.Int) bool { return big.NewInt(0).Mod(n, big.NewInt(17)).Int64() == 0 },
		ThrowToIfTrue:  0,
		ThrowToIfFalse: 1,
		Activity:       0,
		cache:          newCache(),
	},
}
