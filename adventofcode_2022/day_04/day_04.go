package day_04

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
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

type Range struct {
	Min int
	Max int
}

func (r Range) Normalize() Range {
	if r.Min > r.Max {
		return Range{r.Max, r.Min}
	}
	return r
}

func (r Range) ContainsRange(or Range) bool {
	return r.Min <= or.Min && r.Max >= or.Max
}

func (r Range) Contains(n int) bool {
	return n >= r.Min && n <= r.Max
}

func (r Range) Overlaps(or Range) bool {
	return r.Contains(or.Min) || r.Contains(or.Max)
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var cnt int
	for _, line := range lines {
		var r1Min, r1Max, r2Min, r2Max int
		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &r1Min, &r1Max, &r2Min, &r2Max)
		if err != nil {
			fatal("scan %q", line)
		}
		r1 := Range{r1Min, r1Max}.Normalize()
		r2 := Range{r2Min, r2Max}.Normalize()
		if r1.ContainsRange(r2) || r2.ContainsRange(r1) {
			cnt++
		}
	}
	return cnt, nil
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var cnt int
	for _, line := range lines {
		var r1Min, r1Max, r2Min, r2Max int
		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &r1Min, &r1Max, &r2Min, &r2Max)
		if err != nil {
			fatal("scan %q", line)
		}
		r1 := Range{r1Min, r1Max}.Normalize()
		r2 := Range{r2Min, r2Max}.Normalize()
		if r1.Overlaps(r2) || r2.Overlaps(r1) {
			cnt++
		}
	}
	return cnt, nil
}
