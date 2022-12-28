package day_22

import (
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/mathutil"
)

func PointsBetween(p1, p2 grid.Point) []grid.Point {
	var bw []grid.Point
	bw = append(bw, p1)
	if p1.Y == p2.Y {
		d := mathutil.Sign(p2.X - p1.X)
		bp := p1.Add(P(d, 0))
		for {
			if bp == p2 {
				break
			}
			bw = append(bw, bp)
			bp = bp.Add(P(d, 0))
		}
	} else if p1.X == p2.X {
		d := mathutil.Sign(p2.Y - p1.Y)
		bp := p1.Add(P(0, d))
		for {
			if bp == p2 {
				break
			}
			bw = append(bw, bp)
			bp = bp.Add(P(0, d))
		}
	} else {
		fatal("cannot between diagonal points")
	}
	bw = append(bw, p2)
	return bw
}
