package day_18

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/readutil"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := recFreq(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := sendValueCountOf1(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
const (
	snd = "snd"
	set = "set"
	add = "add"
	mul = "mul"
	mod = "mod"
	rcv = "rcv"
	jgz = "jgz"
)

func validName(name string) bool {
	switch name {
	case snd, set, add, mul, mod, rcv, jgz:
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

func makeIVal(s string) ival {
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return ivalNum(int(n))
	}
	return ivalReg(s)
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

	if name == snd {
		return instruction{
			name: name,
			val:  makeIVal(reg),
		}
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

func recFreq(in string) (int, error) {
	iss := mustParseInstructions(in)

	regs := map[string]int{}
	regValue := func(reg string) int {
		return regs[reg]
	}
	setRegValue := func(reg string, v int) {
		regs[reg] = v
	}
	var lastSound int

	pos := 0
outer:
	for {
		if pos < 0 || pos >= len(iss) {
			fatal("pos %d out of range", pos)
		}
		is := iss[pos]
		switch is.name {
		case snd:
			lastSound = is.val.value(regValue)
		case set:
			setRegValue(is.reg, is.val.value(regValue))
		case add:
			setRegValue(is.reg, regValue(is.reg)+is.val.value(regValue))
		case mul:
			setRegValue(is.reg, regValue(is.reg)*is.val.value(regValue))
		case mod:
			setRegValue(is.reg, regValue(is.reg)%is.val.value(regValue))
		case rcv:
			if regValue(is.reg) != 0 {
				return lastSound, nil
			}
		case jgz:
			if regValue(is.reg) > 0 {
				pos += is.val.value(regValue)
				continue outer
			}
		}
		pos++
	}
}

type queue struct {
	values []int
}

func (q *queue) Enqueue(n int) {
	q.values = append(q.values, n)
}

func (q *queue) Dequeue() (int, bool) {
	if len(q.values) == 0 {
		return 0, false
	}
	n := q.values[0]
	q.values = q.values[1:]
	return n, true
}

func sendValueCountOf1(in string) (int, error) {
	iss := mustParseInstructions(in)

	qP0P1 := &queue{}
	qP1P0 := &queue{}
	p0 := newPrg(0, iss, qP1P0, qP0P1)
	p1 := newPrg(1, iss, qP0P1, qP1P0)
	i := 0
	for {
		p0.step()
		p1.step()
		if p0.recvWait && p1.recvWait {
			return p1.sndCount, nil
		}
		i++
		if i%1000000 == 0 {
			log("--------")
			log("%d: q0->1: %d, p0: snd = %d, rec = %d, regs: %s", i, len(qP0P1.values), p0.sndCount, p0.rcvCount, p0.formatRegs())
			log("%d: q1->0: %d, p1: snd = %d, rec = %d, regs: %s", i, len(qP1P0.values), p1.sndCount, p1.rcvCount, p1.formatRegs())
			log("********")
		}
	}
}

type prg struct {
	id       int
	iss      []instruction
	pos      int
	regs     map[string]int
	rcvQ     *queue
	sndQ     *queue
	recvWait bool
	sndCount int
	rcvCount int
}

func newPrg(pid int, iss []instruction, rcvQ *queue, sndQ *queue) *prg {
	p := &prg{
		id:   pid,
		iss:  iss,
		pos:  0,
		regs: map[string]int{},
		rcvQ: rcvQ,
		sndQ: sndQ,
	}
	p.setRegValue("p", pid)
	return p
}

func (p *prg) formatRegs() string {
	regsl := []struct {
		reg string
		val int
	}{}
	for reg, val := range p.regs {
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

func (p *prg) regValue(reg string) int {
	if n, err := strconv.ParseInt(reg, 10, 64); err == nil {
		return int(n)
	}

	return p.regs[reg]
}

func (p *prg) setRegValue(reg string, v int) {
	p.regs[reg] = v
}

func (p *prg) step() {
	if p.pos < 0 || p.pos >= len(p.iss) {
		fatal("pos %d out of range", p.pos)
	}
	is := p.iss[p.pos]
	var rval int
	if is.val != nil {
		rval = is.val.value(p.regValue)
	}
	switch is.name {
	case snd:
		p.sndQ.Enqueue(rval)
		p.sndCount++
		log("p%d: send %d", p.id, rval)
	case set:
		p.setRegValue(is.reg, rval)
		log("p%d: set %s %d", p.id, is.reg, rval)
	case add:
		r := p.regValue(is.reg) + rval
		p.setRegValue(is.reg, r)
		log("p%d: add %s %d", p.id, is.reg, r)
	case mul:
		r := p.regValue(is.reg) * rval
		p.setRegValue(is.reg, r)
		log("p%d: mul %s %d", p.id, is.reg, r)
	case mod:
		r := p.regValue(is.reg) % rval
		p.setRegValue(is.reg, r)
		log("p%d: mod %s %d", p.id, is.reg, r)
	case rcv:
		if n, ok := p.rcvQ.Dequeue(); ok {
			p.setRegValue(is.reg, n)
			p.recvWait = false
			p.rcvCount++
			log("p%d: rcv %s %d", p.id, is.reg, n)
		} else {
			p.recvWait = true
			log("p%d: wait", p.id)
			return
		}

	case jgz:
		if p.regValue(is.reg) > 0 {
			p.pos += rval
			log("p%d: jgz %d", p.id, rval)
			return
		} else {
			log("p%d: jgz nop", p.id)
		}
	}
	p.pos++
}
