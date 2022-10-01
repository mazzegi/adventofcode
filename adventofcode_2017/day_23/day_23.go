package day_23

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := numMulInvoked(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := fixedRegH(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

// //
const (
	set = "set"
	sub = "sub"
	mul = "mul"
	jnz = "jnz"
)

func validName(name string) bool {
	switch name {
	case set, sub, mul, jnz:
		return true
	default:
		return false
	}
}

type regValueFnc func(reg string) int

type ival interface {
	value(rvFnc regValueFnc) int
}

type ivalNum int

func (iv ivalNum) value(rvFnc regValueFnc) int {
	return int(iv)
}

type ivalReg string

func (iv ivalReg) value(rvFnc regValueFnc) int {
	return rvFnc(string(iv))
}

type instruction struct {
	name string
	reg  string
	val  ival
}

func mustParseInstruction(s string) instruction {
	fs := strings.Fields(s)
	if len(fs) < 2 {
		fatal("invalid instruction %q", s)
	}
	name := fs[0]
	reg := fs[1]
	if !validName(name) {
		fatal("invalid name %q", name)
	}

	var val ival
	if len(fs) >= 3 {
		s := fs[2]
		if n, err := strconv.ParseInt(s, 10, 64); err == nil {
			val = ivalNum(int(n))
		} else {
			val = ivalReg(s)
		}
	}

	return instruction{
		name: name,
		reg:  reg,
		val:  val,
	}
}

func mustParseInstructions(in string) []instruction {
	var iss []instruction
	for _, line := range readutil.ReadLines(in) {
		is := mustParseInstruction(line)
		iss = append(iss, is)
	}
	if len(iss) == 0 {
		fatal("no data")
	}
	return iss
}

type regs map[string]int

func (r regs) value(reg string) int {
	if n, err := strconv.ParseInt(reg, 10, 64); err == nil {
		return int(n)
	}
	return r[reg]
}

func (r regs) setValue(reg string, v int) {
	r[reg] = v
}

func (r regs) format() string {
	regsl := []struct {
		reg string
		val int
	}{}
	for reg, val := range r {
		regsl = append(regsl, struct {
			reg string
			val int
		}{reg, val})
	}
	sort.Slice(regsl, func(i, j int) bool {
		return regsl[i].reg < regsl[j].reg
	})
	var sl []string
	for _, rs := range regsl {
		sl = append(sl, fmt.Sprintf("%s:%d", rs.reg, rs.val))
	}

	return strings.Join(sl, ", ")
}

func numMulInvoked(in string) (int, error) {
	iss := mustParseInstructions(in)
	rs := regs{}

	pos := 0
	var numMul int
outer:
	for {
		if pos < 0 || pos >= len(iss) {
			return numMul, nil
		}
		is := iss[pos]
		switch is.name {
		case set:
			rs.setValue(is.reg, is.val.value(rs.value))
		case sub:
			rs.setValue(is.reg, rs.value(is.reg)-is.val.value(rs.value))
		case mul:
			rs.setValue(is.reg, rs.value(is.reg)*is.val.value(rs.value))
			numMul++
		case jnz:
			if rs.value(is.reg) != 0 {
				pos += is.val.value(rs.value)
				continue outer
			}
		}
		pos++
	}
}

func fixedRegH(in string) (int, error) {
	iss := mustParseInstructions(in)
	rs := regs{}
	rs.setValue("a", 1)

	pos := 0
	steps := 0
outer:
	for {
		steps++
		if steps%10000 == 0 {
			log("step %d: reg = %s", steps, rs.format())
		}
		if pos < 0 || pos >= len(iss) {
			return rs.value("h"), nil
		}
		is := iss[pos]
		switch is.name {
		case set:
			rs.setValue(is.reg, is.val.value(rs.value))
		case sub:
			rs.setValue(is.reg, rs.value(is.reg)-is.val.value(rs.value))
		case mul:
			rs.setValue(is.reg, rs.value(is.reg)*is.val.value(rs.value))
		case jnz:
			if rs.value(is.reg) != 0 {
				pos += is.val.value(rs.value)
				continue outer
			}
		}
		pos++
	}
}
