package grid

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"
	"github.com/mazzegi/wasa/errors"
)

type GridPoint struct {
	Col, Row int
}

func GP(col, row int) GridPoint {
	return GridPoint{Col: col, Row: row}
}

func Parse(s string) (*Grid, error) {
	bg := &Grid{}
	ls := readutil.ReadLines(s)
	for _, l := range ls {
		if l == "" {
			continue
		}
		row := GridRow{}
		for _, r := range l {
			row = append(row, r)
		}
		if len(row) == 0 {
			continue
		}
		if len(bg.rows) > 0 && len(row) != len(bg.rows[0]) {
			return nil, errors.Errorf("invalid row size")
		}
		bg.rows = append(bg.rows, row)
	}
	if len(bg.rows) == 0 {
		return nil, errors.Errorf("grid is empty")
	}
	return bg, nil
}

type GridRow []rune

type Grid struct {
	rows []GridRow
}

func (g *Grid) NumRows() int {
	return len(g.rows)
}

func (g *Grid) NumCols() int {
	return len(g.rows[0])
}

func (g *Grid) At(p GridPoint) rune {
	if p.Row < 0 || p.Row >= len(g.rows) {
		panic(fmt.Sprintf("row index %d out of bounds [0, %d]", p.Row, len(g.rows)-1))
	}
	row := g.rows[p.Row]
	if p.Col < 0 || p.Col >= len(row) {
		panic(fmt.Sprintf("col index %d out of bounds [0, %d]", p.Col, len(row)-1))
	}
	return row[p.Col]
}

func (g *Grid) FindFirst(r rune) (GridPoint, bool) {
	for ir, row := range g.rows {
		for ic, rc := range row {
			if rc == r {
				return GP(ic, ir), true
			}
		}
	}
	return GridPoint{}, false
}

func (g *Grid) Contains(p GridPoint) bool {
	if p.Row < 0 || p.Row >= len(g.rows) {
		return false
	}
	row := g.rows[p.Row]
	if p.Col < 0 || p.Col >= len(row) {
		return false
	}
	return true
}
