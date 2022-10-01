package day_22

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2016/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2016/readutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	ns, err := ParseNodes(input)
	errutil.ExitOnErr(err)
	n := ViablePairs(ns)
	log("part1: pairs = %d", n)
}

func Part2() {
	ns, err := ParseNodes(input)
	errutil.ExitOnErr(err)
	//n := fewestSteps(ns)
	n := fewestStepsSimple(ns, 400, 32, 28)
	log("part2: steps = %d", n)
}

type Node struct {
	pt    point
	size  int
	used  int
	avail int
}

func (n *Node) Clone() *Node {
	cn := *n
	return &cn
}

func (n *Node) is(on *Node) bool {
	return n.pt == on.pt
}

// /dev/grid/node-x1-y16    87T   70T    17T   80%
func ParseNode(s string) (*Node, error) {
	sl := strings.Fields(s)
	if len(sl) != 5 {
		return nil, errors.Errorf("invalid number of fields in %q", s)
	}

	var n Node
	var pt point
	_, err := fmt.Sscanf(sl[0], "/dev/grid/node-x%d-y%d", &pt.x, &pt.y)
	if err != nil {
		return nil, errors.Wrapf(err, "scan xy %q", sl[0])
	}
	n.pt = pt

	_, err = fmt.Sscanf(sl[1], "%dT", &n.size)
	if err != nil {
		return nil, errors.Wrapf(err, "scan value %q", sl[1])
	}
	_, err = fmt.Sscanf(sl[2], "%dT", &n.used)
	if err != nil {
		return nil, errors.Wrapf(err, "scan value %q", sl[2])
	}
	_, err = fmt.Sscanf(sl[3], "%dT", &n.avail)
	if err != nil {
		return nil, errors.Wrapf(err, "scan value %q", sl[3])
	}

	return &n, nil
}

func ParseNodes(in string) ([]*Node, error) {
	var ns []*Node
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		n, err := ParseNode(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-node %q", line)
		}
		ns = append(ns, n)
	}
	return ns, nil
}

type NodePair struct {
	A *Node
	B *Node
}

func (p *NodePair) is(op *NodePair) bool {
	return (p.A.is(op.A) && p.B.is(op.B)) ||
		(p.A.is(op.B) && p.B.is(op.A))
}

func ViablePairs(ns []*Node) int {
	var pairs []*NodePair
	containsPair := func(op *NodePair) bool {
		for _, p := range pairs {
			if p.is(op) {
				return true
			}
		}
		return false
	}
	for _, na := range ns {
		for _, nb := range ns {
			if nb.is(na) {
				continue
			}
			if na.used == 0 {
				continue
			}
			if na.used > nb.avail {
				continue
			}
			p := &NodePair{
				A: na,
				B: nb,
			}
			if !containsPair(p) {
				pairs = append(pairs, p)
			}
		}
	}
	return len(pairs)
}

type point struct {
	x, y int
}

func p(x, y int) point {
	return point{x, y}
}

func (pt point) String() string {
	return fmt.Sprintf("%d, %d", pt.x, pt.y)
}

type NodeGrid struct {
	dataAt     point
	dataSize   int
	nodes      map[point]*Node
	xmax, ymax int
}

func (ng *NodeGrid) Clone() *NodeGrid {
	cng := &NodeGrid{
		dataAt:   ng.dataAt,
		dataSize: ng.dataSize,
		nodes:    map[point]*Node{},
		xmax:     ng.xmax,
		ymax:     ng.ymax,
	}
	for k, v := range ng.nodes {
		cng.nodes[k] = v.Clone()
	}
	return cng
}

func NewNodeGrid(ns []*Node) *NodeGrid {
	ng := &NodeGrid{
		nodes: map[point]*Node{},
	}
	var dataNode *Node
	for _, n := range ns {
		ng.nodes[n.pt] = n
		if n.pt.y == 0 {
			if dataNode == nil {
				dataNode = n
			} else if n.pt.x > dataNode.pt.x {
				dataNode = n
			}
		}

		if n.pt.x > ng.xmax {
			ng.xmax = n.pt.x
		}
		if n.pt.y > ng.ymax {
			ng.ymax = n.pt.y
		}
	}
	ng.dataAt = dataNode.pt
	ng.dataSize = dataNode.used

	return ng
}

func (ng *NodeGrid) Adjacents(n *Node) []*Node {
	var adjs []*Node

	addIfViable := func(pt point) {
		if pt.x < 0 || pt.x > ng.xmax {
			return
		}
		if pt.y < 0 || pt.y > ng.ymax {
			return
		}
		an := ng.nodes[pt]
		adjs = append(adjs, an)
	}
	addIfViable(p(n.pt.x-1, n.pt.y))
	addIfViable(p(n.pt.x, n.pt.y-1))
	addIfViable(p(n.pt.x, n.pt.y+1))
	addIfViable(p(n.pt.x+1, n.pt.y))

	return adjs
}

