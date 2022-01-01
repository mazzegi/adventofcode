package intcode

import (
	"github.com/pkg/errors"
)

const (
	Add  int = 1
	Mult int = 2
	Halt int = 99
)

func Exec(in []int) ([]int, error) {
	out := make([]int, len(in))
	copy(out, in)
	pos := 0
	step := 1
	for {
		switch out[pos] {
		case Halt:
			return out, nil
		case Add, Mult:
			p1 := out[pos+1]
			p2 := out[pos+2]
			pr := out[pos+3]
			var r int
			if out[pos] == Add {
				r = out[p1] + out[p2]
			} else if out[pos] == Mult {
				r = out[p1] * out[p2]
			}
			out[pr] = r
			pos += 4
			step++
		default:
			return nil, errors.Errorf("invalid opcode %d", out[pos])
		}
	}
}
