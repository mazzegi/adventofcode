package grid

import "github.com/mazzegi/adventofcode/intutil"

func Pt(x, y int) Point {
	return Point{x, y}
}

type Point struct {
	X, Y int
}

func (p Point) Between(p1, p2 Point) bool {
	if p == p1 || p == p2 {
		return true
	}
	if p1 == p2 {
		return false
	}
	if p1.X == p2.X {
		if p.X != p1.X {
			return false
		}
		return intutil.Between(p.Y, p1.Y, p2.Y)
	}
	if p1.Y == p2.Y {
		if p.Y != p1.Y {
			return false
		}
		return intutil.Between(p.X, p1.X, p2.X)
	}

	c1 := (p2.Y - p1.Y) * (p.X - p1.X)
	c2 := (p.Y - p1.Y) * (p2.X - p1.X)
	if c1 != c2 {
		return false
	}
	alpha := float64(p.X-p1.X) / float64(p2.X-p1.X)
	if alpha < 0 || alpha > 1 {
		return false
	}
	return true
}
