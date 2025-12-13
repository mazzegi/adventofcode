package day_01

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mazzegi/adventofcode/adventofcode_2016/readutil"
	"github.com/mazzegi/adventofcode/errutil"
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

type move struct {
	direction string
	amount    int
}

func parseMove(s string) (move, error) {
	if len(s) < 2 {
		return move{}, fmt.Errorf("invalid move %q", s)
	}
	dirS := s[0]
	amountS := s[1:]
	if dirS != 'R' && dirS != 'L' {
		return move{}, fmt.Errorf("invalid move %q", s)
	}
	amount, err := strconv.ParseInt(amountS, 10, 32)
	if err != nil {
		return move{}, fmt.Errorf("parse int %q: %w", amountS, err)
	}
	return move{direction: string(dirS), amount: int(amount)}, nil
}

func modRotate(a int, mod int) int {
	r := a % mod
	if r < 0 {
		r = mod + r
	}
	return r
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var moves []move
	for i, line := range lines {
		mv, err := parseMove(line)
		if err != nil {
			return 0, fmt.Errorf("parse move %q (line %d): %w", line, i+1, err)
		}
		moves = append(moves, mv)
	}

	curr := 50
	point0Count := 0
	for _, mv := range moves {
		var next int
		switch mv.direction {
		case "L":
			next = curr - mv.amount
		case "R":
			next = curr + mv.amount
		default:
			return 0, fmt.Errorf("invalid direction")
		}
		// modulo rotate
		next = modRotate(next, 100)
		curr = next
		if curr == 0 {
			point0Count++
		}
		//
	}

	return point0Count, nil
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var moves []move
	for i, line := range lines {
		mv, err := parseMove(line)
		if err != nil {
			return 0, fmt.Errorf("parse move %q (line %d): %w", line, i+1, err)
		}
		moves = append(moves, mv)
	}

	curr := 50
	point0Count := 0
	for _, mv := range moves {
		var next int
		var pt0Hits int
		switch mv.direction {
		case "L":
			next, pt0Hits = rotateLeft(curr, mv.amount)
		case "R":
			next, pt0Hits = rotateRight(curr, mv.amount)
		default:
			return 0, fmt.Errorf("invalid direction")
		}
		curr = next
		point0Count += pt0Hits
	}

	return point0Count, nil
}

func rotateLeft(currPos int, amount int) (newPos int, point0Hits int) {
	newPos = currPos
	for range amount {
		newPos--
		if newPos == -1 {
			newPos = 99
		}
		if newPos == 0 {
			point0Hits++
		}
	}
	return newPos, point0Hits
}

func rotateRight(currPos int, amount int) (newPos int, point0Hits int) {
	newPos = currPos
	for range amount {
		newPos++
		if newPos == 100 {
			newPos = 0
		}
		if newPos == 0 {
			point0Hits++
		}
	}
	return newPos, point0Hits
}
