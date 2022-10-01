package day_03

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2017/intutil"
)

func Part1() {
	res := CarrySteps(361527)
	fmt.Printf("part1: res = %d\n", res)
}

func Part2() {
	res := StressTest(361527)
	fmt.Printf("part2: res = %d\n", res)
}

func CarrySteps(square int) int {
	if square <= 0 {
		return 0
	}
	x, y := 0, 0
	dist := func() int {
		return intutil.AbsInt(x) + intutil.AbsInt(y)
	}

	dir := 'r'
	steps := 1
	n := 1
	for {
		switch dir {
		case 'r':
			if n+steps < square {
				x += steps
				n += steps
			} else {
				x += square - n
				return dist()
			}
			dir = 'u'
		case 'u':
			if n+steps < square {
				y += steps
				n += steps
			} else {
				y += square - n
				return dist()
			}
			dir = 'l'
			steps++
		case 'l':
			if n+steps < square {
				x -= steps
				n += steps
			} else {
				x -= square - n
				return dist()
			}
			dir = 'd'
		case 'd':
			if n+steps < square {
				y -= steps
				n += steps
			} else {
				y -= square - n
				return dist()
			}
			dir = 'r'
			steps++
		}
	}
}

type Point struct {
	x, y int
}

func P(x, y int) Point {
	return Point{x: x, y: y}
}

type Grid struct {
	points map[Point]int
}

func NewGrid() *Grid {
	return &Grid{
		points: map[Point]int{},
	}
}

func (g *Grid) AdjacentSum(pt Point) int {
	var sum int
	for x := pt.x - 1; x <= pt.x+1; x++ {
		for y := pt.y - 1; y <= pt.y+1; y++ {
			testpt := P(x, y)
			if testpt == pt {
				continue
			}
			if n, ok := g.points[testpt]; ok {
				sum += n
			}
		}
	}
	return sum
}

func (g *Grid) Add(x, y int, val int) {
	g.points[P(x, y)] = val
}

func StressTest(limit int) int {
	grid := NewGrid()
	x, y := 0, 0

	dir := 'r'
	steps := 1
	grid.Add(0, 0, 1)

	for {
		switch dir {
		case 'r':
			for i := 0; i < steps; i++ {
				x += 1
				p := P(x, y)
				as := grid.AdjacentSum(p)
				if as > limit {
					return as
				}
				grid.Add(x, y, as)
			}
			dir = 'u'
		case 'u':
			for i := 0; i < steps; i++ {
				y += 1
				p := P(x, y)
				as := grid.AdjacentSum(p)
				if as > limit {
					return as
				}
				grid.Add(x, y, as)
			}
			dir = 'l'
			steps++
		case 'l':
			for i := 0; i < steps; i++ {
				x -= 1
				p := P(x, y)
				as := grid.AdjacentSum(p)
				if as > limit {
					return as
				}
				grid.Add(x, y, as)
			}
			dir = 'd'
		case 'd':
			for i := 0; i < steps; i++ {
				y -= 1
				p := P(x, y)
				as := grid.AdjacentSum(p)
				if as > limit {
					return as
				}
				grid.Add(x, y, as)
			}
			dir = 'r'
			steps++
		}
	}
}
