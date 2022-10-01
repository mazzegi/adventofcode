package day_22

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/intutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := onInRegionAfterReboot(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := onAfterReboot(input, zeroCube())
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type step struct {
	op     string
	x0, x1 int
	y0, y1 int
	z0, z1 int
}

func parseStep(in string) (step, error) {
	var s step
	_, err := fmt.Sscanf(in, "%s x=%d..%d,y=%d..%d,z=%d..%d", &s.op,
		&s.x0, &s.x1, &s.y0, &s.y1, &s.z0, &s.z1)
	if err != nil {
		return step{}, errors.Wrapf(err, "scan step %q", in)
	}
	return s, nil
}

func parseSteps(in string) ([]step, error) {
	var steps []step
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		s, err := parseStep(line)
		if err != nil {
			return nil, errors.Wrap(err, "parse-step")
		}
		steps = append(steps, s)
	}
	if len(steps) == 0 {
		return nil, errors.Errorf("no data")
	}
	return steps, nil
}

//

type point struct {
	x, y, z int
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.x, p.y, p.z)
}

func p(x, y, z int) point {
	return point{x: x, y: y, z: z}
}

func pointLess(p1, p2 point) bool {
	if p1.x < p2.x {
		return true
	}
	if p1.x > p2.x {
		return false
	}
	if p1.y < p2.y {
		return true
	}
	if p1.y > p2.y {
		return false
	}
	return p1.z < p2.z
}

//

type grid struct {
	points map[point]bool
}

func newGrid() *grid {
	return &grid{
		points: map[point]bool{},
	}
}

func (g *grid) onPoints(bounds cube) int {
	var count int
	for p, v := range g.points {
		if v && bounds.containsPoint(p) {
			count++
		}
	}
	return count
}

func (g *grid) set(cb cube, v bool) {
	for x := cb.p0.x; x <= cb.p1.x; x++ {
		for y := cb.p0.y; y <= cb.p1.y; y++ {
			for z := cb.p0.z; z <= cb.p1.z; z++ {
				g.points[p(x, y, z)] = v
			}
		}
	}
}

func onInRegionAfterReboot(in string) (int, error) {
	steps, err := parseSteps(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-steps")
	}

	bounds := makeCube(p(-50, -50, -50), p(50, 50, 50))
	g := newGrid()
	for i, step := range steps {
		v := step.op == "on"
		p0 := p(step.x0, step.y0, step.z0)
		p1 := p(step.x1, step.y1, step.z1)
		if !bounds.containsPoint(p0) && !bounds.containsPoint(p1) {
			continue
		}

		cb := makeCube(p0, p1)
		g.set(cb, v)
		log("after step %d: on: %d", i, g.onPoints(bounds))
	}

	on := g.onPoints(bounds)

	return on, nil
}

// Part 2

type cube struct {
	p0 point
	p1 point
}

func zeroCube() cube {
	return makeCube(
		p(0, 0, 0),
		p(-1, -1, -1),
	)
}

func (c cube) String() string {
	if c.isEmpty() {
		return "<empty>"
	}
	return fmt.Sprintf("[%s => %s]", c.p0, c.p1)
}

func (c cube) isEmpty() bool {
	return pointLess(c.p1, c.p0)
}

func makeCube(p0, p1 point) cube {
	return cube{p0: p0, p1: p1}
}

func (c cube) containsPoint(pt point) bool {
	if c.isEmpty() {
		return false
	}
	return pt.x >= c.p0.x && pt.x <= c.p1.x &&
		pt.y >= c.p0.y && pt.y <= c.p1.y &&
		pt.z >= c.p0.z && pt.z <= c.p1.z
}

func (c cube) volume() int {
	if c.isEmpty() {
		return 0
	}
	return (c.p1.x - c.p0.x + 1) *
		(c.p1.y - c.p0.y + 1) *
		(c.p1.z - c.p0.z + 1)
}

/*
on x=-49..-1,y=-11..42,z=-10..38

off x=26..39,y=40..50,z=-2..11
*/

