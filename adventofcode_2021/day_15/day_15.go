package day_15

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"fmt"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := lowestTotalRisk(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d (in %s)\n", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := lowestTotalRiskDup5(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d (in %s)\n", res, time.Since(t0))
}

func parseGraph(in string) (*graph, error) {
	g := &graph{}
	lines := readutil.ReadLines(in)
	for y, line := range lines {
		gr := &graphRow{}
		for x, r := range line {
			n, err := strconv.ParseInt(string(r), 10, 8)
			if err != nil {
				return nil, errors.Wrapf(err, "parse int %q", string(r))
			}
			gr.nodes = append(gr.nodes, &node{
				value:     int(n),
				pt:        p(x, y),
				pathValue: math.MaxInt64,
			})
		}

		if len(g.rows) > 0 && (len(gr.nodes) != len(g.rows[0].nodes)) {
			return nil, errors.Errorf("inconsitent row sizes")
		}
		g.rows = append(g.rows, gr)
	}
	if len(g.rows) > 0 {
		g.ydim = len(g.rows)
		g.xdim = len(g.rows[0].nodes)
	}

	return g, nil
}

type node struct {
	value int
	pt    point

	visited   bool
	pathValue int
	prev      *node
}

func (n *node) String() string {
	return fmt.Sprintf("%d, %d (%d)", n.pt.x, n.pt.y, n.value)
}

type graphRow struct {
	nodes []*node
}

type graph struct {
	rows         []*graphRow
	xdim, ydim   int
	totalVisited int
}

func (g *graph) visit(n *node) {
	if n.visited {
		return
	}
	n.visited = true
	g.totalVisited++
}

func (g *graph) allVisited() bool {
	return g.totalVisited == g.xdim*g.ydim
}

func (g *graph) isEmpty() bool {
	if len(g.rows) == 0 {
		return true
	}
	return len(g.rows[0].nodes) == 0
}

func (g *graph) neighbours(x, y int) []point {
	var nps []point
	if x-1 >= 0 {
		nps = append(nps, p(x-1, y))
	}
	if x+1 < g.xdim {
		nps = append(nps, p(x+1, y))
	}
	if y-1 >= 0 {
		nps = append(nps, p(x, y-1))
	}
	if y+1 < g.ydim {
		nps = append(nps, p(x, y+1))
	}
	return nps
}

func (g *graph) neighbourNodes(n *node) []*node {
	nps := g.neighbours(n.pt.x, n.pt.y)
	nnodes := make([]*node, len(nps))
	for i, np := range nps {
		nnodes[i] = g.mustNode(np.x, np.y)
	}
	return nnodes
}

func (g *graph) mustNode(x, y int) *node {
	if x < 0 || x >= g.xdim || y < 0 || y >= g.ydim {
		panic("invalid pt")
	}
	return g.rows[y].nodes[x]
}

func (g *graph) start() *node {
	return g.mustNode(0, 0)
}

func (g *graph) end() *node {
	return g.mustNode(g.xdim-1, g.ydim-1)
}

func p(x, y int) point {
	return point{x: x, y: y}
}

type point struct {
	x, y int
}

type path struct {
	nodes []*node
}

func (p *path) contains(pt point) bool {
	for _, n := range p.nodes {
		if n.pt == pt {
			return true
		}
	}
	return false
}

func (p *path) value() int {
	var val int
	for _, n := range p.nodes {
		val += n.value
	}
	return val
}

func lowestTotalRisk(in string) (int, error) {
	g, err := parseGraph(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-graph")
	}
	if g.isEmpty() {
		return 0, errors.Errorf("graph is empty")
	}

	notVisitedNodeWithMinDist := func() *node {
		var cand *node
		var candDist int
		for x := 0; x < g.xdim; x++ {
			for y := 0; y < g.ydim; y++ {
				n := g.mustNode(x, y)
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
		}
		return cand
	}

	start := g.start()
	start.pathValue = 0
	for !g.allVisited() {
		n := notVisitedNodeWithMinDist()
		g.visit(n)
		nns := g.neighbourNodes(n)
		for _, nn := range nns {
			if nn.visited {
				continue
			}
			test := n.pathValue + nn.value
			if test < nn.pathValue {
				nn.pathValue = test
				nn.prev = n
			}
		}

	}

	end := g.end()
	log("path value of end: %d", end.pathValue)

	return end.pathValue, nil
}

func lowestTotalRiskDup5(in string) (int, error) {
	g, err := parseGraph(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-graph")
	}
	if g.isEmpty() {
		return 0, errors.Errorf("graph is empty")
	}

	notVisitedNodeWithMinDist := func() *node {
		var cand *node
		var candDist int
		for x := 0; x < g.xdim; x++ {
			for y := 0; y < g.ydim; y++ {
				n := g.mustNode(x, y)
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
		}
		return cand
	}

	g = dupPlus5(g)

	start := g.start()
	start.pathValue = 0
	for !g.allVisited() {
		n := notVisitedNodeWithMinDist()
		g.visit(n)
		nns := g.neighbourNodes(n)
		for _, nn := range nns {
			if nn.visited {
				continue
			}
			test := n.pathValue + nn.value
			if test < nn.pathValue {
				nn.pathValue = test
				nn.prev = n
			}
		}

	}

	end := g.end()
	log("path value of end: %d", end.pathValue)

	return end.pathValue, nil
}

func dupPlus5(g *graph) *graph {
	dg := &graph{
		xdim: 5 * g.xdim,
		ydim: 5 * g.ydim,
	}
	for y := 0; y < dg.ydim; y++ {
		row := &graphRow{
			nodes: make([]*node, dg.xdim),
		}
		for x := 0; x < dg.xdim; x++ {
			n := &node{
				pt:        p(x, y),
				prev:      nil,
				value:     0,
				visited:   false,
				pathValue: math.MaxInt64,
			}
			row.nodes[x] = n
		}

		dg.rows = append(dg.rows, row)
	}

	for dx := 0; dx < 5; dx++ {
		for dy := 0; dy < 5; dy++ {
			xoff := dx * g.xdim
			yoff := dy * g.ydim
			for x := 0; x < g.xdim; x++ {
				for y := 0; y < g.ydim; y++ {
					val := g.mustNode(x, y).value
					val += dx + dy
					if val > 9 {
						val = 1 + (val - 10)
					}
					dg.mustNode(xoff+x, yoff+y).value = val
				}
			}
		}
	}

	return dg
}

func dumpGraph(g *graph) string {
	var sl []string
	for _, row := range g.rows {
		var s string
		for _, n := range row.nodes {
			s += fmt.Sprintf("%d", n.value)
		}
		sl = append(sl, s)
	}
	return strings.Join(sl, "\n")
}
