package day_06

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/set"
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

func isMarker(s string, cnt int) bool {
	if len(s) != cnt {
		return false
	}
	hs := set.New[rune]()
	for _, r := range s {
		if hs.Contains(r) {
			return false
		}
		hs.Insert(r)
	}
	return true
}

func part1MainFunc(in string) (int, error) {
	if len(in) < 4 {
		fatal("to less characters in input")
	}
	for i := 0; i < len(in)-4; i++ {
		probe := in[i : i+4]
		if isMarker(probe, 4) {
			return i + 4, nil
		}
	}
	fatal("no marker found in %q", in)
	return -1, nil
}

func part2MainFunc(in string) (int, error) {
	if len(in) < 14 {
		fatal("to less characters in input")
	}
	for i := 0; i < len(in)-14; i++ {
		probe := in[i : i+14]
		if isMarker(probe, 14) {
			return i + 14, nil
		}
	}
	fatal("no marker found in %q", in)
	return -1, nil
}
