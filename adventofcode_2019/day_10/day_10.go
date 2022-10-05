package day_10

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
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
	g, err := grid.ParseBinaryGrid(in)
	errutil.ExitOnErr(err)

	objs := g.SetPoints()

	canSeeObject := func(o1, o2 grid.Point) bool {
		for _, io := range objs {
			if io == o1 || io == o2 {
				continue
			}
			if io.Between(o1, o2) {
				return false
			}
		}
		return true
	}

	visibleObjCount := func(cobj grid.Point) int {
		var cnt int
		for _, obj := range objs {
			if obj == cobj {
				continue
			}
			if canSeeObject(cobj, obj) {
				cnt++
			}
		}
		return cnt
	}

	var max int
	for _, obj := range objs {
		voc := visibleObjCount(obj)
		if voc > max {
			max = voc
		}
	}

	return max, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
