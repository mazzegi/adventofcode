package day_11

import (
	"adventofcode_2021/errutil"
	"adventofcode_2021/readutil"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := totalFlashes(input, 100)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := stepsToAllFlash(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

type gridRow struct {
	values []int
}

type grid struct {
	rows []*gridRow
}

func (g *grid) isEmpty() bool {
	if len(g.rows) == 0 {
		return true
	}
	return len(g.rows[0].values) == 0
}

func (g *grid) totalElts() int {
	if g.isEmpty() {
		return 0
	}
	rowSize := len(g.rows[0].values)
	return rowSize * len(g.rows)
}

func (g *grid) format() string {
	var sl []string
	for _, row := range g.rows {
		var s string
		for _, v := range row.values {
			if v > 9 {
				s += "#"
			} else {
				s += strconv.Itoa(v)
			}
		}
		sl = append(sl, s)
	}
	return strings.Join(sl, "\n")
}

func parseGridRow(s string) (*gridRow, error) {
	row := &gridRow{}
	for _, r := range s {
		n, err := strconv.ParseInt(string(r), 10, 8)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-int %q", string(r))
		}
		row.values = append(row.values, int(n))
	}
	return row, nil
}

func parseGrid(in string) (*grid, error) {
	g := &grid{}
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		row, err := parseGridRow(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse row %q", line)
		}
		if len(g.rows) > 0 && len(g.rows[0].values) != len(row.values) {
			return nil, errors.Errorf("inconsistent row-len: want %d, have %d", len(g.rows[0].values), len(row.values))
		}
		g.rows = append(g.rows, row)
	}
	return g, nil
}

func totalFlashes(in string, steps int) (int, error) {
	g, err := parseGrid(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse grid")
	}
	if g.isEmpty() {
		return 0, errors.Errorf("grid is empty")
	}

	//fmt.Printf("before any steps:\n%s\n", g.format())
	var total int
	for step := 0; step < steps; step++ {
		flashes := gridStep(g, step+1)
		total += flashes
	}

	return total, nil
}

type point struct {
	x, y int
}

func p(x, y int) point {
	return point{x: x, y: y}
}

func gridStep(g *grid, step int) (flashes int) {
	for _, row := range g.rows {
		for colIdx, val := range row.values {
			row.values[colIdx] = val + 1
		}
	}

	flashed := map[point]bool{}
	for rowIdx, row := range g.rows {
		for colIdx, val := range row.values {
			if val > 9 {
				flash(g, rowIdx, colIdx, flashed)
			}
		}
	}

	flashes = len(flashed)
	//fmt.Printf("\nstep %03d: after flushing (%d):\n%s\n", step, flashes, g.format())

	for _, row := range g.rows {
		for colIdx, val := range row.values {
			if val > 9 {
				row.values[colIdx] = 0
			}
		}
	}

	return flashes
}

func flash(g *grid, rowIdx, colIdx int, flashed map[point]bool) {
	if _, ok := flashed[p(colIdx, rowIdx)]; ok {
		return
	}
	flashed[p(colIdx, rowIdx)] = true

	for y := rowIdx - 1; y <= rowIdx+1; y++ {
		if y < 0 || y >= len(g.rows) {
			continue
		}
		row := g.rows[y]
		for x := colIdx - 1; x <= colIdx+1; x++ {
			if y == rowIdx && x == colIdx {
				continue
			}
			if x < 0 || x >= len(row.values) {
				continue
			}
			val := row.values[x]
			row.values[x] = val + 1
			if row.values[x] > 9 {
				flash(g, y, x, flashed)
			}
		}
	}
}

func stepsToAllFlash(in string) (int, error) {
	g, err := parseGrid(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse grid")
	}
	if g.isEmpty() {
		return 0, errors.Errorf("grid is empty")
	}

	maxSteps := 10000
	gsize := g.totalElts()
	for step := 0; step < maxSteps; step++ {
		flashes := gridStep(g, step+1)
		if flashes == gsize {
			return step + 1, nil
		}
	}

	return 0, errors.Errorf("no flashes after %d steps", maxSteps)
}
