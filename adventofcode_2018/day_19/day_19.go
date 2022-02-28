package day_19

import (
	"adventofcode_2018/errutil"
	"adventofcode_2018/readutil"
	"fmt"
	"strings"

	"github.com/mazzegi/slices"
	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	skip := true
	if skip {
		return
	}
	res, err := part1MainFunc(input, inputIP)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input, inputIP)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//

func parseInstructions(in string) ([]Instr, error) {
	var iss []Instr
	for _, line := range readutil.ReadLines(in) {
		fs := strings.Fields(line)
		if len(fs) != 4 {
			return nil, errors.Errorf("cannot parse %q", line)
		}
		args, err := slices.Convert(fs[1:], slices.ParseInt)
		if err != nil {
			return nil, errors.Wrapf(err, "convert args %q", line)
		}
		iss = append(iss, Instr{fs[0], args})
	}
	return iss, nil
}

func part1MainFunc(in string, ip int) (int, error) {
	iss, err := parseInstructions(in)
	if err != nil {
		return 0, err
	}

	m := NewMachine(ip, [6]int{0, 0, 0, 0, 0, 0}, iss)
	m.Run()

	return m.regs[0], nil
}

func part2MainFunc(in string, ip int) (int, error) {
	iss, err := parseInstructions(in)
	if err != nil {
		return 0, err
	}

	m := NewMachine(ip, [6]int{1, 0, 0, 0, 0, 0}, iss)
	m.Run()

	return m.regs[0], nil
}
