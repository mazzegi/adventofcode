package slices

func Find[T comparable](ts []T, t T) int {
	for i, et := range ts {
		if et == t {
			return i
		}
	}
	return -1
}

func FindFunc[T any](ts []T, match func(mt T) bool) int {
	for i, et := range ts {
		if match(et) {
			return i
		}
	}
	return -1
}

func Contains[T comparable](ts []T, t T) bool {
	return Find(ts, t) > -1
}

func ContainsFunc[T any](ts []T, match func(mt T) bool) bool {
	return FindFunc(ts, match) > -1
}

func DedupFunc[T any](ts []T, eq func(t1, t2 T) bool) []T {
	var ddts []T
	for _, t := range ts {
		contains := ContainsFunc(ddts, func(mt T) bool {
			return eq(t, mt)
		})
		if !contains {
			ddts = append(ddts, t)
		}
	}
	return ddts
}

func Dedup[T comparable](ts []T) []T {
	return DedupFunc(ts, func(t1, t2 T) bool { return t1 == t2 })
}
