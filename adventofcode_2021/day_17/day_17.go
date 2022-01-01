package day_17

import (
	"adventofcode_2021/errutil"
	"fmt"
)

var targetArea = area{
	xmin: 88,
	xmax: 125,
	ymin: -157,
	ymax: -103,
}

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := highestYPos(targetArea)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := part2MainFunc(targetArea)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

//

type area struct {
	xmin, xmax, ymin, ymax int
}

func (a area) contains(pt point) bool {
	return pt.x >= a.xmin &&
		pt.x <= a.xmax &&
		pt.y >= a.ymin &&
		pt.y <= a.ymax
}

type point struct {
	x, y int
}

func p(x, y int) point {
	return point{x: x, y: y}
}

type probe struct {
	pos point
	vel point
}

func (p *probe) step() {
	p.pos.x += p.vel.x
	p.pos.y += p.vel.y

	switch {
	case p.vel.x > 0:
		p.vel.x -= 1
	case p.vel.x < 0:
		p.vel.x += 1
	}
	p.vel.y -= 1
}

func (p *probe) isBeyond(a area) bool {
	return p.pos.x > a.xmax || p.pos.y < a.ymin
}

type probeResult struct {
	hitArea bool
	ymax    int
}

func startProbe(start point, vel point, target area) probeResult {
	prb := &probe{
		pos: start,
		vel: vel,
	}
	ymax := prb.pos.y
	for {
		if prb.pos.y > ymax {
			ymax = prb.pos.y
		}
		if target.contains(prb.pos) {
			return probeResult{hitArea: true, ymax: ymax}
		}
		if prb.isBeyond(target) {
			return probeResult{hitArea: false, ymax: ymax}
		}
		prb.step()
	}
}

func highestYPos(target area) (int, error) {
	velymax := 10000

	start := p(0, 0)
	velxmax := target.xmax + 1
	velymin := target.ymin - 1
	var ymax int
	first := true
	var hitCount int
	for velx := 1; velx <= velxmax; velx++ {
		hit := false
		for vely := velymin; vely <= velymax; vely++ {
			res := startProbe(start, p(velx, vely), target)
			if res.hitArea {
				hit = true
				hitCount++
				if first {
					ymax = res.ymax
					first = false
					continue
				}
				if res.ymax > ymax {
					ymax = res.ymax
				}
			} else {
				_ = hit
				// if hit {
				// 	break
				// }
			}
		}
	}
	log("hitcount = %d", hitCount)

	return ymax, nil
}

func part2MainFunc(in area) (int, error) {
	return 0, nil
}
