package vector

import (
	"fmt"
	"math"
)

type Vector3D struct {
	X, Y, Z float64
}

func V3D(x, y, z float64) Vector3D {
	return Vector3D{x, y, z}
}

func (v Vector3D) String() string {
	return fmt.Sprintf("%f, %f, %f", v.X, v.Y, v.Z)
}

func (v Vector3D) Sub(w Vector3D) Vector3D {
	return V3D(v.X-w.X, v.Y-w.Y, v.Z-w.Z)
}

func (v Vector3D) Add(w Vector3D) Vector3D {
	return V3D(v.X+w.X, v.Y+w.Y, v.Z+w.Z)
}

func (v Vector3D) Mult(f float64) Vector3D {
	return V3D(f*v.X, f*v.Y, f*v.Z)
}

func (v Vector3D) InnerProduct(w Vector3D) float64 {
	return float64(v.X*w.X + v.Y*w.Y + v.Z*w.Z)
}

func (v Vector3D) Norm() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

func (v Vector3D) DistTo(w Vector3D) float64 {
	return v.Sub(w).Norm()
}
