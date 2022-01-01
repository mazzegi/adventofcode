package day_01

import (
	"adventofcode_2021/errutil"
	"adventofcode_2021/readutil"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func Part1() {
	ic, err := incCount(input)
	errutil.ExitOnErr(err)
	fmt.Printf("inc-count: %d\n", ic)
}

func Part2() {
	ic, err := incWindowCount(input)
	errutil.ExitOnErr(err)
	fmt.Printf("inc-win-count: %d\n", ic)
}

func readValues(in string) ([]int, error) {
	var vals []int
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		if line == "" {
			continue
		}
		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parse int %q", line)
		}
		vals = append(vals, int(n))
	}
	return vals, nil
}

func incCount(in string) (int, error) {
	vals, err := readValues(in)
	if err != nil {
		return 0, errors.Wrap(err, "read-values")
	}

	var incCount int
	for i := 1; i < len(vals); i++ {
		if vals[i] > vals[i-1] {
			incCount++
		}
	}
	return incCount, nil
}

func incWindowCount(in string) (int, error) {
	vals, err := readValues(in)
	if err != nil {
		return 0, errors.Wrap(err, "read-values")
	}

	winVal := func(idx int) int {
		return vals[idx-2] + vals[idx-1] + vals[idx]
	}

	var incCount int
	for i := 3; i < len(vals); i++ {
		if winVal(i) > winVal(i-1) {
			incCount++
		}
	}
	return incCount, nil
}
