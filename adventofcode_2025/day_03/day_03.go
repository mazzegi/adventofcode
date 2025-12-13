package day_03

import (
	"fmt"
	"math"
	"slices"
	"strconv"
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

func firstMax(ts []int) (max int, idx int) {
	for i, t := range ts {
		if i == 0 || t > max {
			max = t
			idx = i
		}
	}
	return
}

type bank []int

func maxJoltage(b bank) int {
	max1, maxidx1 := firstMax(b[:len(b)-1])
	max2, _ := firstMax(b[maxidx1+1:])
	return 10*max1 + max2
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var banks []bank
	for _, line := range lines {
		var b bank
		for _, r := range line {
			n, err := strconv.ParseInt(string(r), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("parse int: %w", err)
			}
			b = append(b, int(n))
		}
		banks = append(banks, b)
	}
	var sum int
	for _, b := range banks {
		mj := maxJoltage(b)
		sum += mj
	}
	return sum, nil
}

func maxJoltagePart2(b bank) int {
	var maxSl []int
	bLeft := slices.Clone(b)
	for i := range 12 {
		minLeftCount := 12 - i - 1

		searchUntilIdx := len(bLeft) - minLeftCount
		max1, maxidx1 := firstMax(bLeft[:searchUntilIdx])
		maxSl = append(maxSl, max1)
		bLeft = bLeft[maxidx1+1:]
	}
	var maxJ int
	maxSlRev := slices.Clone(maxSl)
	slices.Reverse(maxSlRev)
	for i, m := range maxSlRev {
		maxJ += m * int(math.Pow10(i))
	}
	// 0 -> 1
	// 1 -> 0
	// 9876543211110
	return maxJ
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var banks []bank
	for _, line := range lines {
		var b bank
		for _, r := range line {
			n, err := strconv.ParseInt(string(r), 10, 8)
			if err != nil {
				return 0, fmt.Errorf("parse int: %w", err)
			}
			b = append(b, int(n))
		}
		banks = append(banks, b)
	}
	var sum int
	for _, b := range banks {
		mj := maxJoltagePart2(b)
		sum += mj
	}
	return sum, nil
}
