package day_09

import (
	"fmt"
	"math"
	"sort"

	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
)

func gpDist(p1, p2 grid.GridPoint) float64 {
	return math.Sqrt(float64((p1.Col-p2.Col)*(p1.Col-p2.Col) + (p1.Row-p2.Row)*(p1.Row-p2.Row)))
}

func MinDistOfInputPoints() (grid.GridPoint, grid.GridPoint, float64) {
	var redPoints []grid.GridPoint
	lines := readutil.ReadLines(input)
	for _, line := range lines {
		var pt grid.GridPoint
		fmt.Sscanf(line, "%d,%d", &pt.Col, &pt.Row)
		redPoints = append(redPoints, pt)
	}
	var min float64
	var minP1, minP2 grid.GridPoint
	first := true
	for _, p1 := range redPoints {
		for _, p2 := range redPoints {
			if p1 == p2 {
				continue
			}
			dist := gpDist(p1, p2)
			if first {
				first = false
				min = dist
				minP1, minP2 = p1, p2
			} else if dist < min {
				min = dist
				minP1, minP2 = p1, p2
			}
		}
	}

	return minP1, minP2, min
}

func CompressInput() (compressedEdge []grid.GridPoint, rowCompressions map[int]int, colCompression map[int]int) {
	var redPoints []grid.GridPoint
	lines := readutil.ReadLines(input)
	for _, line := range lines {
		var pt grid.GridPoint
		fmt.Sscanf(line, "%d,%d", &pt.Col, &pt.Row)
		redPoints = append(redPoints, pt)
	}
	var xVals []int
	var yVals []int
	for _, p := range redPoints {
		xVals = append(xVals, p.Col)
		yVals = append(yVals, p.Row)
	}
	sort.Ints(xVals)
	sort.Ints(yVals)

	//__
	var todo = "TODO"
	_ = todo

	return
}
