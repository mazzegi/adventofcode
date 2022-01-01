package day_12

type visit struct {
	id    string
	count int
}

type Set struct {
	values []*visit
}

func NewSet(ids ...string) *Set {
	set := &Set{}
	set.Insert(ids...)
	return set
}

func (set *Set) Clone() *Set {
	cset := &Set{}
	for _, vi := range set.values {
		cset.values = append(cset.values, &visit{
			id:    vi.id,
			count: vi.count,
		})
	}
	return cset
}

func (set *Set) Append(oset *Set) {
	for _, ov := range oset.values {
		set.Visit(ov.id, ov.count)
	}
}

func (set *Set) Visit(id string, count int) {
	for _, es := range set.values {
		if es.id == id {
			es.count += count
			return
		}
	}
	set.values = append(set.values, &visit{
		id:    id,
		count: count,
	})
}

func (set *Set) Insert(ids ...string) {
	for _, id := range ids {
		if set.Contains(id) {
			continue
		}
		set.values = append(set.values, &visit{
			id:    id,
			count: 0,
		})
	}
}

func (set *Set) Contains(id string) bool {
	for _, es := range set.values {
		if es.id == id {
			return true
		}
	}
	return false
}

func (set *Set) VisitedCount(id string) int {
	for _, es := range set.values {
		if es.id == id {
			return es.count
		}
	}
	return 0
}

func (set *Set) Visits() []*visit {
	return set.values
}
