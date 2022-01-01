package day_20

import (
	"adventofcode_2016/errutil"
	"adventofcode_2016/intutil"
	"adventofcode_2016/readutil"
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

func Part1() {
	n, err := LowestNonBlocked(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: lowest = %d\n", n)
}

func Part2() {
	n, err := AllowedCount(input, maxNum)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: count = %d\n", n)
}

func ParseRange(s string) (Range, error) {
	var r Range
	_, err := fmt.Sscanf(s, "%d-%d", &r.min, &r.max)
	if err != nil {
		return Range{}, errors.Wrapf(err, "scan range %q", s)
	}
	if r.min > r.max {
		return Range{}, errors.Errorf("range %q. min (%d) is greater than max (%d)", s, r.min, r.max)
	}
	return r, nil
}

func ParseRanges(in string) ([]*Range, error) {
	var rs []*Range
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		r, err := ParseRange(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-range %q", line)
		}
		rs = append(rs, &r)
	}
	return rs, nil
}

const maxNum = 4294967295

func LowestNonBlocked(in string) (int, error) {
	rs, err := ParseRanges(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-ranges")
	}

	isBlocked := func(n int) bool {
		for _, r := range rs {
			if r.contains(n) {
				return true
			}
		}
		return false
	}

	for n := 0; n <= maxNum; n++ {
		if !isBlocked(n) {
			return n, nil
		}
	}

	return 0, errors.Errorf("all are blocked")
}

func AllowedCountNaive(in string, max int) (int, error) {
	rs, err := ParseRanges(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-ranges")
	}

	isBlocked := func(n int) bool {
		for _, r := range rs {
			if r.contains(n) {
				return true
			}
		}
		return false
	}

	var cnt int
	for n := 0; n <= max; n++ {
		if !isBlocked(n) {
			cnt++
		}
	}

	return cnt, nil
}

func AllowedCount(in string, max int) (int, error) {
	rs, err := ParseRanges(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-ranges")
	}
	sort.Slice(rs, func(i, j int) bool {
		if rs[i].min == rs[j].min {
			return rs[i].max < rs[j].max
		}
		return rs[i].min < rs[j].min
	})

	nrs := normalizeRanges(rs)
	fmt.Printf("normalized %d to %d\n", len(rs), len(nrs))

	//count blocked content - as they are not intersecting any more
	var ccnt int
	for _, nr := range nrs {
		ccnt += nr.max - nr.min + 1
	}

	allowed := (max + 1) - ccnt
	return allowed, nil
}

type Range struct {
	min int
	max int
}

func (r *Range) clone() *Range {
	return &Range{
		min: r.min,
		max: r.max,
	}
}

func (r *Range) String() string {
	return fmt.Sprintf("%d-%d", r.min, r.max)
}

func (r *Range) contains(n int) bool {
	return r.min <= n && n <= r.max
}

func (r *Range) intersects(or *Range) bool {
	return r.contains(or.min) || r.contains(or.max)
}

func (r *Range) merge(or *Range) {
	r.min = intutil.Min(r.min, or.min)
	r.max = intutil.Max(r.max, or.max)
}

func normalizeRanges(rs []*Range) []*Range {
	var nrs []*Range

	add := func(r *Range) {
		for _, nr := range nrs {
			if nr.intersects(r) {
				fmt.Printf("merge %s to %s\n", r, nr)
				nr.merge(r)

				return
			}
		}
		fmt.Printf("add %s\n", r)
		nrs = append(nrs, r.clone())
	}

	for _, r := range rs {
		add(r)
	}

	if len(nrs) < len(rs) {
		nrs = normalizeRanges(nrs)
	}

	return nrs
}
