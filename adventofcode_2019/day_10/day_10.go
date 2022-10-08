package day_10

import (
	"fmt"
	"sort"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/slices"
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
	center := grid.Pt(14, 17)
	res, err := part2MainFunc(input, center)
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
	var maxObj grid.Point
	for _, obj := range objs {
		voc := visibleObjCount(obj)
		if voc > max {
			max = voc
			maxObj = obj
		}
	}
	fmt.Println("obj", maxObj)

	return max, nil
}

//

func part2MainFunc(in string, center grid.Point) (int, error) {
	g, err := grid.ParseBinaryGrid(in)
	errutil.ExitOnErr(err)
	objs := g.SetPoints()

	//sort objs
	sort.Slice(objs, func(i, j int) bool {
		return center.DistTo(objs[i]) < center.DistTo(objs[j])
	})

	vapCount := 0
	currBeam := grid.Pt(center.X, 0).Sub(center)
	// vaporize first in current beam
	for i, obj := range objs {
		if obj == center {
			continue
		}
		if obj.X == center.X && obj.Y < center.Y {
			objs = slices.DeleteIdx(objs, i)
			vapCount++
			fmt.Println(vapCount, "vaporize", obj)
			break
		}
	}

	nextTarget := func() grid.Point {
		var minAngle float64
		var minAngleObj grid.Point
		first := true
		for _, obj := range objs {
			if obj == center {
				continue
			}
			//transform relative to center
			tobj := obj.Sub(center)
			if tobj.IsMultipleOf(currBeam) {
				continue
			}
			ang := currBeam.LeftAngleTo(tobj)
			//log("    %s => %f", obj, ang)
			if first || (ang > 0 && ang < minAngle) {
				minAngle = ang
				minAngleObj = obj
				first = false
			}
		}
		return minAngleObj
	}

	var res int
	for {
		if len(objs) <= 1 {
			break
		}
		next := nextTarget()
		objs = slices.DeleteFirst(objs, next)
		currBeam = next.Sub(center)
		vapCount++
		if vapCount == 200 {
			res = next.X*100 + next.Y
		}
		//fmt.Println(vapCount, "vaporize", next)
	}

	return res, nil
}
