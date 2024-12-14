package day_04

import (
	"fmt"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/stringutil"
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

func readWord(lines []string, at grid.Point, incX int, incY int, maxLength int) string {
	findRuneAt := func(pt grid.Point) (rune, bool) {
		if pt.Y < 0 || pt.Y >= len(lines) {
			return rune(0), false
		}
		line := lines[pt.Y]
		if pt.X < 0 || pt.X >= len(line) {
			return rune(0), false
		}
		return []rune(line)[pt.X], true
	}

	var word string
	curr := at
	for i := 0; i < maxLength; i++ {
		r, ok := findRuneAt(curr)
		if !ok {
			return word
		}
		word += string(r)
		curr.X += incX
		curr.Y += incY
	}

	return word
}

type direction struct {
	incX, incY int
}

func allDirections() []direction {
	return []direction{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
}

func part1MainFunc(in string) (int, error) {
	findWord := "XMAS"
	numFindWords := 0

	lines := readutil.ReadLines(in)
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			for _, dir := range allDirections() {
				gword := readWord(lines, grid.Pt(x, y), dir.incX, dir.incY, len(findWord))
				if gword == findWord {
					numFindWords++
				}
			}
		}
	}

	return numFindWords, nil
}

func part2MainFunc(in string) (int, error) {
	mas := "MAS"
	numXs := 0

	lines := readutil.ReadLines(in)
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			pt := grid.Pt(x, y)
			// now cross words
			w1 := readWord(lines, pt.Add(grid.Pt(-1, -1)), 1, 1, 3)
			w2 := stringutil.Reverse(w1)

			w3 := readWord(lines, pt.Add(grid.Pt(1, -1)), -1, 1, 3)
			w4 := stringutil.Reverse(w3)

			ok := (w1 == mas || w2 == mas) &&
				(w3 == mas || w4 == mas)
			if ok {
				numXs++
			}
		}
	}

	return numXs, nil
}
