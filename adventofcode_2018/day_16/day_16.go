package day_16

import (
	"adventofcode_2018/errutil"
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/mazzegi/scan"
	"github.com/mazzegi/slices"
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
	res, err := part2MainFunc(input, inputProgram)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
type Sample struct {
	Before      []int
	Instruction []int
	After       []int
}

func parseSamples(in string) ([]Sample, error) {
	in = strings.ReplaceAll(in, "\n", " ")
	in = strings.ReplaceAll(in, "Before", "\nBefore")

	funcs := scan.BuiltinFuncs()
	funcs.Add("intsspaced", func(s string) (any, error) {
		return slices.Convert(strings.Split(s, " "), slices.ParseInt)
	})

	return scan.Lines[Sample]("Before: [{{before:[]int}}] {{instruction:intsspaced}} After:  [{{after:[]int}}]", funcs, bytes.NewBufferString(in))
}

func ints4(ns []int) [4]int {
	return [4]int{ns[0], ns[1], ns[2], ns[3]}
}

func ints3(ns []int) [3]int {
	return [3]int{ns[0], ns[1], ns[2]}
}

func opMatches(s Sample) []string {
	var ms []string
	for _, op := range allOps() {
		m := NewMachine(ints4(s.Before))
		m.Process(op, ints3(s.Instruction[1:]))
		if m.regs == ints4(s.After) {
			ms = append(ms, op)
		}
	}
	return ms
}

func part1MainFunc(in string) (int, error) {
	samples, err := parseSamples(in)
	if err != nil {
		return 0, err
	}
	var num int
	for _, s := range samples {
		ms := opMatches(s)
		if len(ms) >= 3 {
			num++
		}
	}

	return num, nil
}

func intersection(sl1, sl2 []string) []string {
	contains := func(sl []string, s string) bool {
		for _, cs := range sl {
			if cs == s {
				return true
			}
		}
		return false
	}

	var isl []string
	for _, s1 := range sl1 {
		if contains(sl2, s1) {
			isl = append(isl, s1)
		}
	}
	return isl
}

func removeOne(sl []string, rs string) ([]string, bool) {
	for i, s := range sl {
		if s == rs {
			//sl = append(sl[:i], sl[i+1:]...)
			return append(sl[:i], sl[i+1:]...), true
		}
	}
	return sl, false
}

func part2MainFunc(in string, prg string) (int, error) {
	samples, err := parseSamples(in)
	if err != nil {
		return 0, err
	}
	cands := map[int][]string{}
	for _, s := range samples {
		ops := opMatches(s)
		op := s.Instruction[0]

		exops, ok := cands[op]
		if !ok {
			cands[op] = ops
			continue
		}
		isOps := intersection(exops, ops)
		cands[op] = isOps
	}

	dumpCands := func() {
		var keys []int
		for op := range cands {
			keys = append(keys, op)
		}
		sort.Ints(keys)

		log("*** candidates ***")
		for _, k := range keys {
			log("%d => %v", k, cands[k])
		}
	}

	//dumpCands()
	for {
		act := false
		for op, ops := range cands {
			if len(ops) != 1 {
				continue
			}
			for rop, rops := range cands {
				if rop == op {
					continue
				}
				if nops, rem := removeOne(rops, ops[0]); rem {
					cands[rop] = nops
					act = true
				}
			}
		}

		if !act {
			break
		}
		//dumpCands()
	}
	dumpCands()

	funcs := scan.BuiltinFuncs()
	funcs.Add("intsspaced", func(s string) (any, error) {
		return slices.Convert(strings.Split(s, " "), slices.ParseInt)
	})
	type Instr struct {
		Codes []int
	}
	iss, err := scan.Lines[Instr]("{{codes:intsspaced}}", funcs, bytes.NewBufferString(prg))
	if err != nil {
		return 0, err
	}
	log("processing %d instructions ...", len(iss))

	m := NewMachine([4]int{0, 0, 0, 0})
	for _, is := range iss {
		op := cands[is.Codes[0]][0]
		m.Process(op, ints3(is.Codes[1:]))
	}

	return m.regs[0], nil
}
