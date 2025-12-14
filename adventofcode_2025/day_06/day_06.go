package day_06

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
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

func parseIntLine(line string) ([]int, error) {
	sl := strings.Split(line, " ")
	var ns []int
	for _, s := range sl {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse int: %w", err)
		}
		ns = append(ns, int(n))
	}
	return ns, nil
}

func parseOpLine(line string) ([]string, error) {
	sl := strings.Split(line, " ")
	var ops []string
	for _, s := range sl {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		switch s {
		case "+", "*":
			ops = append(ops, s)
		default:
			return nil, fmt.Errorf("invalid operator %q", s)
		}
	}
	return ops, nil
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	if len(lines) < 2 {
		return 0, fmt.Errorf("to less lines %d", len(lines))
	}
	numCols := 0
	var intLines [][]int
	for _, line := range lines[:len(lines)-1] {
		il, err := parseIntLine(line)
		if err != nil {
			return 0, fmt.Errorf("parse_int_line: %w", err)
		}
		intLines = append(intLines, il)
		if numCols == 0 {
			numCols = len(il)
		} else if len(il) != numCols {
			return 0, fmt.Errorf("invalid len: want %d have %d", numCols, len(il))
		}
	}
	opLine, err := parseOpLine(lines[len(lines)-1])
	if err != nil {
		return 0, fmt.Errorf("parse_op_line: %w", err)
	}
	if len(opLine) != numCols {
		return 0, fmt.Errorf("invalid len: want %d have %d", numCols, len(opLine))
	}

	// readyto go
	var sum int
	for c := range numCols {
		var val int
		op := opLine[c]
		switch op {
		case "+":
			val = 0
		default: //*
			val = 1
		}

		for _, il := range intLines {
			v := il[c]
			switch op {
			case "+":
				val += v
			default: //*
				val *= v
			}
		}
		sum += val
	}

	return sum, nil
}

type cephaProblem struct {
	rows [][]rune
	op   rune
}

var (
	zeroRune rune
)

func (cp cephaProblem) isZero() bool {
	return cp.op == zeroRune
}

func (cp cephaProblem) eval() int {
	var cols []int
	// skip last (always 32/space)
	for col := range len(cp.rows[0]) {
		ns := []rune{}
		for _, row := range cp.rows {
			ns = append(ns, row[col])
		}
		str := strings.TrimSpace(string(ns))
		if str == "" {
			continue
		}

		n, _ := strconv.ParseInt(str, 10, 64)
		cols = append(cols, int(n))
	}

	var val int
	switch cp.op {
	case '+':
		val = 0
	default: //*
		val = 1
	}

	for _, c := range cols {
		switch cp.op {
		case '+':
			val += c
		default: //*
			val *= c
		}
	}

	return val
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLinesNoTrim(in)
	if len(lines) < 2 {
		return 0, fmt.Errorf("to less lines %d", len(lines))
	}

	opLine := lines[len(lines)-1]
	// assure all lines have the same len
	for _, line := range lines[:len(lines)-1] {
		if len(line) != len(opLine) {
			return 0, fmt.Errorf("invalid sizes")
		}
	}
	numRows := len(lines) - 1

	var problems []cephaProblem
	curr := cephaProblem{
		rows: make([][]rune, numRows),
	}

	flushCurrent := func() {
		if !curr.isZero() {
			problems = append(problems, curr)
		}
		curr = cephaProblem{
			rows: make([][]rune, numRows),
		}
	}

	// iter operator line
	for iop, op := range opLine {
		if op == '+' || op == '*' {
			flushCurrent()
			curr.op = op
		}
		for il, line := range lines[:len(lines)-1] {
			curr.rows[il] = append(curr.rows[il], []rune(line)[iop])
		}
	}
	flushCurrent()

	var sum int
	for _, p := range problems {
		v := p.eval()
		sum += v
	}

	return sum, nil
}
