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
