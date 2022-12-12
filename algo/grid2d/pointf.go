package grid2d

import (
	"fmt"
	"math"
)

func Ptf(x, y float64) Pointf {
	return Pointf{x, y}
}

type Pointf struct {
	X, Y float64
}

func (p Pointf) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}

func (p Pointf) Sub(q Pointf) Pointf {
	return Ptf(p.X-q.X, p.Y-q.Y)
}

func (p Pointf) Add(q Pointf) Pointf {
	return Ptf(p.X+q.X, p.Y+q.Y)
}

func (p Pointf) InnerProduct(q Pointf) float64 {
	return float64(p.X*q.X + p.Y*q.Y)
}

func (p Pointf) Norm() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p Pointf) DistTo(q Pointf) float64 {
	return p.Sub(q).Norm()
}
