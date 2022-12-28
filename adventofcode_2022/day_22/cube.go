package day_22

import (
	"strings"

	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/mathutil"
)

func P(x, y int) grid.Point {
	return grid.Pt(x, y)
}

type Connection struct {
	From     grid.Point
	To       grid.Point
	FromFace Face
	ToFace   Face
}

func mustParseCubeBoard(in string) *CubeBoard {
	var lines []string
	sl := strings.Split(in, "\n")
	for _, s := range sl {
		s = strings.Trim(s, "\r\n\t")
		if s == "" {
			continue
		}
		lines = append(lines, s)
	}

	cb := &CubeBoard{
		tiles:       map[grid.Point]rune{},
		connections: map[grid.Point]Connection{},
	}
	for iy, line := range lines {
		for ix, r := range line {
			p := grid.Pt(ix, iy)
			switch r {
			case Open, Wall:
				cb.tiles[p] = r
			}
			if ix == 0 && iy == 0 {
				cb.min = p
				cb.max = p
				continue
			}
			if ix < cb.min.X {
				cb.min.X = ix
			}
			if iy < cb.min.Y {
				cb.min.Y = iy
			}

			if ix > cb.max.X {
				cb.max.X = ix
			}
			if iy > cb.max.Y {
				cb.max.Y = iy
			}
		}
	}
	return cb
}

type CubeBoard struct {
	tiles       map[grid.Point]rune
	min         grid.Point
	max         grid.Point
	connections map[grid.Point]Connection
}

func (b *CubeBoard) Connect(from grid.Point, fromFace Face, to grid.Point, toFace Face) {
	b.connections[from] = Connection{
		From:     from,
		To:       to,
		FromFace: fromFace,
		ToFace:   toFace,
	}
	b.connections[to] = Connection{
		From:     to,
		To:       from,
		FromFace: toFace.Opposite(),
		ToFace:   fromFace.Opposite(),
	}
}

/*
cb.Connect(P(8, 0), Up, P(3, 4), Down)
	cb.Connect(P(9, 0), Up, P(2, 4), Down)
	cb.Connect(P(10, 0), Up, P(1, 4), Down)
	cb.Connect(P(11, 0), Up, P(0, 4), Down)

	cb.Connect(P(8, 0), Left, P(4, 4), Down)
	cb.Connect(P(8, 1), Left, P(5, 4), Down)
	cb.Connect(P(8, 2), Left, P(6, 4), Down)
	cb.Connect(P(8, 3), Left, P(7, 4), Down)
*/

func PointsBetween(p1, p2 grid.Point) []grid.Point {
	var bw []grid.Point
	if p1.Y == p2.Y {
		d := mathutil.Sign(p2.X - p1.X)
		bp := p1.Add(P(d, 0))
		for {
			if bp == p2 {
				break
			}
			bw = append(bw, bp)
			bp = bp.Add(P(d, 0))
		}
	} else if p1.X == p2.X {
		d := mathutil.Sign(p2.Y - p1.Y)
		bp := p1.Add(P(0, d))
		for {
			if bp == p2 {
				break
			}
			bw = append(bw, bp)
			bp = bp.Add(P(0, d))
		}
	} else {
		fatal("cannot between diagonal points")
	}
	return bw
}

func (b *CubeBoard) ConnectRange(fromEdgeP1, fromEdgeP2 grid.Point, fromFace Face, toEdgeP1, toEdgeP2 grid.Point, toFace Face) {
	fromBw := PointsBetween(fromEdgeP1, fromEdgeP2)
	toBw := PointsBetween(toEdgeP1, toEdgeP2)
	if len(fromBw) != len(toBw) {
		fatal("points between not equal")
	}
	b.Connect(fromEdgeP1, fromFace, toEdgeP1, toFace)
	for i, fromP := range fromBw {
		b.Connect(fromP, fromFace, toBw[i], toFace)
	}
	b.Connect(fromEdgeP2, fromFace, toEdgeP2, toFace)
}

func (b *CubeBoard) Start() grid.Point {
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

func (b *CubeBoard) firstInRow(y int) grid.Point {
	for x := b.min.X; x <= b.max.X; x++ {
		p := grid.Pt(x, y)
		if _, ok := b.tiles[p]; ok {
			return p
		}
	}
	return grid.Point{}
}

func (b *CubeBoard) lastInRow(y int) grid.Point {
	for x := b.max.X; x >= b.min.X; x-- {
		p := grid.Pt(x, y)
		if _, ok := b.tiles[p]; ok {
			return p
		}
	}
	return grid.Point{}
}

func (b *CubeBoard) firstInCol(x int) grid.Point {
	for y := b.min.Y; y <= b.max.Y; y++ {
		p := grid.Pt(x, y)
		if _, ok := b.tiles[p]; ok {
			return p
		}
	}
	return grid.Point{}
}

func (b *CubeBoard) lastInCol(x int) grid.Point {
	for y := b.max.Y; y >= b.min.Y; y-- {
		p := grid.Pt(x, y)
		if _, ok := b.tiles[p]; ok {
			return p
		}
	}
	return grid.Point{}
}

func (b *CubeBoard) NextPos(p grid.Point, f Face) (grid.Point, rune) {
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
		return next, tile
	}
	// we have to wrap
	switch f {
	case Right:
		next = b.firstInRow(p.Y)
	case Down:
		next = b.firstInCol(p.X)
	case Left:
		next = b.lastInRow(p.Y)
	case Up:
		next = b.lastInCol(p.X)
	}
	tile := b.tiles[next]
	return next, tile
}
