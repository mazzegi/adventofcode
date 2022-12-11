package bigx

import (
	"fmt"
	"math/big"
	"testing"
)

func TestIsDivisible(t *testing.T) {
	tests := []struct {
		x   string
		n   int
		exp bool
	}{
		{"480", 2, true},
		{"481", 2, false},
		{"4804787489230089272", 2, true},
		{"4804787489233568977", 2, false},
		{"543878957438954804787489230089272", 2, true},
		{"1221874893254554804787489233568971", 2, false},

		{"603", 3, true},
		{"226", 3, false},
		{"543878957438954804787489230089272", 3, true},
		{"543878957438954804787489230089273", 3, false},
		{"1221874893254554804787489233568971", 3, false},

		{"480", 5, true},
		{"481", 5, false},
		{"4804787489230089270", 5, true},
		{"4804787489233568977", 5, false},
		{"543878957438954804787489230089275", 5, true},
		{"1221874893254554804787489233568971", 5, false},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			bx := big.NewInt(0)
			bx.SetString(test.x, 10)
			res := IsDivisibleby(bx, test.n)
			if test.exp != res {
				t.Fatalf("want %t, have %t", test.exp, res)
			}
		})
	}
}

func BenchmarkIsDivisible(b *testing.B) {
	//{"1221874893254554804787489233568971", 5, false},
	bx := big.NewInt(0)
	bx.SetString("1221874893254554804787489233568971", 10)
	for i := 0; i < b.N; i++ {
		res := IsDivisibleby(bx, 5)
		if res != false {
			b.FailNow()
		}
	}
}

func BenchmarkIsDivisibleViaMod(b *testing.B) {
	//{"1221874893254554804787489233568971", 5, false},
	bx := big.NewInt(0)
	bx.SetString("1221874893254554804787489233568971", 10)
	for i := 0; i < b.N; i++ {
		res := big.NewInt(0).Mod(bx, big.NewInt(5)).Int64() == 0
		if res != false {
			b.FailNow()
		}
	}
}

func bigMulN(x *big.Int, n int) *big.Int {
	y := big.NewInt(0).Set(x)
	if n <= 0 {
		return y
	}
	for i := 0; i < n-1; i++ {
		y.Add(y, x)
	}
	return y
}

func BenchmarkMul(b *testing.B) {
	//{"1221874893254554804787489233568971", 5, false},
	bx := big.NewInt(0)
	bx.SetString("1221874893254554804787489233568971", 10)
	var n int = 19
	for i := 0; i < b.N; i++ {
		res := bigMulN(bx, n)
		_ = res
	}
}

func BenchmarkMulViaBigMul(b *testing.B) {
	//{"1221874893254554804787489233568971", 5, false},
	bx := big.NewInt(0)
	bx.SetString("1221874893254554804787489233568971", 10)
	var n int64 = 19
	for i := 0; i < b.N; i++ {
		res := big.NewInt(0).Mul(bx, big.NewInt(n))
		_ = res
	}
}
