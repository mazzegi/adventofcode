package day_06

import (
	"fmt"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
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
	res, err := part2MainFunc(input2)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

type RaceResult struct {
	Time     int
	Distance int
}

func CalcDistance(timeAllowed int, hold int) int {
	if hold <= 0 || hold >= timeAllowed {
		return 0
	}
	// speed := hold
	// rem := timeAllowed - hold
	// dist := speed * rem
	dist := hold * (timeAllowed - hold)
	return dist
}

func part1MainFunc(in []RaceResult) (int, error) {
	prod := 1
	for _, rr := range in {
		wins := 0
		for hold := 0; hold <= rr.Time; hold++ {
			dh := CalcDistance(rr.Time, hold)
			if dh > rr.Distance {
				wins++
			}
		}
		prod *= wins
	}
	return prod, nil
}

func part2MainFunc(in RaceResult) (int, error) {
	wins := 0
	for hold := 0; hold <= in.Time; hold++ {
		dh := CalcDistance(in.Time, hold)
		if dh > in.Distance {
			wins++
		}
	}
	return wins, nil
}
