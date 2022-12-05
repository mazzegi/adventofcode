package day_05

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
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
	res, err := part1MainFunc(input, inputStacks)
	errutil.ExitOnErr(err)
	log("part1: result = %s", res)
}

func Part2() {
	res, err := part2MainFunc(input, inputStacks)
	errutil.ExitOnErr(err)
	log("part2: result = %s", res)
}

type Stack []rune

type Move struct {
	Cnt  int
	From int
	To   int
}

func part1MainFunc(in string, stacks []Stack) (string, error) {
	stacks = slices.Clone(stacks)
	for i, stack := range stacks {
		stacks[i] = slices.Reverse(stack)
	}
	lines := readutil.ReadLines(in)
	var moves []Move
	for _, line := range lines {
		var mv Move
		fmt.Sscanf(line, "move %d from %d to %d", &mv.Cnt, &mv.From, &mv.To)
		moves = append(moves, mv)
	}

	for _, mv := range moves {
		from := stacks[mv.From-1]
		to := stacks[mv.To-1]
		for i := 0; i < mv.Cnt; i++ {
			r := from[len(from)-1]
			from = from[:len(from)-1]
			to = append(to, r)
		}
		stacks[mv.From-1] = from
		stacks[mv.To-1] = to
	}
	var msg string
	for _, stack := range stacks {
		if len(stack) > 0 {
			msg += string(stack[len(stack)-1])
		}
	}

	return msg, nil
}

func part2MainFunc(in string, stacks []Stack) (string, error) {
	stacks = slices.Clone(stacks)
	for i, stack := range stacks {
		stacks[i] = slices.Reverse(stack)
	}
	lines := readutil.ReadLines(in)
	var moves []Move
	for _, line := range lines {
		var mv Move
		fmt.Sscanf(line, "move %d from %d to %d", &mv.Cnt, &mv.From, &mv.To)
		moves = append(moves, mv)
	}

	for i, mv := range moves {
		_ = i
		from := stacks[mv.From-1]
		to := stacks[mv.To-1]

		idx := len(from) - mv.Cnt
		if idx < 0 {
			log("ola")
		}

		mvcs := from[idx:]
		from = from[:idx]
		to = append(to, mvcs...)

		stacks[mv.From-1] = from
		stacks[mv.To-1] = to
	}
	var msg string
	for _, stack := range stacks {
		if len(stack) > 0 {
			msg += string(stack[len(stack)-1])
		}
	}

	return msg, nil
}
