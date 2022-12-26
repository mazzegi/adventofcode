package day_19

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/maps"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/slices"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

const skip1 = true

func Part1() {
	if skip1 {
		return
	}
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type Resource int

const (
	None     Resource = 0
	Ore      Resource = 1
	Clay     Resource = 2
	Obsidian Resource = 3
	Geode    Resource = 4
)

func (r Resource) String() string {
	switch r {
	case None:
		return "None"
	case Ore:
		return "Ore"
	case Clay:
		return "Clay"
	case Obsidian:
		return "Obsidian"
	case Geode:
		return "Geode"
	default:
		return "<na>"
	}
}

func (r Resource) Prev() Resource {
	switch r {
	case None:
		return Ore
	case Ore:
		return Clay
	case Clay:
		return Obsidian
	case Obsidian:
		return Geode
	case Geode:
		return None
	default:
		return None
	}
}

type Cost map[Resource]int

func (c Cost) RankDesc() []Resource {
	var rs []Resource
	for r, n := range c {
		if n > 0 {
			rs = append(rs, r)
		}
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] >= rs[j]
	})
	return rs
}

type Blueprint struct {
	ID    int
	Costs map[Resource]Cost
}

func (bp Blueprint) CanBuild(res Resource, with map[Resource]int) bool {
	cost, ok := bp.Costs[res]
	if !ok {
		fatal("invalid resource %d", res)
	}
	for cres, ccnt := range cost {
		wcnt, ok := with[cres]
		if !ok {
			return false
		}
		if ccnt > wcnt {
			return false
		}
	}
	return true
}

func mustParseBlueprint(s string) Blueprint {
	var (
		id            int
		oreOre        int
		clayOre       int
		obsidianOre   int
		obsidianClay  int
		geodeOre      int
		geodeObsidian int
	)
	_, err := fmt.Sscanf(s, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&id, &oreOre, &clayOre, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)
	if err != nil {
		fatal("scan: %v", err)
	}
	return Blueprint{
		ID: id,
		Costs: map[Resource]Cost{
			None:     {},
			Ore:      {Ore: oreOre},
			Clay:     {Ore: clayOre},
			Obsidian: {Ore: obsidianOre, Clay: obsidianClay},
			Geode:    {Ore: geodeOre, Obsidian: geodeObsidian},
		},
	}
}

func hashResMap(m map[Resource]int) string {
	ks := maps.Keys(m)
	slices.Sort(ks)
	var sl []string
	for _, k := range ks {
		sl = append(sl, fmt.Sprintf("%d:%d", k, m[k]))
	}
	return strings.Join(sl, ",")
}

type State struct {
	resources map[Resource]int
	robots    map[Resource]int
	iter      int
}

func NewState() *State {
	return &State{
		resources: map[Resource]int{},
		robots:    map[Resource]int{},
	}
}

func (s *State) Clone() *State {
	return &State{
		resources: maps.Clone(s.resources),
		robots:    maps.Clone(s.robots),
		iter:      s.iter,
	}
}

func (s *State) Hash() string {
	return fmt.Sprintf("%02d#%s#%s", s.iter, hashResMap(s.resources), hashResMap(s.robots))
}

func (s *State) Mine() {
	for res, robcnt := range s.robots {
		s.resources[res] += robcnt
	}
}

func (s *State) Build(blueprint Blueprint, res Resource) {
	if res == None {
		return
	}
	cost := blueprint.Costs[res]
	for res, cnt := range cost {
		s.resources[res] -= cnt
	}
	s.robots[res]++
}

//

type Path struct {
	builds []Resource
}

func NewPath(ress ...Resource) *Path {
	return &Path{
		builds: ress,
	}
}

func (p *Path) Clone() *Path {
	if p == nil {
		return nil
	}
	return &Path{
		builds: slices.Clone(p.builds),
	}
}

//

func NewStateCache() *StateCache {
	return &StateCache{
		states: map[string]int{},
	}
}

type StateCache struct {
	states map[string]int
}

func (c *StateCache) Find(state *State) (int, bool) {
	n, ok := c.states[state.Hash()]
	if !ok {
		return n, false
	}
	return n, true
}

func (c *StateCache) Add(state *State, max int) {
	c.states[state.Hash()] = max
}

type Collector struct {
	blueprint Blueprint
	cache     *StateCache
	maxIter   int
	currMax   int
}

func (c *Collector) heuristicallyExclude(state *State) bool {
	left := c.maxIter - state.iter
	nr := state.robots[Geode]
	maxPossible := state.resources[Geode] + left*nr + (left-1)*(left-2)/2
	if maxPossible < c.currMax {
		return true
	}

	return false
}

func (c *Collector) maxGeodes(state *State) int {
	if state.iter >= c.maxIter {
		n := state.resources[Geode]
		if n > c.currMax {
			c.currMax = n
			log("new-max: %d", c.currMax)
		}
		return n
	}
	if cmax, ok := c.cache.Find(state); ok {
		return cmax
	}
	if c.heuristicallyExclude(state) {
		c.cache.Add(state, state.resources[Geode])
		return state.resources[Geode]
	}
	var max int
	for _, buildRes := range []Resource{Geode, Obsidian, Clay, Ore, None} {
		clonedState := state.Clone()
		if !c.blueprint.CanBuild(buildRes, clonedState.resources) {
			continue
		}
		clonedState.Mine()
		clonedState.Build(c.blueprint, buildRes)
		clonedState.iter++
		subMax := c.maxGeodes(clonedState)
		if subMax > max {
			max = subMax
		}
	}
	c.cache.Add(state, max)
	return max
}

func collectGeodes(blueprint Blueprint, maxIter int) int {
	cache := NewStateCache()
	state := NewState()
	state.robots[Ore] = 1
	state.iter = 0
	coll := &Collector{
		blueprint: blueprint,
		cache:     cache,
		maxIter:   maxIter,
		currMax:   0,
	}
	max := coll.maxGeodes(state)

	//max := maxGeodes(blueprint, cache, state, maxIter)

	return max
}

func part1MainFunc(in string) (int, error) {
	var blueprints []Blueprint
	for _, line := range readutil.ReadLines(in) {
		bp := mustParseBlueprint(line)
		blueprints = append(blueprints, bp)
	}

	var sum int
	for _, bp := range blueprints {
		log("##### blueprint %02d #####", bp.ID)
		num := collectGeodes(bp, 24)
		sum += num * bp.ID
		log("blueprint %02d done: %d (%d)", bp.ID, num, sum)
	}
	return sum, nil
}

// 2138 => too low
// 2327 => too low

func part2MainFunc(in string) (int, error) {
	var blueprints []Blueprint
	for _, line := range readutil.ReadLines(in) {
		bp := mustParseBlueprint(line)
		blueprints = append(blueprints, bp)
	}

	bps := blueprints[:3]
	//bps := blueprints

	for _, bp := range bps {
		log("##### blueprint %02d #####", bp.ID)
		num := collectGeodes(bp, 32)
		log("blueprint %02d done: %d ", bp.ID, num)
	}
	return 0, nil
}
