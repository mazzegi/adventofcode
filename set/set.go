package set

func New[T comparable]() *Set[T] {
	return &Set[T]{
		values: make(map[T]struct{}),
	}
}

type Set[T comparable] struct {
	values map[T]struct{}
}

func (s *Set[T]) Insert(t T) {
	s.values[t] = struct{}{}
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
