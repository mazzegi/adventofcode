package slices

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](ts []T) T {
	var sum T
	for _, t := range ts {
		sum += t
	}
	return sum
}

func Min[T Number](ts []T) T {
	var max T
	for i, t := range ts {
		if i == 0 || t < max {
			max = t
		}
	}
	return max
}
