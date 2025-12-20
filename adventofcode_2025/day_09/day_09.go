package day_09

import (
	"fmt"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
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
	//res, err := part2MainFunc(input)
	res, err := part2MainFuncV2(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

func part1MainFunc(in string) (int, error) {
	var pts []grid.GridPoint
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		var pt grid.GridPoint
		_, err := fmt.Sscanf(line, "%d,%d", &pt.Col, &pt.Row)
		if err != nil {
			return 0, fmt.Errorf("scan line: %w", err)
		}
		pts = append(pts, pt)
	}

	area := func(p1, p2 grid.GridPoint) int {
		dx := mathutil.Abs(p1.Col-p2.Col) + 1
		dy := mathutil.Abs(p1.Row-p2.Row) + 1
		return dx * dy
	}
	var max int
	for _, p1 := range pts {
		for _, p2 := range pts {
			a := area(p1, p2)
			if a > max {
				max = a
			}
		}
	}

	return max, nil
}

func part2MainFunc(in string) (int, error) {
	var redPoints []grid.GridPoint
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
	var greenPoints []grid.GridPoint

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
				greenPoints = append(greenPoints, grid.GP(rp.Col, row))
			}

		case rp.Row == rpNext.Row:
			from := mathutil.Min(rp.Col, rpNext.Col)
			to := mathutil.Max(rp.Col, rpNext.Col)
			if to-from < 2 {
				continue loop_points
			}
			for col := from + 1; col < to; col++ {
				greenPoints = append(greenPoints, grid.GP(col, rp.Row))
			}

		default:
			// neither
			return 0, fmt.Errorf("rp (%s) and rp_next (%s) are not on vert or horz line", rp, rpNext)
		}
	}

	// prepare
	allPointSet := set.New(redPoints...)
	allPointSet.Insert(greenPoints...)
	onEdge := func(pt grid.GridPoint) bool {
		return allPointSet.Contains(pt)
	}

	const (
		stateOut = "out"
		//stateEdge         = "edge"
		stateIn           = "in"
		stateAboutToEnter = "about_to_enter"
		stateAboutToLeave = "about_to_leave"
	)

	scanRow := func(row int) []interval {
		var ivs []interval
		var curr interval
		flush := func() {
			ivs = append(ivs, curr)
			curr = interval{}
		}
		currState := stateOut
		for x := minX; x <= maxX; x++ {
			if onEdge(grid.GP(x, row)) {
				if curr.isZero() {
					// we come from outside and hit the edge
					curr.min = x
					curr.max = x
					currState = stateAboutToEnter
				} else {
					// we are already in or on edge
					curr.max = x
					switch currState {
					case stateAboutToEnter:
						// stays - we are walking on the edge
					case stateIn:
						// we are about to leave
						currState = stateAboutToLeave
					case stateAboutToLeave:
						// in again
						currState = stateAboutToEnter

					default:
						log("oops")
					}
				}
			} else {
				// are we in?
				if !curr.isZero() {
					// now we are out
					switch currState {
					case stateAboutToEnter:
						// we are in
						curr.max = x
						currState = stateIn
					case stateIn:
						// we are staying in
						curr.max = x
					case stateAboutToLeave:
						// we are out
						currState = stateOut
						flush()

					default:
						log("oops")
					}

				} else {
					// we are not in
				}
			}
		}
		if !curr.isZero() {
			// now we are out
			flush()
		}

		return ivs
	}

	//scan horz lines
	rowIntervals := map[int][]interval{}
	for srow := minY - 1; srow <= maxY+1; srow++ {
		ivs := scanRow(srow)
		if len(ivs) > 0 {
			rowIntervals[srow] = ivs
		}
	}

	area := func(p1, p2 grid.GridPoint) int {
		dx := mathutil.Abs(p1.Col-p2.Col) + 1
		dy := mathutil.Abs(p1.Row-p2.Row) + 1
		return dx * dy
	}

	inAnyInterval := func(x int, ivs []interval) bool {
		for _, iv := range ivs {
			if x >= iv.min && x <= iv.max {
				return true
			}
		}
		return false
	}

	rectByRedAndGreen := func(row1, row2, col1, col2 int) bool {
		for row := row1; row <= row2; row++ {
			ivs, ok := rowIntervals[row]
			if !ok {
				return false
			}
			for col := col1; col <= col2; col++ {
				if !inAnyInterval(col, ivs) {
					return false
				}
			}
		}
		return true
	}

	var max int
	for _, p1 := range redPoints {
		for _, p2 := range redPoints {
			a := area(p1, p2)
			if a > max {
				// check if all covered points are red or green
				row1 := mathutil.Min(p1.Row, p2.Row)
				row2 := mathutil.Max(p1.Row, p2.Row)
				col1 := mathutil.Min(p1.Col, p2.Col)
				col2 := mathutil.Max(p1.Col, p2.Col)
				if rectByRedAndGreen(row1, row2, col1, col2) {
					max = a
				}
			}
		}
	}

	return max, nil
}

/*
..............
.......#XXX#..
.......XXXXX..
..OOOOOOOOXX..
..OOOOOOOOXX..
..OOOOOOOOXX..
.........XXX..
.........#X#..
..............
*/
