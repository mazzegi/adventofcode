package day_13

import (
	"adventofcode_2021/errutil"
	"adventofcode_2021/readutil"
	"fmt"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := dotsAfterFirstFold(inputDots, inputFolds)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	_, err := displayCode(inputDots, inputFolds)
	errutil.ExitOnErr(err)
}

//
type point struct {
	x, y int
}

func p(x, y int) point {
	return point{x: x, y: y}
}

const (
	foldX = "x"
	foldY = "y"
)

type fold struct {
	foldAlong string
	foldIndex int
}

func parseFold(s string) (fold, error) {
	var f fold
	_, err := fmt.Sscanf(s, "fold along %s = %d", &f.foldAlong, &f.foldIndex)
	if err != nil {
		return fold{}, errors.Wrapf(err, "scan-fold %q", s)
	}
	return f, nil
}

func parseFolds(in string) ([]fold, error) {
	var fs []fold
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		f, err := parseFold(line)
		if err != nil {
			return nil, errors.Wrap(err, "parse-fold")
		}
		fs = append(fs, f)
	}
	return fs, nil
}

func parsePoint(s string) (point, error) {
	var p point
	_, err := fmt.Sscanf(s, "%d,%d", &p.x, &p.y)
	if err != nil {
		return point{}, errors.Wrapf(err, "scan-point %q", s)
	}
	return p, nil
}

func parsePoints(in string) ([]point, error) {
	var ps []point
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		p, err := parsePoint(line)
		if err != nil {
			return nil, errors.Wrap(err, "parse-point")
		}
		ps = append(ps, p)
	}
	return ps, nil
}

//
type planeRow struct {
	dots []bool
}

type plane struct {
	//rows []*planeRow
	xmax int
	ymax int

	dots map[point]bool
}

func buildPlane(points []point) *plane {
	pl := &plane{
		dots: map[point]bool{},
	}
	for _, pt := range points {
		pl.add(pt)
	}
	return pl
}

func (pl *plane) add(pt point) {
	pl.dots[pt] = true
	if pt.x > pl.xmax {
		pl.xmax = pt.x
	}
	if pt.y > pl.ymax {
		pl.ymax = pt.y
	}
}

func (pl *plane) visibleDots() int {
	return len(pl.dots)
}

func (pl *plane) folded(f fold) (*plane, error) {
	if f.foldAlong == foldX {
		return pl.foldedx(f.foldIndex)
	}
	return pl.foldedy(f.foldIndex)
}

func (pl *plane) foldedx(idx int) (*plane, error) {
	fpl := &plane{
		dots: map[point]bool{},
	}
	for pt := range pl.dots {
		if pt.x < idx {
			fpl.add(pt)
		} else if pt.x > idx {
			mpt := p(idx-(pt.x-idx), pt.y)
			fpl.add(mpt)
		} else {
			return nil, errors.Errorf("dot on fold line")
		}
	}

	return fpl, nil
}

func (pl *plane) foldedy(idx int) (*plane, error) {
	fpl := &plane{
		dots: map[point]bool{},
	}
	for pt := range pl.dots {
		if pt.y < idx {
			fpl.add(pt)
		} else if pt.y > idx {
			mpt := p(pt.x, idx-(pt.y-idx))
			fpl.add(mpt)
		} else {
			return nil, errors.Errorf("dot on fold line")
		}
	}

	return fpl, nil
}

func (pl *plane) display() string {
	var s string
	for y := 0; y <= pl.ymax; y++ {
		for x := 0; x <= pl.xmax; x++ {
			if _, ok := pl.dots[p(x, y)]; ok {
				s += "#"
			} else {
				s += " "
			}
		}
		s += "\n"
	}

	return s
}

func dotsAfterFirstFold(inDots string, inFolds string) (int, error) {
	points, err := parsePoints(inDots)
	if err != nil {
		return 0, errors.Wrap(err, "parse-points")
	}
	folds, err := parseFolds(inFolds)
	if err != nil {
		return 0, errors.Wrap(err, "parse-folds")
	}
	if len(points) == 0 {
		return 0, errors.Errorf("no points")
	}
	if len(folds) == 0 {
		return 0, errors.Errorf("no folds")
	}

	pl := buildPlane(points)
	fpl, err := pl.folded(folds[0])
	if err != nil {
		return 0, errors.Wrap(err, "folded")
	}
	dots := fpl.visibleDots()

	return dots, nil
}

func displayCode(inDots string, inFolds string) (int, error) {
	points, err := parsePoints(inDots)
	if err != nil {
		return 0, errors.Wrap(err, "parse-points")
	}
	folds, err := parseFolds(inFolds)
	if err != nil {
		return 0, errors.Wrap(err, "parse-folds")
	}
	if len(points) == 0 {
		return 0, errors.Errorf("no points")
	}
	if len(folds) == 0 {
		return 0, errors.Errorf("no folds")
	}

	pl := buildPlane(points)
	for _, f := range folds {
		pl, err = pl.folded(f)
		if err != nil {
			return 0, errors.Wrap(err, "folded")
		}
	}
	fmt.Printf("------------------------\n%s\n------------------------\n", pl.display())

	return 0, nil
}
