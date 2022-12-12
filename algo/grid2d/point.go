package grid2d

import (
	"fmt"
	"math"
)

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Pt(x, y int) Point {
	return Point{x, y}
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Point) Sub(q Point) Point {
	return Pt(p.X-q.X, p.Y-q.Y)
}

func (p Point) Add(q Point) Point {
	return Pt(p.X+q.X, p.Y+q.Y)
}

func (p Point) InnerProduct(q Point) float64 {
	return float64(p.X*q.X + p.Y*q.Y)
}

func (p Point) Norm() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p Point) ManhattenNorm() int {
	return AbsInt(p.X) + AbsInt(p.Y)
}

func (p Point) DistTo(q Point) float64 {
	return p.Sub(q).Norm()
}

func (p Point) ManhattenDistTo(q Point) int {
	return p.Sub(q).ManhattenNorm()
}
