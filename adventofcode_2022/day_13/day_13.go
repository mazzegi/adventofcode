package day_13

import (
	"fmt"
	"reflect"
	"sort"
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

type cmpResult int

const (
	ls cmpResult = -1
	eq cmpResult = 0
	gt cmpResult = 1
)

func cmpInts(n1, n2 int) cmpResult {
	switch {
	case n1 < n2:
		return ls
	case n1 > n2:
		return gt
	default:
		return eq
	}
}

func compareLeftInt(v1 int, v2 any) cmpResult {
	switch v2 := v2.(type) {
	case int:
		return cmpInts(v1, v2)
	case []any:
		return compareArrays([]any{v1}, v2)
	default:
		fatal("invalid type %T", v1)
	}
	return 0
}

func compareArrays(v1 []any, v2 []any) cmpResult {
	min := mathutil.Min(len(v1), len(v2))
	for i := 0; i < min; i++ {
		e1 := v1[i]
		e2 := v2[i]
		res := compare(e1, e2)
		switch res {
		case ls:
			return ls
		case gt:
			return gt
		default: // eq
			continue
		}
	}
	if len(v1) < len(v2) {
		return ls
	} else if len(v1) > len(v2) {
		return gt
	}
	return eq
}

func compareLeftArr(v1 []any, v2 any) cmpResult {
	switch v2 := v2.(type) {
	case []any:
		return compareArrays(v1, v2)
	case int:
		return compareArrays(v1, []any{v2})
	default:
		fatal("invalid type %T", v1)
	}
	return 0
}

func compare(v1 any, v2 any) cmpResult {
	switch v1 := v1.(type) {
	case int:
		return compareLeftInt(v1, v2)
	case []any:
		return compareLeftArr(v1, v2)
	default:
		fatal("invalid type %T", v1)
	}
	return 0
}

func part1MainFunc(in []any, inStr string) (int, error) {
	checkInput(in, inStr)
	var sum int
	for i := 0; i < len(in); i += 2 {
		v1 := in[i]
		v2 := in[i+1]
		if compare(v1, v2) == ls {
			pairIdx := i/2 + 1
			sum += pairIdx
			log("%d", pairIdx)
		}
	}
	return sum, nil
}

func part2MainFunc(in []any) (int, error) {
	dec1 := []any{[]any{2}}
	dec2 := []any{[]any{6}}
	//inject decoder packets
	in = append(in, dec1, dec2)
	sort.Slice(in, func(i, j int) bool {
		return compare(in[i], in[j]) == ls
	})
	var ix1 int
	var ix2 int
	for i, v := range in {
		if reflect.DeepEqual(v, dec1) {
			ix1 = i + 1
		}
		if reflect.DeepEqual(v, dec2) {
			ix2 = i + 1
		}
	}
	return ix1 * ix2, nil
}
