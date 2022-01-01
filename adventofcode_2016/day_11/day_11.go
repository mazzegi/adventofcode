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
)

func initEnv() *environment {
	env := &environment{}
	env.floors = append(env.floors, &container{
		chips:      makeSet("thulium", "ruthenium", "cobalt"),
		generators: makeSet("polonium", "thulium", "promethium", "ruthenium", "cobalt"),
	})
	env.floors = append(env.floors, &container{
		chips:      makeSet("polonium", "promethium"),
		generators: makeSet(),
	})
	env.floors = append(env.floors, &container{
		chips:      makeSet(),
		generators: makeSet(),
	})
	env.floors = append(env.floors, &container{
		chips:      makeSet(),
		generators: makeSet(),
	})
	return env
}

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := minSteps(initEnv())
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	// res, err := part2MainFunc(input)
	// errutil.ExitOnErr(err)
	// log("part2: result = %d", res)
}

//
type set map[string]bool

func makeSet(elts ...string) set {
	s := set{}
	for _, elt := range elts {
		s[elt] = true
	}
	return s
}

func (s set) add(elt string) {
	s[elt] = true
}

func (s set) remove(elt string) {
	delete(s, elt)
}

func (s set) contains(elt string) bool {
	_, ok := s[elt]
	return ok
}

func (s set) containsOtherThan(elt string) bool {
	_, ok := s[elt]
	if ok {
		return len(s) >= 2
	}
	return len(s) >= 1
}

func (s set) clone() set {
	cs := set{}
	for elt := range s {
		cs[elt] = true
	}
	return cs
}

func (s set) empty() bool {
	return len(s) == 0
}

//
// type chip struct {
// 	element string
// }

// type generator struct {
// 	element string
// }

type container struct {
	chips      set
	generators set
}

func (c *container) clone() *container {
	return &container{
		chips:      c.chips.clone(),
		generators: c.generators.clone(),
	}
}

type environment struct {
	floors        []*container
	elevator      *container
	elevatorFloor int
}

func (env *environment) clone() *environment {
	cenv := &environment{
		elevator:      env.elevator.clone(),
		elevatorFloor: env.elevatorFloor,
		floors:        make([]*container, len(env.floors)),
	}
	for i, f := range env.floors {
		cenv.floors[i] = f.clone()
	}
	return cenv
}

const (
	typeChip      string = "chip"
	typeGenerator string = "generator"
)

type item struct {
	typ     string
	element string
}

type itemPair struct {
	item1 *item
	item2 *item
}

func (c *container) canRemoveChip(remElt string) bool {
	if c.generators.empty() {
		return true
	}

	for chipElt := range c.chips {
		if chipElt == remElt {
			continue
		}
		if c.generators.contains(chipElt) {
			continue
		}
		return false
	}
	return true
}

func (c *container) canRemoveGenerator(remElt string) bool {
	if c.chips.empty() {
		return true
	}
	if !c.generators.containsOtherThan(remElt) {
		//will be empty after removal
		return true
	}

	willContainGen := func(elt string) bool {
		cont := c.generators.contains(elt)
		if !cont {
			return false
		}
		if elt == remElt {
			return false
		}
		return true
	}

	for chipElt := range c.chips {
		if willContainGen(chipElt) {
			//safe
			continue
		}
		return false
	}
	return true
}

func (env *environment) pairsToBringInElevator() []itemPair {
	var ps []itemPair
	floor := env.floors[env.elevatorFloor]

	//single
	for chipElt := range floor.chips {
		if floor.canRemoveChip(chipElt) {
			ps = append(ps, itemPair{
				item1: &item{typeChip, chipElt},
				item2: nil,
			})
		}
	}
	for genElt := range floor.generators {
		if floor.canRemoveGenerator(genElt) {
			ps = append(ps, itemPair{
				item1: &item{typeGenerator, genElt},
				item2: nil,
			})
		}
	}

	return ps
}

func minSteps(env *environment) (int, error) {

	return 0, nil
}

// func part2MainFunc(in string) (int, error) {
// 	return 0, nil
// }
