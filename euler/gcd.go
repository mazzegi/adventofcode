package euler

import "github.com/mazzegi/adventofcode/mathutil"

func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	a = mathutil.Abs(a)
	b = mathutil.Abs(b)

	for b != 0 {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}
