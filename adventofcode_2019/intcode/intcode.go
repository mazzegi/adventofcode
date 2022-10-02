package intcode

import (
	"github.com/pkg/errors"
)

const (
	Add        int = 1
	Mult       int = 2
	Input      int = 3
	Output     int = 4
	JmpIfTrue  int = 5
	JmpIfFalse int = 6
	Less       int = 7
	Equal      int = 8
	Halt       int = 99
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

func squeezeOpcode(oc int) (code, p1mode, p2mode, p3mode int) {
	d12 := oc - 100*(oc/100)  // 1002 => 2
	pcode := (oc - d12) / 100 // 1002 => 10
	p1 := pcode - 10*(pcode/10)
	pcode = (pcode - p1) / 10
	p2 := pcode - 10*(pcode/10)
	pcode = (pcode - p2) / 10
	p3 := pcode - 10*(pcode/10)
	return d12, p1, p2, p3
}

func Exec2(prg []int, inputs []int) (modin []int, output []int, err error) {
	modin = make([]int, len(prg))
	copy(modin, prg)
	output = []int{}
	pos := 0

	values := func(p1, p2 int, pm1, pm2 int) (int, int) {
		v1 := p1
		v2 := p2
		if pm1 == 0 {
			v1 = modin[p1]
		}
		if pm2 == 0 {
			v2 = modin[p2]
		}
		return v1, v2
	}

	value := func(p1 int, pm1 int) int {
		v1 := p1
		if pm1 == 0 {
			v1 = modin[p1]
		}
		return v1
	}

	for {
		oc := modin[pos]
		code, pm1, pm2, pm3 := squeezeOpcode(oc)
		if pm3 != 0 {
			return nil, nil, errors.Errorf("param-mode-3 is not 0 at pos %d, opcode %d", pos, oc)
		}
		switch code {
		case Halt:
			return modin, output, nil
		case Input:
			pr := modin[pos+1]
			modin[pr] = inputs[0]
			inputs = inputs[1:]
			pos += 2
		case Output:
			pr := modin[pos+1]
			v := value(pr, pm1)
			output = append(output, v)
			pos += 2
		case Add, Mult, Less, Equal:
			p1 := modin[pos+1]
			p2 := modin[pos+2]
			v1, v2 := values(p1, p2, pm1, pm2)
			r := int(0)
			switch code {
			case Add:
				r = v1 + v2
			case Mult:
				r = v1 * v2
			case Less:
				if v1 < v2 {
					r = 1
				}
			case Equal:
				if v1 == v2 {
					r = 1
				}
			}
			pr := modin[pos+3]
			modin[pr] = r
			pos += 4
		case JmpIfTrue:
			p1 := modin[pos+1]
			p2 := modin[pos+2]
			v1, v2 := values(p1, p2, pm1, pm2)
			if v1 != 0 {
				pos = v2
			} else {
				pos += 3
			}
		case JmpIfFalse:
			p1 := modin[pos+1]
			p2 := modin[pos+2]
			v1, v2 := values(p1, p2, pm1, pm2)
			if v1 == 0 {
				pos = v2
			} else {
				pos += 3
			}
		default:
			return nil, nil, errors.Errorf("invalid opcode %d", modin[pos])
		}
	}
}

func ExecC(prg []int, inC chan int, outC chan int) {
	modin := make([]int, len(prg))
	copy(modin, prg)
	//output = []int{}
	pos := 0

	values := func(p1, p2 int, pm1, pm2 int) (int, int) {
		v1 := p1
		v2 := p2
		if pm1 == 0 {
			v1 = modin[p1]
		}
		if pm2 == 0 {
			v2 = modin[p2]
		}
		return v1, v2
	}

	value := func(p1 int, pm1 int) int {
		v1 := p1
		if pm1 == 0 {
			v1 = modin[p1]
		}
		return v1
	}

	for {
		oc := modin[pos]
		code, pm1, pm2, pm3 := squeezeOpcode(oc)
		if pm3 != 0 {
			panic(errors.Errorf("param-mode-3 is not 0 at pos %d, opcode %d", pos, oc))
		}
		switch code {
		case Halt:
			return
		case Input:
			pr := modin[pos+1]
			modin[pr] = <-inC
			pos += 2
		case Output:
			pr := modin[pos+1]
			v := value(pr, pm1)
			outC <- v
			pos += 2
		case Add, Mult, Less, Equal:
			p1 := modin[pos+1]
			p2 := modin[pos+2]
			v1, v2 := values(p1, p2, pm1, pm2)
			r := int(0)
			switch code {
			case Add:
				r = v1 + v2
			case Mult:
				r = v1 * v2
			case Less:
				if v1 < v2 {
					r = 1
				}
			case Equal:
				if v1 == v2 {
					r = 1
				}
			}
			pr := modin[pos+3]
			modin[pr] = r
			pos += 4
		case JmpIfTrue:
			p1 := modin[pos+1]
			p2 := modin[pos+2]
			v1, v2 := values(p1, p2, pm1, pm2)
			if v1 != 0 {
				pos = v2
			} else {
				pos += 3
			}
		case JmpIfFalse:
			p1 := modin[pos+1]
			p2 := modin[pos+2]
			v1, v2 := values(p1, p2, pm1, pm2)
			if v1 == 0 {
				pos = v2
			} else {
				pos += 3
			}
		default:
			panic(errors.Errorf("invalid opcode %d", modin[pos]))
		}
	}
}
