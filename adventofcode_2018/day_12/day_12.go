package day_12

import (
	"bytes"
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2018/errutil"

	"github.com/mazzegi/scan"
	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	skip := true
	if skip {
		return
	}
	res, err := part1MainFunc(inputState, inputRules, 20, 0)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	//res, err := part1MainFunc(inputState, inputRules, 50000000000)
	res, err := part1MainFunc(inputState, inputRules, 1000, 50000000000-1000)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type rule struct {
	Pattern  []byte
	Produces byte
}

const (
	empty byte = '.'
	occ   byte = '#'
)

func part1MainFunc(stateStr string, rulesStr string, gens int, offset int) (int, error) {
	rules, err := scan.Lines[rule]("{{pattern: []byte}} => {{produces: byte}}", scan.BuiltinFuncs(), bytes.NewBufferString(rulesStr))
	if err != nil {
		return 0, errors.Wrap(err, "scan rules")
	}

	findRule := func(pattern []byte) (rule, bool) {
		for _, rule := range rules {
			if bytes.Equal(rule.Pattern, pattern) {
				return rule, true
			}
		}
		return rule{}, false
	}

	state := []byte(stateStr)
	negState := []byte{}

	at := func(i int) byte {
		if i < -len(negState) {
			return empty
		}
		switch {
		case i < -len(negState):
			return empty
		case i < 0:
			return negState[len(negState)+i]
		case i >= len(state):
			return empty
		default:
			return state[i]
		}
	}

	subPattern := func(idx int) []byte {
		var p []byte
		for i := idx - 2; i <= idx+2; i++ {
			p = append(p, at(i))
		}
		return p
	}

	nextState := func() {
		var next []byte
		var negNext []byte
		for i := -(len(negState) + 3); i < len(state)+3; i++ {
			sub := subPattern(i)
			e := empty
			if rule, ok := findRule(sub); ok {
				e = rule.Produces
			}
			if i < 0 {
				negNext = append(negNext, e)
			} else {
				next = append(next, e)
			}
		}
		next = bytes.TrimRight(next, ".")
		negNext = bytes.TrimLeft(negNext, ".")
		state = next
		negState = negNext
	}

	log("00: |%s", string(state))
	for i := 0; i < gens; i++ {
		nextState()
		//log("%02d: %s|%s", i+1, string(negState), string(state))
		log("%s", string(bytes.Trim(state, ".")))

		if i%10000 == 0 {
			fmt.Printf("%d / %d => %.2f %%\n", i, gens, float64(i)/float64(gens)*100.0)
		}
	}

	var sum int
	nidx := -1
	for i := len(negState) - 1; i >= 0; i-- {
		if negState[i] == occ {
			sum += nidx
		}
		nidx--
	}
	for i := 0; i < len(state); i++ {
		if state[i] == occ {
			sum += (i + offset)
		}
	}

	//

	return sum, nil
}

func part2MainFunc(stateStr string, rulesStr string, gens int) (int, error) {
	return 0, nil
}
