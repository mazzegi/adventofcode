package day_05

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
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

type Range struct {
	First   int
	Last    int
	Invalid bool
}

func (r Range) in(n int) bool {
	return r.First <= n && n <= r.Last
}

func parseRange(s string) (Range, error) {
	var r Range
	_, err := fmt.Sscanf(s, "%d-%d", &r.First, &r.Last)
	if err != nil {
		return Range{}, fmt.Errorf("scan: %w", err)
	}
	if r.First > r.Last {
		return Range{}, fmt.Errorf("invalid range %q", s)
	}

	return r, nil
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var ranges []Range
	var ids []int
	for _, line := range lines {
		switch {
		case strings.Contains(line, "-"):
			r, err := parseRange(line)
			if err != nil {
				return 0, fmt.Errorf("parse-range %q: %w", line, err)
			}
			ranges = append(ranges, r)

		default:
			id, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("parse-id %q: %w", line, err)
			}
			ids = append(ids, int(id))
		}
	}
	//
	isFresh := func(id int) bool {
		for _, r := range ranges {
			if r.in(id) {
				return true
			}
		}
		return false
	}
	var numFresh int
	for _, id := range ids {
		if isFresh(id) {
			numFresh++
		}
	}

	return numFresh, nil
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var ranges []Range
	var ids []int
	for _, line := range lines {
		switch {
		case strings.Contains(line, "-"):
			r, err := parseRange(line)
			if err != nil {
				return 0, fmt.Errorf("parse-range %q: %w", line, err)
			}
			ranges = append(ranges, r)

		default:
			id, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("parse-id %q: %w", line, err)
			}
			ids = append(ids, int(id))
		}
	}
	_ = ids

	// sort ranges
	sort.Slice(ranges, func(i, j int) bool {
		r1 := ranges[i]
		r2 := ranges[j]
		return r1.First < r2.First
	})
	//
	for i := 1; i < len(ranges); i++ {
		prev := ranges[i-1]
		curr := ranges[i]

		merged, ok := tryMerge(prev, curr)
		if !ok {
			continue
		}
		ranges[i-1].Invalid = true
		ranges[i] = merged
	}
	// count
	var total int
	for _, r := range ranges {
		if r.Invalid {
			continue
		}
		total += r.Last - r.First + 1
	}

	return total, nil
}

// we already consider ranges are sorted
func tryMerge(r1, r2 Range) (Range, bool) {
	if r2.First > r1.Last {
		// not possible
		return Range{}, false
	}
	return Range{
		First: r1.First,
		Last:  mathutil.Max(r1.Last, r2.Last),
	}, true
}

/*
3-5
10-14
12-18
16-20

3-5
*10-14
*10-18
10-20
*/
