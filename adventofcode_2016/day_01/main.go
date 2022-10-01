package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2016/intutil"

	"github.com/pkg/errors"
)

func main() {
	res := distance(input)
	fmt.Printf("result: %d\n", res)

	res = distanceOfFirstVisitedTwice(input)
	fmt.Printf("result (dofvt): %d\n", res)
}

var input = `L1, R3, R1, L5, L2, L5, R4, L2, R2, R2, L2, R1, L5, R3, L4, L1, L2, R3, R5, L2, R5, L1, R2, L5, R4, R2, R2, L1, L1, R1, L3, L1, R1, L3, R5,
R3, R3, L4, R4, L2, L4, R1, R1, L193, R2, L1, R54, R1, L1, R71, L4, R3, R191, R3, R2, L4, R3, R2, L2, L4, L5, R4, R1, L2, L2, L3,
L2, L1, R4, R1, R5, R3, L5, R3, R4, L2, R3, L1, L3, L3, L5, L1, L3, L3, L1, R3, L3, L2, R1, L3, L1, R5, R4, R3, R2, R3, L1, L2, R4,
L3, R1, L1, L1, R5, R2, R4, R5, L1, L1, R1, L2, L4, R3, L1, L3, R5, R4, R3, R3, L2, R2, L1, R4, R2, L3, L4, L2, R2, R2, L4, R3, R5,
L2, R2, R4, R5, L2, L3, L2, R5, L4, L2, R3, L5, R2, L1, R1, R3, R3, L5, L2, L2, R5`

func exitOnErr(err error) {
	if err == nil {
		return
	}
	fmt.Printf("%v\n", err)
	os.Exit(1)
}

type turn string

const (
	Left  turn = "L"
	Right turn = "R"
)

type direction string

const (
	north direction = "N"
	east  direction = "E"
	south direction = "S"
	west  direction = "W"
)

func (d direction) turned(t turn) direction {
	switch {
	case d == north && t == Left:
		return west
	case d == east && t == Left:
		return north
	case d == south && t == Left:
		return east
	case d == west && t == Left:
		return south

	case d == north && t == Right:
		return east
	case d == east && t == Right:
		return south
	case d == south && t == Right:
		return west
	case d == west && t == Right:
		return north

	default:
		return d
	}
}

type instruction struct {
	turn  turn
	steps int
}

func parseInstruction(in string) (instruction, error) {
	ins := instruction{}
	switch {
	case strings.HasPrefix(in, "L"):
		ins.turn = Left
	case strings.HasPrefix(in, "R"):
		ins.turn = Right
	default:
		return instruction{}, errors.Errorf("instruction starts with neither L nor R")
	}
	n, err := strconv.ParseInt(in[1:], 10, 64)
	if err != nil {
		return instruction{}, err
	}
	ins.steps = int(n)
	return ins, nil
}

func distance(inputs string) int {
	var x, y int
	dir := north

	insl := strings.Split(inputs, ",")
	for _, ins := range insl {
		in, err := parseInstruction(strings.Trim(ins, " \r\n\t"))
		exitOnErr(err)
		dir = dir.turned(in.turn)
		switch dir {
		case north:
			y += in.steps
		case east:
			x += in.steps
		case south:
			y -= in.steps
		case west:
			x -= in.steps
		}
	}

	return intutil.AbsInt(x) + intutil.AbsInt(y)
}

type pt struct {
	x, y int
}

func (p pt) distance() int {
	return intutil.AbsInt(p.x) + intutil.AbsInt(p.y)
}

func pointsBetween(p0 pt, p1 pt) []pt {
	ps := []pt{}

	dirEqual := func(n int, n0 int, to int) bool {
		if n0 == to {
			return n == to
		}

		less := n0 < to
		if less {
			if n <= to {
				return true
			}
			return false
		} else {
			if n >= to {
				return true
			}
			return false
		}
	}

	incX := intutil.SignnumInt(p1.x - p0.x)
	incY := intutil.SignnumInt(p1.y - p0.y)
	for x := p0.x; dirEqual(x, p0.x, p1.x); x += incX {
		for y := p0.y; dirEqual(y, p0.y, p1.y); y += incY {
			ps = append(ps, pt{
				x: x,
				y: y,
			})
		}
	}
	return ps
}

func distanceOfFirstVisitedTwice(inputs string) int {
	var x, y int
	dir := north
	visited := map[pt]bool{}
	visited[pt{}] = true

	insl := strings.Split(inputs, ",")
	for _, ins := range insl {
		in, err := parseInstruction(strings.Trim(ins, " \r\n\t"))
		exitOnErr(err)
		dir = dir.turned(in.turn)

		p0 := pt{x: x, y: y}
		switch dir {
		case north:
			y += in.steps
		case east:
			x += in.steps
		case south:
			y -= in.steps
		case west:
			x -= in.steps
		}
		p1 := pt{x: x, y: y}
		ps := pointsBetween(p0, p1)
		for _, p := range ps[1:] {
			if _, ok := visited[p]; ok {
				return p.distance()
			}
			visited[p] = true
		}
	}

	return 0
}
