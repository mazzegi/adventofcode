package day_11

import (
	"math/big"
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

var inputTest = []*Monkey{
	{
		ID:             0,
		Items:          []int{79, 98},
		Operation:      func(n int) int { return n * 19 },
		Test:           func(n int) bool { return n%23 == 0 },
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 3,
		Activity:       0,
	},
	{
		ID:             1,
		Items:          []int{54, 65, 75, 74},
		Operation:      func(n int) int { return n + 6 },
		Test:           func(n int) bool { return n%19 == 0 },
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 0,
		Activity:       0,
	},
	{
		ID:             2,
		Items:          []int{79, 60, 97},
		Operation:      func(n int) int { return n * n },
		Test:           func(n int) bool { return n%13 == 0 },
		ThrowToIfTrue:  1,
		ThrowToIfFalse: 3,
		Activity:       0,
	},
	{
		ID:             3,
		Items:          []int{74},
		Operation:      func(n int) int { return n + 3 },
		Test:           func(n int) bool { return n%17 == 0 },
		ThrowToIfTrue:  0,
		ThrowToIfFalse: 1,
		Activity:       0,
	},
}

var inputTest2 = []*BigMonkey{
	{
		ID:             0,
		Items:          []*big.Int{big.NewInt(79), big.NewInt(98)},
		Operation:      func(n *big.Int) *big.Int { return big.NewInt(0).Mul(n, big.NewInt(19)) },
		Test:           func(n *big.Int) bool { return big.NewInt(0).Mod(n, big.NewInt(23)).Int64() == 0 },
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 3,
		Activity:       0,
	},
	{
		ID:             1,
		Items:          []*big.Int{big.NewInt(54), big.NewInt(65), big.NewInt(75), big.NewInt(74)},
		Operation:      func(n *big.Int) *big.Int { return big.NewInt(0).Add(n, big.NewInt(6)) },
		Test:           func(n *big.Int) bool { return big.NewInt(0).Mod(n, big.NewInt(19)).Int64() == 0 },
		ThrowToIfTrue:  2,
		ThrowToIfFalse: 0,
		Activity:       0,
	},
	{
		ID:             2,
		Items:          []*big.Int{big.NewInt(79), big.NewInt(60), big.NewInt(97)},
		Operation:      func(n *big.Int) *big.Int { return big.NewInt(0).Mul(n, n) },
		Test:           func(n *big.Int) bool { return big.NewInt(0).Mod(n, big.NewInt(13)).Int64() == 0 },
		ThrowToIfTrue:  1,
		ThrowToIfFalse: 3,
		Activity:       0,
	},
	{
		ID:             3,
		Items:          []*big.Int{big.NewInt(74)},
		Operation:      func(n *big.Int) *big.Int { return big.NewInt(0).Add(n, big.NewInt(3)) },
		Test:           func(n *big.Int) bool { return big.NewInt(0).Mod(n, big.NewInt(17)).Int64() == 0 },
		ThrowToIfTrue:  0,
		ThrowToIfFalse: 1,
		Activity:       0,
	},
}

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 10605
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest2)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 2713310158
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
