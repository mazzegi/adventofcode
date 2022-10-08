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

func DeleteIdx[T comparable](ts []T, i int) []T {
	cts := Clone(ts)
	cts = append(cts[:i], cts[i+1:]...)
	return cts
}

func DeleteFirst[T comparable](ts []T, t T) []T {
	for i, et := range ts {
		if et == t {
			return DeleteIdx(ts, i)
		}
	}
	return Clone(ts)
}
