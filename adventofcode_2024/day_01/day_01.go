package day_01

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

func parseLists(in string) ([]int, []int, error) {
	lines := readutil.ReadLines(in)
	var list1, list2 []int
	for _, line := range lines {
		first, second, ok := strings.Cut(line, " ")
		if !ok {
			return nil, nil, fmt.Errorf("could not parse line: %s", line)
		}
		n1, err := strconv.Atoi(strings.TrimSpace(first))
		if err != nil {
			return nil, nil, fmt.Errorf("could not parse number: %s", first)
		}
		n2, err := strconv.Atoi(strings.TrimSpace(second))
		if err != nil {
			return nil, nil, fmt.Errorf("could not parse number: %s", first)
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	return list1, list2, nil
}

func part1MainFunc(in string) (int, error) {
	list1, list2, err := parseLists(in)
	if err != nil {
		return 0, fmt.Errorf("parse lists: %w", err)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	var sum int
	for i, n1 := range list1 {
		sum += mathutil.Abs(n1 - list2[i])
	}

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	list1, list2, err := parseLists(in)
	if err != nil {
		return 0, fmt.Errorf("parse lists: %w", err)
	}
	// histogramm of list2
	h := map[int]int{}
	for _, n := range list2 {
		h[n]++
	}

	var sum int
	for _, n1 := range list1 {
		sum += n1 * h[n1]
	}

	return sum, nil
}
