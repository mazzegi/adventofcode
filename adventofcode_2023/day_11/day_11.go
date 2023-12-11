package day_11

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/slices"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input, 1)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part1MainFunc(input, 999999)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

func expanded(drows []string) []string {
	rows := slices.Clone(drows)
	rowSize := len(rows[0])
	var expandRowsIdxs []int
	for i, row := range rows {
		if !strings.Contains(row, "#") {
			expandRowsIdxs = append(expandRowsIdxs, i)
		}
	}
	expRowOffset := 0
	for _, eri := range expandRowsIdxs {
		newRow := strings.Repeat(".", rowSize)
		rows = slices.Insert(rows, newRow, eri+expRowOffset)
		expRowOffset++
	}

	isEmptyCol := func(col int) bool {
		for _, row := range rows {
			if row[col] == '#' {
				return false
			}
		}
		return true
	}
	insertSpaceCol := func(col int) {
		for i, row := range rows {
			newRow := slices.Insert([]rune(row), '.', col)
			rows[i] = string(newRow)
		}
	}

	// expand cols
	for x := 0; x < rowSize; x++ {
		if isEmptyCol(x) {
			insertSpaceCol(x)
			rowSize++
			x += 1
		}
	}
	return rows
}

func emptyOnes(rows []string) (emptyRows []int, emptyCols []int) {
	for i, row := range rows {
		if !strings.Contains(row, "#") {
			emptyRows = append(emptyRows, i)
		}
	}
	isEmptyCol := func(col int) bool {
		for _, row := range rows {
			if row[col] == '#' {
				return false
			}
		}
		return true
	}
	rowSize := len(rows[0])
	for x := 0; x < rowSize; x++ {
		if isEmptyCol(x) {
			emptyCols = append(emptyCols, x)
		}
	}

	return
}

func part1MainFunc(in string, expFac int) (int, error) {
	rows := readutil.ReadLines(in)

	//find galaxies
	var galaxies []grid.Point
	for y, row := range rows {
		for x, col := range row {
			if col == '#' {
				galaxies = append(galaxies, grid.Pt(x, y))
			}
		}
	}
	//find empty
	emptyRows, emptyCols := emptyOnes(rows)

	// emptyRows and emptyCols are ordered!
	numBetween := func(vals []int, n1, n2 int) int {
		if n2 < n1 {
			n1, n2 = n2, n1
		}
		var nb int
		for _, v := range vals {
			if v > n1 && v < n2 {
				nb++
			}
		}
		return nb * expFac
	}

	var sum int
	for i := 0; i < len(galaxies); i++ {
		pi := galaxies[i]
		// iterate over all higher
		for j := i + 1; j < len(galaxies); j++ {
			pj := galaxies[j]
			dist := pi.ManhattenDistTo(pj)
			dist += numBetween(emptyRows, pi.Y, pj.Y)
			dist += numBetween(emptyCols, pi.X, pj.X)
			sum += dist
		}
	}
	// expand rows

	return sum, nil
}

// func part2MainFunc(in string, expFac int) (int, error) {
// 	rows := readutil.ReadLines(in)
// 	for i := 0; i < expFac; i++ {
// 		rows = expanded(rows)
// 	}

// 	//find galaxies
// 	var galaxies []grid.Point
// 	for y, row := range rows {
// 		for x, col := range row {
// 			if col == '#' {
// 				galaxies = append(galaxies, grid.Pt(x, y))
// 			}
// 		}
// 	}

// 	var sum int
// 	for i := 0; i < len(galaxies); i++ {
// 		pi := galaxies[i]
// 		// iterate over all higher
// 		for j := i + 1; j < len(galaxies); j++ {
// 			pj := galaxies[j]
// 			dist := pi.ManhattenDistTo(pj)
// 			sum += dist
// 		}
// 	}
// 	// expand rows

// 	return sum, nil
// }
