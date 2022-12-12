package dijkstra

import (
	"fmt"
	"math"
)

type Graph[T any] interface {
	Nodes() []T
	Equal(t1, t2 T) bool
	AreNeighbours(t1, t2 T) bool
	DistanceBetween(t1, t2 T) float64
}

type node[T any] struct {
	value T
	dist  float64
	prev  *node[T]
}

type Path[T any] struct {
	Nodes    []T
	Distance float64
}

func ShortestPath[T any](graph Graph[T], start T, target T) (*Path[T], error) {
	// init
	var startNode *node[T]
	var targetNode *node[T]
	var dnodes []*node[T]
	for _, n := range graph.Nodes() {
		dn := &node[T]{
			value: n,
			prev:  nil,
		}
		if graph.Equal(n, start) {
			dn.dist = 0
			startNode = dn
		} else {
			dn.dist = math.Inf(1)
		}
		if graph.Equal(n, target) {
			targetNode = dn
		}
		dnodes = append(dnodes, dn)
	}

	//
	takeNodeWithMinDist := func() (*node[T], bool) {
		if len(dnodes) == 0 {
			return nil, false
		}
		var minDist float64
		var minDistNode *node[T]
		var minDistNodeIdx int
		for ix, dn := range dnodes {
			if minDistNode == nil || dn.dist < minDist {
				minDist = dn.dist
				minDistNode = dn
				minDistNodeIdx = ix
			}
		}
		dnodes = append(dnodes[:minDistNodeIdx], dnodes[minDistNodeIdx+1:]...)
		return minDistNode, true
	}

	//
	neighbours := func(dn *node[T]) []*node[T] {
		var dnns []*node[T]
		for _, dnn := range dnodes {
			if graph.AreNeighbours(dnn.value, dn.value) {
				dnns = append(dnns, dnn)
			}
		}
		return dnns
	}

	//
	for {
		dn, ok := takeNodeWithMinDist()
		if !ok {
			break
		}
		ns := neighbours(dn)
		for _, dnn := range ns {
			candDist := dn.dist + graph.DistanceBetween(dn.value, dnn.value)
			if candDist < dnn.dist {
				dnn.dist = candDist
				dnn.prev = dn
			}
		}
	}

	//
	var dist float64
	var revPathNodes []T
	tn := targetNode
	revPathNodes = append(revPathNodes, targetNode.value)
	for {
		if tn.prev == nil {
			return nil, fmt.Errorf("no prev for node %v", tn.value)
		}
		dist += tn.dist
		tn = tn.prev
		revPathNodes = append(revPathNodes, tn.value)
		if tn == startNode {
			break
		}
	}

	//reverse path
	path := &Path[T]{
		Distance: dist,
		Nodes:    make([]T, len(revPathNodes)),
	}
	for i := 0; i < len(path.Nodes); i++ {
		path.Nodes[i] = revPathNodes[len(revPathNodes)-1-i]
	}
	return path, nil
}
