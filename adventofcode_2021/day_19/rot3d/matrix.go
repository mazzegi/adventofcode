package rot3d

import (
	"fmt"
	"math"
)

const (
	Rot0   float64 = 0
	Rot90  float64 = 0.5 * math.Pi
	Rot180 float64 = math.Pi
	Rot270 float64 = 1.5 * math.Pi
)

type MatrixRow [3]float64

type Matrix struct {
	Rows [3]MatrixRow
}

func XMatrix(rad float64) Matrix {
	return Matrix{
		Rows: [3]MatrixRow{
			{1, 0, 0},
			{0, math.Cos(rad), -math.Sin(rad)},
			{0, math.Sin(rad), math.Cos(rad)},
		},
	}
}

func YMatrix(rad float64) Matrix {
	return Matrix{
		Rows: [3]MatrixRow{
			{math.Cos(rad), 0, math.Sin(rad)},
			{0, 1, 0},
			{-math.Sin(rad), 0, math.Cos(rad)},
		},
	}
}

func ZMatrix(rad float64) Matrix {
	return Matrix{
		Rows: [3]MatrixRow{
			{math.Cos(rad), -math.Sin(rad), 0},
			{math.Sin(rad), math.Cos(rad), 0},
			{0, 0, 1},
		},
	}
}

type Point struct {
	X, Y, Z float64
}

func P(x, y, z float64) Point {
	return Point{x, y, z}
}

func (p Point) String() string {
	return fmt.Sprintf("%.1f, %.1f, %.1f", p.X, p.Y, p.Z)
}

func (m Matrix) Rotate(pt Point) Point {
	return Point{
		X: m.Rows[0][0]*pt.X + m.Rows[0][1]*pt.Y + m.Rows[0][2]*pt.Z,
		Y: m.Rows[1][0]*pt.X + m.Rows[1][1]*pt.Y + m.Rows[1][2]*pt.Z,
		Z: m.Rows[2][0]*pt.X + m.Rows[2][1]*pt.Y + m.Rows[2][2]*pt.Z,
	}
}
