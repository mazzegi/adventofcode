package day_10

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
	"github.com/mazzegi/adventofcode/slices"
	"github.com/mazzegi/adventofcode/stringutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

type indicators []bool

func indicatorsEqual(i1, i2 indicators) bool {
	return slices.Equal(i1, i2)
}

type button []int

func buttonsEqual(b1, b2 button) bool {
	return slices.Equal(b1, b2)
}

type joltages []int

type machine struct {
	targetIndicators indicators
	buttons          []button
	joltages         joltages
}

func parseMachine(s string) (machine, error) {
	fields := strings.Fields(s)
	if len(s) < 3 {
		return machine{}, fmt.Errorf("invalid machine string %q", s)
	}
	indicatorsStr := fields[0]
	joltagesStr := fields[len(fields)-1]
	buttonStrs := fields[1 : len(fields)-1]

	// parse indicator
	indicatorsStr = strings.TrimPrefix(indicatorsStr, "[")
	indicatorsStr = strings.TrimSuffix(indicatorsStr, "]")
	is := make(indicators, len(indicatorsStr))
	for i, r := range indicatorsStr {
		switch r {
		case '.':
			is[i] = false
		case '#':
			is[i] = true
		default:
			return machine{}, fmt.Errorf("invalid indicator string in %q", s)
		}
	}

	// joltages
	joltagesStr = strings.TrimPrefix(joltagesStr, "{")
	joltagesStr = strings.TrimSuffix(joltagesStr, "}")
	js, err := stringutil.StringsToInts(strings.Split(joltagesStr, ","))
	if err != nil {
		return machine{}, fmt.Errorf("joltages: strings_to_ints %q: %w", joltagesStr, err)
	}

	// buttons
	m := machine{
		targetIndicators: is,
		joltages:         js,
	}
	for _, bstr := range buttonStrs {
		bstr = strings.TrimPrefix(bstr, "(")
		bstr = strings.TrimSuffix(bstr, ")")
		b, err := stringutil.StringsToInts(strings.Split(bstr, ","))
		if err != nil {
			return machine{}, fmt.Errorf("buttons: strings_to_ints %q: %w", bstr, err)
		}
		//validate button
		for i, n := range b {
			if i < len(b)-1 {
				if n >= b[i+1] {
					return machine{}, fmt.Errorf("invalid button %q", bstr)
				}
			}
		}

		//
		m.buttons = append(m.buttons, b)
	}

	return m, nil
}

func applyButton(isIn indicators, btn button) indicators {
	isOut := slices.Clone(isIn)
	for _, ti := range btn {
		isOut[ti] = !isOut[ti]
	}
	return isOut
}

func hashIndicators(is indicators) string {
	isHash := make([]rune, len(is))
	for i, v := range is {
		if v {
			isHash[i] = '#'
		} else {
			isHash[i] = '.'
		}
	}
	return string(isHash)
}

func hashIndicatorsButtonTrial(is indicators, btnIdx int) string {
	isHash := make([]rune, len(is))
	for i, v := range is {
		if v {
			isHash[i] = '#'
		} else {
			isHash[i] = '.'
		}
	}
	return fmt.Sprintf("%s:%d", string(isHash), btnIdx)
}

type result struct {
	ok         bool
	numPresses int
}

// 484 is to high :) for 168 machines - so set max pressed to 484 - 168 ~ 320 (at least 1 for each machine)

// 488 to high

