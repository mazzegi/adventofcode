package main

import (
	"fmt"

	"github.com/mazzegi/algo/dijkstra"
	. "github.com/mazzegi/algo/grid2d"
)

/*
				|---(10,10)------|			|---------(20,10)T
				|				 |			|
				|				 |--(15,8)--
				|						|
   |---------(5,5)					(15,5)
   |									|
   |									|
(1,1)S--------------(10,1)--------------

*/

type Graph struct {
	edges []Edge
}

func (g *Graph) init() {

}

func (g *Graph) Nodes() []Pointf {
	var ps []Pointf
	m := map[Pointf]bool{}
	for _, e := range g.edges {
		if _, ok := m[e.P1]; !ok {
			m[e.P1] = true
			ps = append(ps, e.P1)
		}
		if _, ok := m[e.P2]; !ok {
			m[e.P2] = true
			ps = append(ps, e.P2)
		}
	}
	return ps
}

func (g *Graph) Equal(t1, t2 Pointf) bool {
	return t1 == t2
}

func (g *Graph) AreNeighbours(t1, t2 Pointf) bool {
	for _, e := range g.edges {
		if (e.P1 == t1 && e.P2 == t2) ||
			(e.P1 == t2 && e.P2 == t1) {
			return true
		}
	}
	return false
}

func (g *Graph) DistanceBetween(t1, t2 Pointf) float64 {
	return t1.DistTo(t2)
}

func main() {

	p1 := Ptf(1, 1)
	p2 := Ptf(5, 5)
	p3 := Ptf(10, 1)
	p4 := Ptf(10, 10)
	p5 := Ptf(15, 5)
	p6 := Ptf(15, 8)
	p7 := Ptf(20, 10)

	g := &Graph{
		edges: []Edge{
			E(p1, p2),
			E(p1, p3),
			E(p2, p4),
			E(p4, p6),
			E(p6, p7),
			E(p6, p5),
			E(p3, p5),
			E(p5, p6),
		},
	}
	g.init()

	path, err := dijkstra.ShortestPath[Pointf](g, p1, p7)
	if err != nil {
		panic(err)
	}
	fmt.Printf("path: %v\n", path.Nodes)
	fmt.Printf("dist: %f\n", path.Distance)
}
