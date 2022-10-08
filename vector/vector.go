package vector

import (
	"fmt"
	"math"
)

func EpsEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 1e-10
}

type Vector2D struct {
	X, Y float64
}

func V2D(x, y float64) Vector2D {
	return Vector2D{x, y}
}

func (v Vector2D) String() string {
	return fmt.Sprintf("%f, %f", v.X, v.Y)
}

func (v Vector2D) Sub(w Vector2D) Vector2D {
	return V2D(v.X-w.X, v.Y-w.Y)
}

func (v Vector2D) Add(w Vector2D) Vector2D {
	return V2D(v.X+w.X, v.Y+w.Y)
}

func (v Vector2D) Mult(f float64) Vector2D {
	return V2D(f*v.X, f*v.Y)
}

func (v Vector2D) InnerProduct(w Vector2D) float64 {
	return float64(v.X*w.X + v.Y*w.Y)
}

func (v Vector2D) Norm() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vector2D) DistTo(w Vector2D) float64 {
	return v.Sub(w).Norm()
}

func (v Vector2D) RotateBy(a float64) Vector2D {
	return V2D(v.X*math.Cos(a)-v.Y*math.Sin(a), v.X*math.Sin(a)+v.Y*math.Cos(a))
}

func (v Vector2D) RightAngleTo(w Vector2D) float64 {
	//av := v.RightAngleToXAxis()
	aw := w.RightAngleToXAxis()
	//rotate both, so that w lies on x -axis
	rv := v.RotateBy(-aw)
	//rw := w.RotateBy(aw)
	return rv.RightAngleToXAxis()
}

func (v Vector2D) LeftAngleTo(w Vector2D) float64 {
	return 2*math.Pi - v.RightAngleTo(w)
}

func (v Vector2D) RightAngleToXAxis() float64 {
	nv := v.Norm()
	if EpsEqual(nv, 0) {
		return 0
	}
	vu := v.Mult(1 / nv)
	switch {
	case vu.X >= 0 && vu.Y >= 0:
		return math.Asin(vu.Y)
	case vu.X >= 0 && vu.Y < 0:
		return 2*math.Pi + math.Asin(vu.Y)
	case vu.X < 0 && vu.Y < 0:
		return math.Pi - math.Asin(vu.Y)
	case vu.X < 0 && vu.Y >= 0:
		return math.Pi - math.Asin(vu.Y)
	default:
		panic("oops - unhandled case")
	}
}
