package day_12

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/algo/dijkstra"
	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

///

type Row struct {
	Cols []int
}

type Grid struct {
	Rows  []*Row
	Start grid.Point
	End   grid.Point
}

// impl Diskstra
func (g *Grid) Nodes() []grid.Point {
	var ps []grid.Point
	for ir, row := range g.Rows {
		for ic := range row.Cols {
			ps = append(ps, grid.Pt(ic, ir))
		}
	}
	return ps
}

func (g *Grid) Equal(t1, t2 grid.Point) bool {
	return t1 == t2
}

// func (g *Grid) AreNeighbours(t1, t2 grid.Point) bool {
// 	if t1.ManhattenDistTo(t2) != 1 {
// 		return false
// 	}
// 	e1 := g.Rows[t1.Y].Cols[t1.X]
// 	e2 := g.Rows[t2.Y].Cols[t2.X]
// 	return mathutil.Abs(e1-e2) <= 1
// }

// func (g *Grid) DistanceBetween(t1, t2 grid.Point) float64 {
// 	return float64(t1.ManhattenDistTo(t2))
// }

func (g *Grid) AreNeighbours(t1, t2 grid.Point) bool {
	return t1.ManhattenDistTo(t2) == 1
	// if t1.ManhattenDistTo(t2) != 1 {
	// 	return false
	// }
	// e1 := g.Rows[t1.Y].Cols[t1.X]
	// e2 := g.Rows[t2.Y].Cols[t2.X]
	// if e1+1 >= e2 {
	// 	return true
	// }
	// return false
}

func (g *Grid) DistanceBetween(t1, t2 grid.Point) float64 {
	e1 := g.Rows[t1.Y].Cols[t1.X]
	e2 := g.Rows[t2.Y].Cols[t2.X]
	if e1+1 >= e2 {
		return float64(t1.ManhattenDistTo(t2))
	}
	return float64(t1.ManhattenDistTo(t2)) + 10000000
}

//

func mustParse(in string) *Grid {
	g := &Grid{}
	for ir, line := range readutil.ReadLines(in) {
		row := &Row{}
		for ic, b := range []byte(line) {
			switch b {
			case 'S':
				g.Start = grid.Pt(ic, ir)
				row.Cols = append(row.Cols, int('a')-97)
			case 'E':
				g.End = grid.Pt(ic, ir)
				row.Cols = append(row.Cols, int('z')-97)
			default:
				e := int(b) - 97
				if e < 0 || e > 25 {
					fatal("invalid elevation value")
				}
				row.Cols = append(row.Cols, e)
			}
		}
		if len(g.Rows) > 0 && (len(g.Rows[0].Cols) != len(row.Cols)) {
			fatal("inconsistent row size")
		}
		g.Rows = append(g.Rows, row)
	}
	if len(g.Rows) == 0 {
		fatal("grid is empty")
	}
	return g
}

func part1MainFunc(in string) (int, error) {
	g := mustParse(in)
	path, err := dijkstra.ShortestPath[grid.Point](g, g.Start, g.End)
	if err != nil {
		fatal(err.Error())
	}
	//dumpGrid(g, path.Nodes)
	return len(path.Nodes) - 1, nil
}

func validPath(g *Grid, path []grid.Point) bool {
	for i := 0; i < len(path)-1; i++ {
		pc := path[i]
		pn := path[i+1]
		ec := g.Rows[pc.Y].Cols[pc.X]
		en := g.Rows[pn.Y].Cols[pn.X]
		if en > ec+1 {
			return false
		}
	}
	return true
}

func part2MainFunc(in string) (int, error) {
	g := mustParse(in)
	//collect 'a's
	as := []grid.Point{}
	for ir, row := range g.Rows {
		for ic, col := range row.Cols {
			if col == 0 {
				as = append(as, grid.Pt(ic, ir))
			}
		}
	}

	log("find shortest of %d", len(as))
	var shortest int
	for i, a := range as {
		path, err := dijkstra.ShortestPath[grid.Point](g, a, g.End)
		if err != nil {
			fatal(err.Error())
		}
		if !validPath(g, path.Nodes) {
			log("%d: -> invalid", i)
			continue
		}

		steps := len(path.Nodes) - 1
		if i == 0 || steps < shortest {
			shortest = steps
		}
		log("%d: -> %d", i, steps)
	}
	return shortest, nil
}

func dumpGrid(g *Grid, path []grid.Point) {
	onPath := func(p grid.Point) bool {
		for _, pp := range path {
			if pp == p {
				return true
			}
		}
		return false
	}

	for ir, row := range g.Rows {
		var srow string
		for ic, col := range row.Cols {
			sc := string(rune(col + 97))
			if onPath(grid.Pt(ic, ir)) {
				sc = strings.ToUpper(sc)
			}
			srow += sc
		}

		log("%s", srow)
	}
}