func (ng *NodeGrid) MustNode(pt point) *Node {
	n, ok := ng.nodes[pt]
	if !ok {
		fatal("no node at %s", pt)
	}
	return n
}

//

// func fewestSteps(ns []*Node) int {
// 	ing := NewNodeGrid(ns)

// 	cng := ing.Clone()
// 	for {
// 		if cng.dataAt == p(0, 0) {
// 			log("did it")
// 			return 1
// 		}

// 		ok := false
// 		minSteps := 0
// 		var minGrid *NodeGrid

// 		dataNode := cng.MustNode(cng.dataAt)
// 		for _, adj := range cng.Adjacents(dataNode) {

// 			cadj := cng.MustNode(adj.pt)
// 			visited := map[point]bool{}
// 			visited[cadj.pt] = true
// 			inNg := cng.Clone()
// 			tryNg, steps, tryOk := tryFree(inNg, visited, cadj, dataNode)
// 			_ = tryNg
// 			_ = steps
// 			if tryOk && (!ok || steps+1 < minSteps) {
// 				log("move data from %s  to %s", dataNode.pt, cadj.pt)

// 				tadj := tryNg.MustNode(cadj.pt)
// 				tdataNode := tryNg.MustNode(dataNode.pt)

// 				tadj.used += tdataNode.used
// 				tadj.avail -= tdataNode.used
// 				tdataNode.used = 0
// 				tdataNode.avail = tdataNode.size

// 				tryNg.dataAt = tadj.pt
// 				ok = true
// 				minSteps = steps + 1
// 				minGrid = tryNg
// 			}

// 			// //
// 			// log("move data from %s  to %s", dataNode.pt, cadj.pt)
// 			// cadj.used += dataNode.used
// 			// cadj.avail -= dataNode.used
// 			// dataNode.used = 0
// 			// dataNode.avail = dataNode.size
// 			// cng.dataAt = cadj.pt
// 			// break
// 		}
// 		cng = minGrid
// 	}
// }

// func cloneVisited(v map[point]bool) map[point]bool {
// 	cv := map[point]bool{}
// 	for pt := range v {
// 		cv[pt] = true
// 	}
// 	return cv
// }

// func tryFree(ng *NodeGrid, visited map[point]bool, n *Node, f *Node) (*NodeGrid, int, bool) {
// 	if n.used == 0 {
// 		return ng, 0, true
// 	}

// 	ok := false
// 	minSteps := 0
// 	var minGrid *NodeGrid
// 	for _, adj := range ng.Adjacents(n) {
// 		if n.used > adj.size {
// 			continue
// 		}
// 		if adj.is(f) {
// 			continue
// 		}
// 		if adj.pt == ng.dataAt {
// 			continue
// 		}
// 		if _, ok := visited[adj.pt]; ok {
// 			continue
// 		}

// 		cvisited := cloneVisited(visited)
// 		cvisited[adj.pt] = true

// 		log("try %s (%d / %d) => %s (%d / %d)", n.pt, n.used, n.size, adj.pt, adj.used, adj.size)
// 		inNg := ng.Clone()
// 		tryNg, steps, tryOk := tryFree(inNg, cvisited, adj, n)
// 		if tryOk && (!ok || steps+1 < minSteps) {
// 			log("free: move data from %s to %s", n.pt, adj.pt)
// 			tadj := tryNg.MustNode(adj.pt)
// 			tn := tryNg.MustNode(n.pt)

// 			tadj.used += tn.used
// 			tadj.avail -= tn.used
// 			tn.used = 0
// 			tn.avail = tn.size

// 			ok = true
// 			minSteps = steps + 1
// 			minGrid = tryNg
// 		}
// 	}

// 	return minGrid, minSteps, ok
// }

type state string

const (
	movable state = "."
	fixed   state = "#"
	empty   state = "_"
	data    state = "G"
)

type grid struct {
	states [][]state
}

func initGrid(ns []*Node, bigNodeMin int, xsize, ysize int) *grid {
	g := &grid{
		states: make([][]state, ysize),
	}
	for y := 0; y < ysize; y++ {
		g.states[y] = make([]state, xsize)
	}
	for _, n := range ns {
		var s state
		if n.used == 0 {
			s = empty
		} else if n.used > bigNodeMin {
			s = fixed
		} else if n.pt.y == 0 && n.pt.x == xsize-1 {
			s = data
		} else {
			s = movable
		}
		g.states[n.pt.y][n.pt.x] = s
	}

	return g
}

func (g *grid) dump() string {
	var sl []string
	for _, row := range g.states {
		var sr string
		for _, s := range row {
			sr += string(s)
		}
		sl = append(sl, sr)
	}
	return strings.Join(sl, "\n")
}

func fewestStepsSimple(ns []*Node, bigNodeMin int, xsize, ysize int) int {
	g := initGrid(ns, bigNodeMin, xsize, ysize)
	log("***\n%s", g.dump())
	return 0
}
