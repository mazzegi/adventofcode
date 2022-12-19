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
