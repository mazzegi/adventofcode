package day_08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	width := 25
	height := 6
	res, err := part1MainFunc(input, width, height)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	width := 25
	height := 6
	res, err := part2MainFunc(input, width, height)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func part1MainFunc(in string, width, height int) (int, error) {
	img := parseImage(in, width, height)

	var min0DigCount int
	var min0DigLayer *layer
	for i, l := range img.layers {
		nullCount := l.digitCount(0)
		if i == 0 || nullCount < min0DigCount {
			min0DigCount = nullCount
			min0DigLayer = l
		}
	}
	cnt1 := min0DigLayer.digitCount(1)
	cnt2 := min0DigLayer.digitCount(2)

	return cnt1 * cnt2, nil
}

func part2MainFunc(in string, width, height int) (int, error) {
	img := parseImage(in, width, height)
	img.print()
	return 0, nil
}

func parseImage(in string, width, height int) *image {

	in = strings.TrimSpace(in)
	img := &image{}
	currRow := row{}
	currLayer := &layer{}
	for _, r := range in {
		n, err := strconv.Atoi(string(r))
		errutil.ExitOnErr(err)
		currRow = append(currRow, n)
		if len(currRow) >= width {
			currLayer.rows = append(currLayer.rows, currRow)
			currRow = row{}
			if len(currLayer.rows) >= height {
				img.layers = append(img.layers, currLayer)
				currLayer = &layer{}
			}
		}
	}

	if len(currRow) > 0 {
		currLayer.rows = append(currLayer.rows, currRow)
		img.layers = append(img.layers, currLayer)
	}
	return img
}

type row []int

type layer struct {
	rows []row
}

type image struct {
	layers []*layer
}

func (l *layer) digitCount(d int) int {
	var cnt int
	for _, r := range l.rows {
		for _, c := range r {
			if c == d {
				cnt++
			}
		}
	}
	return cnt
}

func (img *image) print() {
	pixelAt := func(ri, ci int) int {
		for _, l := range img.layers {
			p := l.rows[ri][ci]
			if p == 0 || p == 1 {
				return p
			}
		}
		return 2
	}

	prows := []string{}
	fl := img.layers[0]
	for ri, r := range fl.rows {
		prow := []string{}
		for ci := range r {
			p := pixelAt(ri, ci)
			switch p {
			case 0:
				prow = append(prow, " ")
			case 1:
				prow = append(prow, "#")
			default:
				prow = append(prow, "?")
			}
		}
		prows = append(prows, strings.Join(prow, ""))
	}
	for _, pr := range prows {
		fmt.Println(pr)
	}
}
