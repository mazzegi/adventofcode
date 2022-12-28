package set

func New[T comparable](vs ...T) *Set[T] {
	s := &Set[T]{
		values: make(map[T]struct{}),
	}
	for _, v := range vs {
		s.Insert(v)
	}
	return s
}

type Set[T comparable] struct {
	values map[T]struct{}
}

func (s *Set[T]) Insert(ts ...T) {
	for _, t := range ts {
		s.values[t] = struct{}{}
	}
}

func (s *Set[T]) Remove(t T) {
	delete(s.values, t)
}

func (s *Set[T]) Contains(t T) bool {
	_, ok := s.values[t]
	return ok
}

func (s *Set[T]) Count() int {
	return len(s.values)
}

func (s *Set[T]) Clone() *Set[T] {
	cs := New[T]()
	for t := range s.values {
		cs.Insert(t)
	}
	return cs
}

func (s *Set[T]) Values() []T {
	var ts []T
	for t := range s.values {
		ts = append(ts, t)
	}
	return ts
}
