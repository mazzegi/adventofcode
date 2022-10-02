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

//
type OrderedSet struct {
	values []string
}

func NewOrderedSet(elts ...string) *OrderedSet {
	return &OrderedSet{
		values: elts,
	}
}

func (set *OrderedSet) Insert(values ...string) {
	for _, v := range values {
		if set.Contains(v) {
			continue
		}
		set.values = append(set.values, v)
	}
}

func (set *OrderedSet) Contains(s string) bool {
	for _, es := range set.values {
		if s == es {
			return true
		}
	}
	return false
}

func (set *OrderedSet) Values() []string {
	return set.values
}
