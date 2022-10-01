package day_18

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2016/errutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := safeTiles(input, 40)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := safeTiles(input, 400000)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type tile rune

const (
	safeTile tile = '.'
	trapTile tile = '^'
)

type roomRow struct {
	tiles []tile
}

type room struct {
	rows []*roomRow
}

func mustParseRow(s string) *roomRow {
	var row roomRow
	for _, r := range s {
		switch r {
		case rune(safeTile):
			row.tiles = append(row.tiles, safeTile)
		case rune(trapTile):
			row.tiles = append(row.tiles, trapTile)
		default:
			fatal("invalid tile %q", string(r))
		}
	}

	if len(row.tiles) == 0 {
		fatal("empty row")
	}
	return &row
}

func (r *room) safeTiles() int {
	var st int
	for _, row := range r.rows {
		for _, t := range row.tiles {
			if t == safeTile {
				st++
			}
		}
	}
	return st
}

func nextRow(row *roomRow) *roomRow {
	nr := &roomRow{
		tiles: make([]tile, len(row.tiles)),
	}
	left := func(idx int) tile {
		if idx-1 >= 0 {
			return row.tiles[idx-1]
		}
		return safeTile
	}
	right := func(idx int) tile {
		if idx+1 < len(row.tiles) {
			return row.tiles[idx+1]
		}
		return safeTile
	}

	for i := 0; i < len(row.tiles); i++ {
		lt := left(i)
		rt := right(i)
		ct := row.tiles[i]

		nextIsTrap := false
		switch {
		case lt == trapTile && ct == trapTile && rt == safeTile:
			nextIsTrap = true
		case lt == safeTile && ct == trapTile && rt == trapTile:
			nextIsTrap = true
		case lt == trapTile && ct == safeTile && rt == safeTile:
			nextIsTrap = true
		case lt == safeTile && ct == safeTile && rt == trapTile:
			nextIsTrap = true
		}
		if nextIsTrap {
			nr.tiles[i] = trapTile
		} else {
			nr.tiles[i] = safeTile
		}
	}

	return nr
}

func safeTiles(in string, rows int) (int, error) {
	row0 := mustParseRow(in)
	r := &room{
		rows: []*roomRow{row0},
	}
	for len(r.rows) < rows {
		nr := nextRow(r.rows[len(r.rows)-1])
		r.rows = append(r.rows, nr)
	}

	return r.safeTiles(), nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
