package day_07

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

type equation struct {
	result   int
	operants []int
}

func parseEquations(in string) ([]equation, error) {
	lines := readutil.ReadLines(in)
	var eqs []equation
	for _, line := range lines {
		left, right, ok := strings.Cut(line, ":")
		if !ok {
			return nil, fmt.Errorf("invalid equation %q", line)
		}
		leftNum, err := strconv.Atoi(strings.TrimSpace(left))
		if err != nil {
			return nil, fmt.Errorf("atoi %q: %w", left, err)
		}
		rightNums, err := readutil.ReadInts(right, " ")
		if err != nil {
			return nil, fmt.Errorf("read-ints %v: %w", right, err)
		}
		if len(rightNums) <= 1 {
			return nil, fmt.Errorf("too less operants in %q", line)
		}

		eqs = append(eqs, equation{
			result:   leftNum,
			operants: rightNums,
		})
	}

	return eqs, nil
}

type opFunc func(a, b int) int

func allOpFuncsPart1() []opFunc {
	return []opFunc{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },
	}
}

func allOpFuncsPart2() []opFunc {
	return []opFunc{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },

		func(a, b int) int {
			sa := strconv.Itoa(a)
			sb := strconv.Itoa(b)
			res, _ := strconv.Atoi(sa + sb)
			return res
		},
	}
}

func canEquationBeTruePart1(eq equation) bool {
	opIdxs := make([]int, len(eq.operants)-1)
	for i := range len(opIdxs) {
		opIdxs[i] = 0
	}
	allOps := allOpFuncsPart1()
	numAllOps := len(allOps)

	incOpIdxs := func() bool {
		for i := 0; i < len(opIdxs); i++ {
			if opIdxs[i] < numAllOps-1 {
				opIdxs[i]++
				return true
			}
			opIdxs[i] = 0
		}
		return false
	}

	calc := func() int {
		op := allOps[opIdxs[0]]
		res := op(eq.operants[0], eq.operants[1])
		for i := 2; i < len(eq.operants); i++ {
			op = allOps[opIdxs[i-1]]
			res = op(res, eq.operants[i])
		}
		return res
	}

	for {
		// test
		res := calc()
		if res == eq.result {
			return true
		}
		// increment opidxs
		ok := incOpIdxs()
		if !ok {
			break
		}
	}

	return false
}

func canEquationBeTruePart2(eq equation) bool {
	opIdxs := make([]int, len(eq.operants)-1)
	for i := range len(opIdxs) {
		opIdxs[i] = 0
	}
	allOps := allOpFuncsPart2()
	numAllOps := len(allOps)

	incOpIdxs := func() bool {
		for i := 0; i < len(opIdxs); i++ {
			if opIdxs[i] < numAllOps-1 {
				opIdxs[i]++
				return true
			}
			opIdxs[i] = 0
		}
		return false
	}

	calc := func() int {
		op := allOps[opIdxs[0]]
		res := op(eq.operants[0], eq.operants[1])
		for i := 2; i < len(eq.operants); i++ {
			op = allOps[opIdxs[i-1]]
			res = op(res, eq.operants[i])
		}
		return res
	}

	for {
		// test
		res := calc()
		if res == eq.result {
			return true
		}
		// increment opidxs
		ok := incOpIdxs()
		if !ok {
			break
		}
	}

	return false
}

func part1MainFunc(in string) (int, error) {
	eqs, err := parseEquations(in)
	if err != nil {
		return 0, fmt.Errorf("parse-equations: %w", err)
	}

	var sum int
	for _, eq := range eqs {
		if canEquationBeTruePart1(eq) {
			sum += eq.result
		}
	}

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	eqs, err := parseEquations(in)
	if err != nil {
		return 0, fmt.Errorf("parse-equations: %w", err)
	}

	var sum int
	for _, eq := range eqs {
		if canEquationBeTruePart2(eq) {
			sum += eq.result
		}
	}

	return sum, nil
}
