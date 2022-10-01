package day_25

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := stepsToNoMoves(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type point struct {
	x, y int
}

func p(x, y int) point {
	return point{x: x, y: y}
}

const (
	faceEast  = "east"
	faceSouth = "south"
)

type cucumber struct {
	face string
	pos  point
	next point
}

type gridRow struct {
	positions []*cucumber
}

type grid struct {
	rows []*gridRow
}

func parseGridRow(in string, idx int) (*gridRow, error) {
	row := &gridRow{}
	for x, r := range in {
		switch r {
		case '>':
			row.positions = append(row.positions, &cucumber{
				face: faceEast,
				pos:  p(x, idx),
				next: p(-1, -1),
			})
		case 'v':
			row.positions = append(row.positions, &cucumber{
				face: faceSouth,
				pos:  p(x, idx),
				next: p(-1, -1),
			})
		case '.':
			row.positions = append(row.positions, nil)
		default:
			return nil, errors.Errorf("invalid rune %q", string(r))
		}
	}
	return row, nil
}

func parseGrid(in string) (*grid, error) {
	g := &grid{}
	lines := readutil.ReadLines(in)
	rowIdx := 0
	for _, line := range lines {
		row, err := parseGridRow(line, rowIdx)
		if err != nil {
			return nil, errors.Wrap(err, "parse-row")
		}
		if len(g.rows) > 0 && (len(g.rows[0].positions) != len(row.positions)) {
			return nil, errors.Errorf("invalid row len")
		}
		g.rows = append(g.rows, row)
		rowIdx++
	}
	if len(g.rows) == 0 || len(g.rows[0].positions) == 0 {
		return nil, errors.Errorf("no data")
	}
	return g, nil
}

func (g *grid) isFree(pt point) bool {
	if pt.y < 0 || pt.y >= len(g.rows) {
		panic("free")
	}
	row := g.rows[pt.y]
	if pt.x < 0 || pt.x >= len(row.positions) {
		panic("free")
	}
	return row.positions[pt.x] == nil
}

func (g *grid) set(pt point, c *cucumber) {
	if pt.y < 0 || pt.y >= len(g.rows) {
		panic("free")
	}
	row := g.rows[pt.y]
	if pt.x < 0 || pt.x >= len(row.positions) {
		panic("free")
	}
	row.positions[pt.x] = c
}

func (g *grid) step() (changes int) {
	moving := []*cucumber{}
	for _, row := range g.rows {
		for _, c := range row.positions {
			if c == nil {
				continue
			}
			if c.face != faceEast {
				continue
			}
			nextCand := p(c.pos.x+1, c.pos.y)
			if nextCand.x >= len(row.positions) {
				nextCand.x = 0
			}
			if !g.isFree(nextCand) {
				continue
			}
			c.next = nextCand
			moving = append(moving, c)
		}
	}
	for _, c := range moving {
		g.set(c.pos, nil)
		g.set(c.next, c)
		c.pos = c.next
		c.next = p(-1, -1)
	}
	changes += len(moving)

	//
	moving = []*cucumber{}
	for _, row := range g.rows {
		for _, c := range row.positions {
			if c == nil {
				continue
			}
			if c.face != faceSouth {
				continue
			}
			nextCand := p(c.pos.x, c.pos.y+1)
			if nextCand.y >= len(g.rows) {
				nextCand.y = 0
			}
			if !g.isFree(nextCand) {
				continue
			}
			c.next = nextCand
			moving = append(moving, c)
		}
	}
	for _, c := range moving {
		g.set(c.pos, nil)
		g.set(c.next, c)
		c.pos = c.next
		c.next = p(-1, -1)
	}
	changes += len(moving)

	return changes
}

func (g *grid) dump() string {
	var sl []string
	for _, row := range g.rows {
		var sr string
		for _, c := range row.positions {
			if c == nil {
				sr += "."
			} else if c.face == faceEast {
				sr += ">"
			} else {
				sr += "v"
			}
		}
		sl = append(sl, sr)
	}
	return strings.Join(sl, "\n")
}

//

func stepsToNoMoves(in string) (int, error) {
	g, err := parseGrid(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-grid")
	}
	steps := 0
	for {
		changes := g.step()
		log("**** after step %d *****\n%s", steps, g.dump())
		steps++
		if changes == 0 {
			break
		}
	}
	log("stalled after %d steps", steps)

	return steps, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
