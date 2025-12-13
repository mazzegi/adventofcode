package day_04

import (
	"fmt"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
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

func numSetNeighbours(g *grid.BinaryGrid, col, row int) int {
	var num int
	for ir := row - 1; ir <= row+1; ir++ {
		if ir < 0 || ir >= g.NumRows() {
			continue
		}
		for ic := col - 1; ic <= col+1; ic++ {
			if ic < 0 || ic >= g.NumCols() {
				continue
			}
			if ir == row && ic == col {
				// dont count center itself
				continue
			}
			if g.IsSet(ic, ir) {
				num++
			}
		}
	}
	return num
}

func part1MainFunc(in string) (int, error) {
	g, err := grid.ParseBinaryGrid(in)
	if err != nil {
		return 0, fmt.Errorf("parse-grid: %w", err)
	}
	var numAccessible int
	for r := range g.NumRows() {
		for c := range g.NumCols() {
			if !g.IsSet(c, r) {
				continue
			}
			ns := numSetNeighbours(g, c, r)
			if ns < 4 {
				numAccessible++
			}
		}
	}

	return numAccessible, nil
}

func part2MainFunc(in string) (int, error) {
	g, err := grid.ParseBinaryGrid(in)
	if err != nil {
		return 0, fmt.Errorf("parse-grid: %w", err)
	}

	removeAllAccessible := func() int {
		var numRemoved int
		for r := range g.NumRows() {
			for c := range g.NumCols() {
				if !g.IsSet(c, r) {
					continue
				}
				ns := numSetNeighbours(g, c, r)
				if ns < 4 {
					g.Unset(c, r)
					numRemoved++
				}
			}
		}
		return numRemoved
	}

	var numRemoved int
	for {
		rem := removeAllAccessible()
		if rem == 0 {
			break
		}
		numRemoved += rem
	}

	return numRemoved, nil
}
