package day_09

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := sumOfRiskLevels(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := productOf3LargestBasinSizes(input)
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

func (g *grid) findValue(rowIdx int, colIdx int) (int, bool) {
	if rowIdx < 0 || rowIdx >= len(g.rows) {
		return 0, false
	}
	row := g.rows[rowIdx]
	if colIdx < 0 || colIdx >= len(row.values) {
		return 0, false
	}
	return row.values[colIdx], true
}

func (g *grid) adjacents(rowIdx int, colIdx int) []int {
	var as []int
	if v, ok := g.findValue(rowIdx, colIdx-1); ok {
		as = append(as, v)
	}
	if v, ok := g.findValue(rowIdx, colIdx+1); ok {
		as = append(as, v)
	}
	if v, ok := g.findValue(rowIdx-1, colIdx); ok {
		as = append(as, v)
	}
	if v, ok := g.findValue(rowIdx+1, colIdx); ok {
		as = append(as, v)
	}
	return as
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

func sumOfRiskLevels(in string) (int, error) {
	g, err := parseGrid(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse grid")
	}
	if g.isEmpty() {
		return 0, errors.Errorf("grid is empty")
	}

	isLow := func(val int, adjs []int) bool {
		for _, a := range adjs {
			if val >= a {
				return false
			}
		}
		return true
	}

	var sum int
	for rowIdx, row := range g.rows {
		for colIdx, val := range row.values {
			adjs := g.adjacents(rowIdx, colIdx)
			if len(adjs) < 2 {
				return 0, errors.Errorf("invalid len of adjacents %d", len(adjs))
			}
			if isLow(val, adjs) {
				sum += val + 1
			}
		}
	}

	return sum, nil
}

////

type point struct {
	x, y int
}

func p(x, y int) point {
	return point{
		x: x,
		y: y,
	}
}

type basin struct {
	points []point
}

func (b *basin) size() int {
	return len(b.points)
}

func (b *basin) isEmpty() bool {
	return b.size() == 0
}

func (b *basin) add(pts ...point) {
	b.points = append(b.points, pts...)
}

func (b *basin) purge() {
	b.points = []point{}
}

func productOf3LargestBasinSizes(in string) (int, error) {
	g, err := parseGrid(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse grid")
	}
	if g.isEmpty() {
		return 0, errors.Errorf("grid is empty")
	}

	basins := findBasins(g)
	if len(basins) < 3 {
		return 0, errors.Errorf("found only %d basins", len(basins))
	}
	sort.Slice(basins, func(i, j int) bool {
		return basins[i].size() > basins[j].size()
	})
	prod := basins[0].size() * basins[1].size() * basins[2].size()
	return prod, nil
}

func findBasins(g *grid) []*basin {
	var basins []*basin

	handlePoint := func(pt point) {
		for _, b := range basins {
			if canMergePointToBasin(b, pt) {
				b.add(pt)
				return
			}
		}
		//no basin matched - create new one
		basins = append(basins, &basin{
			points: []point{pt},
		})
	}

	fmt.Printf("find basins (rows = %d, cols = %d)\n", len(g.rows), len(g.rows[0].values))
	for rowIdx, row := range g.rows {
		fmt.Printf("process row %d\n", rowIdx)
		for colIdx, val := range row.values {
			if val == 9 {
				continue
			}
			pt := p(colIdx, rowIdx)
			handlePoint(pt)
			//			basins = mergedBasins(basins)
		}
	}

	fmt.Printf("merge basins ...\n")
	basins = mergedBasins(basins)
	fmt.Printf("merge basins ... done\n")

	return basins
}

func arePointsAdjacent(p1, p2 point) bool {
	if p1.x == p2.x {
		return p2.y == p1.y+1 || p2.y == p1.y-1
	}
	if p1.y == p2.y {
		return p2.x == p1.x+1 || p2.x == p1.x-1
	}
	return false
}

func canMergePointToBasin(b *basin, pt point) bool {
	for _, bpt := range b.points {
		if arePointsAdjacent(bpt, pt) {
			return true
		}
	}
	return false
}

func canMergeBasins(b1, b2 *basin) bool {
	for _, p2 := range b2.points {
		if canMergePointToBasin(b1, p2) {
			return true
		}
	}
	return false
}

func mergedBasins(basins []*basin) []*basin {
	for i1, b1 := range basins {
		if b1.isEmpty() {
			continue
		}
		for i2, b2 := range basins {
			if i1 == i2 {
				continue
			}
			if b2.isEmpty() {
				continue
			}
			if canMergeBasins(b1, b2) {
				b1.add(b2.points...)
				b2.purge()
			}
		}
	}
	var mbs []*basin
	for _, b := range basins {
		if b.isEmpty() {
			continue
		}
		mbs = append(mbs, b)
	}
	return mbs
}
