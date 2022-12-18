package maps

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
