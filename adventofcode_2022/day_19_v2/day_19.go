package day_19

import (
	"fmt"
	"reflect"
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
	newRobot  Resource
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
		newRobot:  s.newRobot,
	}
}

func (s *State) Hash() string {
	return fmt.Sprintf("%s#%s#%d", hashResMap(s.resources), hashResMap(s.robots), s.newRobot)
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
		states: map[string][]*Path{},
	}
}

type StateCache struct {
	states map[string][]*Path
}

func (c *StateCache) Find(state *State) ([]*Path, bool) {
	ps, ok := c.states[state.Hash()]
	if !ok {
		return nil, false
	}
	return ps, true
}

func (c *StateCache) Add(state *State, paths []*Path) {
	c.states[state.Hash()] = paths
}

func shortestPathsToNextGeodeRobot(blueprint Blueprint, cache *StateCache, badCache *StateCache, state *State, minGeodeAtIter *int, iter int, maxIter int) ([]*Path, bool) {
	if iter >= maxIter {
		return nil, false
	}
	if *minGeodeAtIter > -1 && iter > *minGeodeAtIter {
		return nil, false
	}
	if _, ok := badCache.Find(state); ok {
		return nil, false
	}
	if pc, ok := cache.Find(state); ok {
		if len(pc[0].builds) <= maxIter-iter {
			return pc, true
		}
	}
	var shortestPaths []*Path
	shortestPathsLen := func() int {
		if len(shortestPaths) == 0 {
			return -1
		}
		return len(shortestPaths[0].builds)
	}

	for _, res := range []Resource{Geode, Obsidian, Clay, Ore, None} {
		cloneState := state.Clone()
		//cloneState.Mine()
		if !blueprint.CanBuild(res, cloneState.resources) {
			continue
		}
		if res == Geode {
			//log("Geode at iter %d", iter)
			path := NewPath(Geode)
			cache.Add(cloneState, []*Path{path})
			return []*Path{path}, true
		}
		cloneState.Mine()
		cloneState.Build(blueprint, res)
		foundPaths, found := shortestPathsToNextGeodeRobot(blueprint, cache, badCache, cloneState, minGeodeAtIter, iter+1, maxIter)
		if !found {
			badCache.Add(cloneState, nil)
			continue
		}

		fpl := len(foundPaths[0].builds)
		for _, fp := range foundPaths {
			if len(fp.builds) != fpl {
				panic("oops - thta shouldn't happen at all")
			}
		}

		foundTotalIter := iter + fpl
		if *minGeodeAtIter < 0 {
			*minGeodeAtIter = foundTotalIter
			//log("min-geode: %d", *minGeodeAtIter)
		} else if foundTotalIter < *minGeodeAtIter {
			*minGeodeAtIter = foundTotalIter
			//log("min-geode: %d", *minGeodeAtIter)
		}

		if fpl > maxIter-iter {
			log("örks %d > %d", fpl, maxIter-iter)
		}

		spl := shortestPathsLen()

		thisPathLen := fpl + 1
		var thisPaths []*Path
		for _, fp := range foundPaths {
			thisPath := NewPath(res)
			thisPath.builds = append(thisPath.builds, fp.builds...)
			thisPaths = append(thisPaths, thisPath)
		}
		if spl < 0 {
			shortestPaths = thisPaths
		} else if thisPathLen < spl {
			shortestPaths = thisPaths
		} else if thisPathLen == spl {
			shortestPaths = append(shortestPaths, thisPaths...)
		}
	}
	spl := shortestPathsLen()
	if spl > 0 {
		cache.Add(state, shortestPaths)
		return shortestPaths, true
	}
	return nil, false
}

func shortestPaths(blueprint Blueprint, cache *StateCache, badCache *StateCache, state *State, iter int, maxIter int) ([]*Path, bool) {
	totalPaths := []*Path{}
	minGeodeAtIter := -1
	paths, ok := shortestPathsToNextGeodeRobot(blueprint, cache, badCache, state, &minGeodeAtIter, iter, maxIter)
	if !ok {
		return nil, false
	}
	for _, path := range paths {
		cloneState := state.Clone()
		newIter := applyPath(blueprint, cloneState, path)
		newIter += iter
		thisPaths, ok := shortestPaths(blueprint, cache, badCache, cloneState, newIter, maxIter)
		if ok {
			for _, tp := range thisPaths {
				newPath := NewPath(path.builds...)
				newPath.builds = append(newPath.builds, tp.builds...)
				totalPaths = append(totalPaths, newPath)
			}
		} else {
			totalPaths = append(totalPaths, path)
		}
	}
	return totalPaths, true
}

