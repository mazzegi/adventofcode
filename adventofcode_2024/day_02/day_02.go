package day_02

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/slices"
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

func parseInput(in string) ([][]int, error) {
	lines := readutil.ReadLines(in)
	var res [][]int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		ns, err := readutil.ReadInts(line, " ")
		if err != nil {
			return nil, err
		}
		if len(ns) < 2 {
			return nil, fmt.Errorf("not enough numbers: %v", ns)
		}
		res = append(res, ns)
	}
	return res, nil
}

func isIncreasingReportSafe(rep []int) bool {
	for i := 1; i < len(rep); i++ {
		n1, n2 := rep[i-1], rep[i]
		if n1 >= n2 {
			// not increasing
			return false
		}
		if n2-n1 > 3 {
			// they differ to much
			return false
		}
	}
	return true
}

func isDecreasingReportSafe(rep []int) bool {
	for i := 1; i < len(rep); i++ {
		n1, n2 := rep[i-1], rep[i]
		if n1 <= n2 {
			// not decreasing
			return false
		}
		if n1-n2 > 3 {
			// they differ to much
			return false
		}
	}
	return true
}

func isReportSafe(rep []int) bool {
	if len(rep) < 2 {
		// neither increasing nor decreasing
		return true
	}
	switch {
	case rep[0] < rep[1]:
		return isIncreasingReportSafe(rep)
	case rep[0] > rep[1]:
		return isDecreasingReportSafe(rep)
	default:
		// equal not safe
		return false
	}
}

func part1MainFunc(in string) (int, error) {
	reports, err := parseInput(in)
	if err != nil {
		return 0, fmt.Errorf("parse input: %w", err)
	}

	var safeCount int
	for _, rep := range reports {
		if isReportSafe(rep) {
			safeCount++
		}
	}

	return safeCount, nil
}

func part2MainFunc(in string) (int, error) {
	reports, err := parseInput(in)
	if err != nil {
		return 0, fmt.Errorf("parse input: %w", err)
	}

	var safeCount int
loop_reports:
	for _, rep := range reports {
		if isReportSafe(rep) {
			safeCount++
			continue loop_reports
		}
		// try removing levels
		for i := 0; i < len(rep); i++ {
			probe := slices.DeleteIdx(rep, i)
			if isReportSafe(probe) {
				safeCount++
				continue loop_reports
			}
		}
	}

	return safeCount, nil
}
