package day_22

import (
	"strings"

	"github.com/mazzegi/adventofcode/grid"
)

type PosFace struct {
	Pos  grid.Point
	Face Face
}

type Connection struct {
	FromPosFace PosFace
	ToPosFace   PosFace
}

func P(x, y int) grid.Point {
	return grid.Pt(x, y)
}

type Cube struct {
	tiles       map[grid.Point]rune
	min         grid.Point
	max         grid.Point
	connections map[PosFace]PosFace
}

func (c *Cube) Connect(from PosFace, to PosFace) {
	c.connections[from] = to

	rfrom := to
	rfrom.Face = to.Face.Opposite()
	rto := from
	rto.Face = from.Face.Opposite()
	c.connections[rfrom] = rto
}

func (c *Cube) ConnectRange(fromFace Face, fromP1, fromP2 grid.Point, toFace Face, toP1, toP2 grid.Point) {
	fromBw := PointsBetween(fromP1, fromP2)
	toBw := PointsBetween(toP1, toP2)
	if len(fromBw) != len(toBw) {
		fatal("points between not equal size")
	}
	for i, fromP := range fromBw {
		c.Connect(PosFace{fromP, fromFace}, PosFace{toBw[i], toFace})
	}
}

func (b *Cube) Start() grid.Point {
	y := 0
	for x := b.min.X; x <= b.max.X; x++ {
		p := grid.Pt(x, y)
		if t, ok := b.tiles[p]; ok && t == Open {
			return p
		}
	}
	fatal("found no start")
	return grid.Point{}
}

func (b *Cube) NextPos(p grid.Point, f Face) (grid.Point, Face, rune) {
	var next grid.Point
	switch f {
	case Right:
		next = p.Add(grid.Pt(1, 0))
	case Down:
		next = p.Add(grid.Pt(0, 1))
	case Left:
		next = p.Add(grid.Pt(-1, 0))
	case Up:
		next = p.Add(grid.Pt(0, -1))
	default:
		fatal("invalid face %q", f)
	}
	if tile, ok := b.tiles[next]; ok {
		return next, f, tile
	}

	//do we have a connection
	nextPF, ok := b.connections[PosFace{p, f}]
	if !ok {
		panic("outside board, but no connection found")
	}
	tile := b.tiles[nextPF.Pos]
	return nextPF.Pos, nextPF.Face, tile
}

//

func mustParseCube(in string) *Cube {
	var lines []string
	sl := strings.Split(in, "\n")
	for _, s := range sl {
		s = strings.Trim(s, "\r\n\t")
		if s == "" {
			continue
		}
		lines = append(lines, s)
	}

	c := &Cube{
		tiles:       map[grid.Point]rune{},
		connections: map[PosFace]PosFace{},
	}
	for iy, line := range lines {
		for ix, r := range line {
			p := grid.Pt(ix, iy)
			switch r {
			case Open, Wall:
				c.tiles[p] = r
			}
			if ix == 0 && iy == 0 {
				c.min = p
				c.max = p
				continue
			}
			if ix < c.min.X {
				c.min.X = ix
			}
			if iy < c.min.Y {
				c.min.Y = iy
			}

			if ix > c.max.X {
				c.max.X = ix
			}
			if iy > c.max.Y {
				c.max.Y = iy
			}
		}
	}
	return c
}
