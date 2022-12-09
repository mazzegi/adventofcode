package day_09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
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

type Rope struct {
	Head grid.Point
	Tail grid.Point
}

type RopeXL struct {
	Head  grid.Point
	Inter []grid.Point
	Tail  grid.Point
}

type Dir string

const (
	Up    Dir = "U"
	Down  Dir = "D"
	Right Dir = "R"
	Left  Dir = "L"
)

func ParseDir(s string) (Dir, error) {
	d := Dir(s)
	switch d {
	case Up, Down, Right, Left:
		return d, nil
	default:
		return Dir(""), fmt.Errorf("invalid dir %q", s)
	}
}

type Move struct {
	Dir   Dir
	Count int
}

func ParseMove(s string) (Move, error) {
	sd, sc, ok := strings.Cut(strings.TrimSpace(s), " ")
	if !ok {
		return Move{}, fmt.Errorf("invalid move %q", s)
	}
	d, err := ParseDir(sd)
	if err != nil {
		return Move{}, err
	}
	c, err := strconv.Atoi(sc)
	if err != nil {
		return Move{}, err
	}
	return Move{d, c}, nil
}

func ParseMoves(in string) ([]Move, error) {
	lines := readutil.ReadLines(in)
	var mvs []Move
	for _, line := range lines {
		mv, err := ParseMove(line)
		if err != nil {
			return nil, err
		}
		mvs = append(mvs, mv)
	}
	if len(mvs) == 0 {
		return nil, fmt.Errorf("no moves")
	}
	return mvs, nil
}

func moveTail(head grid.Point, tail grid.Point) grid.Point {
	if head.Sub(tail).Norm() < 2 {
		return tail
	}
	tail.Y = tail.Y + mathutil.Sign(head.Y-tail.Y)
	tail.X = tail.X + mathutil.Sign(head.X-tail.X)
	return tail
}

func part1MainFunc(in string) (int, error) {
	rope := &Rope{
		Head: grid.Pt(0, 0),
		Tail: grid.Pt(0, 0),
	}
	coveredByTail := set.New[grid.Point]()
	coveredByTail.Insert(rope.Tail)
	moves, err := ParseMoves(in)
	if err != nil {
		return 0, err
	}
	for _, mv := range moves {
		for i := 0; i < mv.Count; i++ {
			switch mv.Dir {
			case Up:
				rope.Head.Y--
			case Down:
				rope.Head.Y++
			case Right:
				rope.Head.X++
			case Left:
				rope.Head.X--
			}
			rope.Tail = moveTail(rope.Head, rope.Tail)
			coveredByTail.Insert(rope.Tail)
		}
	}

	return coveredByTail.Count(), nil
}

func part2MainFunc(in string) (int, error) {
	rope := &RopeXL{
		Head:  grid.Pt(0, 0),
		Inter: make([]grid.Point, 8),
		Tail:  grid.Pt(0, 0),
	}

	coveredByTail := set.New[grid.Point]()
	coveredByTail.Insert(rope.Tail)
	moves, err := ParseMoves(in)
	if err != nil {
		return 0, err
	}
	for _, mv := range moves {
		for i := 0; i < mv.Count; i++ {
			switch mv.Dir {
			case Up:
				rope.Head.Y--
			case Down:
				rope.Head.Y++
			case Right:
				rope.Head.X++
			case Left:
				rope.Head.X--
			}

			h := rope.Head
			for i, ip := range rope.Inter {
				np := moveTail(h, ip)
				rope.Inter[i] = np
				h = np
			}
			rope.Tail = moveTail(h, rope.Tail)

			coveredByTail.Insert(rope.Tail)
		}
	}

	return coveredByTail.Count(), nil
}
