package day_22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatalIfErr(err error) {
	if err == nil {
		return
	}
	fatal("err not nil: %v", err)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input, inputPath)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	cube := mustParseCube(input)
	res, err := part2MainFunc(cube, inputPath)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type PathPair struct {
	Moves int
	Turn  string
}

type Path []PathPair

// "10R5L5R10L4R5L5"
func mustParsePath(in string) Path {
	in = strings.TrimSpace(in)
	var p Path
	for {
		tix := strings.IndexAny(in, "RL")
		if tix < 0 {
			n, err := strconv.ParseInt(in, 10, 64)
			fatalIfErr(err)
			p = append(p, PathPair{Moves: int(n)})
			break
		}
		smv := in[:tix]
		strn := string(in[tix])
		n, err := strconv.ParseInt(smv, 10, 64)
		fatalIfErr(err)
		p = append(p, PathPair{Moves: int(n), Turn: strn})
		in = in[tix+1:]
	}
	return p
}

const (
	Open rune = '.'
	Wall rune = '#'
)

type Board struct {
	tiles map[grid.Point]rune
	min   grid.Point
	max   grid.Point
}

func (b *Board) Start() grid.Point {
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

func (b *Board) firstInRow(y int) grid.Point {
	for x := b.min.X; x <= b.max.X; x++ {
		p := grid.Pt(x, y)
		if _, ok := b.tiles[p]; ok {
			return p
		}
	}
	return grid.Point{}
}

func (b *Board) lastInRow(y int) grid.Point {
	for x := b.max.X; x >= b.min.X; x-- {
		p := grid.Pt(x, y)
		if _, ok := b.tiles[p]; ok {
			return p
		}
	}
	return grid.Point{}
}

func (b *Board) firstInCol(x int) grid.Point {
	for y := b.min.Y; y <= b.max.Y; y++ {
		p := grid.Pt(x, y)
		if _, ok := b.tiles[p]; ok {
			return p
		}
	}
	return grid.Point{}
}

func (b *Board) lastInCol(x int) grid.Point {
	for y := b.max.Y; y >= b.min.Y; y-- {
		p := grid.Pt(x, y)
		if _, ok := b.tiles[p]; ok {
			return p
		}
	}
	return grid.Point{}
}

func (b *Board) NextPos(p grid.Point, f Face) (grid.Point, rune) {
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

//

func mustParseBoard(in string) *Board {
	var lines []string
	sl := strings.Split(in, "\n")
	for _, s := range sl {
		s = strings.Trim(s, "\r\n\t")
		if s == "" {
			continue
		}
		lines = append(lines, s)
	}

	m := &Board{
		tiles: map[grid.Point]rune{},
	}
	for iy, line := range lines {
		for ix, r := range line {
			p := grid.Pt(ix, iy)
			switch r {
			case Open, Wall:
				m.tiles[p] = r
			}
			if ix == 0 && iy == 0 {
				m.min = p
				m.max = p
				continue
			}
			if ix < m.min.X {
				m.min.X = ix
			}
			if iy < m.min.Y {
				m.min.Y = iy
			}

			if ix > m.max.X {
				m.max.X = ix
			}
			if iy > m.max.Y {
				m.max.Y = iy
			}
		}
	}
	return m
}

type Face string

const (
	Right Face = "R"
	Down  Face = "D"
	Left  Face = "L"
	Up    Face = "U"
)

func (f Face) Score() int {
	switch f {
	case Right:
		return 0
	case Down:
		return 1
	case Left:
		return 2
	case Up:
		return 3
	default:
		fatal("invalid face %q", f)
	}
	return 0
}

func (f Face) Opposite() Face {
	switch f {
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	case Up:
		return Down
	default:
		fatal("invalid face %q", f)
	}
	return ""
}

func (f Face) Turn(t string) Face {
	if t == "R" {
		switch f {
		case Right:
			return Down
		case Down:
			return Left
		case Left:
			return Up
		case Up:
			return Right
		default:
			fatal("invalid face %q", f)
		}
	} else { //L
		switch f {
		case Right:
			return Up
		case Down:
			return Right
		case Left:
			return Down
		case Up:
			return Left
		default:
			fatal("invalid face %q", f)
		}
	}
	fatal("invalid turn %q", t)
	return ""
}

func part1MainFunc(in string, inPath string) (int, error) {
	path := mustParsePath(inPath)
	board := mustParseBoard(in)

	currPos := board.Start()
	currFace := Right
	for _, pp := range path {
		// go ahead number of moves
		for imv := 0; imv < pp.Moves; imv++ {
			nextPos, nextTile := board.NextPos(currPos, currFace)
			if nextTile == Open {
				currPos = nextPos
			} else {
				break
			}
		}
		//
		if pp.Turn != "" {
			currFace = currFace.Turn(pp.Turn)
		}
	}
	pwd := 1000*(currPos.Y+1) + 4*(currPos.X+1) + currFace.Score()

	return pwd, nil
}

func part2MainFunc(cube *Cube, inPath string) (int, error) {
	path := mustParsePath(inPath)

	currPos := cube.Start()
	currFace := Right
	for _, pp := range path {
		// go ahead number of moves
		for imv := 0; imv < pp.Moves; imv++ {
			nextPos, nextFace, nextTile := cube.NextPos(currPos, currFace)
			if nextTile == Open {
				currPos = nextPos
				currFace = nextFace
			} else {
				break
			}
		}
		//
		if pp.Turn != "" {
			currFace = currFace.Turn(pp.Turn)
		}
	}
	pwd := 1000*(currPos.Y+1) + 4*(currPos.X+1) + currFace.Score()

	return pwd, nil
}
