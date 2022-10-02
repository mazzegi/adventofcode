package slices

func Clone[T any](ts []T) []T {
	cts := make([]T, len(ts))
	copy(cts, ts)
	return cts
}
