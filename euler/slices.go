package euler

import "golang.org/x/exp/constraints"

func SliceOf[T constraints.Integer](min, max T) []T {
	var ts []T
	if min > max {
		min, max = max, min
	}
	for t := min; t <= max; t++ {
		ts = append(ts, t)
	}
	return ts
}

func SliceSum[T constraints.Integer | constraints.Float](ts []T) T {
	var sum T
	for _, t := range ts {
		sum += t
	}
	return sum
}

func SliceSumFunc[T constraints.Integer | constraints.Float](ts []T, transform func(t T) T) T {
	var sum T
	for _, t := range ts {
		sum += transform(t)
	}
	return sum
}

func SliceProduct[T constraints.Integer | constraints.Float](ts []T) T {
	if len(ts) == 0 {
		var zt T
		return zt
	}
	var prod T = 1
	for _, t := range ts {
		prod *= t
	}
	return prod
}

func SliceContainsOne[T comparable](ts []T, t T) bool {
	for _, et := range ts {
		if et == t {
			return true
		}
	}
	return false
}

func SliceContainsMany[T comparable](ts []T, ots []T) bool {
	for _, ot := range ots {
		if !SliceContainsOne(ts, ot) {
			return false
		}
	}
	return true
}

func SliceClone[T any](ts []T) []T {
	cts := make([]T, len(ts))
	copy(cts, ts)
	return cts
}

func SliceFindIdx[T comparable](ts []T, t T) int {
	for i, et := range ts {
		if et == t {
			return i
		}
	}
	return -1
}

func SliceRemoveIdx[T any](ts []T, ri int) []T {
	var rts []T
	for i, et := range ts {
		if i == ri {
			continue
		}
		rts = append(rts, et)
	}
	return rts
}

func SliceRemain[T comparable](from []T, in []T) []T {
	var rem []T
	fromC := SliceClone(from)
	inC := SliceClone(in)
	for _, t := range fromC {
		i := SliceFindIdx(inC, t)
		if i < 0 {
			rem = append(rem, t)
		} else {
			inC = SliceRemoveIdx(inC, i)
		}
	}
	return rem
}