func applyPath(blueprint Blueprint, state *State, path *Path) int {
	iter := 0
	for _, bres := range path.builds {
		if !blueprint.CanBuild(bres, state.resources) {
			panic("oops")
		}
		state.Mine()
		state.Build(blueprint, bres)
		iter++
	}
	return iter
}

func collectGeodes(blueprint Blueprint, maxIter int) int {
	cache := NewStateCache()
	badCache := NewStateCache()
	state := NewState()
	state.robots[Ore] = 1
	iter := 0
	paths, ok := shortestPaths(blueprint, cache, badCache, state, iter, maxIter)
	if !ok {
		panic("oh my god")
	}

	max := 0
	for _, totalPath := range paths {
		//log("total-path (%d): %v", len(totalPath.builds), totalPath)
		finalState := NewState()
		finalState.robots[Ore] = 1

		for fi := 0; fi < maxIter; fi++ {
			//for _, bres := range totalPath {
			var bres Resource
			if fi < len(totalPath.builds) {
				bres = totalPath.builds[fi]
			} else {
				bres = None
			}
			if !blueprint.CanBuild(bres, finalState.resources) {
				log("oops")
			}
			finalState.Mine()
			// log("== Minute %02d ==", fi+1)
			// log("resources: ore=%d, clay=%d, obs=%d, geode=%d", finalState.resources[Ore], finalState.resources[Clay], finalState.resources[Obsidian], finalState.resources[Geode])
			// log("robots   : ore=%d, clay=%d, obs=%d, geode=%d", finalState.robots[Ore], finalState.robots[Clay], finalState.robots[Obsidian], finalState.robots[Geode])
			// log("")
			finalState.Build(blueprint, bres)
		}
		if finalState.resources[Geode] > max {
			max = finalState.resources[Geode]
		}
	}

	return max
}

func collectGeodes_depr(blueprint Blueprint, maxIter int) int {
	cache := NewStateCache()
	badCache := NewStateCache()
	state := NewState()
	state.robots[Ore] = 1
	state.newRobot = None
	iter := 0
	totalPath := []Resource{}
	for {
		min := -1
		paths, ok := shortestPathsToNextGeodeRobot(blueprint, cache, badCache, state, &min, iter, maxIter)
		if !ok {
			break
		}
		log("%d paths", len(paths))
		paths = slices.DedupFunc(paths, func(t1, t2 *Path) bool {
			return reflect.DeepEqual(t1.builds, t2.builds)
		})
		log("%d dedup paths", len(paths))
		// totalPath = append(totalPath, path.builds...)
		// //apply path
		// for _, bres := range path.builds {
		// 	if !blueprint.CanBuild(bres, state.resources) {
		// 		log("oops")
		// 	}
		// 	state.Mine()
		// 	state.Build(blueprint, bres)
		// 	iter++
		// }
	}

	log("total-path (%d): %v", len(totalPath), totalPath)
	finalState := NewState()
	finalState.robots[Ore] = 1

	for fi := 0; fi < maxIter; fi++ {
		//for _, bres := range totalPath {
		var bres Resource
		if fi < len(totalPath) {
			bres = totalPath[fi]
		} else {
			bres = None
		}
		if !blueprint.CanBuild(bres, finalState.resources) {
			log("oops")
		}
		finalState.Mine()
		// log("== Minute %02d ==", fi+1)
		// log("resources: ore=%d, clay=%d, obs=%d, geode=%d", finalState.resources[Ore], finalState.resources[Clay], finalState.resources[Obsidian], finalState.resources[Geode])
		// log("robots   : ore=%d, clay=%d, obs=%d, geode=%d", finalState.robots[Ore], finalState.robots[Clay], finalState.robots[Obsidian], finalState.robots[Geode])
		// log("")
		finalState.Build(blueprint, bres)
	}

	return finalState.resources[Geode]
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

	var sum int
	for _, bp := range blueprints {
		log("##### blueprint %02d #####", bp.ID)
		num := collectGeodes(bp, 32)
		sum += num * bp.ID
		log("blueprint %02d done: %d (%d)", bp.ID, num, sum)
	}
	return sum, nil
}
