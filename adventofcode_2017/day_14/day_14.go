package day_14

import (
	"fmt"
	"math/bits"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2017/day_10/knot"
	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := usedSquares(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := numRegions(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func countOnes(bs []byte) int {
	var ones int
	for _, b := range bs {
		ones += bits.OnesCount8(b)
	}
	return ones
}

type bitfield []bool

func (bf bitfield) String() string {
	var s string
	for _, b := range bf {
		if b {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}

func byteToBits(b byte) bitfield {
	bits := make(bitfield, 8)
	for i := 0; i < 8; i++ {
		mask := byte(1 << i)
		if b&mask == mask {
			bits[len(bits)-1-i] = true
		}
	}
	return bits
}

func bytesToBits(bs []byte) bitfield {
	var bits bitfield
	for _, b := range bs {
		bits = append(bits, byteToBits(b)...)
	}
	return bits
}

func usedSquares(in string) (int, error) {
	rows := 128
	//hashes := make([]byte, rows)
	var onesTotal int
	for i := 0; i < rows; i++ {
		hin := fmt.Sprintf("%s-%d", in, i)
		hash, err := knot.Hash(hin)
		if err != nil {
			return 0, errors.Wrapf(err, "knot-hash %q", hin)
		}
		ones := countOnes(hash)
		onesTotal += ones
	}

	return onesTotal, nil
}

type gridCell struct {
	group int
	x, y  int
}

type gridRow struct {
	cells []*gridCell
}

type grid struct {
	rows []*gridRow
}

func (g *grid) findFirstCellWithoutGroup() (*gridCell, bool) {
	for _, row := range g.rows {
		for _, cell := range row.cells {
			if cell == nil {
				continue
			}
			if cell.group < 0 {
				return cell, true
			}
		}
	}
	return nil, false
}

func (g *grid) adjacents(x0, y0 int) []*gridCell {
	var as []*gridCell
	for y := y0 - 1; y <= y0+1; y++ {
		if y < 0 || y >= len(g.rows) || y == y0 {
			continue
		}
		row := g.rows[y]
		c := row.cells[x0]
		if c != nil {
			as = append(as, c)
		}
	}

	row := g.rows[y0]
	for x := x0 - 1; x <= x0+1; x++ {
		if x < 0 || x >= len(row.cells) || x == x0 {
			continue
		}
		c := row.cells[x]
		if c != nil {
			as = append(as, c)
		}
	}

	return as
}

func (g *grid) dumpGroups() string {
	var sl []string
	for _, row := range g.rows {
		var sr string
		for _, cell := range row.cells {
			if cell == nil {
				sr += "."
				continue
			}
			db := byte(65 + cell.group%24)
			sr += string(db)
		}
		sl = append(sl, sr)
	}
	return strings.Join(sl, "\n")
}

func numRegions(in string) (int, error) {

	g := &grid{}
	rows := 128
	for i := 0; i < rows; i++ {
		hin := fmt.Sprintf("%s-%d", in, i)
		hash, err := knot.Hash(hin)
		if err != nil {
			return 0, errors.Wrapf(err, "knot-hash %q", hin)
		}
		bf := bytesToBits(hash)
		if len(bf) != 128 {
			return 0, errors.Errorf("invalid size of bitfield")
		}
		row := &gridRow{}
		for x, b := range bf {
			if b {
				row.cells = append(row.cells, &gridCell{
					group: -1,
					x:     x,
					y:     i,
				})
			} else {
				row.cells = append(row.cells, nil)
			}
		}
		g.rows = append(g.rows, row)
	}

	var infectAdjacents func(cell *gridCell) (changed bool)
	infectAdjacents = func(cell *gridCell) (changed bool) {
		as := g.adjacents(cell.x, cell.y)
		for _, a := range as {
			if a.group == cell.group {
				continue
			}
			a.group = cell.group
			changed = true
			infectAdjacents(a)
		}
		return
	}

	group := 0
	for {
		cell, ok := g.findFirstCellWithoutGroup()
		if !ok {
			break
		}
		group++
		cell.group = group
		infectAdjacents(cell)
	}

	//
	log("*** map ***\n%s", g.dumpGroups())

	return group, nil
}
