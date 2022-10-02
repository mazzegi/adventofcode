package day_06

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/slices"
	"github.com/pkg/errors"
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
	objs, err := parseObjects(in)
	errutil.ExitOnErr(err)

	objMap := map[string]object{}
	for _, obj := range objs {
		objMap[obj.name] = obj
	}

	orbitCount := func(obj object) int {
		var cnt int
		cobj := obj
		var ok bool
		for {
			if cobj.name == COM {
				return cnt
			}
			cnt++
			if cobj.orbitsAround == COM {
				return cnt
			}
			cobj, ok = objMap[cobj.orbitsAround]
			if !ok {
				fatal("no such object %q", cobj.orbitsAround)
			}
		}
	}

	var sum int
	for _, obj := range objs {
		sum += orbitCount(obj)
	}

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	objs, err := parseObjects(in)
	errutil.ExitOnErr(err)

	isOrbitedBy := func(objName string) []string {
		var oos []string
		for _, obj := range objs {
			if obj.orbitsAround == objName {
				oos = append(oos, obj.name)
			}
		}
		return oos
	}

	objMap := map[string]object{}
	for _, obj := range objs {
		obj.isOrbitedBy = isOrbitedBy(obj.name)
		objMap[obj.name] = obj
	}

	mustObj := func(name string) object {
		obj, ok := objMap[name]
		if !ok {
			fatal("no such object %q", name)
		}
		return obj
	}

	pathToCOM := func(obj object) []string {
		var path []string
		cobj := obj
		for {
			if cobj.orbitsAround == COM {
				return path
			}
			path = append(path, cobj.orbitsAround)
			cobj = mustObj(cobj.orbitsAround)
		}
	}

	dest := mustObj("SAN")
	curr := mustObj("YOU")

	currPath := pathToCOM(curr)
	destPath := pathToCOM(dest)

	// now look for the less common elt
	var transfers int
	for ci, o := range currPath {
		di := slices.Find(destPath, o)
		if di < 0 {
			continue
		}
		// its in the path
		transfers = ci + di

		break
	}

	return transfers, nil
}

const COM = "COM"

func parseObjects(in string) ([]object, error) {
	var objs []object
	for i, l := range readutil.ReadLines(in) {
		b, a, ok := strings.Cut(l, ")")
		if !ok {
			return nil, errors.Errorf("line %d, cannot parse %q as object", i, l)
		}
		if b == "" && a == "" {
			return nil, errors.Errorf("empty object at %d", i)
		}
		objs = append(objs, object{
			name:         a,
			orbitsAround: b,
		})
	}
	return objs, nil
}

type object struct {
	name         string
	orbitsAround string
	isOrbitedBy  []string
}
