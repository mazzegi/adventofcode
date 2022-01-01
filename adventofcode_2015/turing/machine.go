package turing

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/wasa/errors"
)

type Instruction interface{}

type Half struct {
	reg string
}

type Triple struct {
	reg string
}

type Inc struct {
	reg string
}

type Jump struct {
	offset int
}

type JumpIfEven struct {
	reg    string
	offset int
}

type JumpIfOdd struct {
	reg    string
	offset int
}

func isValidRegister(r string) bool {
	switch r {
	case "a", "b":
		return true
	default:
		return false
	}
}

func ParseInstruction(s string) (Instruction, error) {
	fs := strings.Fields(s)
	if len(fs) < 2 {
		return nil, errors.Errorf("invalid instruction %q", s)
	}
	switch fs[0] {
	case "hlf":
		if !isValidRegister(fs[1]) {
			return nil, errors.Errorf("invalid instruction %q", s)
		}
		return Half{reg: fs[1]}, nil
	case "tpl":
		if !isValidRegister(fs[1]) {
			return nil, errors.Errorf("invalid instruction %q", s)
		}
		return Triple{reg: fs[1]}, nil
	case "inc":
		if !isValidRegister(fs[1]) {
			return nil, errors.Errorf("invalid instruction %q", s)
		}
		return Inc{reg: fs[1]}, nil
	case "jmp":
		n, err := strconv.ParseInt(fs[1], 10, 64)
		if err != nil {
			return nil, err
		}
		return Jump{offset: int(n)}, nil
	case "jie":
		if len(fs) != 3 {
			return nil, errors.Errorf("invalid instruction %q", s)
		}
		r := strings.Trim(fs[1], ",")
		if !isValidRegister(r) {
			return nil, errors.Errorf("invalid instruction %q", s)
		}
		n, err := strconv.ParseInt(fs[2], 10, 64)
		if err != nil {
			return nil, err
		}
		return JumpIfEven{reg: r, offset: int(n)}, nil
	case "jio":
		if len(fs) != 3 {
			return nil, errors.Errorf("invalid instruction %q", s)
		}
		r := strings.Trim(fs[1], ",")
		if !isValidRegister(r) {
			return nil, errors.Errorf("invalid instruction %q", s)
		}
		n, err := strconv.ParseInt(fs[2], 10, 64)
		if err != nil {
			return nil, err
		}
		return JumpIfOdd{reg: r, offset: int(n)}, nil
	default:
		return nil, errors.Errorf("invalid instruction %q", s)
	}
}

//

type Machine struct {
	instructions []Instruction
	regA         int
	regB         int
}

func NewMachine(is []Instruction) *Machine {
	return &Machine{
		instructions: is,
		regA:         1,
		regB:         0,
	}
}

func (m *Machine) Run() {
	i := 0

	inRange := func() bool {
		return i >= 0 && i < len(m.instructions)
	}

	updateReg := func(s string, updateFunc func(n *int)) {
		var reg *int
		switch s {
		case "a":
			reg = &m.regA
		case "b":
			reg = &m.regB
		default:
			panic("unknown register: " + s)
		}
		updateFunc(reg)
	}

	regValue := func(s string) int {
		switch s {
		case "a":
			return m.regA
		case "b":
			return m.regB
		default:
			panic("unknown register: " + s)
		}
	}

	for inRange() {
		inst := m.instructions[i]
		switch inst := inst.(type) {
		case Half:
			updateReg(inst.reg, func(n *int) { *n = *n / 2 })
			i += 1
		case Triple:
			updateReg(inst.reg, func(n *int) { *n = *n * 3 })
			i += 1
		case Inc:
			updateReg(inst.reg, func(n *int) { *n = *n + 1 })
			i += 1
		case Jump:
			i += inst.offset
		case JumpIfEven:
			if regValue(inst.reg)%2 == 0 {
				i += inst.offset
			} else {
				i += 1
			}
		case JumpIfOdd:
			if regValue(inst.reg) == 1 {
				i += inst.offset
			} else {
				i += 1
			}
		}
	}
	fmt.Printf("A: %d\n", m.regA)
	fmt.Printf("B: %d\n", m.regB)
}
