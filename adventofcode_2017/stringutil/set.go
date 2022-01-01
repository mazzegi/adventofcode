package stringutil

type Set struct {
	values map[string]bool
}

func NewSet() *Set {
	return &Set{
		values: map[string]bool{},
	}
}

func (set *Set) Insert(s string) {
	set.values[s] = true
}

func (set *Set) Contains(s string) bool {
	_, ok := set.values[s]
	return ok
}
