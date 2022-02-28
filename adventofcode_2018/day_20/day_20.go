package day_20

import (
	"adventofcode_2018/errutil"
	"fmt"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
const (
	Location rune = 'X'
	Space    rune = '.'
	Wall     rune = '#'
	VDoor    rune = '|'
	HDoor    rune = '-'
	Door     rune = '+'
	Unknown  rune = '?'
)

type point struct {
	x, y int
}

func (p point) dir(r rune) point {
	switch r {
	case 'E':
		return point{p.x + 1, p.y}
	case 'S':
		return point{p.x, p.y + 1}
	case 'W':
		return point{p.x - 1, p.y}
	case 'N':
		return point{p.x, p.y - 1}
	default:
		return p
	}
}

func (pt point) diags() []point {
	var dps []point
	dps = append(dps,
		point{pt.x - 1, pt.y - 1},
		point{pt.x - 1, pt.y + 1},
		point{pt.x + 1, pt.y - 1},
		point{pt.x + 1, pt.y + 1},
	)
	return dps
}

func (pt point) perps() []point {
	var dps []point
	dps = append(dps,
		point{pt.x, pt.y - 1},
		point{pt.x, pt.y + 1},
		point{pt.x + 1, pt.y},
		point{pt.x - 1, pt.y},
	)
	return dps
}

func markUnknown(m map[point]rune, ps []point) {
	for _, p := range ps {
		if _, ok := m[p]; ok {
			continue
		}
		m[p] = Unknown
	}
}

func markWall(m map[point]rune, ps []point) {
	for _, p := range ps {
		m[p] = Wall
	}
}

func buildMap(rex *Regex, m map[point]rune, pos point) {
	curr := pos
	for _, elt := range rex.Elts {
		markWall(m, curr.diags())
		markUnknown(m, curr.perps())
		if m[curr] != Location {
			m[curr] = Space
		}

		next := curr.dir(elt.Value)
		switch elt.Value {
		case 'E', 'W':
			m[next] = VDoor
		default:
			m[next] = HDoor
		}
		curr = next.dir(elt.Value)
		for _, br := range elt.Branches {
			buildMap(br, m, curr)
		}
	}
}

func part1MainFunc(in string) (int, error) {
	rex, err := ParseRegex(in)
	if err != nil {
		return 0, err
	}
	m := map[point]rune{}
	pos := point{0, 0}
	m[pos] = Location
	buildMap(rex, m, pos)
	xmin, xmax, ymin, ymax := 0, 0, 0, 0
	for p, v := range m {
		if v == Unknown {
			m[p] = Wall
		}
		if p.x < xmin {
			xmin = p.x
		}
		if p.x > xmax {
			xmax = p.x
		}
		if p.y < ymin {
			ymin = p.y
		}
		if p.y > ymax {
			ymax = p.y
		}
	}
	log("x: %d..%d, y: %d..%d", xmin, xmax, ymin, ymax)

	for y := ymin; y <= ymax; y++ {
		var sr string
		for x := xmin; x <= xmax; x++ {
			if v, ok := m[point{x, y}]; ok {
				sr += string(v)
			} else {
				sr += "?"
			}
		}
		log(sr)
	}

	return 0, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
