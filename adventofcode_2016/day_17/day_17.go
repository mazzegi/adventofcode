package day_17

import (
	"crypto/md5"
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2016/errutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := shortestPath("qljzarfv", 4, 4)
	errutil.ExitOnErr(err)
	log("part1: result = %q", res)
}

func Part2() {
	res, err := longestPath("qljzarfv", 4, 4)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type dir string

const (
	up    dir = "U"
	down  dir = "D"
	left  dir = "L"
	right dir = "R"
)

type dirs []dir

func (ds dirs) String() string {
	var s string
	for _, d := range ds {
		s += string(d)
	}
	return s
}

func (ds dirs) clone() dirs {
	cds := make(dirs, len(ds))
	copy(cds, ds)
	return cds
}

func isOpenChar(b byte) bool {
	switch b {
	case 'b', 'c', 'd', 'e', 'f':
		return true
	default:
		return false
	}
}

func openDoors(passcode string, path string) dirs {
	var openDirs dirs
	hash := fmt.Sprintf("%x", md5.Sum([]byte(passcode+path)))
	if isOpenChar(hash[0]) {
		openDirs = append(openDirs, up)
	}
	if isOpenChar(hash[1]) {
		openDirs = append(openDirs, down)
	}
	if isOpenChar(hash[2]) {
		openDirs = append(openDirs, left)
	}
	if isOpenChar(hash[3]) {
		openDirs = append(openDirs, right)
	}
	return openDirs
}

type point struct {
	x, y int
}

func p(x, y int) point {
	return point{x, y}
}

type grid struct {
	passcode     string
	sizeX, sizeY int
}

type option struct {
	dir dir
	pt  point
}

func (g *grid) possibleDirs(path string, pt point) []option {
	open := openDoors(g.passcode, path)
	var possible []option
	for _, o := range open {
		add := false
		opt := pt
		switch o {
		case up:
			add = pt.y > 0
			opt.y--
		case down:
			add = pt.y < g.sizeY-1
			opt.y++
		case left:
			add = pt.x > 0
			opt.x--
		case right:
			add = pt.x < g.sizeX-1
			opt.x++
		default:
			fatal("invalid direction %q", o)
		}
		if add {
			possible = append(possible, option{
				dir: o,
				pt:  opt,
			})
		}
	}

	return possible
}

func shortestPath(passcode string, sizeX, sizeY int) (string, error) {
	curr := p(0, 0)
	dest := p(sizeX-1, sizeY-1)
	g := &grid{passcode, sizeX, sizeY}

	minDist, minPath, ok := walkShortest(g, dirs{}, curr, dest)
	if !ok {
		fatal("no path")
	}
	log("min = %d, path = %q", minDist, minPath.String())

	return minPath.String(), nil
}

func walkShortest(g *grid, path dirs, curr point, dest point) (minDist int, minPath dirs, ok bool) {
	//
	minDist = 0
	minPath = dirs{}
	ok = false
	options := g.possibleDirs(path.String(), curr)
	for _, opt := range options {
		if opt.pt == dest {
			return 1, dirs{opt.dir}, true
		}

		cpath := path.clone()
		cpath = append(cpath, opt.dir)

		subDist, subPath, subOk := walkShortest(g, cpath, opt.pt, dest)
		if !subOk {
			continue
		}

		if !ok || subDist+1 < minDist {
			minDist = subDist + 1
			minPath = append(dirs{opt.dir}, subPath...)
			ok = true
		}
	}

	return
}

func longestPath(passcode string, sizeX, sizeY int) (int, error) {
	curr := p(0, 0)
	dest := p(sizeX-1, sizeY-1)
	g := &grid{passcode, sizeX, sizeY}

	maxDist, maxPath, ok := walkLongest(g, dirs{}, curr, dest)
	if !ok {
		fatal("no path")
	}
	log("min = %d, path = %q", maxDist, maxPath.String())

	return maxDist, nil
}

func walkLongest(g *grid, path dirs, curr point, dest point) (maxDist int, maxPath dirs, ok bool) {
	//
	maxDist = 0
	maxPath = dirs{}
	ok = false
	options := g.possibleDirs(path.String(), curr)
outer:
	for _, opt := range options {
		if opt.pt == dest {
			if !ok || 1 > maxDist {
				maxDist = 1
				maxPath = dirs{opt.dir}
				ok = true
			}
			continue outer
		}

		cpath := path.clone()
		cpath = append(cpath, opt.dir)

		subDist, subPath, subOk := walkLongest(g, cpath, opt.pt, dest)
		if !subOk {
			continue
		}

		if !ok || subDist+1 > maxDist {
			maxDist = subDist + 1
			maxPath = append(dirs{opt.dir}, subPath...)
			ok = true
		}
	}

	return
}
