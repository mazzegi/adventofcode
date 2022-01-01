package day_11

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/intutil"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := distance(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type direction string

const (
	north     direction = "n"
	northEast direction = "ne"
	southEast direction = "se"
	south     direction = "s"
	southWest direction = "sw"
	northWest direction = "nw"
)

//
type hexPoint struct {
	x, y int
}

func hp(x, y int) hexPoint {
	return hexPoint{x: x, y: y}
}

func (p hexPoint) nextIn(dir direction) hexPoint {
	switch dir {
	case north:
		return hp(p.x, p.y+1)
	case northEast:
		return hp(p.x+1, p.y)
	case southEast:
		return hp(p.x+1, p.y-1)
	case south:
		return hp(p.x, p.y-1)
	case southWest:
		return hp(p.x-1, p.y)
	case northWest:
		return hp(p.x-1, p.y+1)
	default:
		panic("invalid dir")
	}
}

func dist(p1, p2 hexPoint) int {
	if p1 == p2 {
		return 0
	}
	if p1.x == p2.x {
		return intutil.AbsInt(p2.y - p1.y)
	}
	if p1.y == p2.y {
		return intutil.AbsInt(p2.x - p1.x)
	}

	// reach the desired 1 from 2
	var d int
	for p2.x != p1.x {
		var xstep int
		if p2.x > p1.x {
			p2.x--
			xstep = -1
		} else if p2.x < p1.x {
			p2.x++
			xstep = 1
		}
		// mv y towards desired y
		dy := p1.y - p2.y
		if dy > 0 {
			if xstep == 0 || xstep == -1 {
				p2.y++
			}
		} else if dy < 0 {
			if xstep == 0 || xstep == 1 {
				p2.y--
			}
		}
		//otherwise equal - no y step
		d++
	}

	//now x are equal
	d += intutil.AbsInt(p2.y - p1.y)

	return d
}

func parseDirections(in string) ([]direction, error) {
	in = strings.Trim(in, " \r\n\t")
	sl := strings.Split(in, ",")
	var dirs []direction
	for _, s := range sl {
		d := direction(s)
		dirs = append(dirs, d)
	}
	return dirs, nil
}

func distance(in string) (int, error) {
	dirs, err := parseDirections(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-directions")
	}
	if len(dirs) == 0 {
		return 0, errors.Errorf("no data")
	}

	p := hp(0, 0)
	var maxDist int
	for _, dir := range dirs {
		p = p.nextIn(dir)
		d := dist(hp(0, 0), p)
		if d > maxDist {
			maxDist = d
		}
	}
	d := dist(hp(0, 0), p)

	log("dist: %d", d)
	log("max : %d", maxDist)

	return d, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
