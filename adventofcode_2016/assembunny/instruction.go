package assembunny

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2016/readutil"

	"github.com/pkg/errors"
)

type Instruction interface{}

type Cpy struct {
	Src string
	Dst string
}

type Inc struct {
	Reg string
}

type Dec struct {
	Reg string
}

type Jnz struct {
	Cond  string
	Steps string
}

type Tgl struct {
	Dst string
}

type Out struct {
	Src string
}

func ParseInstruction(s string) (Instruction, error) {
	switch {
	case strings.HasPrefix(s, "cpy"):
		var p1, p2 string
		_, err := fmt.Sscanf(s, "cpy %s %s", &p1, &p2)
		if err != nil {
			return nil, err
		}
		return Cpy{
			Src: p1,
			Dst: p2,
		}, nil

	case strings.HasPrefix(s, "inc"):
		var p1 string
		_, err := fmt.Sscanf(s, "inc %s", &p1)
		if err != nil {
			return nil, err
		}
		return Inc{
			Reg: p1,
		}, nil
	case strings.HasPrefix(s, "dec"):
		var p1 string
		_, err := fmt.Sscanf(s, "dec %s", &p1)
		if err != nil {
			return nil, err
		}
		return Dec{
			Reg: p1,
		}, nil
	case strings.HasPrefix(s, "jnz"):
		var p1, p2 string
		_, err := fmt.Sscanf(s, "jnz %s %s", &p1, &p2)
		if err != nil {
			return nil, err
		}
		return Jnz{
			Cond:  p1,
			Steps: p2,
		}, nil
	case strings.HasPrefix(s, "tgl"):
		var p1 string
		_, err := fmt.Sscanf(s, "tgl %s", &p1)
		if err != nil {
			return nil, err
		}
		return Tgl{
			Dst: p1,
		}, nil
	case strings.HasPrefix(s, "out"):
		var p1 string
		_, err := fmt.Sscanf(s, "out %s", &p1)
		if err != nil {
			return nil, err
		}
		return Out{
			Src: p1,
		}, nil
	default:
		return nil, errors.Errorf("invalid instruction %q", s)
	}
}

func ParseInstructions(in string) ([]Instruction, error) {
	var iss []Instruction
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		is, err := ParseInstruction(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-instruction %q", line)
		}
		iss = append(iss, is)
	}
	return iss, nil
}
