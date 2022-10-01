package day_05

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := StepsToExit(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := StepsToExitDec(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

func StepsToExit(in string) (int, error) {
	ns, err := readutil.ReadInts(in, "\n")
	if err != nil {
		return 0, errors.Wrap(err, "read-ints")
	}
	if len(ns) == 0 {
		return 0, errors.Errorf("list is empty")
	}

	pos := 0
	step := 0
	for {
		v := ns[pos]
		ns[pos]++
		pos += v
		step++
		if pos < 0 || pos >= len(ns) {
			return step, nil
		}
	}
}

func StepsToExitDec(in string) (int, error) {
	ns, err := readutil.ReadInts(in, "\n")
	if err != nil {
		return 0, errors.Wrap(err, "read-ints")
	}
	if len(ns) == 0 {
		return 0, errors.Errorf("list is empty")
	}

	pos := 0
	step := 0
	for {
		v := ns[pos]
		if v >= 3 {
			ns[pos]--
		} else {
			ns[pos]++
		}
		pos += v
		step++
		if pos < 0 || pos >= len(ns) {
			return step, nil
		}
	}
}
