package day_19

import (
	"fmt"
	"sort"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/maps"
	"github.com/mazzegi/adventofcode/readutil"
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

func collectGeodes(blueprint Blueprint, minutes int) int {
	resources := map[Resource]int{
		Ore:      0,
		Clay:     0,
		Obsidian: 0,
		Geode:    0,
	}
	robots := map[Resource]int{
		Ore:      1,
		Clay:     0,
		Obsidian: 0,
		Geode:    0,
	}

	dump := func() {
		log("resources: ore=%d, clay=%d, obs=%d, geode=%d", resources[Ore], resources[Clay], resources[Obsidian], resources[Geode])
		log("robots   : ore=%d, clay=%d, obs=%d, geode=%d", robots[Ore], robots[Clay], robots[Obsidian], robots[Geode])
	}

	buildRobot := func(which Resource) {
		cost := blueprint.Costs[which]
		for cres, ccnt := range cost {
			resources[cres] -= ccnt
		}
		robots[which]++
	}

	canBuildIn := func(res Resource, resources map[Resource]int, robots map[Resource]int) int {
		cost := blueprint.Costs[res]
		var maxDT int
		for costRes, costCount := range cost {
			left := costCount - resources[costRes]
			if left <= 0 {
				continue
			}
			if robots[costRes] <= 0 {
				return -1
			}
			var dt int
			if left%robots[costRes] == 0 {
				dt = left / robots[costRes]
			} else {
				dt = left/robots[costRes] + 1
			}
			if dt > maxDT {
				maxDT = dt
			}
		}
		return maxDT
	}

	var goForRobot func(gres Resource) Resource
	goForRobot = func(gres Resource) Resource {
		if blueprint.CanBuild(gres, resources) {
			return gres
		}
		cost := blueprint.Costs[gres]
		rank := cost.RankDesc()
		for _, nr := range rank {
			if nr == gres {
				continue
			}
			// check, if its better to build this, or to wait, until the next more worth robot can be build

			if resources[nr] < cost[nr] {
				prev := nr.Prev()
				if prev == None {
					return goForRobot(nr)
				}
				canBuildPrevInCurr := canBuildIn(prev, resources, robots)
				if canBuildPrevInCurr < 0 {
					return goForRobot(nr)
				}
				testRobots := maps.Clone(robots)
				testRobots[nr]++

				testResources := maps.Clone(resources)
				testCost := blueprint.Costs[nr]
				for r, rc := range testCost {
					testResources[r] -= rc
				}

				canBuildPrevInNext := canBuildIn(prev, testResources, testRobots)
				if canBuildPrevInNext < canBuildPrevInCurr {
					return goForRobot(nr)
				}
				return None
			}
		}
		return None
	}

	for i := 0; i < minutes; i++ {
		buildRobotRes := goForRobot(Geode)

		for res, robcnt := range robots {
			resources[res] += robcnt
		}

		log("== Minute %d ==", i+1)
		dump()
		if buildRobotRes != None {
			buildRobot(buildRobotRes)
			log("build robot: %s", buildRobotRes)
		}
		log("")
	}

	return resources[Geode]
}

var earliestClay int
var earliestObsidian int
var earliestGeode int
var gmax int

// new approach
func collectGeodesRecursive(blueprint Blueprint, resources map[Resource]int, robots map[Resource]int, newRobotRes Resource, iter int, maxIter int) int {
	if iter >= maxIter {
		gcnt := resources[Geode]
		if gcnt > gmax {
			gmax = gcnt
			log("new max: %d", gmax)
			log("robots   : ore=%d, clay=%d, obs=%d, geode=%d", robots[Ore], robots[Clay], robots[Obsidian], robots[Geode])
		}

		return resources[Geode]
	}

	// check heuristically if max can be reached
	// if robots[Ore] > 6 {
	// 	return 0
	// }
	if earliestGeode > -1 && robots[Geode] == 0 && iter > earliestGeode+1 {
		return 0
	}
	if earliestObsidian > -1 && robots[Obsidian] == 0 && iter > earliestObsidian+2 {
		return 0
	}
	// if earliestClay > -1 && robots[Clay] == 0 && iter > earliestClay {
	// 	return 0
	// }

	// if iter >= 20 && robots[Geode] == 0 {
	// 	return 0
	// }
	// if iter >= 15 && robots[Obsidian] == 0 {
	// 	return 0
	// }
	// if iter >= 5 && robots[Clay] == 0 {
	// 	return 0
	// }

	cres := maps.Clone(resources)
	crobs := maps.Clone(robots)
	for res, robcnt := range crobs {
		cres[res] += robcnt
	}
	if newRobotRes != None {
		cost := blueprint.Costs[newRobotRes]
		for res, cnt := range cost {
			cres[res] -= cnt
		}
		crobs[newRobotRes]++
		if newRobotRes == Clay && (earliestClay < 0 || iter < earliestClay) {
			earliestClay = iter
			log("earliest clay at: %d", iter)
		}
		if newRobotRes == Obsidian && (earliestObsidian < 0 || iter < earliestObsidian) {
			earliestObsidian = iter
			log("earliest obsidian at: %d", iter)
		}
		if newRobotRes == Geode && (earliestGeode < 0 || iter < earliestGeode) {
			earliestGeode = iter
			log("earliest geode at: %d", iter)
		}
	}

	var maxGeodes int
	for _, res := range []Resource{Geode, Obsidian, Clay, Ore, None} {
		if blueprint.CanBuild(res, cres) {
			n := collectGeodesRecursive(blueprint, cres, crobs, res, iter+1, maxIter)
			if n > maxGeodes {
				maxGeodes = n
			}
			if res == Geode {
				break
			}
		}
	}
	return maxGeodes
}

func part1MainFunc(in string) (int, error) {
	var blueprints []Blueprint
	for _, line := range readutil.ReadLines(in) {
		bp := mustParseBlueprint(line)
		blueprints = append(blueprints, bp)
	}

	var sum int
	for _, bp := range blueprints {
		//num := collectGeodes(bp, 24)
		log("##### blueprint %02d #####", bp.ID)
		gmax = 0
		earliestClay = -1
		earliestObsidian = -1
		earliestGeode = -1
		resources := map[Resource]int{
			Ore:      0,
			Clay:     0,
			Obsidian: 0,
			Geode:    0,
		}
		robots := map[Resource]int{
			Ore:      1,
			Clay:     0,
			Obsidian: 0,
			Geode:    0,
		}
		num := collectGeodesRecursive(bp, resources, robots, None, 0, 24)
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
		//num := collectGeodes(bp, 24)
		log("##### blueprint %02d #####", bp.ID)
		gmax = 0
		earliestClay = -1
		earliestObsidian = -1
		earliestGeode = -1
		resources := map[Resource]int{
			Ore:      0,
			Clay:     0,
			Obsidian: 0,
			Geode:    0,
		}
		robots := map[Resource]int{
			Ore:      1,
			Clay:     0,
			Obsidian: 0,
			Geode:    0,
		}
		num := collectGeodesRecursive(bp, resources, robots, None, 0, 32)
		sum += num * bp.ID
		log("blueprint %02d done: %d (%d)", bp.ID, num, sum)
	}
	return sum, nil
}
