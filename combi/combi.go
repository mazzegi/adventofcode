package combi

import "github.com/mazzegi/adventofcode/slices"

func IntPermutations(ns []int) [][]int {
	var perms [][]int
	if len(ns) == 0 {
		return perms
	}
	if len(ns) == 1 {
		perms = append(perms, slices.Clone(ns))
		return perms
	}
	s0 := ns[0]
	for _, sub := range IntPermutations(ns[1:]) {
		var subPerms [][]int
		perm := append([]int{s0}, sub...)
		subPerms = append(subPerms, perm)
		for i := 0; i < len(sub); i++ {
			perm := append([]int{}, sub[:i+1]...)
			perm = append(perm, s0)
			perm = append(perm, sub[i+1:]...)
			subPerms = append(subPerms, perm)
		}
		perms = append(perms, subPerms...)
	}
	return perms
}

func Permutations[T any](ts []T) [][]T {
	var perms [][]T
	if len(ts) == 0 {
		return perms
	}
	if len(ts) == 1 {
		perms = append(perms, slices.Clone(ts))
		return perms
	}
	t0 := ts[0]
	for _, sub := range Permutations(ts[1:]) {
		var subPerms [][]T
		perm := append([]T{t0}, sub...)
		subPerms = append(subPerms, perm)
		for i := 0; i < len(sub); i++ {
			perm := append([]T{}, sub[:i+1]...)
			perm = append(perm, t0)
			perm = append(perm, sub[i+1:]...)
			subPerms = append(subPerms, perm)
		}
		perms = append(perms, subPerms...)
	}
	return perms
}
