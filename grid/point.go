package grid

import (
	"fmt"
	"math"
	"slices"

	"github.com/mazzegi/adventofcode/intutil"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/vector"
)

func EpsEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 1e-10
}

func Pt(x, y int) Point {
	return Point{x, y}
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d, %d", p.X, p.Y)
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

func (p Point) DistTo(q Point) float64 {
	return p.Sub(q).Norm()
}

func (p Point) ManhattenNorm() int {
	return mathutil.Abs(p.X) + mathutil.Abs(p.Y)
}

func (p Point) ManhattenDistTo(q Point) int {
	return p.Sub(q).ManhattenNorm()
}

func (p Point) RightAngleTo(q Point) float64 {
	// a1 := p.RightAngleToXAxis()
	// a2 := q.RightAngleToXAxis()
	vp := vector.V2D(float64(p.X), float64(p.Y))
	vq := vector.V2D(float64(q.X), float64(q.Y))

	return vp.RightAngleTo(vq)
}

func (p Point) LeftAngleTo(q Point) float64 {
	// a1 := p.RightAngleToXAxis()
	// a2 := q.RightAngleToXAxis()
	vp := vector.V2D(float64(p.X), float64(p.Y))
	vq := vector.V2D(float64(q.X), float64(q.Y))

	return vp.LeftAngleTo(vq)
}

func (p Point) RightAngleToXAxis() float64 {
	return vector.V2D(float64(p.X), float64(p.Y)).RightAngleToXAxis()
}

func (p Point) IsZero() bool {
	return p.X == 0 && p.Y == 0
}

func (p Point) IsMultipleOf(q Point) bool {
	if p.IsZero() && q.IsZero() {
		return true
	}
	if p.IsZero() || q.IsZero() {
		return false
	}
	if q.X == 0 {
		return p.X == 0
	}
	if q.Y == 0 {
		return p.Y == 0
	}
	return EpsEqual(float64(p.X)/float64(q.X), float64(p.Y)/float64(q.Y))
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

func SortPoints(pts []Point) {
	slices.SortFunc(pts, func(a, b Point) int {
		if a.X < b.X {
			return -1
		} else if a.X > b.X {
			return 1
		}

		if a.Y < b.Y {
			return -1
		} else if a.Y > b.Y {
			return 1
		}

		return 0
	})
}
