package day_08

import (
	"adventofcode_2017/readutil"
	"fmt"

	"github.com/pkg/errors"
)

type op string

const (
	opInc op = "inc"
	opDec op = "dec"
)

type comparator string

const (
	compEq   comparator = "=="
	compNeq  comparator = "!="
	compLs   comparator = "<"
	compGt   comparator = ">"
	compLsEq comparator = "<="
	compGtEq comparator = ">="
)

type instruction struct {
	reg      string
	op       op
	opValue  int
	condLeft string
	comparator
	condRight int
}

func parseInstruction(s string) (instruction, error) {
	var ins instruction
	_, err := fmt.Sscanf(s, "%s %s %d if %s %s %d", &ins.reg, &ins.op, &ins.opValue, &ins.condLeft, &ins.comparator, &ins.condRight)
	if err != nil {
		return instruction{}, errors.Wrapf(err, "scan instruction %q", s)
	}
	return ins, nil
}

func parseInstructions(in string) ([]instruction, error) {
	var inss []instruction
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		ins, err := parseInstruction(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse instruction %q", line)
		}
		inss = append(inss, ins)
	}
	return inss, nil
}

type CPU struct {
	instructions []instruction
	registers    map[string]int
	highestEver  int
}

func NewCPU(inss []instruction) *CPU {
	return &CPU{
		instructions: inss,
		registers:    map[string]int{},
	}
}

func (cpu *CPU) value(reg string) int {
	v, ok := cpu.registers[reg]
	if !ok {
		cpu.registers[reg] = 0
		return 0
	}
	return v
}

func (cpu *CPU) set(reg string, v int) {
	cpu.registers[reg] = v
}

func (cpu *CPU) Run() {
	for i, ins := range cpu.instructions {
		lv := cpu.value(ins.condLeft)
		compRes := evalComp(lv, ins.comparator, ins.condRight)
		if !compRes {
			continue
		}

		v := cpu.value(ins.reg)
		var setv int
		switch ins.op {
		case opInc:
			setv = v + ins.opValue
		case opDec:
			setv = v - ins.opValue
		}
		cpu.set(ins.reg, setv)

		if i == 0 {
			cpu.highestEver = setv
		} else if setv > cpu.highestEver {
			cpu.highestEver = setv
		}
	}
}

func evalComp(left int, comp comparator, right int) bool {
	switch comp {
	case compEq:
		return left == right
	case compNeq:
		return left != right
	case compLs:
		return left < right
	case compGt:
		return left > right
	case compLsEq:
		return left <= right
	case compGtEq:
		return left >= right
	default:
		return false
	}
}

func (cpu *CPU) LargestRegisterValue() int {
	first := true
	var max int
	for _, val := range cpu.registers {
		if first {
			max = val
			first = false
			continue
		}
		if val > max {
			max = val
		}
	}
	return max
}
