package day_14

import (
	"fmt"
	"math"
	"reflect"

	"github.com/mazzegi/adventofcode/adventofcode_2018/errutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(3, 7, 236021)
	errutil.ExitOnErr(err)
	log("part1: result = %v", res)
}

func Part2() {
	res, err := part2MainFunc(3, 7, []int{2, 3, 6, 0, 2, 1})
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//

func digits(n int) []int {
	var ds []int
	power := 1
	for {
		base := int(math.Floor(float64(n)/math.Pow10(power)) * math.Pow10(power))
		ds = append(ds, (n-base)/int(math.Pow10(power-1)))
		if base == 0 {
			break
		}
		n = base
		power++
	}
	for i := 0; i < len(ds)/2; i++ {
		ds[i], ds[len(ds)-1-i] = ds[len(ds)-1-i], ds[i]
	}
	return ds
}

type elf struct {
	current int
	pos     int
}

type board struct {
	scores []int
}

func part1MainFunc(score1, score2, numRecipes int) (string, error) {
	e1 := elf{score1, 0}
	e2 := elf{score2, 1}
	board := &board{scores: []int{score1, score2}}

	moveElf := func(e *elf) {
		for i := 0; i < e.current+1; i++ {
			e.pos++
			if e.pos >= len(board.scores) {
				e.pos = 0
			}
		}
	}

	tick := func() {
		sum := e1.current + e2.current
		board.scores = append(board.scores, digits(sum)...)
		moveElf(&e1)
		moveElf(&e2)
		e1.current = board.scores[e1.pos]
		e2.current = board.scores[e2.pos]
	}

	for {
		tick()
		if len(board.scores) >= numRecipes+10 {
			break
		}
	}
	var s string
	for _, d := range board.scores[numRecipes : numRecipes+10] {
		s += fmt.Sprintf("%d", d)
	}

	return s, nil
}

func hasSuffix(ns []int, suffix []int) bool {
	if len(ns) < len(suffix) {
		return false
	}
	return reflect.DeepEqual(ns[len(ns)-len(suffix):], suffix)
}

func part2MainFunc(score1, score2 int, target []int) (int, error) {
	e1 := elf{score1, 0}
	e2 := elf{score2, 1}
	board := &board{scores: []int{score1, score2}}

	moveElf := func(e *elf) {
		for i := 0; i < e.current+1; i++ {
			e.pos++
			if e.pos >= len(board.scores) {
				e.pos = 0
			}
		}
	}

	tick := func() {
		sum := e1.current + e2.current
		board.scores = append(board.scores, digits(sum)...)
		moveElf(&e1)
		moveElf(&e2)
		e1.current = board.scores[e1.pos]
		e2.current = board.scores[e2.pos]
	}

	for {
		tick()
		if hasSuffix(board.scores, target) {
			return len(board.scores) - len(target), nil
		}
		if hasSuffix(board.scores[:len(board.scores)-1], target) {
			return len(board.scores) - len(target) - 1, nil
		}
	}
}
