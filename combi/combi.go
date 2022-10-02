package combi

import "github.com/mazzegi/adventofcode/slices"

func Permutations(ns []int) [][]int {
	var perms [][]int
	if len(ns) == 0 {
		return perms
	}
	if len(ns) == 1 {
		perms = append(perms, slices.Clone(ns))
		return perms
	}
	s0 := ns[0]
	for _, sub := range Permutations(ns[1:]) {
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
