package day_24

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
	"github.com/mazzegi/adventofcode/slices"
)

func P(x, y int) grid.Point {
	return grid.Pt(x, y)
}

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatalIfErr(err error) {
	if err == nil {
		return
	}
	fatal("err not nil: %v", err)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

const skip1 = true

func Part1() {
	if skip1 {
		return
	}
	res, err := part1MainFunc(input, 900)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input, 500)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

const (
	right = '>'
	left  = '<'
	up    = '^'
	down  = 'v'
)

func mirrorDirection(d rune) rune {
	switch d {
	case right:
		return left
	case left:
		return right
	case up:
		return down
	case down:
		return up
	default:
		panic("invalid direction " + string(d))
	}
}

type Blizzard struct {
	Direction rune
	Pos       grid.Point
}

type Grid struct {
	dimX, dimY int
	occupied   *set.Set[grid.Point]
	blizzards  []Blizzard
}

func (g *Grid) Clone() *Grid {
	return &Grid{
		dimX:      g.dimX,
		dimY:      g.dimY,
		blizzards: slices.Clone(g.blizzards),
	}
}

func (g *Grid) Hash(pos grid.Point, iter int) string {
	var sl []string
	sl = append(sl, fmt.Sprintf("%d=%d:%d", iter, pos.X, pos.Y))
	for _, blz := range g.blizzards {
		sl = append(sl, fmt.Sprintf("%s:%d:%d", string(blz.Direction), blz.Pos.X, blz.Pos.Y))
	}
	return strings.Join(sl, ",")
}

func mustParseGrid(in string) *Grid {
	g := &Grid{
		occupied: set.New[grid.Point](),
	}
	lines := readutil.ReadLines(in)
	lines = lines[1 : len(lines)-1]
	g.dimX = len(lines[0])
	g.dimY = len(lines) + 2
	for li, line := range lines {
		for i := 1; i < len(line)-1; i++ {
			r := rune(line[i])
			if r == '.' {
				continue
			}
			blz := Blizzard{
				Direction: r,
				Pos:       P(i, li+1),
			}
			g.blizzards = append(g.blizzards, blz)
			g.occupied.Insert(blz.Pos)
		}
	}
	return g
}

func (g *Grid) Move(n int) {
	occ := set.New[grid.Point]()
	for i := 0; i < n; i++ {
		for i, blz := range g.blizzards {
			var newPos grid.Point
			switch blz.Direction {
			case left:
				newPos = blz.Pos.Add(P(-1, 0))
				if newPos.X == 0 {
					newPos.X = g.dimX - 2
				}
			case right:
				newPos = blz.Pos.Add(P(1, 0))
				if newPos.X == g.dimX-1 {
					newPos.X = 1
				}
			case up:
				newPos = blz.Pos.Add(P(0, -1))
				if newPos.Y == 0 {
					newPos.Y = g.dimY - 2
				}
			case down:
				newPos = blz.Pos.Add(P(0, 1))
				if newPos.Y == g.dimY-1 {
					newPos.Y = 1
				}
			default:
				fatal("invalid direction %q", string(blz.Direction))
			}
			if newPos.Y >= g.dimY-1 || newPos.X >= g.dimX-1 || newPos.Y <= 0 || newPos.X <= 0 {
				log("örks")
			}
			g.blizzards[i].Pos = newPos
			occ.Insert(newPos)
		}
	}
	g.occupied = occ
}

func (g *Grid) Mirrored() *Grid {
	mg := &Grid{
		dimX:     g.dimX,
		dimY:     g.dimY,
		occupied: set.New[grid.Point](),
	}
	for _, blz := range g.blizzards {
		mblz := Blizzard{
			Direction: mirrorDirection(blz.Direction),
			Pos:       P(mg.dimX-blz.Pos.X-1, mg.dimY-blz.Pos.Y-1),
		}
		mg.blizzards = append(mg.blizzards, mblz)
		mg.occupied.Insert(mblz.Pos)
	}
	return mg
}

func (g *Grid) StartEnd() (start, end grid.Point) {
	start = P(1, 0)
	end = P(g.dimX-2, g.dimY-1)
	return
}

func (g *Grid) isOpen(p grid.Point) bool {
	s, e := g.StartEnd()
	if p == s || p == e {
		return true
	}

	if p.X <= 0 || p.X >= g.dimX-1 ||
		p.Y <= 0 || p.Y >= g.dimY-1 {
		return false
	}
	if g.occupied.Contains(p) {
		return false
	}
	return true
}

func (g *Grid) OpenNeighbours(p grid.Point) []grid.Point {
	var ons []grid.Point
	addIfOpen := func(p grid.Point) {
		if g.isOpen(p) {
			ons = append(ons, p)
		}
	}
	addIfOpen(p.Add(P(0, 1)))
	addIfOpen(p.Add(P(1, 0)))
	addIfOpen(p)
	addIfOpen(p.Add(P(-1, 0)))
	addIfOpen(p.Add(P(0, -1)))

	return ons
}

//

func part1MainFunc(in string, maxIter int) (int, error) {
	g := mustParseGrid(in)
	finder := NewFinder(g, maxIter)
	res := finder.Find()
	return res, nil
}

func part2MainFunc(in string, maxIter int) (int, error) {
	log("part1 ...")
	g := mustParseGrid(in)
	finder1 := NewFinder(g, 300)
	res1 := finder1.Find()
	log("part1: %d", res1)

	log("part2 ...")
	g.Move(res1)
	mg2 := g.Mirrored()
	finder2 := NewFinder(mg2, maxIter)
	//finder2.cache = finder1.cache
	res2 := finder2.Find()
	log("part2: %d", res2)

	log("part3 ...")
	mg2.Move(res2)
	mg3 := mg2.Mirrored()
	finder3 := NewFinder(mg3, maxIter)
	//finder3.cache = finder2.cache
	res3 := finder3.Find()
	log("part3: %d", res3)

	res := res1 + res2 + res3
	return res, nil
}
