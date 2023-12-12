package slices

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Clone[T any](ts []T) []T {
	cts := make([]T, len(ts))
	copy(cts, ts)
	return cts
}

func Repeat[T any](t T, count int) []T {
	var sl []T
	for i := 0; i < count; i++ {
		sl = append(sl, t)
	}
	return sl
}

func DeleteIdx[T comparable](ts []T, i int) []T {
	var dts []T
	for ix, t := range ts {
		if ix == i {
			continue
		}
		dts = append(dts, t)
	}
	return dts
}

func DeleteFirst[T comparable](ts []T, t T) []T {
	for i, et := range ts {
		if et == t {
			return DeleteIdx(ts, i)
		}
	}
	return Clone(ts)
}

func Reverse[T any](ts []T) []T {
	sz := len(ts)
	rts := make([]T, sz)
	for i, t := range ts {
		rts[sz-i-1] = t
	}
	return rts
}

func Sort[T constraints.Ordered](ts []T) {
	sort.Slice(ts, func(i, j int) bool {
		return ts[i] < ts[j]
	})
}

func Insert[T any](ts []T, t T, atIdx int) []T {
	its := Clone(ts[:atIdx])
	its = append(its, t)
	its = append(its, ts[atIdx:]...)
	return its
}

func Equal[T comparable](ts1 []T, ts2 []T) bool {
	if len(ts1) != len(ts2) {
		return false
	}
	for i, t1 := range ts1 {
		if t1 != ts2[i] {
			return false
		}
	}
	return true
}
