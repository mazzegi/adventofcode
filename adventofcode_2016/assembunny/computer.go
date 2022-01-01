package assembunny

import "strconv"

type Register struct {
	ID    string
	Value int
}

type Computer struct {
	registers map[string]*Register
	//out       int
}

func NewComputer(regs []string) *Computer {
	c := &Computer{
		registers: map[string]*Register{},
	}
	for _, reg := range regs {
		c.registers[reg] = &Register{
			ID:    reg,
			Value: 0,
		}
	}
	return c
}

func (c *Computer) ValueOf(reg string) int {
	if reg, ok := c.registers[reg]; ok {
		return reg.Value
	}
	return 0
}

func (c *Computer) SetReg(reg string, value int) {
	if reg, ok := c.registers[reg]; ok {
		reg.Value = value
	}
}

func (c *Computer) IncReg(reg string) {
	if reg, ok := c.registers[reg]; ok {
		reg.Value++
	}
}

func (c *Computer) DecReg(reg string) {
	if reg, ok := c.registers[reg]; ok {
		reg.Value--
	}
}

func (c *Computer) resolve(s string) (int, bool) {
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(n), true
	}
	if reg, ok := c.registers[s]; ok {
		return reg.Value, true
	}
	return 0, false
}

type CbResult string

const (
	CbProceed CbResult = "proceed"
	CbQuit    CbResult = "quit"
)

func (c *Computer) Execute(iss []Instruction, outFnc func(out int) CbResult) {
	pos := 0
	toggle := func(n int) {
		tidx := pos + n
		if tidx < 0 || tidx >= len(iss) {
			return
		}
		is := iss[tidx]
		var tis Instruction
		switch is := is.(type) {
		case Inc:
			tis = Dec{Reg: is.Reg}
		case Dec:
			tis = Inc{Reg: is.Reg}
		case Tgl:
			tis = Inc{Reg: is.Dst}
		case Jnz:
			tis = Cpy{
				Src: is.Cond,
				Dst: is.Steps,
			}
		case Cpy:
			tis = Jnz{
				Cond:  is.Src,
				Steps: is.Dst,
			}
		}

		iss[tidx] = tis
	}
loop:
	for {
		if pos < 0 || pos >= len(iss) {
			return
		}

		is := iss[pos]
		switch is := is.(type) {
		case Cpy:
			if val, ok := c.resolve(is.Src); ok {
				c.SetReg(is.Dst, val)
			}
		case Inc:
			c.IncReg(is.Reg)
		case Dec:
			c.DecReg(is.Reg)
		case Jnz:
			condVal, ok := c.resolve(is.Cond)
			if !ok || condVal == 0 {
				pos++
				continue loop
			}
			steps, ok := c.resolve(is.Steps)
			if !ok {
				pos++
				continue loop
			}
			pos += steps
			continue loop

		case Tgl:
			if tval, ok := c.resolve(is.Dst); ok {
				toggle(tval)
			}
		case Out:
			if v, ok := c.resolve(is.Src); ok {
				if outFnc != nil {
					res := outFnc(v)
					switch res {
					case CbQuit:
						return
					}
				}
			}
		}

		pos++
	}
}
