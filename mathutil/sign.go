package mathutil

import "golang.org/x/exp/constraints"

func Sign[T constraints.Integer | constraints.Float](t T) T {
	var one int = 1
	var mone int = -1
	switch {
	case t < 0:
		return T(mone)
	case t > 0:
		return T(one)
	default:
		return 0
	}
}
