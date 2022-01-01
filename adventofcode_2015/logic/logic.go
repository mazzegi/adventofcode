package logic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/wasa/errors"
)

func ParseBlock(s string) (BlockScheme, error) {
	sl := strings.Split(s, "->")
	if len(sl) != 2 {
		return BlockScheme{}, errors.Errorf("invalid block %q", s)
	}
	gateS := strings.Trim(sl[0], " ")
	output := strings.Trim(sl[1], " ")

	gateSl := strings.Fields(gateS)
	if len(gateSl) == 1 {
		n, err := strconv.ParseUint(gateSl[0], 10, 16)
		if err == nil {
			return BlockScheme{
				Input:  []string{},
				Gate:   ConstGate{Value: uint16(n)},
				Output: output,
			}, nil
		} else {
			return BlockScheme{
				Input:  []string{gateSl[0]},
				Gate:   RouteGate{},
				Output: output,
			}, nil
		}
	}
	//not a const or route - check NOT
	if len(gateSl) == 2 {
		if gateSl[0] != "NOT" {
			return BlockScheme{}, errors.Errorf("invalid block %q", s)
		}
		return BlockScheme{
			Input:  []string{gateSl[1]},
			Gate:   NotGate{},
			Output: output,
		}, nil
	}
	if len(gateSl) > 3 {
		return BlockScheme{}, errors.Errorf("invalid block %q", s)
	}
	switch gateSl[1] {
	case And:
		return BlockScheme{
			Input:  []string{gateSl[0], gateSl[2]},
			Gate:   AndGate{},
			Output: output,
		}, nil
	case Or:
		return BlockScheme{
			Input:  []string{gateSl[0], gateSl[2]},
			Gate:   OrGate{},
			Output: output,
		}, nil
	case LShift:
		by, err := strconv.ParseInt(gateSl[2], 10, 16)
		if err != nil {
			return BlockScheme{}, errors.Errorf("invalid block %q", s)
		}
		return BlockScheme{
			Input:  []string{gateSl[0]},
			Gate:   LShiftGate{By: uint16(by)},
			Output: output,
		}, nil
	case RShift:
		by, err := strconv.ParseInt(gateSl[2], 10, 16)
		if err != nil {
			return BlockScheme{}, errors.Errorf("invalid block %q", s)
		}
		return BlockScheme{
			Input:  []string{gateSl[0]},
			Gate:   RShiftGate{By: uint16(by)},
			Output: output,
		}, nil
	default:
		return BlockScheme{}, errors.Errorf("invalid block %q", s)
	}
}

const (
	Const  = "CONST"
	Route  = "ROUTE"
	And    = "AND"
	Or     = "OR"
	LShift = "LSHIFT"
	RShift = "RSHIFT"
	Not    = "NOT"
)

type Gate interface {
	fmt.Stringer
	Transform(in ...uint16) (uint16, error)
}

type BlockScheme struct {
	Input  []string
	Gate   Gate
	Output string
}

func (b BlockScheme) String() string {
	return fmt.Sprintf("%v -> %s -> %s", b.Input, b.Gate, b.Output)
}

//
type Block struct {
	Scheme BlockScheme
	value  *uint16
}

func NewBlock(scheme BlockScheme) *Block {
	b := &Block{
		Scheme: scheme,
	}
	if cb, ok := scheme.Gate.(ConstGate); ok {
		v := cb.Value
		b.value = &v
	}
	return b
}

func (b *Block) Evaluate(in ...uint16) (uint16, error) {
	v, err := b.Scheme.Gate.Transform(in...)
	if err != nil {
		return 0, err
	}
	b.value = &v
	return v, nil
}

func (b *Block) Value() (uint16, bool) {
	if b.value == nil {
		return 0, false
	}
	return *b.value, true
}

//
type ConstGate struct {
	Value uint16
}

type RouteGate struct {
}

type AndGate struct {
}

type OrGate struct {
}

type LShiftGate struct {
	By uint16
}

type RShiftGate struct {
	By uint16
}

type NotGate struct {
}

//
func (g ConstGate) String() string {
	return fmt.Sprintf("CONST: %d", g.Value)
}

func (g RouteGate) String() string {
	return fmt.Sprintf("ROUTE")
}

func (g AndGate) String() string {
	return fmt.Sprintf("AND")
}

func (g OrGate) String() string {
	return fmt.Sprintf("OR")
}

func (g LShiftGate) String() string {
	return fmt.Sprintf("LSHIFT: %d", g.By)
}

func (g RShiftGate) String() string {
	return fmt.Sprintf("RSHIFT: %d", g.By)
}

func (g NotGate) String() string {
	return fmt.Sprintf("NOT")
}

//
func (g ConstGate) Transform(in ...uint16) (uint16, error) {
	return g.Value, nil
}

func (g RouteGate) Transform(in ...uint16) (uint16, error) {
	if len(in) == 0 {
		return 0, errors.Errorf("ROUTE: no input")
	}
	return in[0], nil
}

func (g AndGate) Transform(in ...uint16) (uint16, error) {
	if len(in) < 2 {
		return 0, errors.Errorf("ADD: insufficient input")
	}
	r := in[0] & in[1]
	return r, nil
}

func (g OrGate) Transform(in ...uint16) (uint16, error) {
	if len(in) < 2 {
		return 0, errors.Errorf("OR: insufficient input")
	}
	r := in[0] | in[1]
	return r, nil
}

func (g LShiftGate) Transform(in ...uint16) (uint16, error) {
	if len(in) < 1 {
		return 0, errors.Errorf("LSHIFT: insufficient input")
	}
	r := in[0] << g.By
	return r, nil
}

func (g RShiftGate) Transform(in ...uint16) (uint16, error) {
	if len(in) < 1 {
		return 0, errors.Errorf("RSHIFT: insufficient input")
	}
	r := in[0] >> g.By
	return r, nil
}

func (g NotGate) Transform(in ...uint16) (uint16, error) {
	if len(in) < 1 {
		return 0, errors.Errorf("NOT: insufficient input")
	}
	r := ^in[0]
	return r, nil
}