func cubeIntersection(c1, c2 cube) cube {
	var ip0x int
	var ip1x int
	if c2.p0.x >= c1.p0.x {
		ip0x = c2.p0.x
		ip1x = intutil.MinInt(c1.p1.x, c2.p1.x)
	} else {
		ip0x = c1.p0.x
		ip1x = intutil.MinInt(c2.p1.x, c1.p1.x)
	}
	if ip0x > ip1x {
		return zeroCube()
	}

	var ip0y int
	var ip1y int
	if c2.p0.y >= c1.p0.y {
		ip0y = c2.p0.y
		ip1y = intutil.MinInt(c1.p1.y, c2.p1.y)
	} else {
		ip0y = c1.p0.y
		ip1y = intutil.MinInt(c2.p1.y, c1.p1.y)
	}
	if ip0y > ip1y {
		return zeroCube()
	}

	var ip0z int
	var ip1z int
	if c2.p0.z >= c1.p0.z {
		ip0z = c2.p0.z
		ip1z = intutil.MinInt(c1.p1.z, c2.p1.z)
	} else {
		ip0z = c1.p0.z
		ip1z = intutil.MinInt(c2.p1.z, c1.p1.z)
	}
	if ip0z > ip1z {
		return zeroCube()
	}

	ip0 := p(ip0x, ip0y, ip0z)
	ip1 := p(ip1x, ip1y, ip1z)

	return makeCube(ip0, ip1)
}

//

type cubeGrid struct {
	cubes []cube
}

func (g *cubeGrid) onPoints() int {
	var sum int
	for _, cb := range g.cubes {
		sum += cb.volume()
	}
	return sum
}

type interval struct {
	v0, v1 int
}

func (i interval) normalized() interval {
	if i.v1 >= i.v0 {
		return i
	}
	return interval{i.v1, i.v0}
}

func (i interval) isEmpty() bool {
	return i.v1 < i.v0
}

func (c cube) normalized() cube {
	ix := interval{c.p0.x, c.p1.x}.normalized()
	iy := interval{c.p0.y, c.p1.y}.normalized()
	iz := interval{c.p0.z, c.p1.z}.normalized()

	return makeCube(
		p(ix.v0, iy.v0, iz.v0),
		p(ix.v1, iy.v1, iz.v1),
	)
}

func (c cube) is(oc cube) bool {
	return c == oc
}

func (c cube) splitAtIntersection(ci cube) []cube {
	ixs := []interval{
		{c.p0.x, ci.p0.x - 1},
		{ci.p0.x, ci.p1.x},
		{ci.p1.x + 1, c.p1.x},
	}
	iys := []interval{
		{c.p0.y, ci.p0.y - 1},
		{ci.p0.y, ci.p1.y},
		{ci.p1.y + 1, c.p1.y},
	}
	izs := []interval{
		{c.p0.z, ci.p0.z - 1},
		{ci.p0.z, ci.p1.z},
		{ci.p1.z + 1, c.p1.z},
	}

	var cs []cube
	for _, ix := range ixs {
		if ix.isEmpty() {
			continue
		}
		for _, iy := range iys {
			if iy.isEmpty() {
				continue
			}
			for _, iz := range izs {
				if iz.isEmpty() {
					continue
				}
				sc := makeCube(
					p(ix.v0, iy.v0, iz.v0),
					p(ix.v1, iy.v1, iz.v1),
				).normalized()
				if sc.isEmpty() {
					continue
				}
				if sc.is(ci) {
					continue
				}
				cs = append(cs, sc)
			}
		}
	}
	return cs
}

func (g *cubeGrid) setOn(oncb cube) {
	cbs := []cube{oncb}
	//outer:
	for _, c := range g.cubes {
		newCbs := []cube{}
		for _, cb := range cbs {
			ci := cubeIntersection(c, cb)
			if ci.isEmpty() {
				newCbs = append(newCbs, cb)
				continue
			}

			newCbs = append(newCbs, cb.splitAtIntersection(ci)...)

			//continue outer
		}
		cbs = newCbs
	}

	// no more intersections
	// just  add rest
	g.cubes = append(g.cubes, cbs...)
}

func (g *cubeGrid) setOff(offcb cube) {
	newCubes := []cube{}
	for _, c := range g.cubes {
		ci := cubeIntersection(c, offcb)
		if ci.isEmpty() {
			newCubes = append(newCubes, c)
			continue
		}
		cbs := c.splitAtIntersection(ci)
		newCubes = append(newCubes, cbs...)
	}
	g.cubes = newCubes
}

func onAfterReboot(in string, bounds cube) (int, error) {
	steps, err := parseSteps(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-steps")
	}

	g := &cubeGrid{}
	for i, step := range steps {
		v := step.op == "on"
		p0 := p(step.x0, step.y0, step.z0)
		p1 := p(step.x1, step.y1, step.z1)

		if !bounds.isEmpty() && !bounds.containsPoint(p0) && !bounds.containsPoint(p1) {
			continue
		}

		if v {
			g.setOn(makeCube(p0, p1))
		} else {
			g.setOff(makeCube(p0, p1))
		}
		log("after %d: on: %d", i, g.onPoints())
	}

	on := g.onPoints()

	return on, nil
}
