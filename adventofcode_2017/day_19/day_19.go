package day_19

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/readutil"
	"fmt"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := letters(input)
	errutil.ExitOnErr(err)
	log("part1: result = %q", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

/*
const inputTest = `
     |
     |  +--+
     A  |  C
 F---|----E|--+
     |  |  |  D
     +B-+  +--+
`
*/

//
type point struct {
	x, y int
}

func (pt point) add(op point) point {
	return p(pt.x+op.x, pt.y+op.y)
}

func (pt point) sub(op point) point {
	return p(pt.x-op.x, pt.y-op.y)
}

func p(x, y int) point {
	return point{x: x, y: y}
}

const (
	symBlank rune = ' '
	symVert  rune = '|'
	symHorz  rune = '-'
	symTurn  rune = '+'
)

type diagram struct {
	symbols    [][]rune
	xdim, ydim int
}

func (d *diagram) sym(x, y int) rune {
	if y < 0 || y >= d.ydim {
		fatal("invalid diagram point")
	}
	if x < 0 || x >= d.xdim {
		fatal("invalid diagram point")
	}
	return d.symbols[y][x]
}

func (d *diagram) inRange(pt point) bool {
	if pt.y < 0 || pt.y >= d.ydim {
		return false
	}
	if pt.x < 0 || pt.x >= d.xdim {
		return false
	}
	return true
}

func (d *diagram) adjacent(pt point) []point {
	var ps []point
	addIfInRange := func(pt point) {
		if d.inRange(pt) {
			ps = append(ps, pt)
		}
	}
	addIfInRange(p(pt.x-1, pt.y))
	addIfInRange(p(pt.x+1, pt.y))
	addIfInRange(p(pt.x, pt.y-1))
	addIfInRange(p(pt.x, pt.y+1))
	return ps
}

func mustParseDiagram(in string) *diagram {
	d := &diagram{}
	lines := readutil.ReadLinesUntrimmed(in)
	for _, line := range lines {
		var lrs []rune
		for _, r := range line {
			lrs = append(lrs, r)
		}
		if len(d.symbols) > 0 && len(d.symbols[0]) != len(lrs) {
			fatal("inconsistent row size")
		}

		d.symbols = append(d.symbols, lrs)
	}
	if len(d.symbols) == 0 || len(d.symbols[0]) == 0 {
		fatal("no data")
	}
	d.ydim = len(d.symbols)
	d.xdim = len(d.symbols[0])
	return d
}

func letters(in string) (string, error) {
	d := mustParseDiagram(in)
	//find start
	var p0 point
	dir := p(0, 1)
	for x := 0; x < d.xdim; x++ {
		if d.sym(x, 0) == symVert {
			p0 = p(x, 0)
			break
		}
	}

	var res string
	var turn func(pturn point, pcurr point) (next point, nextDir point, nextSym rune)
	turn = func(pturn point, pcurr point) (next point, nextDir point, nextSym rune) {
		ads := d.adjacent(pturn)
		for _, ad := range ads {
			if ad == pcurr {
				continue
			}
			sym := d.sym(ad.x, ad.y)
			if sym == symBlank {
				continue
			}
			if sym == symTurn {
				return turn(ad, pturn)
			}
			if sym == symHorz {
				dir := ad.sub(pturn)
				return ad, dir, symHorz
			}
			if sym == symVert {
				dir := ad.sub(pturn)
				return ad, dir, symVert
			}
			// a letter
			res += string(sym)
			dir := ad.sub(pturn)
			if dir.x != 0 {
				nextSym = symHorz
			} else {
				nextSym = symVert
			}
			return ad, dir, nextSym
		}
		fatal("found no adjacent turn")
		return
	}

	currSym := symVert
	var steps int
	for {
		//same sym in direction?
		pn := p0.add(dir)
		steps++
		sym := d.sym(pn.x, pn.y)
		switch {
		case sym == symBlank:
			log("went %d steps total", steps)
			return res, nil
		case currSym == symVert && sym == symVert:
			p0 = pn
		case currSym == symHorz && sym == symHorz:
			p0 = pn
		case currSym == symVert && sym == symHorz:
			pn = pn.add(dir)
			steps++
			p0 = pn
			if symJ := d.sym(p0.x, p0.y); symJ != symVert {
				res += string(symJ)
			}
		case currSym == symHorz && sym == symVert:
			pn = pn.add(dir)
			steps++
			p0 = pn
			if symJ := d.sym(p0.x, p0.y); symJ != symHorz {
				res += string(symJ)
			}
		case sym == symTurn:
			next, nextDir, nextSym := turn(pn, p0)
			p0 = next
			dir = nextDir
			currSym = nextSym
			steps += 1
		default:
			// a letter
			res += string(sym)
			p0 = pn
		}
	}
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
