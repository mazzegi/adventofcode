package grid

import (
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"
	"github.com/mazzegi/wasa/errors"
)

const (
	Set   rune = '#'
	Unset rune = '.'
)

func ParseBinaryGrid(s string) (*BinaryGrid, error) {
	bg := &BinaryGrid{}
	ls := readutil.ReadLines(s)
	for _, l := range ls {
		row := []bool{}
		for _, r := range l {
			switch r {
			case Set:
				row = append(row, true)
			case Unset:
				row = append(row, false)
			default:
				return nil, errors.Errorf("cannot handle rune %q", string(r))
			}
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

type BinaryGridRow []bool

type BinaryGrid struct {
	rows [][]bool
}

func (g *BinaryGrid) NumRows() int {
	return len(g.rows)
}

func (g *BinaryGrid) NumCols() int {
	return len(g.rows[0])
}

func (g *BinaryGrid) IsSet(col, row int) bool {
	return g.rows[row][col]
}

func (g *BinaryGrid) SetPoints() []Point {
	var sps []Point
	for ri, row := range g.rows {
		for ci, col := range row {
			if col {
				sps = append(sps, Pt(ci, ri))
			}
		}
	}
	return sps
}
