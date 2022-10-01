package day_24

import (
	"fmt"
	"math"
	"strconv"

	"github.com/mazzegi/adventofcode/adventofcode_2016/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2016/readutil"
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

const skip1 = true

func Part1() {
	if skip1 {
		return
	}
	res, err := fewestSteps(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := fewestStepsWithReturn(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func mustParseGrid(in string) *dijkstraGrid {
	g := &dijkstraGrid{
		walls:   map[point]bool{},
		numbers: map[point]int{},
		nodes:   map[point]*node{},
	}
	lines := readutil.ReadLines(in)
	for y, line := range lines {
		if y == 0 {
			g.xsize = len(line)
		}
		for x, r := range line {
			pt := point{x, y}
			switch r {
			case '.':
				g.nodes[pt] = &node{
					pt:        pt,
					pathValue: math.MaxInt64,
				}
			case '#':
				g.walls[pt] = true
			default:
				n, err := strconv.ParseInt(string(r), 10, 8)
				fatalOnErr(err)
				g.numbers[pt] = int(n)
				g.nodes[pt] = &node{
					pt:        pt,
					pathValue: math.MaxInt64,
				}
			}
		}
	}
	g.ysize = len(lines)
	g.numNodes = len(g.nodes)

	return g
}

type point struct {
	x, y int
}

type node struct {
	pt        point
	visited   bool
	pathValue int
	prev      *node
}

func (pt point) String() string {
	return fmt.Sprintf("%d, %d", pt.x, pt.y)
}

type dijkstraGrid struct {
	walls        map[point]bool
	numbers      map[point]int
	xsize, ysize int
	nodes        map[point]*node
	numNodes     int
	numVisited   int
}

func (g *dijkstraGrid) resetNodes() {
	for _, n := range g.nodes {
		n.visited = false
		n.pathValue = math.MaxInt64
		n.prev = nil
	}
	g.numVisited = 0
}

func (g *dijkstraGrid) adjacents(pt point) []point {
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

func (g *dijkstraGrid) findPositionOf(num int) (point, bool) {
	for pt, pnum := range g.numbers {
		if pnum == num {
			return pt, true
		}
	}
	return point{}, false
}

func (g *dijkstraGrid) mustFindPositionOf(num int) point {
	if pt, ok := g.findPositionOf(num); ok {
		return pt
	}
	fatal("didn't find position of %d", num)
	return point{}
}

func (g *dijkstraGrid) mustNode(pt point) *node {
	n, ok := g.nodes[pt]
	if !ok {
		fatal("no node at %s", pt)
	}
	return n
}

func (g *dijkstraGrid) visit(n *node) {
	if n.visited {
		return
	}
	n.visited = true
	g.numVisited++
}

func (g *dijkstraGrid) allVisited() bool {
	return g.numVisited == g.numNodes
}

func (g *dijkstraGrid) maxNum() int {
	max := 0
	for _, num := range g.numbers {
		if num > max {
			max = num
		}
	}
	return max
}

//

func fewestSteps(in string) (int, error) {
	g := mustParseGrid(in)

	min := 1
	max := g.maxNum()

	//search all permutations of [1...max]
	test := make([]int, max)
	for i := 0; i < len(test); i++ {
		test[i] = min
	}

	containsDupl := func(ns []int) bool {
		m := map[int]bool{}
		for _, n := range ns {
			if _, ok := m[n]; ok {
				return true
			}
			m[n] = true
		}
		return false
	}

	next := func() bool {
		for i := 0; i < len(test); i++ {
			if test[i] < max {
				test[i]++
				return true
			} else {
				test[i] = min
			}
		}
		return false
	}

	type pair struct {
		a, b int
	}
	cache := map[pair]int{}

	minSteps := 0
	for {
		ok := next()
		if !ok {
			break
		}
		if containsDupl(test) {
			continue
		}
		currSteps := 0
		prevNum := 0
		for _, num := range test {
			var subSteps int
			if s, ok := cache[pair{prevNum, num}]; ok {
				subSteps = s
			} else {
				subSteps = fewestStepsBetween(g, prevNum, num)
				cache[pair{prevNum, num}] = subSteps
			}

			currSteps += subSteps
			prevNum = num
		}
		if minSteps == 0 || currSteps < minSteps {
			minSteps = currSteps
		}
		log("%v => %d", test, currSteps)
	}

	return minSteps, nil
}

func fewestStepsBetween(g *dijkstraGrid, num1, num2 int) int {
	g.resetNodes()
	curr := g.mustFindPositionOf(num1)
	dest := g.mustFindPositionOf(num2)

	notVisitedNodeWithMinDist := func() *node {
		var cand *node
		var candDist int
		for _, n := range g.nodes {
			if n.visited {
				continue
			}
			if cand == nil {
				cand = n
				candDist = n.pathValue
				continue
			}
			if n.pathValue < candDist {
				cand = n
				candDist = n.pathValue
			}
		}
		return cand
	}

	steps := 0
	start := g.mustNode(curr)
	start.pathValue = 0
	for !g.allVisited() {
		n := notVisitedNodeWithMinDist()
		g.visit(n)

		adjs := g.adjacents(n.pt)
		for _, adj := range adjs {
			if adj == dest {
				return n.pathValue + 1
			}

			an := g.mustNode(adj)
			if an.visited {
				continue
			}
			test := n.pathValue + 1
			if test < an.pathValue {
				an.pathValue = test
				an.prev = n
			}
		}
	}

	return steps
}

func fewestStepsWithReturn(in string) (int, error) {
	g := mustParseGrid(in)

	min := 1
	max := g.maxNum()

	//search all permutations of [1...max]
	test := make([]int, max)
	for i := 0; i < len(test); i++ {
		test[i] = min
	}

	containsDupl := func(ns []int) bool {
		m := map[int]bool{}
		for _, n := range ns {
			if _, ok := m[n]; ok {
				return true
			}
			m[n] = true
		}
		return false
	}

	next := func() bool {
		for i := 0; i < len(test); i++ {
			if test[i] < max {
				test[i]++
				return true
			} else {
				test[i] = min
			}
		}
		return false
	}

	type pair struct {
		a, b int
	}
	cache := map[pair]int{}

	cloneTestWith0 := func() []int {
		ct := make([]int, len(test))
		copy(ct, test)
		ct = append(ct, 0)
		return ct
	}

	minSteps := 0
	for {
		ok := next()
		if !ok {
			break
		}
		if containsDupl(test) {
			continue
		}
		ctest0 := cloneTestWith0()

		currSteps := 0
		prevNum := 0
		for _, num := range ctest0 {
			var subSteps int
			if s, ok := cache[pair{prevNum, num}]; ok {
				subSteps = s
			} else {
				subSteps = fewestStepsBetween(g, prevNum, num)
				cache[pair{prevNum, num}] = subSteps
			}

			currSteps += subSteps
			prevNum = num
		}
		if minSteps == 0 || currSteps < minSteps {
			minSteps = currSteps
		}
		log("%v => %d", ctest0, currSteps)
	}

	return minSteps, nil
}
