package day_08

import (
	"fmt"
	"strconv"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
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

type row struct {
	cols []int
}

type grid struct {
	rows []*row
}

func part1MainFunc(in string) (int, error) {
	g := &grid{}
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		row := &row{}
		for _, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				return -1, fmt.Errorf("nan in %q", line)
			}
			row.cols = append(row.cols, n)
		}
		if len(g.rows) > 0 && len(row.cols) != len(g.rows[0].cols) {
			return -1, fmt.Errorf("invalid row size")
		}
		g.rows = append(g.rows, row)
	}
	rowSize := len(g.rows[0].cols)

	isVisible := func(atRow, atCol int, height int) bool {
		visible := true
		for ir := 0; ir < atRow; ir++ {
			v := g.rows[ir].cols[atCol]
			if v >= height {
				visible = false
				break
			}
		}
		if visible {
			return true
		}
		//
		visible = true
		for ir := atRow + 1; ir < len(g.rows); ir++ {
			v := g.rows[ir].cols[atCol]
			if v >= height {
				visible = false
				break
			}
		}
		if visible {
			return true
		}
		//
		visible = true
		for ic := 0; ic < atCol; ic++ {
			v := g.rows[atRow].cols[ic]
			if v >= height {
				visible = false
				break
			}
		}
		if visible {
			return true
		}
		//
		visible = true
		for ic := atCol + 1; ic < rowSize; ic++ {
			v := g.rows[atRow].cols[ic]
			if v >= height {
				visible = false
				break
			}
		}
		return visible
	}

	var totalVisible int
	for ir, row := range g.rows {
		for ic, height := range row.cols {
			if isVisible(ir, ic, height) {
				totalVisible++
			}
		}
	}
	return totalVisible, nil
}

func part2MainFunc(in string) (int, error) {
	g := &grid{}
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		row := &row{}
		for _, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				return -1, fmt.Errorf("nan in %q", line)
			}
			row.cols = append(row.cols, n)
		}
		if len(g.rows) > 0 && len(row.cols) != len(g.rows[0].cols) {
			return -1, fmt.Errorf("invalid row size")
		}
		g.rows = append(g.rows, row)
	}
	rowSize := len(g.rows[0].cols)

	score := func(atRow, atCol int, height int) int {
		var scores [4]int
		s := 0
		for ir := atRow - 1; ir >= 0; ir-- {
			v := g.rows[ir].cols[atCol]
			s++
			if v >= height {
				break
			}
		}
		scores[0] = s
		//
		s = 0
		for ir := atRow + 1; ir < len(g.rows); ir++ {
			v := g.rows[ir].cols[atCol]
			s++
			if v >= height {
				break
			}
		}
		scores[1] = s
		//
		s = 0
		for ic := atCol - 1; ic >= 0; ic-- {
			v := g.rows[atRow].cols[ic]
			s++
			if v >= height {
				break
			}
		}
		scores[2] = s
		//
		s = 0
		for ic := atCol + 1; ic < rowSize; ic++ {
			v := g.rows[atRow].cols[ic]
			s++
			if v >= height {
				break
			}
		}
		scores[3] = s
		return scores[0] * scores[1] * scores[2] * scores[3]
	}

	var maxScore int
	for ir, row := range g.rows {
		for ic, height := range row.cols {
			s := score(ir, ic, height)
			if s > maxScore {
				maxScore = s
			}
		}
	}
	return maxScore, nil
}
