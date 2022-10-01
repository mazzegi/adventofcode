package day_18

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2018/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2018/readutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	skip := false
	if skip {
		return
	}
	res, err := part1MainFunc(input, 1000)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

const (
	openGround rune = '.'
	trees      rune = '|'
	lumberYard rune = '#'
)

type point struct {
	x, y int
}

type Cell struct {
	value rune
	next  rune
}

type GridRow struct {
	cells []*Cell
}

type Grid struct {
	rows []*GridRow
}

func parseGrid(in string) (*Grid, error) {
	g := &Grid{}
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		row := &GridRow{}
		for _, r := range line {
			row.cells = append(row.cells, &Cell{value: r})
		}
		if len(g.rows) > 0 && len(g.rows[0].cells) != len(row.cells) {
			return nil, errors.Errorf("invalid row size. want %d, have %d", len(g.rows[0].cells), len(row.cells))
		}
		g.rows = append(g.rows, row)
	}
	return g, nil
}

func (g *Grid) adjacentCells(x, y int) []*Cell {
	var acs []*Cell
	for iy := y - 1; iy <= y+1; iy++ {
		if iy < 0 || iy >= len(g.rows) {
			continue
		}
		row := g.rows[iy]
		for ix := x - 1; ix <= x+1; ix++ {
			if ix == x && iy == y {
				continue
			}
			if ix < 0 || ix >= len(row.cells) {
				continue
			}
			acs = append(acs, row.cells[ix])
		}
	}
	return acs
}

func countValues(cs []*Cell, val rune) int {
	var cnt int
	for _, c := range cs {
		if c.value == val {
			cnt++
		}
	}
	return cnt
}

func (g *Grid) String() string {
	var sl []string
	for _, row := range g.rows {
		var rs string
		for _, cell := range row.cells {
			rs += string(cell.value)
		}
		sl = append(sl, rs)
	}
	return strings.Join(sl, "\n")
}

func (g *Grid) dump() {
	log("\n%s\n", g.String())
}

func (g *Grid) numOf(typ rune) int {
	var num int
	for _, row := range g.rows {
		for _, cell := range row.cells {
			if cell.value == typ {
				num++
			}
		}
	}
	return num
}

func (g *Grid) tick() {
	for y, row := range g.rows {
		for x, cell := range row.cells {
			acs := g.adjacentCells(x, y)
			cell.next = cell.value
			switch cell.value {
			case openGround:
				if countValues(acs, trees) >= 3 {
					cell.next = trees
				}
			case trees:
				if countValues(acs, lumberYard) >= 3 {
					cell.next = lumberYard
				}
			case lumberYard:
				als := countValues(acs, lumberYard)
				ats := countValues(acs, trees)
				if !(als >= 1 && ats >= 1) {
					cell.next = openGround
				}
			}
		}
	}
	for _, row := range g.rows {
		for _, cell := range row.cells {
			cell.value = cell.next
		}
	}
}

func (g *Grid) hash() string {
	return g.String()
	// var hs []byte
	// for _, row := range g.rows {
	// 	for _, cell := range row.cells {
	// 		hs = append(hs, byte(cell.value))
	// 	}
	// }
	// return fmt.Sprintf("%x", md5.Sum(hs))
}

func part1MainFunc(in string, ticks int) (int, error) {
	g, err := parseGrid(in)
	if err != nil {
		return 0, err
	}
	//g.dump()
	for i := 0; i < ticks; i++ {
		g.tick()
		//g.dump()
	}

	return g.numOf(trees) * g.numOf(lumberYard), nil
}

func part2MainFunc(in string) (int, error) {
	g, err := parseGrid(in)
	if err != nil {
		return 0, err
	}
	hashes := map[string]int{}
	num := 0
	var repeated int
	hashes[g.hash()] = num
	for {
		g.tick()
		num++
		hs := g.hash()
		if i, ok := hashes[hs]; ok {
			repeated = i
			break
		}
		hashes[g.hash()] = num
	}
	period := num - repeated
	log("repeated %d => %d; period = %d", num, repeated, period)

	total := 1000000000
	for {
		if num+period < total {
			num += period
		} else {
			break
		}
	}
	for num < total {
		num++
		g.tick()
	}

	//total := 1000000000
	// total := 1000
	// rem := total - (num-repeated)*(total/(num-repeated))
	// log("rem: %d", rem)
	// for i := 0; i < rem; i++ {
	// 	g.tick()
	// }

	return g.numOf(trees) * g.numOf(lumberYard), nil
}
