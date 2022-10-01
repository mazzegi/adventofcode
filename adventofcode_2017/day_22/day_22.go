package day_22

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := numInfections(input, 10000)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := numInfectionsExt(input, 10000000)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type point struct {
	x, y int
}

func p(x, y int) point {
	return point{x: x, y: y}
}

func (pt point) add(op point) point {
	return p(pt.x+op.x, pt.y+op.y)
}

type state uint8

const (
	clean    state = 0
	weakened state = 1
	infected state = 2
	flagged  state = 3
)

type grid struct {
	infected map[point]state
	center   point
}

func mustParseGrid(in string) *grid {
	g := &grid{
		infected: map[point]state{},
	}

	lines := readutil.ReadLines(in)
	if len(lines) == 0 {
		fatal("no data")
	}
	dimy := len(lines)
	dimx := len(lines[0])
	g.center = p(dimx/2, dimy/2)

	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				g.infected[p(x, y)] = infected
			}
		}
	}
	return g
}

func (g *grid) get(p point) state {
	s, ok := g.infected[p]
	if !ok {
		return clean
	}
	return s
}

func (g *grid) set(p point, v state) {
	if v == clean {
		delete(g.infected, p)
	} else {
		g.infected[p] = v
	}
}

// func (g *grid) dump(curr point) string {

// 	minx, maxx, miny, maxy := curr.x, curr.x, curr.y, curr.y

// 	for pt, v := range g.infected {
// 		if !v {
// 			fatal("uninfected in map")
// 		}

// 		if pt.x < minx {
// 			minx = pt.x
// 		}
// 		if pt.x > maxx {
// 			maxx = pt.x
// 		}

// 		if pt.y < miny {
// 			miny = pt.y
// 		}
// 		if pt.y > maxy {
// 			maxy = pt.y
// 		}
// 	}

// 	var sl []string
// 	for y := miny; y <= maxy; y++ {
// 		var sr string
// 		for x := minx; x <= maxx; x++ {
// 			pt := p(x, y)
// 			if g.get(pt) {
// 				if pt == curr {
// 					sr += "+"
// 				} else {
// 					sr += "#"
// 				}
// 			} else {
// 				if pt == curr {
// 					sr += "o"
// 				} else {
// 					sr += "."
// 				}
// 			}
// 		}
// 		sl = append(sl, sr)
// 	}

// 	return strings.Join(sl, "\n")
// }

func numInfections(in string, bursts int) (int, error) {
	g := mustParseGrid(in)

	curr := g.center
	// y increases downwards
	dir := p(0, -1)
	turnLeft := func() {
		dir = p(dir.y, -dir.x)
	}
	turnRight := func() {
		dir = p(-dir.y, dir.x)
	}

	//log("*** initial ***\n%s", g.dump(curr))

	infections := 0
	for i := 0; i < bursts; i++ {
		//turn
		if g.get(curr) == infected {
			turnRight()
		} else {
			turnLeft()
		}

		// in/desin fect
		if g.get(curr) == clean {
			g.set(curr, infected)
			infections++
		} else {
			g.set(curr, clean)
		}

		//move
		curr = curr.add(dir)
		//log("\n*** after %d ***\n%s", i+1, g.dump(curr))
	}

	return infections, nil
}

func numInfectionsExt(in string, bursts int) (int, error) {
	g := mustParseGrid(in)

	curr := g.center
	// y increases downwards
	dir := p(0, -1)
	turnLeft := func() {
		dir = p(dir.y, -dir.x)
	}
	turnRight := func() {
		dir = p(-dir.y, dir.x)
	}

	//log("*** initial ***\n%s", g.dump(curr))

	infections := 0
	for i := 0; i < bursts; i++ {
		//turn
		s := g.get(curr)
		switch s {
		case clean:
			turnLeft()
			g.set(curr, weakened)
		case weakened:
			g.set(curr, infected)
			infections++
		case infected:
			turnRight()
			g.set(curr, flagged)
		case flagged:
			turnLeft()
			turnLeft()
			g.set(curr, clean)
		default:
			fatal("invalid state %d", s)
		}

		//move
		curr = curr.add(dir)
		//log("\n*** after %d ***\n%s", i+1, g.dump(curr))
	}

	return infections, nil
}
