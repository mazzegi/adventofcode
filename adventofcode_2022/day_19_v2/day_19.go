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
	}
}

func (s *State) Hash() string {
	return fmt.Sprintf("%s#%s", hashResMap(s.resources), hashResMap(s.robots))
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
	return &Path{
		builds: slices.Clone(p.builds),
	}
}

//

func NewStateCache() *StateCache {
	return &StateCache{
		states: map[string]*Path{},
	}
}

type StateCache struct {
	states map[string]*Path
}

func (c *StateCache) Find(state *State) (*Path, bool) {
	p, ok := c.states[state.Hash()]
	if !ok {
		return nil, false
	}
	return p.Clone(), true
}

func (c *StateCache) Add(state *State, path *Path) {
	c.states[state.Hash()] = path.Clone()
}

func shortestPathToNextGeodeRobot(blueprint Blueprint, cache *StateCache, state *State, iter int, maxIter int) (*Path, bool) {
	if iter >= maxIter-1 {
		return nil, false
	}
	if pc, ok := cache.Find(state); ok {
		return pc, true
	}
	var shortestPath *Path
	for _, res := range []Resource{Geode, Obsidian, Clay, Ore, None} {
		cloneState := state.Clone()
		cloneState.Mine()
		if !blueprint.CanBuild(res, cloneState.resources) {
			continue
		}
		if res == Geode {
			path := NewPath(Geode)
			cache.Add(cloneState, path)
			return path, true
		}
		cloneState.Build(blueprint, res)
		foundPath, found := shortestPathToNextGeodeRobot(blueprint, cache, cloneState, iter+1, maxIter)
		if !found {
			continue
		}
		if len(foundPath.builds) > maxIter-iter {
			log("örks")
		}
		thisPath := NewPath(res)
		thisPath.builds = append(thisPath.builds, foundPath.builds...)
		if shortestPath == nil || len(thisPath.builds) < len(shortestPath.builds) {
			cache.Add(cloneState, thisPath)
			shortestPath = thisPath
		}
	}
	if shortestPath != nil && len(shortestPath.builds) > 0 {
		return shortestPath, true
	}
	return nil, false
}

func collectGeodes(blueprint Blueprint, maxIter int) int {
	cache := NewStateCache()
	state := NewState()
	state.robots[Ore] = 1
	iter := 0
	for {
		path, ok := shortestPathToNextGeodeRobot(blueprint, cache, state, iter, maxIter)
		if !ok {
			return state.resources[Geode]
		}
		//apply path
		for _, bres := range path.builds {
			state.Mine()
			state.Build(blueprint, bres)
			iter++
		}
	}
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
	// var blueprints []Blueprint
	// for _, line := range readutil.ReadLines(in) {
	// 	bp := mustParseBlueprint(line)
	// 	blueprints = append(blueprints, bp)
	// }

	// var sum int
	// for _, bp := range blueprints {
	// 	//num := collectGeodes(bp, 24)
	// 	log("##### blueprint %02d #####", bp.ID)
	// 	gmax = 0
	// 	earliestClay = -1
	// 	earliestObsidian = -1
	// 	earliestGeode = -1
	// 	resources := map[Resource]int{
	// 		Ore:      0,
	// 		Clay:     0,
	// 		Obsidian: 0,
	// 		Geode:    0,
	// 	}
	// 	robots := map[Resource]int{
	// 		Ore:      1,
	// 		Clay:     0,
	// 		Obsidian: 0,
	// 		Geode:    0,
	// 	}
	// 	num := collectGeodesRecursive(bp, resources, robots, None, 0, 32)
	// 	sum += num * bp.ID
	// 	log("blueprint %02d done: %d (%d)", bp.ID, num, sum)
	// }
	return 0, nil
}
