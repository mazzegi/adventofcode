package hex

import (
	"strings"

	"github.com/pkg/errors"
)

const (
	East      = "e"
	SouthEast = "se"
	SouthWest = "sw"
	West      = "w"
	NorthWest = "nw"
	NorthEast = "ne"
)

func AllDirs() []string {
	return []string{
		East,
		SouthEast,
		SouthWest,
		West,
		NorthWest,
		NorthEast,
	}
}

func ParsePath(s string) (Path, error) {
	p := Path{}
	for len(s) > 0 {
		switch {
		case strings.HasPrefix(s, East):
			p.dirs = append(p.dirs, East)
			s = s[1:]
		case strings.HasPrefix(s, West):
			p.dirs = append(p.dirs, West)
			s = s[1:]
		case strings.HasPrefix(s, SouthEast):
			p.dirs = append(p.dirs, SouthEast)
			s = s[2:]
		case strings.HasPrefix(s, SouthWest):
			p.dirs = append(p.dirs, SouthWest)
			s = s[2:]
		case strings.HasPrefix(s, NorthWest):
			p.dirs = append(p.dirs, NorthWest)
			s = s[2:]
		case strings.HasPrefix(s, NorthEast):
			p.dirs = append(p.dirs, NorthEast)
			s = s[2:]
		default:
			return Path{}, errors.Errorf("invalid path seq %q", s)
		}
	}
	return p, nil
}

type Path struct {
	dirs []string
}

type Pos struct {
	X, Y int
}

func (p Pos) Step(dir string) Pos {
	switch dir {
	case East:
		return Pos{X: p.X + 1, Y: p.Y}
	case West:
		return Pos{X: p.X - 1, Y: p.Y}
	case SouthEast:
		if p.Y%2 == 0 {
			return Pos{X: p.X, Y: p.Y - 1}
		} else {
			return Pos{X: p.X + 1, Y: p.Y - 1}
		}
	case NorthEast:
		if p.Y%2 == 0 {
			return Pos{X: p.X, Y: p.Y + 1}
		} else {
			return Pos{X: p.X + 1, Y: p.Y + 1}
		}
	case SouthWest:
		if p.Y%2 == 0 {
			return Pos{X: p.X - 1, Y: p.Y - 1}
		} else {
			return Pos{X: p.X, Y: p.Y - 1}
		}
	case NorthWest:
		if p.Y%2 == 0 {
			return Pos{X: p.X - 1, Y: p.Y + 1}
		} else {
			return Pos{X: p.X, Y: p.Y + 1}
		}
	default:
		panic("invalid direction: " + dir)
	}
}

func (p Pos) Path(path Path) Pos {
	for _, dir := range path.dirs {
		p = p.Step(dir)
	}
	return p
}

type Grid struct {
	blackTiles map[Pos]bool
}

func NewGrid() *Grid {
	return &Grid{
		blackTiles: map[Pos]bool{},
	}
}

func (g *Grid) Flip(path Path) {
	fp := Pos{}.Path(path)
	if _, contains := g.blackTiles[fp]; contains {
		delete(g.blackTiles, fp)
	} else {
		g.blackTiles[fp] = true
	}
}

func (g *Grid) NumBlackTiles() int {
	return len(g.blackTiles)
}

func (g *Grid) AdjacentBlack(pos Pos) []Pos {
	var abs []Pos
	for _, dir := range AllDirs() {
		apos := pos.Step(dir)
		if _, contains := g.blackTiles[apos]; contains {
			abs = append(abs, apos)
		}
	}
	return abs
}

func (g *Grid) AdjacentWhite(pos Pos) []Pos {
	var aws []Pos
	for _, dir := range AllDirs() {
		apos := pos.Step(dir)
		if _, contains := g.blackTiles[apos]; !contains {
			aws = append(aws, apos)
		}
	}
	return aws
}

func (g *Grid) WhiteTiles() map[Pos]bool {
	wts := map[Pos]bool{}
	for pos := range g.blackTiles {
		aws := g.AdjacentWhite(pos)
		for _, wpos := range aws {
			wts[wpos] = true
		}
	}
	return wts
}

func (g *Grid) Next() {
	newBlackTiles := map[Pos]bool{}
	for pos := range g.blackTiles {
		adjBlacks := g.AdjacentBlack(pos)
		if len(adjBlacks) == 0 || len(adjBlacks) > 2 {
			//this flips back to white
			continue
		}
		newBlackTiles[pos] = true
	}
	//check white
	for pos := range g.WhiteTiles() {
		adjBlacks := g.AdjacentBlack(pos)
		if len(adjBlacks) == 2 {
			newBlackTiles[pos] = true
		}
	}
	g.blackTiles = newBlackTiles
}
