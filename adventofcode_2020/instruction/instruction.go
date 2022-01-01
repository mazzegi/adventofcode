package instruction

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Operation string

const (
	Acc Operation = "acc"
	Jmp Operation = "jmp"
	Nop Operation = "nop"
)

type Instruction struct {
	Op    Operation
	Value int
}

func (i Instruction) IsValid() bool {
	switch i.Op {
	case Acc, Jmp, Nop:
		return true
	default:
		return false
	}
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s: %d", i.Op, i.Value)
}

func Parse(s string) (Instruction, error) {
	var i Instruction
	n, err := fmt.Sscanf(s, "%s %d", &i.Op, &i.Value)
	if err != nil {
		return Instruction{}, errors.Wrapf(err, "scan instruction %q: %v", s, err)
	}
	if n != 2 {
		return Instruction{}, errors.Errorf("scan instruction %q count %d", s, n)
	}
	if !i.IsValid() {
		return Instruction{}, errors.Errorf("invalid instruction %q", s)
	}
	return i, nil
}

func ParseMany(r io.Reader) ([]Instruction, error) {
	var is []Instruction
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := strings.Trim(scanner.Text(), " \n\r\t")
		if s == "" {
			continue
		}
		i, err := Parse(s)
		if err != nil {
			return nil, errors.Wrapf(err, "parse %q", s)
		}
		is = append(is, i)
	}
	return is, nil
}

type Machine struct {
	is  []Instruction
	acc int
	ptr int
}

func NewMachine(is []Instruction) *Machine {
	return &Machine{
		is:  is,
		acc: 0,
		ptr: 0,
	}
}

func (m *Machine) Reset() {
	m.ptr = 0
	m.acc = 0
}

func (m *Machine) Acc() int {
	return m.acc
}

func (m *Machine) Exec() (int, error) {
	visited := map[int]bool{}
	var cnt int
	for {
		if m.ptr == len(m.is) {
			return cnt, nil
		}
		if _, contains := visited[m.ptr]; contains {
			return cnt, errors.Errorf("already executed %d - infinite loop", m.ptr)
		}
		visited[m.ptr] = true

		err := m.ExecNext()
		if err != nil {
			return 0, errors.Wrapf(err, "exec-next (%d)", cnt)
		}
		cnt++
	}
}

func (m *Machine) ExecOneLoop() (int, error) {
	visited := map[int]bool{}
	var cnt int
	for {
		if _, contains := visited[m.ptr]; contains {
			return cnt, nil
		}
		visited[m.ptr] = true

		err := m.ExecNext()
		if err != nil {
			return 0, errors.Wrapf(err, "exec-next (%d)", cnt)
		}
		cnt++
	}
}

func (m *Machine) ExecNext() error {
	if m.ptr < 0 || m.ptr >= len(m.is) {
		return errors.Errorf("index %d out of bounds [0,%d)", m.ptr, len(m.is))
	}
	i := m.is[m.ptr]
	switch i.Op {
	case Nop:
		m.ptr++
	case Acc:
		m.acc += i.Value
		m.ptr++
	case Jmp:
		m.ptr += i.Value
	}
	return nil
}
