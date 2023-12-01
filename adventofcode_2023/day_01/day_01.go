package day_01

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
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

func digits(sl string) []int {
	var ns []int
	for _, r := range sl {
		if unicode.IsDigit(r) {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				panic("what?")
			}
			ns = append(ns, n)
		}
	}
	return ns
}

func part1MainFunc(in string) (int, error) {
	ls := readutil.ReadLines(in)
	sum := 0
	for _, l := range ls {
		ds := digits(l)
		if len(ds) < 1 {
			return 0, fmt.Errorf("no digits in %q", l)
		}
		num := 10*ds[0] + ds[len(ds)-1]
		sum += num
	}

	return sum, nil
}

// var digReplacer = strings.NewReplacer(
//
//	"one", "1",
//	"two", "2",
//	"three", "3",
//	"four", "4",
//	"five", "5",
//	"six", "6",
//	"seven", "7",
//	"eight", "8",
//	"nine", "9",
//
// )
var spelledDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func digitsWithSpelled(sl string) []int {
	var ns []int
	spelledDigitPrefix := func(s string) (int, bool) {
		for sd, n := range spelledDigits {
			if strings.HasPrefix(s, sd) {
				return n, true
			}
		}
		return 0, false
	}

	for i, r := range sl {
		if unicode.IsDigit(r) {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				panic("what?")
			}
			ns = append(ns, n)
		} else if n, ok := spelledDigitPrefix(sl[i:]); ok {
			ns = append(ns, n)
		}
	}
	return ns
}

func part2MainFunc(in string) (int, error) {
	ls := readutil.ReadLines(in)
	sum := 0
	for _, l := range ls {
		ds := digitsWithSpelled(l)
		if len(ds) < 1 {
			return 0, fmt.Errorf("no digits in %q", l)
		}
		num := 10*ds[0] + ds[len(ds)-1]
		sum += num
	}

	return sum, nil
}
