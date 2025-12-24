package rat

import (
	"fmt"

	"github.com/mazzegi/adventofcode/euler"
)

var Zero = R(0, 1)

func R(num, den int) Number {
	return Number{Num: num, Den: den}
}

type Number struct {
	Num int // zähler
	Den int // nenner
}

func IsZero(n Number) bool {
	return n.Num == 0
}

func IsInf(n Number) bool {
	return n.Den == 0
}

func Format(n Number) string {
	n = Simplify(n)
	if n.Den != 1 {
		return fmt.Sprintf("%d/%d", n.Num, n.Den)
	}
	return fmt.Sprintf("%d", n.Num)
}

func Neg(n Number) Number {
	return R(-n.Num, n.Den)
}

func Inv(n Number) Number {
	return R(n.Den, n.Num)
}

func Add(n1, n2 Number) Number {
	return R(n1.Num*n2.Den+n2.Num*n1.Den, n1.Den*n2.Den)
}

func Sub(n1, n2 Number) Number {
	return Add(n1, Neg(n2))
}

func Mult(n Number, by Number) Number {
	return R(n.Num*by.Num, n.Den*by.Den)
}

func Div(n Number, by Number) Number {
	return Mult(n, Inv(by))
}

func SimplifySign(n Number) Number {
	switch {
	case n.Num < 0 && n.Den < 0:
		return R(-n.Num, -n.Den)

	case n.Num >= 0 && n.Den < 0:
		return R(-n.Num, -n.Den)

	default:
		return n
	}
}

func SimplifyZero(n Number) Number {
	if IsZero(n) {
		return R(0, 1)
	}
	return n
}

func Simplify(n Number) Number {
	if IsZero(n) {
		return R(0, 1)
	}
	n = SimplifySign(n)

	gcd := euler.GCD(n.Num, n.Den)
	if gcd != 1 {
		return R(n.Num/gcd, n.Den/gcd)
	}
	return n
}
