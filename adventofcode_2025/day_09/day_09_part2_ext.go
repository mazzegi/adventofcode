package day_09

import (
	"fmt"
	"sort"
	"time"

	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/listv2"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
)

func iv(min, max int) interval {
	return interval{min, max}
}

type interval struct {
	min int
	max int
}

func (i interval) contains(v int) bool {
	return v >= i.min && v <= i.max
}

func (i interval) isZero() bool {
	return i.min == 0 && i.max == 0
}

func intsToIntervals(ns []int) []interval {
	if len(ns) == 0 {
		return []interval{}
	}
	sort.Ints(ns)
	var ivs []interval
	var curr interval

	for i, n := range ns {
		if i == 0 {
			curr = iv(n, n)
			continue
		}
		// can merge
		if n <= curr.max+1 { // either max or max+1 (!!! sorted)
			curr.max = n
		} else {
			ivs = append(ivs, curr)
			curr = iv(n, n)
		}
	}
	ivs = append(ivs, curr)
	return ivs
}

func part2MainFuncV2(in string) (int, error) {
	t0 := time.Now()
	var redPoints []grid.GridPoint
	edge := set.New[grid.GridPoint]()
	edgeRows := map[int][]int{} // row -> []col
	insertIntoEdge := func(pt grid.GridPoint) {
		edge.Insert(pt)
		edgeRows[pt.Row] = append(edgeRows[pt.Row], pt.Col)
	}

	lines := readutil.ReadLines(in)
	var (
		minX int
		minY int
		maxX int
		maxY int
	)
	for i, line := range lines {
		var pt grid.GridPoint
		_, err := fmt.Sscanf(line, "%d,%d", &pt.Col, &pt.Row)
		if err != nil {
			return 0, fmt.Errorf("scan line: %w", err)
		}
		redPoints = append(redPoints, pt)
		insertIntoEdge(pt)
		if i == 0 {
			minX, minY = pt.Col, pt.Row
			maxX, maxY = pt.Col, pt.Row
		} else {
			if pt.Col < minX {
				minX = pt.Col
			} else if pt.Col > maxX {
				maxX = pt.Col
			}

			if pt.Row < minY {
				minY = pt.Row
			} else if pt.Row > maxY {
				maxY = pt.Row
			}
		}
	}
	log("scan red points (%d): x:[%d, %d] - %d; x:[%d, %d] - %d; in %s", len(redPoints), minX, maxX, maxX-minX, minY, maxY, maxY-minY, time.Since(t0).Round(time.Microsecond))
	var greenPoints []grid.GridPoint

	t0 = time.Now()
loop_points:
	for i, rp := range redPoints {
		var rpNext grid.GridPoint
		if i < len(redPoints)-1 {
			rpNext = redPoints[i+1]
		} else {
			rpNext = redPoints[0]
		}
		if rp == rpNext {
			continue
		}
		switch {
		case rp.Col == rpNext.Col:
			from := mathutil.Min(rp.Row, rpNext.Row)
			to := mathutil.Max(rp.Row, rpNext.Row)
			if to-from < 2 {
				continue loop_points
			}
			for row := from + 1; row < to; row++ {
				pt := grid.GP(rp.Col, row)
				greenPoints = append(greenPoints, pt)
				insertIntoEdge(pt)
			}

		case rp.Row == rpNext.Row:
			from := mathutil.Min(rp.Col, rpNext.Col)
			to := mathutil.Max(rp.Col, rpNext.Col)
			if to-from < 2 {
				continue loop_points
			}
			for col := from + 1; col < to; col++ {
				pt := grid.GP(col, rp.Row)
				greenPoints = append(greenPoints, pt)
				insertIntoEdge(pt)
			}

		default:
			// neither
			return 0, fmt.Errorf("rp (%s) and rp_next (%s) are not on vert or horz line", rp, rpNext)
		}
	}
	log("determine green points (%d) in %s", len(greenPoints), time.Since(t0).Round(time.Microsecond))

	//
	t0 = time.Now()
	edgeRowIntervals := map[int][]interval{}
	for row, cols := range edgeRows {
		ivs := intsToIntervals(cols)
		edgeRowIntervals[row] = ivs
	}
	log("built edge-row-intervals in %s", time.Since(t0).Round(time.Microsecond))

	firstOnEdge := func(row, fromColIncl int) (int, bool) {
		rowIvs, ok := edgeRowIntervals[row]
		if !ok {
			return 0, false
		}
		//
		for _, riv := range rowIvs {
			if fromColIncl > riv.max {
				continue
			}
			//col is <= riv-max
			if fromColIncl < riv.min {
				return riv.min, true
			}
			// its >= riv.min, so from col is already on edge
			return fromColIncl, true
		}
		return 0, false
	}

	scanRow := func(row int) []interval {
		var ivs []interval
		fromCol := minX - 1
		for {
			col, ok := firstOnEdge(row, fromCol)
			if !ok {
				// no more on edge
				break
			}
			_ = col
		}
		return ivs
	}

	_ = scanRow
	insideRowIntervals := map[int][]interval{}
	_ = insideRowIntervals

	return 0, nil
}

type rowIntervals map[int]*listv2.List[interval]

func (rivs rowIntervals) counts() (numIntervals, numPoints int) {
	for _, ls := range rivs {
		for n := ls.Front(); n != nil; n = n.Next() {
			numIntervals++
			iv := n.Data()
			numPoints += iv.max - iv.min + 1
		}
	}
	return
}

func (rivs rowIntervals) contains(pt grid.GridPoint) bool {
	ls, ok := rivs[pt.Row]
	if !ok {
		return false
	}
	for n := ls.Front(); n != nil; n = n.Next() {
		if n.Data().contains(pt.Col) {
			return true
		}
	}
	return false
}

func (rivs rowIntervals) insert(pt grid.GridPoint) {
	ls, ok := rivs[pt.Row]
	if !ok {
		ls := listv2.New[interval]()
		ls.PushBack(interval{pt.Col, pt.Col})
		rivs[pt.Row] = ls
		return
	}
	for n := ls.Front(); n != nil; n = n.Next() {
		iv := n.Data()
		switch {
		case pt.Col < iv.min-1: // must insert before
			ls.InsertBefore(n, interval{pt.Col, pt.Col})
			return

		case pt.Col == iv.min-1:
			// we can merge
			iv.min = pt.Col
			n.SetData(iv)
			return

		case iv.contains(pt.Col):
			return

		case pt.Col == iv.max+1:
			// maybe it can be merged with next interval
			if n.Next() != nil && pt.Col == n.Next().Data().min-1 {
				iv.max = n.Next().Data().max
				n.SetData(iv)
				ls.Remove(n.Next())
				return
			} else {
				iv.max = pt.Col
				n.SetData(iv)
				return
			}
		}
	}
	// no matches - append
	ls.PushBack(interval{pt.Col, pt.Col})
}

/*
..............
.......#XXX#..
.......X...X..
..#XXXX#...X..
..X........X..
..#XXXXXX#.X..
.........X.X..
.........#X#..
..............
*/
