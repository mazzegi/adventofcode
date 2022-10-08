package grid

import "fmt"

func P3D(x, y, z int) Point3D {
	return Point3D{x, y, z}
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("%d, %d, %d", p.X, p.Y, p.Z)
}

func (p Point3D) Sub(q Point3D) Point3D {
	return P3D(p.X-q.X, p.Y-q.Y, p.Z-q.Z)
}

func (p Point3D) Add(q Point3D) Point3D {
	return P3D(p.X+q.X, p.Y+q.Y, p.Z+q.Z)
}
