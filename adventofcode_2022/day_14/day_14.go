package day_14

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/mathutil"
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

type Path []grid.Point

type Material string

const (
	Rock Material = "#"
	Sand Material = "o"
)

type Cave struct {
	Paths            []Path
	Occ              map[grid.Point]Material
	maxY, minX, maxX int
	Floor            int
}

func validatePath(path Path) {
	if len(path) == 0 {
		fatal("path is empty")
	}
	for i := 0; i < len(path)-1; i++ {
		p1 := path[i]
		p2 := path[i+1]
		if p1.X != p2.X && p1.Y != p2.Y {
			fatal("no straight line")
		}
	}
}

func mustParsePath(in string) Path {
	var path []grid.Point
	sl := strings.Split(in, "->")
	for _, s := range sl {
		s = strings.TrimSpace(s)
		var p grid.Point
		_, err := fmt.Sscanf(s, "%d,%d", &p.X, &p.Y)
		if err != nil {
			fatal("scan point %q: %v", s, err)
		}
		path = append(path, p)
	}
	validatePath(path)
	return path
}

func pointsBetween(p1, p2 grid.Point) []grid.Point {
	if p1 == p2 {
		return []grid.Point{}
	}
	bw := []grid.Point{}
	if p1.Y == p2.Y {
		dx := mathutil.Sign(p2.X - p1.X)
		for x := p1.X + dx; x != p2.X; x += dx {
			bw = append(bw, grid.Pt(x, p1.Y))
		}
		return bw
	}
	if p1.X == p2.X {
		dy := mathutil.Sign(p2.Y - p1.Y)
		for y := p1.Y + dy; y != p2.Y; y += dy {
			bw = append(bw, grid.Pt(p1.X, y))
		}
		return bw
	}
	fatal("no straight line")
	return bw
}

func expandPath(path Path) Path {
	var exPath Path
	for i := 0; i < len(path)-1; i++ {
		p1 := path[i]
		p2 := path[i+1]
		bps := pointsBetween(p1, p2)
		if i == 0 {
			exPath = append(exPath, p1)
		}
		for _, p := range bps {
			exPath = append(exPath, p)
		}
		exPath = append(exPath, p2)
	}
	return exPath
}

func mustParseCave(in string) *Cave {
	cave := &Cave{
		Occ: map[grid.Point]Material{},
	}
	first := true
	for _, line := range readutil.ReadLines(in) {
		path := mustParsePath(line)
		path = expandPath(path)
		cave.Paths = append(cave.Paths, path)
		for _, p := range path {
			cave.Occ[p] = Rock
			if first {
				cave.maxY = p.Y
				cave.minX = p.X
				cave.maxX = p.X
				first = false
				continue
			}
			if p.Y > cave.maxY {
				cave.maxY = p.Y
			}
			if p.X < cave.minX {
				cave.minX = p.X
			}
			if p.X > cave.maxX {
				cave.maxX = p.X
			}
		}
	}
	if len(cave.Paths) == 0 {
		fatal("empty cave")
	}
	return cave
}

func (c *Cave) Free(p grid.Point) bool {
	if c.Floor > 0 && p.Y >= c.Floor {
		return false
	}

	_, ok := c.Occ[p]
	return !ok
}

func (c *Cave) pour() bool {
	start := grid.Pt(500, 0)
	curr := start
	for {
		if c.Floor == 0 && curr.Y >= c.maxY {
			return false
		}

		down := curr.Add(grid.Pt(0, 1))
		if c.Free(down) {
			curr = down
			continue
		}
		downLeft := curr.Add(grid.Pt(-1, 1))
		if c.Free(downLeft) {
			curr = downLeft
			continue
		}
		downRight := curr.Add(grid.Pt(1, 1))
		if c.Free(downRight) {
			curr = downRight
			continue
		}
		//
		//block
		c.Occ[curr] = Sand
		if curr.X < c.minX {
			c.minX = curr.X
		}
		if curr.X > c.maxX {
			c.maxX = curr.X
		}
		if curr.Y > c.maxY {
			c.maxY = curr.Y
		}

		if curr == start {
			return false
		}
		return true
	}
}

func part1MainFunc(in string) (int, error) {
	cave := mustParseCave(in)
	//dumpCave(cave)
	rest := 0
	for {
		r := cave.pour()
		//dumpCave(cave)
		if !r {
			break
		}
		rest++
	}
	return rest, nil
}

func part2MainFunc(in string) (int, error) {
	cave := mustParseCave(in)
	cave.Floor = cave.maxY + 2
	//dumpCave(cave)
	rest := 0
	for {
		r := cave.pour()
		//dumpCave(cave)
		if !r {
			break
		}
		rest++
	}
	//dumpCave(cave)
	return rest + 1, nil
}

func dumpCave(cave *Cave) {
	for y := 0; y <= cave.maxY; y++ {
		var sr string
		for x := cave.minX; x <= cave.maxX; x++ {
			if mat, ok := cave.Occ[grid.Pt(x, y)]; ok {
				sr += string(mat)
			} else {
				sr += "."
			}
		}
		log(sr)
	}
}
