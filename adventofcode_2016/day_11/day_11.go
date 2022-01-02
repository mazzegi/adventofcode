package day_11

//a polonium generator,
//a thulium generator,
//a promethium generator,
//a ruthenium generator,
//a cobalt generator,

//a thulium-compatible microchip,
//a ruthenium-compatible microchip,
//and a cobalt-compatible microchip

import (
	"adventofcode_2016/errutil"
	"fmt"
	"math"
	"strings"
)

func initTestState() *state {
	s := &state{
		elevatorFloor: 0,
		floors:        4,
		chips:         chips{{"hydrogen", 0}, {"lithium", 0}},
		generators:    generators{{"hydrogen", 1}, {"lithium", 2}},
	}
	return s
}

func initState() *state {
	s := &state{
		elevatorFloor: 0,
		floors:        4,
		chips:         chips{{"thulium", 0}, {"ruthenium", 0}, {"cobalt", 0}, {"polonium", 1}, {"promethium", 1}},
		generators:    generators{{"polonium", 0}, {"thulium", 0}, {"promethium", 0}, {"ruthenium", 0}, {"cobalt", 0}},
	}
	return s
}

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

const skip1 = false
const skip2 = true

func Part1() {
	if skip1 {
		return
	}
	res, err := minSteps(initState())
	//res, err := minSteps(initTestState())
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	if skip2 {
		return
	}
	s := initState()
	s.chips = append(s.chips, chip{"elerium", 0}, chip{"dilithium", 0})
	s.generators = append(s.generators, generator{"elerium", 0}, generator{"dilithium", 0})

	res, err := minSteps(s)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
type chip struct {
	elt   string
	floor int
}

func (c chip) String() string {
	return fmt.Sprintf("chip %q", c.elt)
}

func (c chip) hash() string {
	return fmt.Sprintf("chip:%s:%d", c.elt, c.floor)
}

type chips []chip

func (sl chips) clone() chips {
	csl := make(chips, len(sl))
	copy(csl, sl)
	return csl
}

func (sl chips) has(elt string) bool {
	for _, e := range sl {
		if e.elt == elt {
			return true
		}
	}
	return false
}
func (cs chips) sub(subcs ...chip) chips {
	mustSub := func(c chip) bool {
		for _, sc := range subcs {
			if sc.elt == c.elt {
				return true
			}
		}
		return false
	}

	var scs chips
	for _, c := range cs {
		if mustSub(c) {
			continue
		}
		scs = append(scs, c)
	}
	return scs
}

//

type generator struct {
	elt   string
	floor int
}

func (g generator) String() string {
	return fmt.Sprintf("gen %q", g.elt)
}

func (g generator) hash() string {
	return fmt.Sprintf("gen:%s:%d", g.elt, g.floor)
}

type generators []generator

func (sl generators) clone() generators {
	csl := make(generators, len(sl))
	copy(csl, sl)
	return csl
}

func (sl generators) has(elt string) bool {
	for _, e := range sl {
		if e.elt == elt {
			return true
		}
	}
	return false
}

func (gs generators) sub(subgs ...generator) generators {
	mustSub := func(g generator) bool {
		for _, sg := range subgs {
			if sg.elt == g.elt {
				return true
			}
		}
		return false
	}

	var sgs generators
	for _, g := range gs {
		if mustSub(g) {
			continue
		}
		sgs = append(sgs, g)
	}
	return sgs
}

type state struct {
	chips         chips
	generators    generators
	elevatorFloor int
	floors        int

	visited   bool
	pathValue int
	prev      *state
}

func (s *state) clone() *state {
	return &state{
		chips:         s.chips.clone(),
		generators:    s.generators.clone(),
		elevatorFloor: s.elevatorFloor,
		floors:        s.floors,
	}
}

func (s *state) hash() string {

	// sort.SliceStable(s.chips, func(i, j int) bool {
	// 	return s.chips[i].elt < s.chips[j].elt
	// })
	// sort.SliceStable(s.generators, func(i, j int) bool {
	// 	return s.generators[i].elt < s.generators[j].elt
	// })
	var sl []string
	for _, c := range s.chips {
		sl = append(sl, c.hash())
	}
	for _, g := range s.generators {
		sl = append(sl, g.hash())
	}

	return fmt.Sprintf("%d-%s", s.elevatorFloor, strings.Join(sl, "-"))
}

//
type pair struct {
	item1 interface{}
	item2 interface{}
}

type move struct {
	pair    pair
	toFloor int
}

func (mv move) String() string {
	var s1 string
	var s2 string
	if mv.pair.item1 != nil {
		s1 = fmt.Sprintf("%s", mv.pair.item1)
	} else {
		s1 = "nil"
	}
	if mv.pair.item2 != nil {
		s2 = fmt.Sprintf("%s", mv.pair.item2)
	} else {
		s2 = "nil"
	}

	return fmt.Sprintf("%s, %s => %d", s1, s2, mv.toFloor)
}

func makeMove(item1, item2 interface{}) move {
	return move{pair: pair{item1, item2}}
}

func (mv move) forFloor(floor int) move {
	return move{pair: mv.pair, toFloor: floor}
}

func (s *state) execMove(mv move) {
	moveItem := func(item interface{}, toFloor int) {
		if item == nil {
			return
		}
		switch item := item.(type) {
		case chip:
			for i, e := range s.chips {
				if e.elt == item.elt {
					s.chips[i].floor = toFloor
					return
				}
			}

		case generator:
			for i, e := range s.generators {
				if e.elt == item.elt {
					s.generators[i].floor = toFloor
					return
				}
			}
		default:
			fatal("invalid item %T", item)
		}
	}

	moveItem(mv.pair.item1, mv.toFloor)
	moveItem(mv.pair.item2, mv.toFloor)
	s.elevatorFloor = mv.toFloor
}

func (s *state) complete() bool {
	lastFloor := s.floors - 1
	for _, c := range s.chips {
		if c.floor != lastFloor {
			return false
		}
	}
	for _, g := range s.generators {
		if g.floor != lastFloor {
			return false
		}
	}
	return true
}

func (s *state) in(floor int) (chips, generators) {
	var cs chips
	var gs generators
	for _, c := range s.chips {
		if c.floor == floor {
			cs = append(cs, c)
		}
	}
	for _, g := range s.generators {
		if g.floor == floor {
			gs = append(gs, g)
		}
	}
	return cs, gs
}

func allowed(cs chips, gs generators) bool {
	if len(gs) == 0 {
		// no generators - no danger
		return true
	}
	if len(cs) == 0 {
		// no chips - no danger
		return true
	}

	for _, c := range cs {
		if !gs.has(c.elt) {
			// chip is not protected - but there is at least one generator (not matching) - forbidden.
			return false
		}
	}

	return true
}

func (s *state) pairAllowedAt(p pair, floor int) bool {
	fcs, fgs := s.in(floor)

	addItem := func(item interface{}) {
		if item == nil {
			return
		}

		switch item := item.(type) {
		case chip:
			fcs = append(fcs, item)
		case generator:
			fgs = append(fgs, item)
		default:
			fatal("invalid item %T", item)
		}
	}
	addItem(p.item1)
	addItem(p.item2)

	return allowed(fcs, fgs)
}

func (s *state) allowedMoves() []move {
	var pickMvs []move

	fcs, fgs := s.in(s.elevatorFloor)

	// take from current floor

	// one or two chips can always be removed
	for i1, c1 := range fcs {
		pickMvs = append(pickMvs, makeMove(c1, nil))
		for i2, c2 := range fcs {
			if i2 <= i1 {
				continue
			}
			pickMvs = append(pickMvs, makeMove(c1, c2))
		}
	}

	// check single and paired generators
	for i1, g1 := range fgs {
		if allowed(fcs, fgs.sub(g1)) {
			pickMvs = append(pickMvs, makeMove(g1, nil))
		}
		for i2, g2 := range fgs {
			if i2 <= i1 {
				continue
			}
			if allowed(fcs, fgs.sub(g1, g2)) {
				pickMvs = append(pickMvs, makeMove(g1, g2))
			}
		}
	}

	//chip and generator pairs
	for _, c := range fcs {
		for _, g := range fgs {
			// can only pair corresponding types
			if c.elt != g.elt {
				continue
			}

			if !allowed(fcs.sub(c), fgs.sub(g)) {
				continue
			}
			pickMvs = append(pickMvs, makeMove(c, g))
		}
	}

	// now test to which floor they can be moved
	canMoveTo := func(mv move, destFloor int) bool {
		if destFloor == s.elevatorFloor {
			return false
		}
		var inc int
		var start int
		var end int
		if destFloor > s.elevatorFloor {
			inc = 1
			start = s.elevatorFloor + 1
			end = destFloor + 1
		} else {
			inc = -1
			start = s.elevatorFloor - 1
			end = destFloor - 1
		}
		for f := start; f != end; f += inc {
			if !s.pairAllowedAt(mv.pair, f) {
				return false
			}
		}

		return true
	}

	var availFloors []int
	if s.elevatorFloor > 0 {
		availFloors = append(availFloors, s.elevatorFloor-1)
	}
	if s.elevatorFloor < s.floors-1 {
		availFloors = append(availFloors, s.elevatorFloor+1)
	}

	var mvs []move
	for _, mv := range pickMvs {
		//for f := 0; f < s.floors; f++ {
		for _, f := range availFloors {
			if !canMoveTo(mv, f) {
				continue
			}
			mvs = append(mvs, mv.forFloor(f))
		}
	}

	return mvs
}

func dumpMoves(mvs []move) {
	for _, mv := range mvs {
		log("%s", mv.String())
	}
}

func minSteps(st *state) (int, error) {
	// visited := map[string]int{}
	// steps, mvs, ok := walk(0, st, visited)
	// if !ok {
	// 	fatal("not solved")
	// }
	// dumpMoves(mvs)
	// return steps, nil
	return minStepsDijkstra(st), nil
}

func cloneVisited(vs map[string]int) map[string]int {
	cvs := map[string]int{}
	for h, l := range vs {
		cvs[h] = l
	}
	return cvs
}

func walk(level int, st *state, visited map[string]int) (int, []move, bool) {

	var minSteps int
	solved := false
	var minMoves []move

	mvs := st.allowedMoves()
	for _, mv := range mvs {
		cst := st.clone()
		cst.execMove(mv)

		hash := cst.hash()
		if vlevel, ok := visited[hash]; ok {
			if level >= vlevel {
				//log("%d: try move %s => already visited", level, mv.String())
				continue
			}
			visited[hash] = level
		}

		// cvisited := cloneVisited(visited)
		// cvisited[hash] = true
		// subSteps, subOk := walk(level+1, cst, cvisited)

		//log("%d: try move %s", level, mv.String())
		visited[hash] = level
		if cst.complete() {
			log("%d:%s ==> complete", level, hash)
			return 1, []move{mv}, true
		}
		//log("%d:%s", level, hash)

		subSteps, subMoves, subOk := walk(level+1, cst, visited)

		if !subOk {
			continue
		}
		if !solved {
			minSteps = subSteps + 1
			solved = true
			minMoves = append([]move{mv}, subMoves...)
		} else if subSteps+1 < minSteps {
			minSteps = subSteps + 1
			minMoves = append([]move{mv}, subMoves...)
		}
	}

	return minSteps, minMoves, solved
}

// func part2MainFunc(in string) (int, error) {
// 	return 0, nil
// }

type orderedListElt struct {
	state *state
	next  *orderedListElt
}

type orderedList struct {
	first *orderedListElt
	last  *orderedListElt
	sz    int
}

func (l *orderedList) size() int {
	return l.sz
}

func (l *orderedList) add(s *state) {
	l.sz++
	elt := &orderedListElt{s, nil}
	if l.first == nil {
		l.first = elt
		l.last = l.first
		return
	}

	if l.first == l.last {
		if elt.state.pathValue <= l.first.state.pathValue {
			elt.next = l.first
			l.first = elt
		} else {
			l.first.next = elt
			l.last = elt
		}
		return
	}

	if elt.state.pathValue <= l.first.state.pathValue {
		elt.next = l.first
		l.first = elt
		return
	}
	if elt.state.pathValue >= l.last.state.pathValue {
		l.last.next = elt
		l.last = elt
		return
	}

	prev := l.first
	curr := l.first.next
	for curr != nil {
		if elt.state.pathValue <= curr.state.pathValue {
			prev.next = elt
			elt.next = curr
			return
		}

		prev = curr
		curr = curr.next
	}
	fatal("should not come here")
}

func (l *orderedList) mustTakeFirst() *state {
	if l.first == nil {
		fatal("list is empty")
	}
	l.sz--
	elt := l.first
	if l.first == l.last {
		l.first = nil
		l.last = nil
	} else {
		l.first = elt.next
	}
	return elt.state
}

func (l *orderedList) dump() string {
	var sl []string
	curr := l.first
	for curr != nil {
		sl = append(sl, fmt.Sprintf("%d", curr.state.pathValue))
		curr = curr.next
	}
	return strings.Join(sl, ", ")
}

func minStepsDijkstra(st *state) int {

	//nodes := map[string]*state{}
	nodes := map[string]bool{}
	notVisited := map[string]*state{}

	notVisitedNodeWithMinDist := func() *state {
		var cand *state
		var candDist int
		for _, n := range notVisited {
			// if n.visited {
			// 	continue
			// }
			if cand == nil {
				cand = n
				candDist = n.pathValue
				continue
			}
			if n.pathValue < candDist {
				cand = n
				candDist = n.pathValue
			}
		}
		return cand
	}

	visit := func(s *state) {
		hash := s.hash()
		//nodes[hash] = s
		nodes[hash] = true
		delete(notVisited, hash)
	}

	start := st.clone()
	start.pathValue = 0
	notVisited[start.hash()] = start
	for {
		n := notVisitedNodeWithMinDist()
		//n.visited = true
		visit(n)

		mvs := n.allowedMoves()
		for _, mv := range mvs {
			cs := n.clone()
			cs.execMove(mv)
			if cs.complete() {
				return n.pathValue + 1
			}

			hash := cs.hash()
			if _, ok := nodes[hash]; ok {
				// already visited
				continue
			}

			cs.pathValue = math.MaxInt64
			notVisited[hash] = cs

			test := n.pathValue + 1
			if test < cs.pathValue {
				cs.pathValue = test
				cs.prev = n
			}
		}
		lenNs := len(nodes)
		lenNotVs := len(notVisited)
		if lenNs%1000 == 0 {
			log("nodes = %d, not_visited = %d", lenNs, lenNotVs)
		}
	}
}
