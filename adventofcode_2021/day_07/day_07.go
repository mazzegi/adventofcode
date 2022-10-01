package day_07

import (
	"fmt"
	"sort"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/intutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := LeastFuel(input, linearCost)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := LeastFuel(input, progressiveCost)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

func linearCost(pos int, ns []int) int {
	var sum int
	for _, n := range ns {
		sum += intutil.AbsInt(n - pos)
	}
	return sum
}

func progressiveCost(pos int, ns []int) int {
	var sum int
	for _, n := range ns {
		d := intutil.AbsInt(n - pos)
		sum += d * (d + 1) / 2
	}
	return sum
}

func LeastFuel(in string, costFunc func(pos int, ns []int) int) (int, error) {
	ns, err := readutil.ReadInts(in, ",")
	if err != nil {
		return 0, errors.Wrap(err, "read-ints")
	}
	if len(ns) == 0 {
		return 0, errors.Errorf("no entries")
	}
	sort.Ints(ns)
	maxPos := ns[len(ns)-1]
	minPos := ns[0]

	var lessCost int
	var lessCostPos int
	first := true
	for pos := minPos; pos <= maxPos; pos++ {
		c := costFunc(pos, ns)
		if first {
			lessCost = c
			lessCostPos = pos
			first = false
			continue
		}
		if c < lessCost {
			lessCost = c
			lessCostPos = pos
		}
	}
	fmt.Printf("less: pos=%d, cost=%d\n", lessCostPos, lessCost)
	return lessCost, nil
}
