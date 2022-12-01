package day_16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/slices"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input, 100)
	errutil.ExitOnErr(err)
	log("part1: result = %v", res)
}

func Part2() {
	res, err := part2MainFunc(input, 100)
	errutil.ExitOnErr(err)
	log("part2: result = %v", res)
}

func parseInput(in string) []int {
	var ns []int
	in = strings.TrimSpace(in)
	for _, r := range in {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			fatal("parsing %q: %v", string(r), err)
		}
		ns = append(ns, n)
	}
	return ns
}

func part1MainFunc(in string, phases int) (string, error) {
	list := parseInput(in)
	basePattern := []int{0, 1, 0, -1}
	nthPattern := func(n int) []int {
		var np []int
		for _, bpv := range basePattern {
			np = append(np, slices.Repeat(bpv, n+1)...)
		}
		return np
	}

	for p := 0; p < phases; p++ {
		out := make([]int, len(list))
		for i := 0; i < len(out); i++ {
			pattern := nthPattern(i)
			var outv int
			ipattern := 1
			for _, inv := range list {
				outv += inv * pattern[ipattern]
				ipattern++
				if ipattern >= len(pattern) {
					ipattern = 0
				}
			}
			d := mathutil.Abs(outv)
			d = d - 10*(d/10)
			out[i] = d
		}
		list = out
	}

	var s8 string
	for i := 0; i < 8; i++ {
		s8 += fmt.Sprintf("%d", list[i])
	}
	return s8, nil
}

func part2MainFunc(in string, phases int) (string, error) {
	listBase := parseInput(in)
	//repeat 10000 times
	var list []int
	for i := 0; i < 10000; i++ {
		list = append(list, listBase...)
	}

	basePattern := []int{0, 1, 0, -1}
	nthPattern := func(n int) []int {
		var np []int
		for _, bpv := range basePattern {
			np = append(np, slices.Repeat(bpv, n+1)...)
		}
		return np
	}

	for p := 0; p < phases; p++ {
		out := make([]int, len(list))
		for i := 0; i < len(out); i++ {
			pattern := nthPattern(i)
			var outv int
			ipattern := 1
			for _, inv := range list {
				outv += inv * pattern[ipattern]
				ipattern++
				if ipattern >= len(pattern) {
					ipattern = 0
				}
			}
			d := mathutil.Abs(outv)
			d = d - 10*(d/10)
			out[i] = d
		}
		list = out
	}

	var s8 string
	for i := 0; i < 8; i++ {
		s8 += fmt.Sprintf("%d", list[i])
	}
	return s8, nil
}
