package day_25

import (
	"adventofcode_2017/errutil"
	"fmt"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := checksum(inputBlueprint)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	// res, err := part2MainFunc(input)
	// errutil.ExitOnErr(err)
	// log("part2: result = %d", res)
}

type turing struct {
	tape map[int]bool
}

func (t *turing) checksum() int {
	var cs int
	for _, v := range t.tape {
		if v {
			cs++
		}
	}
	return cs
}

func (t *turing) get(x int) bool {
	if v, ok := t.tape[x]; ok {
		return v
	}
	return false
}

func (t *turing) set(x int) {
	t.tape[x] = true
}

func (t *turing) reset(x int) {
	delete(t.tape, x)
}

func checksum(in blueprint) (int, error) {
	tm := &turing{
		tape: map[int]bool{},
	}

	currState := in.states[in.startStateID]
	currX := 0
	for i := 0; i < in.steps; i++ {
		tv := tm.get(currX)
		var ins instruction
		if !tv {
			ins = currState.if0
		} else {
			ins = currState.if1
		}

		if ins.write == 0 {
			tm.reset(currX)
		} else {
			tm.set(currX)
		}

		if ins.move == moveLeft {
			currX--
		} else {
			currX++
		}

		currState = in.states[ins.continueState]
	}

	cs := tm.checksum()
	return cs, nil
}

// func part2MainFunc(in string) (int, error){
// 	return 0, nil
// }
