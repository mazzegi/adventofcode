package fuel

import "math"

func ByMass(m int) int {
	f := int(math.Floor(float64(m)/3.0)) - 2
	if f < 0 {
		f = 0
	}
	return f
}

func FuelMass(m int) int {
	f := ByMass(m)
	pf := f
	for {
		ff := ByMass(pf)
		if ff <= 0 {
			return f
		}
		f += ff
		pf = ff
	}
}
