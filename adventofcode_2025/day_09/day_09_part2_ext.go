package day_09

import (
	"fmt"
	"time"

	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
)

func part2MainFuncV2(in string) (int, error) {
	t0 := time.Now()
	var redPoints []grid.GridPoint
	rowEdges := map[int][]int{} // row -> cols
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
		rowEdges[pt.Row] = append(rowEdges[pt.Row], pt.Col)
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
				rowEdges[pt.Row] = append(rowEdges[pt.Row], pt.Col)
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
				rowEdges[pt.Row] = append(rowEdges[pt.Row], pt.Col)
			}

		default:
			// neither
			return 0, fmt.Errorf("rp (%s) and rp_next (%s) are not on vert or horz line", rp, rpNext)
		}
	}
	log("determine green points (%d) in %s", len(greenPoints), time.Since(t0).Round(time.Microsecond))

	//

	return 0, nil
}
