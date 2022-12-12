package siman

import (
	"math"
	"math/rand"
)

// Temp decreasing func
type TempFnc func(t uint64) float64

// where D represents the solution-space
func Minimize[D any](fnc func(D) float64, tempFnc TempFnc, neighbourFnc func(D, uint64) D, start D, tmax uint64) (D, float64) {
	t := uint64(0)
	xopt := start
	fxopt := fnc(xopt)
	x := start
	fx := fxopt
	//stepsBeforeAnneal := 20
	stepsBeforeAnneal := 1
	for t <= tmax {
		for i := 0; i < stepsBeforeAnneal; i++ {
			T := tempFnc(t)
			y := neighbourFnc(x, t)
			fy := fnc(y)
			if fy <= fx {
				x = y
				fx = fy
			} else {
				ptest := math.Exp(-(fy - fx) / T)
				rnd := rand.Float64()
				if rnd <= ptest {
					//accept
					x = y
					fx = fy
				}
			}
			if fx < fxopt {
				xopt = x
				fxopt = fx
			}
		}
		t++
	}
	return xopt, fxopt
}
