package day_13

import (
	"fmt"
	"math/bits"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2016/errutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

const skip1 = true

func Part1() {
	if skip1 {
		return
	}
	res, err := minSteps(1352, p(31, 39))
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := locsIn50(1352)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type point struct {
	x, y int
}

func (pt point) String() string {
	return fmt.Sprintf("%d,%d", pt.x, pt.y)
}

func p(x, y int) point {
	return point{x, y}
}

func isWall(x, y int, magic int) bool {
	if x < 0 || y < 0 {
		fatal("invalid coordinates")
	}

	n := uint(x*x + 3*x + 2*x*y + y + y*y + magic)
	oc := bits.OnesCount(n)
	return oc%2 == 1
}

type maze struct {
	magic int
	walls map[point]bool
}

func (m *maze) isWall(pt point) bool {
	if w, ok := m.walls[pt]; ok {
		return w
	}
	is := isWall(pt.x, pt.y, m.magic)
	m.walls[pt] = is
	return is
}

func makeMaze(magic int) *maze {
	m := &maze{
		magic: magic,
		walls: map[point]bool{},
	}
	return m
}

func (m *maze) dump(xmax, ymax int) string {
	var sl []string
	for y := 0; y <= ymax; y++ {
		var sr string
		for x := 0; x <= xmax; x++ {
			if m.isWall(p(x, y)) {
				sr += "#"
			} else {
				sr += "."
			}
		}
		sl = append(sl, sr)
	}

	return strings.Join(sl, "\n")
}

func minSteps(magic int, dest point) (int, error) {
	m := makeMaze(magic)
	curr := p(1, 1)

	route, ok := findShortest(m, map[point]bool{}, curr, dest)
	if !ok {
		log("found no route")
	}

	return route, nil
}

func cloneVisited(v map[point]bool) map[point]bool {
	cv := map[point]bool{}
	for k, v := range v {
		cv[k] = v
	}
	return cv
}

func findShortest(m *maze, visited map[point]bool, curr point, dest point) (int, bool) {
	var adjs []point
	testAndAddAdj := func(pt point) {
		if pt.x < 0 || pt.y < 0 {
			return
		}

		if _, ok := visited[pt]; ok {
			return
		}
		if m.isWall(pt) {
			return
		}
		adjs = append(adjs, pt)
	}

	testAndAddAdj(p(curr.x, curr.y-1))
	testAndAddAdj(p(curr.x, curr.y+1))
	testAndAddAdj(p(curr.x-1, curr.y))
	testAndAddAdj(p(curr.x+1, curr.y))

	found := false
	var minRoute int
	for _, adj := range adjs {
		if adj == dest {
			return 1, true
		}

		cvisited := cloneVisited(visited)
		cvisited[adj] = true

		route, ok := findShortest(m, cvisited, adj, dest)
		if ok {
			if !found {
				found = true
				minRoute = route + 1
			} else if route+1 < minRoute {
				minRoute = route + 1
			}
		}
	}

	return minRoute, found
}

type pointSteps struct {
	pt    point
	steps int
}

func locsIn50(magic int) (int, error) {
	m := makeMaze(magic)
	curr := p(1, 1)

	cache := map[pointSteps]map[point]bool{}

	reached := testAny(cache, m, curr, 50)
	total := len(reached)

	return total, nil
}

func testAny(cache map[pointSteps]map[point]bool, m *maze, curr point, steps int) map[point]bool {
	var adjs []point
	testAndAddAdj := func(pt point) {
		if pt.x < 0 || pt.y < 0 {
			return
		}
		if m.isWall(pt) {
			return
		}
		adjs = append(adjs, pt)
	}
	testAndAddAdj(p(curr.x, curr.y-1))
	testAndAddAdj(p(curr.x, curr.y+1))
	testAndAddAdj(p(curr.x-1, curr.y))
	testAndAddAdj(p(curr.x+1, curr.y))

	reached := map[point]bool{}
	for _, adj := range adjs {
		leftSteps := steps - 1
		reached[adj] = true
		if leftSteps == 0 {
			continue
		}

		var subReached map[point]bool
		pss := pointSteps{
			adj,
			leftSteps,
		}
		if cr, ok := cache[pss]; ok {
			subReached = cr
		} else {
			subReached = testAny(cache, m, adj, leftSteps)
			cache[pss] = subReached
		}

		for pt := range subReached {
			reached[pt] = true
		}
	}
	return reached
}
