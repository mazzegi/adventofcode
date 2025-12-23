package day_10

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
	"github.com/mazzegi/adventofcode/slices"
)

// max joltage: 278

func hashJoltages(js joltages) string {
	sl := make([]string, len(js))
	for i, j := range js {
		sl[i] = fmt.Sprintf("%d", j)
	}
	return strings.Join(sl, ",")
}

func hashJoltagesAndButtonIdx(js joltages, btnIdx int) string {
	sl := make([]string, len(js))
	for i, j := range js {
		sl[i] = fmt.Sprintf("%d", j)
	}
	return fmt.Sprintf("%s:%d", strings.Join(sl, ","), btnIdx)
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var ms []machine
	var ajs []int
	for _, line := range lines {
		m, err := parseMachine(line)
		if err != nil {
			return 0, fmt.Errorf("parse machine %q: %w", line, err)
		}
		ms = append(ms, m)
		ajs = append(ajs, m.targetJoltages...)
	}
	maxJ := slices.Max(ajs)
	fmt.Println("max joltage: ", maxJ)

	var sum int
	for i, m := range ms {
		// sort buttons with more effect to front
		sort.Slice(m.buttons, func(i, j int) bool {
			bs1 := m.buttons[i]
			bs2 := m.buttons[j]
			return len(bs1) > len(bs2)
		})

		np, err := findFewestButtonPressesForMachineJoltage(m)
		if err != nil {
			return 0, fmt.Errorf("find_for_machine: %w", err)
		}
		log("part2: machine %d -> %d presses", i+1, np)
		sum += np
	}

	return sum, nil
}

func applyButtonJoltages(jsIn joltages, btn button) joltages {
	jsOut := slices.Clone(jsIn)
	for _, ti := range btn {
		jsOut[ti]++
	}
	return jsOut
}

func findFewestButtonPressesForMachineJoltage(m machine) (int, error) {
	var min int
	foundOne := false
	dontTryThis := set.New[string]()
	thisWasOk := map[string]int{} // hash -> num_presses
	for _, btn := range m.buttons {
		startJs := make(joltages, len(m.targetJoltages))
		numPresses, ok := findFewestButtonPressesJoltage(m, startJs, btn, 0, min, dontTryThis, thisWasOk)
		if !ok {
			continue
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
	return min, nil
}

func findFewestButtonPressesJoltage(m machine, startJs joltages, startBtn button, currPresses int, currMin int, dontTryThis *set.Set[string], thisWasOk map[string]int) (int, bool) {
	if currMin > 0 && currPresses+1 >= currMin {
		// we dont get better
		return 0, false
	}
	currJs := applyButtonJoltages(startJs, startBtn)
	if joltagesEqual(currJs, m.targetJoltages) {
		return 1, true
	}
	currJsHash := hashJoltages(currJs)
	// is any of joltages higher than target? if yes break.
	for i, cj := range currJs {
		if cj > m.targetJoltages[i] {
			return 0, false
		}
	}
	if dontTryThis.Contains(currJsHash) {
		return 0, false
	}
	if np, ok := thisWasOk[currJsHash]; ok {
		return np, true
	}

	currPresses++

	var minfromHere int
	newMin := currMin
	foundOne := false
	for i, btn := range m.buttons {
		currJsAndBtnHash := hashJoltagesAndButtonIdx(currJs, i)
		// if dontTryThis.Contains(currJsAndBtnHash) {
		// 	continue
		// }
		_ = i
		var presses int
		var ok bool
		presses, ok = findFewestButtonPressesJoltage(m, currJs, btn, currPresses, newMin, dontTryThis, thisWasOk)

		// var presses int
		// var ok bool
		// if np, cacheok := thisWasOk[currJsAndBtnHash]; cacheok {
		// 	presses = np
		// 	ok = true
		// } else {
		// 	presses, ok = findFewestButtonPressesJoltage(m, currJs, btn, currPresses, newMin, dontTryThis, thisWasOk)
		// }
		if !ok {
			//dontTryThis.Insert(currJsAndBtnHash)
			continue
		}
		thisWasOk[currJsAndBtnHash] = presses
		if !foundOne {
			minfromHere = presses
			if newMin == 0 {
				newMin = currPresses + presses
			}
			foundOne = true
		} else if presses < minfromHere {
			minfromHere = presses
			if newMin == 0 {
				newMin = currPresses + presses
			}
		}
	}
	if !foundOne {
		dontTryThis.Insert(currJsHash)
		return 0, false
	}
	thisWasOk[currJsHash] = minfromHere + 1
	return minfromHere + 1, foundOne // +1 for the first press
}
