package day_18

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
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

type cacheValue struct {
	doneIn  int
	pending bool
}

func part1MainFunc(in string) (int, error) {
	m, err := ParseMap(in)
	errutil.ExitOnErr(err)

	cache := map[string]cacheValue{}
	state := NewState(m)

	steps, ok := stepsUntilDone(state, 0, cache)
	if !ok {
		return 0, fmt.Errorf("didnt make it")
	}

	return steps, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}

func stepsUntilDone(state *State, steps int, cache map[string]cacheValue) (int, bool) {
	hash := hashState(state)
	if state.done {
		cache[hash] = cacheValue{doneIn: steps}
		return steps, true
	}
	if val, ok := cache[hash]; ok {
		// we reached that state before
		if val.doneIn > 0 {
			return val.doneIn, true
		}
		return 0, false
	}
	cache[hash] = cacheValue{pending: true}

	nbs := accessibleNeighbours(state)
	var min *int
	for _, nb := range nbs {
		cstate := cloneState(state)
		moveTo(cstate, nb)
		substeps, ok := stepsUntilDone(cstate, steps+1, cache)
		if !ok {
			cache[hashState(cstate)] = cacheValue{doneIn: -1}
			continue
		}
		if min == nil {
			min = &substeps
		} else if substeps < *min {
			*min = substeps
		}
	}
	if min == nil {
		return 0, false
	}
	cache[hash] = cacheValue{
		doneIn:  *min,
		pending: false,
	}
	return *min, true
}
