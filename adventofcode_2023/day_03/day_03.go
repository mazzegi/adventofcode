package day_03

import (
	"fmt"
	"math"
	"strconv"
	"unicode"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
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

type NumberItem struct {
	Number    int
	NumDigits int
	Position  grid.Point
}

func NumberItemAdjacents(ni NumberItem) []grid.Point {
	var as []grid.Point
	for x := ni.Position.X - 1; x <= ni.Position.X+ni.NumDigits; x++ {
		for y := ni.Position.Y - 1; y <= ni.Position.Y+1; y++ {
			if y == ni.Position.Y && x >= ni.Position.X && x < ni.Position.X+ni.NumDigits {
				continue
			}
			as = append(as, grid.Pt(x, y))
		}
	}
	return as
}

type SymbolItem struct {
	Symbol   rune
	Position grid.Point
}

func mustParseLine(l string, ypos int) ([]NumberItem, []SymbolItem) {
	var numItems []NumberItem
	var symItems []SymbolItem

	var (
		currNum  []rune
		currNumX int
	)
	flushCurr := func() {
		if len(currNum) == 0 {
			return
		}
		n, err := strconv.Atoi(string(currNum))
		errutil.FatalWhen(err)

		numItems = append(numItems, NumberItem{
			Number:    n,
			NumDigits: len(currNum),
			Position:  grid.Pt(currNumX, ypos),
		})
		currNum = []rune{}
		currNumX = -1
	}

	for x, r := range l {
		switch {
		case r == '.':
			flushCurr()
		case unicode.IsDigit(r):
			if len(currNum) == 0 {
				currNumX = x
			}
			currNum = append(currNum, r)
			// symItems = append(symItems, SymbolItem{
			// 	Symbol:   r,
			// 	Position: grid.Pt(x, ypos),
			// })
		default: // a symbol
			flushCurr()
			symItems = append(symItems, SymbolItem{
				Symbol:   r,
				Position: grid.Pt(x, ypos),
			})
		}
	}
	flushCurr()
	return numItems, symItems
}

func mustParseItems(lines []string) ([]NumberItem, []SymbolItem) {
	var numItems []NumberItem
	var symItems []SymbolItem
	for y, line := range lines {
		lnumImtes, lsymItems := mustParseLine(line, y)
		numItems = append(numItems, lnumImtes...)
		symItems = append(symItems, lsymItems...)
	}
	return numItems, symItems
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	numItems, symItems := mustParseItems(lines)

	symPosMap := map[grid.Point]SymbolItem{}
	for _, si := range symItems {
		symPosMap[si.Position] = si
	}
	symAtAnyOf := func(pts []grid.Point) bool {
		for _, pt := range pts {
			if _, ok := symPosMap[pt]; ok {
				return true
			}
		}
		return false
	}

	var sum int
	for _, ni := range numItems {
		as := NumberItemAdjacents(ni)
		if symAtAnyOf(as) {
			sum += ni.Number
		}
	}

	return sum, nil
}

func pointsAdjacent(p1, p2 grid.Point) bool {
	d := p1.Sub(p2)
	nm := d.Norm()
	return mathutil.FloatsEqual(nm, 1.0, 1e-6) || mathutil.FloatsEqual(nm, math.Sqrt2, 1e-6)
}

func isNumItemAdjacentTo(ni NumberItem, pt grid.Point) bool {
	for i := 0; i < ni.NumDigits; i++ {
		test := grid.Pt(ni.Position.X+i, ni.Position.Y)
		if pointsAdjacent(test, pt) {
			return true
		}
	}
	return false
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	numItems, symItems := mustParseItems(lines)

	numItemsAdjacentTo := func(pt grid.Point) []NumberItem {
		var nis []NumberItem
		for _, ni := range numItems {
			if isNumItemAdjacentTo(ni, pt) {
				nis = append(nis, ni)
			}
		}
		return nis
	}

	//find gears
	var sum int
	for _, si := range symItems {
		if si.Symbol != '*' {
			continue
		}
		// a gear
		nis := numItemsAdjacentTo(si.Position)
		if len(nis) != 2 {
			continue
		}
		rat := nis[0].Number * nis[1].Number
		sum += rat
	}

	return sum, nil
}