func findFewestButtonPresses(m machine, startIs indicators, startButtonIdx int, currMin int, currPresses int, maxPresses int,
	dontTryItAgainCache *set.Set[string], resultCache map[string]result, useResultCache bool) (int, bool) {

	if useResultCache {
		resHash := hashIndicatorsButtonTrial(startIs, startButtonIdx)
		if res, ok := resultCache[resHash]; ok {
			return res.numPresses + 1, res.ok
		}
	}

	if currPresses > maxPresses {
		return 0, false
	}
	if currMin > 0 && currPresses+1 >= currMin {
		// we dont get better
		//log("we dont get better")
		return 0, false
	}

	currIs := applyButton(startIs, m.buttons[startButtonIdx])
	if indicatorsEqual(currIs, m.targetIndicators) {
		return 1, true
	}
	isHash := hashIndicators(currIs)
	if dontTryItAgainCache.Contains(isHash) {
		// this indicator state we had already before - we dont get better, when we proceed further here
		return 0, false
	}
	clonedCache := dontTryItAgainCache.Clone()
	clonedCache.Insert(isHash)
	currPresses++
	//log("curr: [%s]", isHash)

	var min int
	foundOne := false
	for ib := range m.buttons {
		if ib == startButtonIdx {
			// applying the same button would just reverse the previous action
			continue
		}

		subMin, ok := findFewestButtonPresses(m, currIs, ib, currMin, currPresses, maxPresses, clonedCache, resultCache, useResultCache)
		resHash := hashIndicatorsButtonTrial(startIs, startButtonIdx)
		if !ok {
			resultCache[resHash] = result{
				ok:         false,
				numPresses: 0,
			}
			continue
		}
		resultCache[resHash] = result{
			ok:         true,
			numPresses: subMin,
		}
		if !foundOne {
			min = subMin
			foundOne = true
		} else if subMin < min {
			min = subMin
		}
	}
	if !foundOne {
		return 0, false
	}
	return min + 1, foundOne // +1 for the first press
}

func findFewestButtonPressesForMachine(m machine) (int, error) {
	var min int
	foundOne := false
	resultCache := map[string]result{}
	for ib := range m.buttons {
		currIs := make(indicators, len(m.targetIndicators))
		dontTryItAgainCache := set.New[string]()
		dontTryItAgainCache.Insert(hashIndicators(currIs))
		numPresses, ok := findFewestButtonPresses(m, currIs, ib, min, 0, 320, dontTryItAgainCache, resultCache, true)
		resHash := hashIndicatorsButtonTrial(currIs, ib)
		if !ok {
			resultCache[resHash] = result{
				ok:         false,
				numPresses: 0,
			}
			continue
		}
		resultCache[resHash] = result{
			ok:         true,
			numPresses: numPresses,
		}
		if !foundOne {
			min = numPresses
			foundOne = true
		} else if numPresses < min {
			min = numPresses
		}
	}
	if !foundOne {
		return 0, fmt.Errorf("found no pressing sequence")
	}
	// and another run with not using the result cache
	for ib := range m.buttons {
		currIs := make(indicators, len(m.targetIndicators))
		dontTryItAgainCache := set.New[string]()
		dontTryItAgainCache.Insert(hashIndicators(currIs))
		numPresses, ok := findFewestButtonPresses(m, currIs, ib, min, 0, 320, dontTryItAgainCache, resultCache, false)
		//resHash := hashIndicatorsButtonTrial(currIs, ib)
		if !ok {
			// resultCache[resHash] = result{
			// 	ok:         false,
			// 	numPresses: 0,
			// }
			continue
		}
		// resultCache[resHash] = result{
		// 	ok:         true,
		// 	numPresses: numPresses,
		// }
		if numPresses < min {
			min = numPresses
		}
	}

	return min, nil
}

func part1MainFunc(in string) (int, error) {
	t0 := time.Now()
	lines := readutil.ReadLines(in)
	var ms []machine
	for _, line := range lines {
		m, err := parseMachine(line)
		if err != nil {
			return 0, fmt.Errorf("parse machine %q: %w", line, err)
		}
		ms = append(ms, m)
	}
	log("parsed %d machines in %s", len(ms), time.Since(t0).Round(time.Microsecond))

	var sum int
	for i, m := range ms {
		np, err := findFewestButtonPressesForMachine(m)
		if err != nil {
			return 0, fmt.Errorf("find_for_machine: %w", err)
		}
		log("machine %d -> %d presses", i+1, np)
		sum += np
	}

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
