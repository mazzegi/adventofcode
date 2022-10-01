package day_01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := Solve(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := SolveHalfWayRound(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

func Parse(in string) ([]int, error) {
	var ns []int
	in = strings.Trim(in, " \r\n\t")
	for _, r := range in {
		n, err := strconv.ParseUint(string(r), 10, 8)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-uint %q", string(r))
		}
		ns = append(ns, int(n))
	}
	return ns, nil
}

func Solve(in string) (int, error) {
	ns, err := Parse(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse")
	}
	if len(ns) == 0 {
		return 0, errors.Errorf("list is empty")
	}
	var sum int
	for i, n := range ns {
		var next int
		if i+1 < len(ns) {
			next = ns[i+1]
		} else {
			next = ns[0]
		}
		if n == next {
			sum += n
		}
	}

	return sum, nil
}

func SolveHalfWayRound(in string) (int, error) {
	ns, err := Parse(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse")
	}
	if len(ns) == 0 {
		return 0, errors.Errorf("list is empty")
	}
	fw := len(ns) / 2
	var sum int
	for i, n := range ns {
		var next int
		if i+fw < len(ns) {
			next = ns[i+fw]
		} else {
			next = ns[(i+fw)-len(ns)]
		}
		if n == next {
			sum += n
		}
	}

	return sum, nil
}
