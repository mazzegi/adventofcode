package day_16

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/algo/dijkstra"
	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/maps"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
	"github.com/mazzegi/adventofcode/slices"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type Valve struct {
	ID        string
	FlowRate  int
	TunnelsTo []string
	Open      bool
}

func (v Valve) HasTunnelTo(id string) bool {
	return slices.Contains(v.TunnelsTo, id)
}

func mustParseValve(s string) *Valve {
	v := Valve{}
	_, err := fmt.Sscanf(s, "Valve %s has flow rate=%d;", &v.ID, &v.FlowRate)
	if err != nil {
		fatal(err.Error())
	}
	sl := strings.Split(s, " ")
	for i := 9; i < len(sl); i++ {
		id := strings.Trim(sl[i], ", ")
		v.TunnelsTo = append(v.TunnelsTo, id)
	}
	return &v
}

type Action interface{}

type Move struct {
	ToValve string
}

type Open struct{}

type Rest struct{}

type Path struct {
	actions []Action
	closed  *set.Set[string]
}

func (p *Path) String() string {
	var sl []string
	for _, a := range p.actions {
		switch a := a.(type) {
		case Move:
			sl = append(sl, fmt.Sprintf("M(%s)", a.ToValve))
		case Open:
			sl = append(sl, "O")
		}
	}
	cs := p.closed.Values()
	sort.Strings(cs)
	return fmt.Sprintf("C:%v - ", cs) + strings.Join(sl, ", ")
}

func (p *Path) Clone() *Path {
	return &Path{
		actions: slices.Clone(p.actions),
		closed:  p.closed.Clone(),
	}
}

type ValveSet struct {
	Valves map[string]*Valve
}

func (vs *ValveSet) Nodes() []string {
	return maps.Keys(vs.Valves)
}

func (vs *ValveSet) Equal(t1, t2 string) bool {
	return t1 == t2
}

func (vs *ValveSet) AreNeighbours(t1, t2 string) bool {
	return slices.Contains(vs.Valves[t1].TunnelsTo, t2)
}

func (vs *ValveSet) DistanceBetween(t1, t2 string) float64 {
	if t1 == t2 {
		return 0
	}
	if vs.AreNeighbours(t1, t2) {
		return 1
	}
	return math.Inf(1)
}

func (vs *ValveSet) ShortestPath(from string, to string) ([]string, int) {
	path, err := dijkstra.ShortestPath[string](vs, from, to)
	if err != nil {
		panic(err)
	}
	return path.Nodes, len(path.Nodes) - 1
}

//

func (vs *ValveSet) Valve(id string) *Valve {
	v, ok := vs.Valves[id]
	if !ok {
		fatal("no such valve %q", id)
	}
	return v
}

func (vs *ValveSet) Fitness(start string, path *Path) int {
	size := len(path.actions)
	curr := vs.Valve(start)
	var fn int
	for i, a := range path.actions {
		switch a := a.(type) {
		case Move:
			if !curr.HasTunnelTo(a.ToValve) {
				fatal("theres no tunnle to %q", a.ToValve)
			}
			curr = vs.Valve(a.ToValve)
		case Open:
			fn += (size - 1 - i) * curr.FlowRate
		case Rest:
		}
	}
	return fn
}

//

func mustParseValveSet(in string) *ValveSet {
	vs := &ValveSet{
		Valves: map[string]*Valve{},
	}
	for _, line := range readutil.ReadLines(in) {
		v := mustParseValve(line)
		vs.Valves[v.ID] = v
	}
	return vs
}

func part1MainFunc(in string) (int, error) {
	vs := mustParseValveSet(in)
	count := 30

	startPath := &Path{
		closed: set.New[string](),
	}
	for _, v := range vs.Valves {
		if v.FlowRate > 0 {
			startPath.closed.Insert(v.ID)
		}
	}
	curr := vs.Valve("AA")
	var maxFn int

	type route struct {
		from string
		to   string
	}
	cache := map[route][]string{}
	var cacheHits int

	var step func(path *Path, valve *Valve, level int, currFn int)
	step = func(path *Path, valve *Valve, level int, currFn int) {
		if len(path.actions) >= count || path.closed.Count() == 0 {
			for len(path.actions) < count {
				path.actions = append(path.actions, Rest{})
			}
			//fn := vs.Fitness("AA", path)
			if currFn > maxFn {
				log("path %v: currfn=%d", path.String(), currFn)
				maxFn = currFn
			}
			return
		}

		closed := path.closed.Values()
		sort.Strings(closed)
		for _, cl := range closed {
			rt := route{valve.ID, cl}
			sp, ok := cache[rt]
			if !ok {
				sp, _ = vs.ShortestPath(valve.ID, cl)
				sp = sp[1:]
				cache[rt] = sp
			} else {
				cacheHits++
			}

			sub := path.Clone()
			currID := valve.ID
			for _, s := range sp {
				sub.actions = append(sub.actions, Move{s})
				currID = s
			}
			sub.actions = append(sub.actions, Open{})
			curr := vs.Valve(currID)
			newFn := currFn + (count-len(sub.actions))*curr.FlowRate
			sub.closed.Remove(currID)
			step(sub, curr, level+1, newFn)
		}
	}
	step(startPath, curr, 0, 0)
	log("with %d cache hits", cacheHits)
	return maxFn, nil
}

func part2MainFunc(in string) (int, error) {
	vs := mustParseValveSet(in)
	count := 26

	startPath := &Path{
		closed: set.New[string](),
	}
	for _, v := range vs.Valves {
		if v.FlowRate > 0 {
			startPath.closed.Insert(v.ID)
		}
	}
	curr := vs.Valve("AA")
	var maxFn int

	type route struct {
		from string
		to   string
	}
	cache := map[route][]string{}
	var cacheHits int

	var step func(path *Path, valve *Valve, level int, currFn int)
	step = func(path *Path, valve *Valve, level int, currFn int) {
		if len(path.actions) >= count || path.closed.Count() == 0 {
			for len(path.actions) < count {
				path.actions = append(path.actions, Rest{})
			}
			if currFn > maxFn {
				log("path %v: currfn=%d", path.String(), currFn)
				maxFn = currFn
			}
			return
		}

		closed := path.closed.Values()
		sort.Strings(closed)
		for _, cl := range closed {
			rt := route{valve.ID, cl}
			sp, ok := cache[rt]
			if !ok {
				sp, _ = vs.ShortestPath(valve.ID, cl)
				sp = sp[1:]
				cache[rt] = sp
			} else {
				cacheHits++
			}

			sub := path.Clone()
			currID := valve.ID
			for _, s := range sp {
				sub.actions = append(sub.actions, Move{s})
				currID = s
			}
			sub.actions = append(sub.actions, Open{})
			curr := vs.Valve(currID)
			newFn := currFn + (count-len(sub.actions))*curr.FlowRate
			sub.closed.Remove(currID)
			step(sub, curr, level+1, newFn)
		}
	}
	step(startPath, curr, 0, 0)
	log("with %d cache hits", cacheHits)
	return maxFn, nil
}
