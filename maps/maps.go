package maps

import (
	stdmaps "maps"
)

func Clone[K comparable, T any](m map[K]T) map[K]T {
	return stdmaps.Clone(m)
}

func Values[K comparable, T any](m map[K]T) []T {
	var ts []T
	for _, t := range m {
		ts = append(ts, t)
	}
	return ts
}

func Keys[K comparable, T any](m map[K]T) []K {
	var ks []K
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
