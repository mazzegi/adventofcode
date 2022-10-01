package day_05

import (
	"fmt"
	"math"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func Part1() {
	cnt, err := OverlappingCount(input, LinePointsHV)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: overlapping = %d\n", cnt)
}

func Part2() {
	cnt, err := OverlappingCount(input, LinePointsHVD)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: overlapping = %d\n", cnt)
}

type Point struct {
	X, Y int
}

func (p Point) Len() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

type Line struct {
	P1, P2 Point
}

func Inc(a, b int) int {
	if a <= b {
		return 1
	}
	return -1
}

func LinePointsHV(l Line) []Point {
	var ps []Point

	if l.P1.X == l.P2.X {
		inc := Inc(l.P1.Y, l.P2.Y)
		y := l.P1.Y
		for {
			ps = append(ps, Point{
				X: l.P1.X,
				Y: y,
			})
			if y == l.P2.Y {
				return ps
			}
			y += inc
		}

	} else if l.P1.Y == l.P2.Y {
		inc := Inc(l.P1.X, l.P2.X)
		x := l.P1.X
		for {
			ps = append(ps, Point{
				X: x,
				Y: l.P1.Y,
			})
			if x == l.P2.X {
				return ps
			}
			x += inc
		}
	}

	return ps
}

func LinePointsHVD(l Line) []Point {
	if l.P1.X == l.P2.X || l.P1.Y == l.P2.Y {
		return LinePointsHV(l)
	}
	var ps []Point

	incX := Inc(l.P1.X, l.P2.X)
	incY := Inc(l.P1.Y, l.P2.Y)
	x := l.P1.X
	y := l.P1.Y
	for {
		ps = append(ps, Point{
			X: x,
			Y: y,
		})
		if x == l.P2.X {
			return ps
		}
		x += incX
		y += incY
	}
}

func ParseLines(in string) ([]Line, error) {
	var ls []Line
	inlines := readutil.ReadLines(in)
	for _, inline := range inlines {
		var l Line
		_, err := fmt.Sscanf(inline, "%d,%d -> %d,%d", &l.P1.X, &l.P1.Y, &l.P2.X, &l.P2.Y)
		if err != nil {
			return nil, errors.Wrapf(err, "scan-line %q", inline)
		}
		ls = append(ls, l)
	}
	return ls, nil
}

func OverlappingCount(in string, linePointsFunc func(l Line) []Point) (int, error) {
	lines, err := ParseLines(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-lines")
	}

	om := map[Point]int{}
	for _, line := range lines {
		for _, p := range linePointsFunc(line) {
			om[p]++
		}
	}

	var ocnt int
	for _, cnt := range om {
		if cnt >= 2 {
			ocnt++
		}
	}

	return ocnt, nil
}
