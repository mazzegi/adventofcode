package day_24

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := strongestBridge(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := strongestLongestBridge(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type component struct {
	id   int
	pin1 int
	pin2 int
}

func (c component) flipped() component {
	return component{
		id:   c.id,
		pin1: c.pin2,
		pin2: c.pin1,
	}
}

func mustParseComponents(in string) []component {
	var cs []component
	lines := readutil.ReadLines(in)
	for id, line := range lines {
		c := component{id: id}
		_, err := fmt.Sscanf(line, "%d/%d", &c.pin1, &c.pin2)
		if err != nil {
			fatal("scan component %q: %v", line, err)
		}
		cs = append(cs, c)
	}
	if len(cs) == 0 {
		fatal("no data")
	}
	return cs
}

func findComponentsWithPinsAndAlign(cs []component, usedIDs map[int]bool, pins int) []component {
	var fcs []component
	for _, c := range cs {
		if _, ok := usedIDs[c.id]; ok {
			continue
		}
		if c.pin1 == pins {
			fcs = append(fcs, c)
		} else if c.pin2 == pins {
			fcs = append(fcs, c.flipped())
		}
	}

	return fcs
}

func strongestBridge(in string) (int, error) {
	cs := mustParseComponents(in)
	usedIDs := map[int]bool{}
	strength := buildStrongest(cs, usedIDs, 0)
	return strength, nil
}

func cloneUsedIDs(usedIDs map[int]bool) map[int]bool {
	cuids := map[int]bool{}
	for id := range usedIDs {
		cuids[id] = true
	}
	return cuids
}

func buildStrongest(cs []component, usedIDs map[int]bool, pins int) (maxStrength int) {
	fcs := findComponentsWithPinsAndAlign(cs, usedIDs, pins)
	maxStrength = 0
	for _, c := range fcs {
		cusedIDs := cloneUsedIDs(usedIDs)
		cusedIDs[c.id] = true
		strength := c.pin1 + c.pin2
		ms := buildStrongest(cs, cusedIDs, c.pin2)
		if strength+ms > maxStrength {
			maxStrength = strength + ms
		}
	}
	return maxStrength
}

func strongestLongestBridge(in string) (int, error) {
	cs := mustParseComponents(in)
	usedIDs := map[int]bool{}
	length, strength := buildStrongestLongest(cs, usedIDs, 0)
	log("len = %d, str= %d", length, strength)
	return strength, nil
}

func buildStrongestLongest(cs []component, usedIDs map[int]bool, pins int) (maxLength int, maxStrength int) {
	fcs := findComponentsWithPinsAndAlign(cs, usedIDs, pins)
	maxLength = 0
	maxStrength = 0
	for _, c := range fcs {
		cusedIDs := cloneUsedIDs(usedIDs)
		cusedIDs[c.id] = true
		length := 1
		strength := c.pin1 + c.pin2
		ml, ms := buildStrongestLongest(cs, cusedIDs, c.pin2)
		if ml+length > maxLength {
			maxLength = ml + length
			maxStrength = ms + strength
		} else if ml+length == maxLength && strength+ms > maxStrength {
			maxLength = ml + length
			maxStrength = ms + strength
		}
	}
	return maxLength, maxStrength
}
