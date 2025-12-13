package day_02

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
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
	First int
	Last  int
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

func isInvalidID(id int) bool {
	s := fmt.Sprintf("%d", id)
	if len(s)%2 != 0 {
		return false
	}
	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]
	return s1 == s2
}

func part1MainFunc(in string) (int, error) {
	var rs []Range
	sl := readutil.ReadStrings(in, ",")
	for _, s := range sl {
		r, err := parseRange(s)
		if err != nil {
			return 0, fmt.Errorf("parse range %q: %w", s, err)
		}
		rs = append(rs, r)
	}
	//
	var sum int
	for _, r := range rs {
		for n := r.First; n <= r.Last; n++ {
			if isInvalidID(n) {
				sum += n
			}
		}
	}

	return sum, nil
}

func isInvalidIDPart2(id int) bool {
	s := fmt.Sprintf("%d", id)

	for i := 1; i <= len(s)/2; i++ {
		start := s[:i]
		patternCount := len(s) / len(start)
		repeated := strings.Repeat(start, patternCount)
		if s == repeated {
			return true
		}
	}
	return false
}

func part2MainFunc(in string) (int, error) {
	var rs []Range
	sl := readutil.ReadStrings(in, ",")
	for _, s := range sl {
		r, err := parseRange(s)
		if err != nil {
			return 0, fmt.Errorf("parse range %q: %w", s, err)
		}
		rs = append(rs, r)
	}
	//
	var sum int
	for _, r := range rs {
		for n := r.First; n <= r.Last; n++ {
			if isInvalidIDPart2(n) {
				sum += n
			}
		}
	}

	return sum, nil
}
