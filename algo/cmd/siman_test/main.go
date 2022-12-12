package main

import (
	"fmt"
	. "math"
	"math/rand"

	"github.com/mazzegi/algo/siman"
)

type params [2]float64

func f(ps params) float64 {
	x := ps[0]
	y := ps[1]
	_ = y
	//return Pow(x, 4)*Pow(y, 4) - 2*Pow(x, 2)*y + Pow(y, 2)*x - 5*x*y + 4
	return 0.02*x*x + Sin(x*5)*Sin(x*5) + 0.02*y*y + Sin(y*4)*Sin(y*4)
	//return x*x + y*y
}

func tf(tmax uint64) siman.TempFnc {
	return func(t uint64) float64 {
		return 1 - Sqrt(float64(t)/float64(tmax))
	}
}

func nf(tmax uint64) func(params, uint64) params {
	return func(ps params, t uint64) params {
		rndX := -1 + 2*rand.Float64()
		rndY := -1 + 2*rand.Float64()
		scale := 1.0 * (1.0 - Sqrt(float64(t)/float64(tmax))*(1-0.001))
		return [2]float64{ps[0] + scale*rndX, ps[1] + scale*rndY}
	}
}

func main() {
	//rand.Seed(time.Now().UnixMilli())
	tmax := uint64(1e7)
	sol, val := siman.Minimize(f, tf(tmax), nf(tmax), params{20, 30}, tmax)
	fmt.Printf("opt: %v => %f\n", sol, val)
}
