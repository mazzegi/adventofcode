package day_24

import (
	"adventofcode_2016/errutil"
	"adventofcode_2016/readutil"
	"fmt"
	"sort"
	"strconv"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func fatalOnErr(err error) {
	if err == nil {
		return
	}
	fatal("%v", err)
}

func Part1() {
	res, err := fewestSteps(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
func mustParseGrid(in string) *grid {
	g := &grid{
		walls:   map[point]bool{},
		numbers: map[point]int{},
	}
	lines := readutil.ReadLines(in)
	for y, line := range lines {
		if y == 0 {
			g.xsize = len(line)
		}
		for x, r := range line {
			switch r {
			case '.':
			case '#':
				g.walls[point{x, y}] = true
			default:
				n, err := strconv.ParseInt(string(r), 10, 8)
				fatalOnErr(err)
				g.numbers[point{x, y}] = int(n)
			}
		}
	}
	g.ysize = len(lines)

	return g
}

type point struct {
	x, y int
}

func (pt point) String() string {
	return fmt.Sprintf("%d, %d", pt.x, pt.y)
}

func (pt point) hash() string {
	return fmt.Sprintf("%d:%d", pt.x, pt.y)
}

func (pt point) less(opt point) bool {
	if pt.x == opt.x {
		return pt.y < opt.y
	}
	return pt.x < opt.x
}

type grid struct {
	walls        map[point]bool
	numbers      map[point]int
	xsize, ysize int
}

func (g *grid) adjacents(pt point) []point {
	var adjs []point

	addIfValid := func(apt point) {
		if apt.x < 0 || apt.x >= g.xsize || apt.y < 0 || apt.y >= g.ysize {
			return
		}
		if _, ok := g.walls[apt]; ok {
			return
		}
		adjs = append(adjs, apt)
	}

	addIfValid(point{pt.x, pt.y - 1})
	addIfValid(point{pt.x, pt.y + 1})
	addIfValid(point{pt.x - 1, pt.y})
	addIfValid(point{pt.x + 1, pt.y})

	return adjs
}

func (g *grid) findPositionOf(num int) (point, bool) {
	for pt, pnum := range g.numbers {
		if pnum == num {
			return pt, true
		}
	}
	return point{}, false
}

func (g *grid) mustFindPositionOf(num int) point {
	if pt, ok := g.findPositionOf(num); ok {
		return pt
	}
	fatal("didn't find position of %d", num)
	return point{}
}

//
type pointSet map[point]bool

func (ps pointSet) contains(pt point) bool {
	_, ok := ps[pt]
	return ok
}

func (ps pointSet) add(pt point) {
	ps[pt] = true
}

func (ps pointSet) clone() pointSet {
	cps := pointSet{}
	for pt := range ps {
		cps.add(pt)
	}
	return cps
}

func (ps pointSet) hash() string {
	pts := make([]point, len(ps))
	i := 0
	for pt := range ps {
		pts[i] = pt
		i++
	}
	sort.Slice(pts, func(i, j int) bool {
		return pts[i].less(pts[j])
	})
	var s string
	for _, pt := range pts {
		s += pt.hash() + "|"
	}
	return s
}

//
type numberSet map[int]bool

func (ns numberSet) contains(n int) bool {
	_, ok := ns[n]
	return ok
}

func (ns numberSet) add(n int) {
	ns[n] = true
}

func (ns numberSet) clone() numberSet {
	cns := numberSet{}
	for n := range ns {
		cns.add(n)
	}
	return cns
}

func (ns numberSet) remove(n int) {
	delete(ns, n)
}

func (ns numberSet) empty() bool {
	return len(ns) == 0
}

func (ns numberSet) hash() string {
	nsl := make([]int, len(ns))
	i := 0
	for n := range ns {
		nsl[i] = n
		i++
	}
	sort.Ints(nsl)
	var s string
	for _, n := range nsl {
		s += strconv.FormatInt(int64(n), 10)
	}
	return s
}

//

// func fewestSteps(in string) (int, error) {
// 	g := mustParseGrid(in)

// 	numbersLeft := numberSet{}
// 	for _, n := range g.numbers {
// 		numbersLeft.add(n)
// 	}
// 	numbersLeft.remove(0)
// 	pt0 := g.mustFindPositionOf(0)
// 	visited := pointSet{pt0: true}

// 	cache := map[string]cacheResult{}
// 	steps, ok := walk(cache, g, visited, numbersLeft, pt0)
// 	if !ok {
// 		fatal("didn't find way")
// 	}

// 	return steps, nil
// }

func fewestSteps(in string) (int, error) {
	g := mustParseGrid(in)

	steps := fewestStepsBetween(g, 0, 5)

	return steps, nil
}

func fewestStepsBetween(g *grid, num1, num2 int) int {
	curr := g.mustFindPositionOf(num1)
	dest := g.mustFindPositionOf(num2)

	cache := map[string]cacheResult{}
	visited := pointSet{curr: true}
	steps, ok := walkShortestTo(0, cache, g, visited, curr, dest)
	if !ok {
		fatal("walk from %s to %s", curr, dest)
	}

	return steps
}

func walkShortestTo(level int, cache map[string]cacheResult, g *grid, visited pointSet, curr point, dest point) (int, bool) {
	minSteps := 0
	log("%d: try %s (visit = %d)", level, curr, len(visited))
	found := false
	adjs := g.adjacents(curr)
	for _, adj := range adjs {
		if visited.contains(adj) {
			continue
		}
		if adj == dest {
			return 1, true
		}

		cvisited := visited.clone()
		cvisited.add(adj)

		// hash := makeCacheKey(visited, adj, dest)
		// res, ok := cache[hash]
		// if !ok {
		// 	subSteps, subOk := walkShortestTo(level+1, cache, g, cvisited, adj, dest)
		// 	res.steps = subSteps
		// 	res.ok = subOk
		// 	cache[hash] = res
		// } else {
		// 	log("cache hit")
		// }
		var res cacheResult

		subSteps, subOk := walkShortestTo(level+1, cache, g, cvisited, adj, dest)
		res.steps = subSteps
		res.ok = subOk

		if !res.ok {
			continue
		}
		if found && (res.steps+1 >= minSteps) {
			continue
		}

		minSteps = res.steps + 1
		found = true
	}

	return minSteps, found
}

func makeCacheKey(vs pointSet, pt1 point, pt2 point) string {
	return vs.hash() + "-" + pt1.hash() + "-" + pt2.hash()
}

type cacheResult struct {
	steps int
	ok    bool
}

//

// func makeCacheKey(vs pointSet, ns numberSet, pt point) string {
// 	return vs.hash() + "-" + ns.hash() + "-" + pt.hash()
// }

// type cacheResult struct {
// 	steps int
// 	ok    bool
// }

// func walk(cache map[string]cacheResult, g *grid, visited pointSet, left numberSet, curr point) (int, bool) {
// 	minSteps := 0
// 	found := false
// 	adjs := g.adjacents(curr)
// 	for _, adj := range adjs {
// 		if visited.contains(adj) {
// 			continue
// 		}

// 		cvisited := visited.clone()
// 		cvisited.add(adj)
// 		cleft := left.clone()

// 		if num, ok := g.numbers[adj]; ok && cleft.contains(num) {
// 			log("found %d", num)
// 			cleft.remove(num)
// 			if cleft.empty() {
// 				log("found path %d", num)
// 				return 1, true
// 			}
// 			// reset visited
// 			cvisited = pointSet{adj: true}
// 		}

// 		hash := makeCacheKey(visited, cleft, adj)
// 		res, ok := cache[hash]
// 		if !ok {
// 			subSteps, subOk := walk(cache, g, cvisited, cleft, adj)
// 			res.steps = subSteps
// 			res.ok = subOk
// 			cache[hash] = res
// 		}

// 		if !res.ok {
// 			continue
// 		}
// 		if found && (res.steps+1 >= minSteps) {
// 			continue
// 		}

// 		minSteps = res.steps + 1
// 		found = true

// 	}

// 	return minSteps, found
// }

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
