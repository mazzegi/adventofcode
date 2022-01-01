package lights

import (
	"bufio"
	"bytes"
	"strings"
)

type Cell struct {
	value     int
	nextValue int
}

type Grid struct {
	rows [][]*Cell
}

func NewGrid(dimX, dimY int) *Grid {
	g := &Grid{
		rows: make([][]*Cell, dimY),
	}
	for ri := 0; ri < dimY; ri++ {
		cells := make([]*Cell, dimX)
		for ci := 0; ci < dimX; ci++ {
			cells[ci] = &Cell{value: 0}
		}
		g.rows[ri] = cells
	}
	return g
}

func (g *Grid) PopulateFromString(s string) {
	scanner := bufio.NewScanner(bytes.NewBufferString(s))
	ri := 0
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		for ci, r := range l {
			if r == '#' {
				g.rows[ri][ci].value = 1
			} else {
				g.rows[ri][ci].value = 0
			}
		}
		ri++
	}
}

func (g *Grid) LigthingNeighbors(cri, cci int) int {
	var n int
	for ri := cri - 1; ri <= cri+1; ri++ {
		if ri < 0 || ri >= len(g.rows) {
			continue
		}
		row := g.rows[ri]
		for ci := cci - 1; ci <= cci+1; ci++ {
			if ci < 0 || ci >= len(row) {
				continue
			}
			if ri == cri && ci == cci {
				continue
			}
			cell := row[ci]
			if cell.value > 0 {
				n++
			}
		}
	}
	return n
}

func (g *Grid) Next() {
	for ri, row := range g.rows {
		for ci, cell := range row {
			ns := g.LigthingNeighbors(ri, ci)
			if cell.value > 0 {
				if ns == 2 || ns == 3 {
					//stays on
					cell.nextValue = 1
				} else {
					cell.nextValue = 0
				}
			} else {
				if ns == 3 {
					cell.nextValue = 1
				} else {
					cell.nextValue = 0
				}
			}
		}
	}

	for _, row := range g.rows {
		for _, cell := range row {
			cell.value = cell.nextValue
		}
	}
	g.TurnOnEdges()
}

func (g *Grid) TurnOnEdges() {
	row0 := g.rows[0]
	row0[0].value = 1
	row0[len(row0)-1].value = 1

	rowN := g.rows[len(g.rows)-1]
	rowN[0].value = 1
	rowN[len(row0)-1].value = 1
}

func (g *Grid) LightsOn() int {
	var lo int
	for _, row := range g.rows {
		for _, cell := range row {
			if cell.value > 0 {
				lo++
			}
		}
	}
	return lo
}
