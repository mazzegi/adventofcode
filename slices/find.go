package slices

func Find[T comparable](ts []T, t T) int {
	for i, et := range ts {
		if et == t {
			return i
		}
	}
	return -1
}

func Contains[T comparable](ts []T, t T) bool {
	return Find(ts, t) > -1
}
