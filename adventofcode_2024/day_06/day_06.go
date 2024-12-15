package day_06

import (
	"fmt"
	"slices"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

type Grid struct {
	rows [][]bool
}

func (g *Grid) clone() *Grid {
	cg := &Grid{
		rows: make([][]bool, len(g.rows)),
	}
	for i, row := range g.rows {
		cg.rows[i] = slices.Clone(row)
	}
	return cg
}

func (g *Grid) contains(pos grid.Point) bool {
	return pos.X >= 0 && pos.X < len(g.rows[0]) &&
		pos.Y >= 0 && pos.Y < len(g.rows)
}

func (g *Grid) isObstacleAt(pt grid.Point) bool {
	if !g.contains(pt) {
		return false
	}
	return g.rows[pt.Y][pt.X]
}

func parseGrid(in string) (g *Grid, guardPos grid.Point, err error) {
	g = &Grid{}
	lines := readutil.ReadLines(in)
	for y, line := range lines {
		row := []bool{}
		for x, r := range line {
			switch r {
			case '#':
				row = append(row, true)
			case '.':
				row = append(row, false)
			case '^':
				guardPos = grid.Pt(x, y)
				row = append(row, false)
			default:
				return nil, grid.Point{}, fmt.Errorf("cannot handle rune %q", string(r))
			}
		}
		if len(row) == 0 {
			continue
		}
		if len(g.rows) > 0 && len(row) != len(g.rows[0]) {
			return nil, grid.Point{}, fmt.Errorf("invalid row size")
		}
		g.rows = append(g.rows, row)
	}
	return
}

const (
	north = "N"
	south = "S"
	west  = "W"
	east  = "E"
)

func part1MainFunc(in string) (int, error) {
	g, gpos, err := parseGrid(in)
	if err != nil {
		return 0, fmt.Errorf("parse-grid: %w", err)
	}
	visited := set.New[grid.Point]()
	visited.Insert(gpos)

	dir := north
	move := func() {
		var nextInDir grid.Point
		var dirBy90Deg string
		switch dir {
		case north:
			nextInDir = gpos.Add(grid.Pt(0, -1))
			dirBy90Deg = east
		case south:
			nextInDir = gpos.Add(grid.Pt(0, 1))
			dirBy90Deg = west
		case west:
			nextInDir = gpos.Add(grid.Pt(-1, 0))
			dirBy90Deg = north
		case east:
			nextInDir = gpos.Add(grid.Pt(1, 0))
			dirBy90Deg = south
		default:
			panic("invalid direction " + dir)
		}
		if !g.isObstacleAt(nextInDir) {
			gpos = nextInDir
			return
		}
		// turn right by 90 degree
		dir = dirBy90Deg
	}

	for {
		move()
		if g.contains(gpos) {
			visited.Insert(gpos)
		} else {
			break
		}
	}
	dposs := visited.Count()

	return dposs, nil
}

func part2MainFunc(in string) (int, error) {
	g, gpos, err := parseGrid(in)
	if err != nil {
		return 0, fmt.Errorf("parse-grid: %w", err)
	}

	var loopPoss int
	for y := 0; y < len(g.rows); y++ {
		for x := 0; x < len(g.rows[y]); x++ {
			pt := grid.Pt(x, y)
			if pt == gpos {
				continue
			}
			if g.isObstacleAt(pt) {
				continue
			}
			// put an obstacle and see if guard loops
			cg := g.clone()
			cg.rows[y][x] = true
			// probe
			if guardLoops(cg, gpos, north) {
				log("ok: obstacle at %s", pt)
				loopPoss++
			}
		}
	}

	return loopPoss, nil
}

type posDir struct {
	pos grid.Point
	dir string
}

func guardLoops(g *Grid, gpos grid.Point, dir string) bool {
	visitedPosDirs := set.New[posDir]()
	visitedPosDirs.Insert(posDir{
		pos: gpos,
		dir: dir,
	})

	move := func() {
		var nextInDir grid.Point
		var dirBy90Deg string
		switch dir {
		case north:
			nextInDir = gpos.Add(grid.Pt(0, -1))
			dirBy90Deg = east
		case south:
			nextInDir = gpos.Add(grid.Pt(0, 1))
			dirBy90Deg = west
		case west:
			nextInDir = gpos.Add(grid.Pt(-1, 0))
			dirBy90Deg = north
		case east:
			nextInDir = gpos.Add(grid.Pt(1, 0))
			dirBy90Deg = south
		default:
			panic("invalid direction " + dir)
		}
		if !g.isObstacleAt(nextInDir) {
			gpos = nextInDir
			return
		}
		// turn right by 90 degree
		dir = dirBy90Deg
	}

	for {
		move()
		if !g.contains(gpos) {
			return false
		}
		pd := posDir{
			pos: gpos,
			dir: dir,
		}
		if visitedPosDirs.Contains(pd) {
			return true
		}
		visitedPosDirs.Insert(pd)
	}
}
