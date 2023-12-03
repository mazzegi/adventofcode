package mathutil

import "math"

func FloatsEqual(f1, f2 float64, eps float64) bool {
	return math.Abs(f1-f2) <= eps
}
