package day_13

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input, inputCheck)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func str(v any) string {
	switch v := v.(type) {
	case []any:
		var sl []string
		for _, e := range v {
			sl = append(sl, str(e))
		}
		return fmt.Sprintf("[%s]", strings.Join(sl, ","))
	case int:
		return fmt.Sprintf("%d", v)
	default:
		fatal("invalid type %T", v)
	}
	return ""
}

func checkInput(arr []any, in string) {
	lines := readutil.ReadLines(in)
	if len(lines) != len(arr) {
		fatal("not matching sizes")
	}
	for i, line := range lines {
		e := arr[i]
		se := str(e)
		if se != line {
			fatal("%q != %q", se, line)
		}
	}
}

////

func inRightOrderLeftInt(v1 int, v2 any) bool {
	switch v2 := v2.(type) {
	case int:
		return v1 <= v2
	case []any:
		return inRightOrderArrays([]any{v1}, v2)
	default:
		fatal("invalid type %T", v1)
	}
	return false
}

func inRightOrderArrays(v1 []any, v2 []any) bool {
	min := mathutil.Min(len(v1), len(v2))
	for i := 0; i < min; i++ {
		e1 := v1[i]
		e2 := v2[i]
		if !inRightOrder(e1, e2) {
			return false
		}
	}
	if len(v1) < len(v2) {
		return true
	} else if len(v1) > len(v2) {
		return false
	}
	return true
}

func inRightOrderLeftArr(v1 []any, v2 any) bool {
	switch v2 := v2.(type) {
	case []any:
		return inRightOrderArrays(v1, v2)
	case int:
		return inRightOrderArrays(v1, []any{v2})
	default:
		fatal("invalid type %T", v1)
	}
	return false
}

func inRightOrder(v1 any, v2 any) bool {
	switch v1 := v1.(type) {
	case int:
		return inRightOrderLeftInt(v1, v2)
	case []any:
		return inRightOrderLeftArr(v1, v2)
	default:
		fatal("invalid type %T", v1)
	}

	return false
}

func part1MainFunc(in []any, inStr string) (int, error) {
	checkInput(in, inStr)
	var sum int
	for i := 0; i < len(in); i += 2 {
		v1 := in[i]
		v2 := in[i+1]
		if inRightOrder(v1, v2) {
			pairIdx := i/2 + 1
			sum += pairIdx
			log("%d", pairIdx)
		}
	}
	return sum, nil
}

func part2MainFunc(in []any) (int, error) {
	return 0, nil
}
