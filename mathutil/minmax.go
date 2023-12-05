package mathutil

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](t1, t2 T) T {
	if t1 <= t2 {
		return t1
	}
	return t2
}

func Max[T constraints.Ordered](t1, t2 T) T {
	if t1 >= t2 {
		return t1
	}
	return t2
}

func MinOfSlice[T constraints.Ordered](ts []T) T {
	if len(ts) == 0 {
		panic("slice is empty")
	}
	min := ts[0]
	for _, t := range ts {
		if t < min {
			min = t
		}
	}
	return min
}

func MaxOfSlice[T constraints.Ordered](ts []T) T {
	if len(ts) == 0 {
		panic("slice is empty")
	}
	max := ts[0]
	for _, t := range ts {
		if t > max {
			max = t
		}
	}
	return max
}
