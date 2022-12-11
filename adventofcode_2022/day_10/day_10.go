package day_10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
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

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	cycle := 0
	reg := 1
	acc := 0
	step := func() {
		cycle++
		switch cycle {
		case 20, 60, 100, 140, 180, 220:
			acc += cycle * reg
		}
	}

	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "noop"):
			step()
			continue
		case strings.HasPrefix(line, "addx"):
			_, vs, ok := strings.Cut(line, " ")
			if !ok {
				return 0, fmt.Errorf("invalid input %q", line)
			}
			v, err := strconv.Atoi(vs)
			if err != nil {
				return 0, fmt.Errorf("invalid input %q: %v", line, err)
			}
			step()
			step()
			reg += v
		}
	}
	return acc, nil
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	cycle := 0
	reg := 1
	sprite := [3]int{}
	out := ""
	crt := []string{}
	step := func() {
		sprite[0] = reg - 1
		sprite[1] = reg
		sprite[2] = reg + 1

		if cycle >= sprite[0] && cycle <= sprite[2] {
			out += "#"
		} else {
			out += "."
		}
		cycle++
		if cycle >= 40 {
			crt = append(crt, out)
			out = ""
			cycle = 0
		}
	}

	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "noop"):
			step()
			continue
		case strings.HasPrefix(line, "addx"):
			_, vs, ok := strings.Cut(line, " ")
			if !ok {
				return 0, fmt.Errorf("invalid input %q", line)
			}
			v, err := strconv.Atoi(vs)
			if err != nil {
				return 0, fmt.Errorf("invalid input %q: %v", line, err)
			}
			step()
			step()
			reg += v
		}
	}
	fmt.Println(strings.Join(crt, "\n"))
	return 0, nil
}
