package day_18

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2021/day_18/sfish"
	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := magnitudeOfFinalSum(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := largestSumMagnitude(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

func parsePairs(in string) ([]*sfish.Pair, error) {
	var ps []*sfish.Pair
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		p, err := sfish.Parse(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse %q", line)
		}
		ps = append(ps, p)
	}
	return ps, nil
}

func magnitudeOfFinalSum(in string) (int, error) {
	pairs, err := parsePairs(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-pairs")
	}
	if len(pairs) == 0 {
		return 0, errors.Errorf("no data")
	}

	sum := pairs[0]
	for _, pair := range pairs[1:] {
		sum = sfish.Add(sum, pair)
	}

	magn := sum.Magnitude()
	return magn, nil
}

func largestSumMagnitude(in string) (int, error) {
	pairs, err := parsePairs(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-pairs")
	}
	if len(pairs) == 0 {
		return 0, errors.Errorf("no data")
	}

	var largest int
	var lp1, lp2 *sfish.Pair
	for i1, p1 := range pairs {
		for i2, p2 := range pairs {
			if i1 == i2 {
				continue
			}
			sum := sfish.Add(p1, p2)
			magn := sum.Magnitude()
			if magn > largest {
				largest = magn
				lp1 = p1
				lp2 = p2
			}
		}
	}

	log("p1 : %q", lp1.String())
	log("p2 : %q", lp2.String())
	log("sum: %q", sfish.Add(lp1, lp2).String())

	return largest, nil
}
