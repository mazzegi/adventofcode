package day_18

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
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

//

func neighbours(c grid.Point3D) []grid.Point3D {
	return []grid.Point3D{
		c.Add(grid.P3D(1, 0, 0)),
		c.Add(grid.P3D(-1, 0, 0)),
		c.Add(grid.P3D(0, 1, 0)),
		c.Add(grid.P3D(0, -1, 0)),
		c.Add(grid.P3D(0, 0, 1)),
		c.Add(grid.P3D(0, 0, -1)),
	}
}

func part1MainFunc(in string) (int, error) {
	cubes := set.New[grid.Point3D]()
	for _, line := range readutil.ReadLines(in) {
		var p grid.Point3D
		_, err := fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		if err != nil {
			fatal("scan cube: %v", err)
		}
		cubes.Insert(p)
	}
	var unconnected int
	for _, c := range cubes.Values() {
		for _, n := range neighbours(c) {
			if !cubes.Contains(n) {
				unconnected++
			}
		}
	}
	return unconnected, nil
}

func part2MainFunc(in string) (int, error) {
	var min, max grid.Point3D
	inBounds := func(p grid.Point3D) bool {
		return p.X >= min.X && p.X <= max.X &&
			p.Y >= min.Y && p.Y <= max.Y &&
			p.Z >= min.Z && p.Z <= max.Z
	}

	cubes := set.New[grid.Point3D]()
	for i, line := range readutil.ReadLines(in) {
		var p grid.Point3D
		_, err := fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		if err != nil {
			fatal("scan cube: %v", err)
		}
		cubes.Insert(p)
		if i == 0 {
			min, max = p, p
			continue
		}
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Z < min.Z {
			min.Z = p.Z
		}

		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
		if p.Z > max.Z {
			max.Z = p.Z
		}
	}

	min = min.Add(grid.P3D(-1, -1, -1))
	max = max.Add(grid.P3D(1, 1, 1))

	//flood
	flooded := set.New[grid.Point3D]()
	start := min
	flooded.Insert(start)

	var extent func(p grid.Point3D)
	extent = func(p grid.Point3D) {
		for _, n := range neighbours(p) {
			if !inBounds(n) {
				continue
			}
			if cubes.Contains(n) || flooded.Contains(n) {
				continue
			}
			flooded.Insert(n)
			extent(n)
		}
	}
	extent(start)

	var unconnected int
	for _, c := range cubes.Values() {
		for _, n := range neighbours(c) {
			if flooded.Contains(n) {
				unconnected++
			}
		}
	}
	return unconnected, nil
}
