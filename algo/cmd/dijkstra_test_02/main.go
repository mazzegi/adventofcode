package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/algo/dijkstra"
	. "github.com/mazzegi/adventofcode/algo/grid2d"
)

const maze = `
##########################################
#S#######.................######.........#
#....####.###############.######.#######.#
####........#############........#######.#
####.######.############################.#
####...####.############################.#
######.####.#########.........##########.#
######.####.#########.#######.##########.#
###....####.......###.#######.####.......#
###.#######.#####.###.#######.#######.####
###.........###...###.#####...###.....####
###.#################.#####.#####.########
###...................#####............T##
##########################################
`

type Graph struct {
	openPoints map[Point]bool
}

func (g *Graph) Nodes() []Point {
	var ps []Point
	for p := range g.openPoints {
		ps = append(ps, p)
	}
	return ps
}

func (g *Graph) Equal(t1, t2 Point) bool {
	return t1 == t2
}

func (g *Graph) AreNeighbours(t1, t2 Point) bool {
	return t1.ManhattenDistTo(t2) == 1
}

func (g *Graph) DistanceBetween(t1, t2 Point) float64 {
	return float64(t1.ManhattenDistTo(t2))
}

func main() {
	var start, target Point
	g := &Graph{
		openPoints: map[Point]bool{},
	}

	buf := bytes.NewBufferString(maze)
	scanner := bufio.NewScanner(buf)
	rix := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		for cix, r := range line {
			p := Pt(cix, rix)
			switch r {
			case 'S':
				g.openPoints[p] = true
				start = p
			case 'T':
				g.openPoints[p] = true
				target = p
			case '.':
				g.openPoints[p] = true
			}
		}
		rix++
	}

	path, err := dijkstra.ShortestPath[Point](g, start, target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("path: %v\n", path.Nodes)
	fmt.Printf("dist: %f\n", path.Distance)
}
