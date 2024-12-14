package day_03

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

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

func part1MainFunc(in string) (int, error) {
	//log("parsing:\n%s\n", in)
	pos := 0
	rs := []rune(in)

	var sum int

outer_loop:
	for {
		if !strings.HasPrefix(string(rs[pos:]), "mul(") {
			pos++
			if pos >= len(rs) {
				break outer_loop
			}
			continue outer_loop
		}
		// it starts with "mul("
		pos += 4
		if pos >= len(rs) {
			break outer_loop
		}

		// first number until ","
		firstNumberStr := ""
		for {
			if rs[pos] == ',' {
				break
			}
			if !unicode.IsDigit(rs[pos]) {
				continue outer_loop
			}
			firstNumberStr += string(rs[pos])
			pos++
			if pos >= len(rs) {
				break outer_loop
			}
		}

		pos++
		if pos >= len(rs) {
			break outer_loop
		}
		firstNumber, err := strconv.Atoi(firstNumberStr)
		if err != nil {
			continue outer_loop
		}

		// second number until ")"
		secondNumberStr := ""
		for {
			if rs[pos] == ')' {
				break
			}
			if !unicode.IsDigit(rs[pos]) {
				continue outer_loop
			}
			secondNumberStr += string(rs[pos])
			pos++
			if pos >= len(rs) {
				break outer_loop
			}
		}
		secondNumber, err := strconv.Atoi(secondNumberStr)
		if err != nil {
			continue outer_loop
		}
		log("detected: mul(%d,%d)", firstNumber, secondNumber)
		sum += firstNumber * secondNumber

		pos++
		if pos >= len(rs) {
			break outer_loop
		}
	}

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	pos := 0
	rs := []rune(in)

	var sum int
	mulEnabled := true

outer_loop:
	for {
		//log("%d: %q", pos, string(rs[pos:]))

		if strings.HasPrefix(string(rs[pos:]), "don't()") {
			pos += 7
			if pos >= len(rs) {
				break outer_loop
			}
			mulEnabled = false
		}
		if strings.HasPrefix(string(rs[pos:]), "do()") {
			pos += 4
			if pos >= len(rs) {
				break outer_loop
			}
			mulEnabled = true
		}

		if !strings.HasPrefix(string(rs[pos:]), "mul(") {
			pos++
			if pos >= len(rs) {
				break outer_loop
			}
			continue outer_loop
		}
		// it starts with "mul("
		pos += 4
		if pos >= len(rs) {
			break outer_loop
		}

		// first number until ","
		firstNumberStr := ""
		for {
			if rs[pos] == ',' {
				break
			}
			if !unicode.IsDigit(rs[pos]) {
				continue outer_loop
			}
			firstNumberStr += string(rs[pos])
			pos++
			if pos >= len(rs) {
				break outer_loop
			}
		}

		pos++
		if pos >= len(rs) {
			break outer_loop
		}
		firstNumber, err := strconv.Atoi(firstNumberStr)
		if err != nil {
			continue outer_loop
		}

		// second number until ")"
		secondNumberStr := ""
		for {
			if rs[pos] == ')' {
				break
			}
			if !unicode.IsDigit(rs[pos]) {
				continue outer_loop
			}
			secondNumberStr += string(rs[pos])
			pos++
			if pos >= len(rs) {
				break outer_loop
			}
		}
		secondNumber, err := strconv.Atoi(secondNumberStr)
		if err != nil {
			continue outer_loop
		}
		//log("detected: mul(%d,%d)", firstNumber, secondNumber)
		if mulEnabled {
			sum += firstNumber * secondNumber
		}

		pos++
		if pos >= len(rs) {
			break outer_loop
		}
	}

	return sum, nil
}
