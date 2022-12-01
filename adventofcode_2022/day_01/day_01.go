package day_01

import (
	"fmt"
	"sort"
	"strconv"

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

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLinesKeepEmpty(in)
	var max int
	var curr int
	for _, line := range lines {
		if line != "" {
			n, err := strconv.Atoi(line)
			if err != nil {
				fatal("atoi: %v", err)
			}
			curr += n
			continue
		}
		//empty
		if curr > 0 {
			if curr > max {
				max = curr
			}
			curr = 0
		}
	}
	if curr > 0 && curr > max {
		max = curr
	}
	return max, nil
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLinesKeepEmpty(in)
	var cals []int
	var curr int
	for _, line := range lines {
		if line != "" {
			n, err := strconv.Atoi(line)
			if err != nil {
				fatal("atoi: %v", err)
			}
			curr += n
			continue
		}
		//empty
		if curr > 0 {
			cals = append(cals, curr)
			curr = 0
		}
	}
	if curr > 0 {
		cals = append(cals, curr)
	}
	sort.Slice(cals, func(i, j int) bool {
		return cals[i] > cals[j]
	})
	var sum int
	for i := 0; i < 3; i++ {
		if i < len(cals) {
			sum += cals[i]
		}
	}
	return sum, nil
}
