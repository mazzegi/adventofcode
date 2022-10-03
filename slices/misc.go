package slices

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
