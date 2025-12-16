package day_07

import (
	"fmt"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/set"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

type Beam struct {
	start grid.GridPoint
	curr  grid.GridPoint
	done  bool
}

func part1MainFunc(in string) (int, error) {
	g, err := grid.Parse(in)
	if err != nil {
		return 0, fmt.Errorf("parse_grid: %w", err)
	}
	startPt, ok := g.FindFirst('S')
	if !ok {
		return 0, fmt.Errorf("start point not found")
	}
	var beams []*Beam
	beam0 := &Beam{
		start: startPt,
		curr:  startPt,
	}
	beams = append(beams, beam0)

	splitCount := 0
	advanceBeam := func(b *Beam) (newBeams []*Beam) {
		// advance down
		newCurr := b.curr
		newCurr.Row++
		if !g.Contains(newCurr) {
			b.done = true
			return
		}
		b.curr = newCurr

		if g.At(b.curr) == '^' {
			b.done = true
			splitCount++

			// create 2 other beams
			b1Pt := grid.GP(b.curr.Col-1, b.curr.Row)
			if g.Contains(b1Pt) {
				newBeams = append(newBeams, &Beam{
					start: b1Pt,
					curr:  b1Pt,
				})
			}
			b2Pt := grid.GP(b.curr.Col+1, b.curr.Row)
			if g.Contains(b2Pt) {
				newBeams = append(newBeams, &Beam{
					start: b2Pt,
					curr:  b2Pt,
				})
			}
		}
		return
	}

	tidyBeams := func(bs []*Beam) []*Beam {
		occPts := set.New[grid.GridPoint]()
		var tbs []*Beam
		for _, b := range bs {
			if occPts.Contains(b.curr) {
				continue
			}
			tbs = append(tbs, b)
			occPts.Insert(b.curr)
		}
		return tbs
	}

	// process beams
	iter := 0
	for {

		advCount := 0
		var newBeamsAll []*Beam
		for _, b := range beams {
			if b.done {
				continue
			}
			newBeams := advanceBeam(b)
			newBeamsAll = append(newBeamsAll, newBeams...)
			advCount++
		}
		if advCount == 0 {
			break
		}
		iter++
		beams = append(beams, newBeamsAll...)
		beams = tidyBeams(beams)
		//log("iter %d: %d active beams", iter, numActiveBeams())
	}

	return splitCount, nil
}

func part2MainFunc(in string) (int, error) {
	g, err := grid.Parse(in)
	if err != nil {
		return 0, fmt.Errorf("parse_grid: %w", err)
	}
	startPt, ok := g.FindFirst('S')
	if !ok {
		return 0, fmt.Errorf("start point not found")
	}

	var totalDone int

	var next func(p grid.GridPoint)
	next = func(p grid.GridPoint) {
		p.Row++
		if !g.Contains(p) {
			totalDone++
			return
		}
		if g.At(p) == '^' {
			//left
			pLeft := p
			pLeft.Col--
			if !g.Contains(pLeft) {
				totalDone++
				return
			}
			next(pLeft)
			//right
			pRight := p
			pRight.Col++
			if !g.Contains(pRight) {
				totalDone++
				return
			}
			next(pRight)
		} else {
			next(p)
		}
	}

	next(startPt)

	return totalDone, nil
}
