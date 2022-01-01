package bitmask

import (
	"github.com/pkg/errors"
)

type Processor struct {
	Programm []Command
	mem      map[uint64]uint64
	mask     Mask
}

func NewProcessor(cmds []Command) *Processor {
	return &Processor{
		Programm: cmds,
		mem:      map[uint64]uint64{},
	}
}

func (p *Processor) Sum() uint64 {
	var sum uint64
	for _, v := range p.mem {
		sum += v
	}
	return sum
}

func (p *Processor) Process() error {
	if len(p.Programm) == 0 {
		return nil
	}
	firstMask, ok := p.Programm[0].(ChangeMask)
	if !ok {
		return errors.Errorf("first command must be a mask-command")
	}
	p.mask = firstMask.mask
	for i := 1; i < len(p.Programm); i++ {
		switch c := p.Programm[i].(type) {
		case ChangeMask:
			p.mask = c.mask
		case Poke:
			p.Poke(c.addr, c.val)
		default:
			return errors.Errorf("invalid command %T", c)
		}
	}
	return nil
}

func (p *Processor) Poke(addr, value uint64) {
	mval := value
	for i := 0; i < len(p.mask.raw); i++ {
		maskBit := rune(p.mask.raw[len(p.mask.raw)-1-i])
		switch maskBit {
		case '0':
			mval = mval &^ (1 << i)
		case '1':
			mval = mval | (1 << i)
		case 'X':
			continue
		}
	}
	p.mem[addr] = mval
}

func (p *Processor) ProcessV2() error {
	if len(p.Programm) == 0 {
		return nil
	}
	firstMask, ok := p.Programm[0].(ChangeMask)
	if !ok {
		return errors.Errorf("first command must be a mask-command")
	}
	p.mask = firstMask.mask
	for i := 1; i < len(p.Programm); i++ {
		switch c := p.Programm[i].(type) {
		case ChangeMask:
			p.mask = c.mask
		case Poke:
			p.PokeV2(c.addr, c.val)
		default:
			return errors.Errorf("invalid command %T", c)
		}
	}
	return nil
}

func (p *Processor) PokeV2(addr, value uint64) {
	rmask := make([]rune, len(p.mask.raw))
	var floatBits []int
	for i := 0; i < len(p.mask.raw); i++ {
		revi := len(p.mask.raw) - 1 - i
		maskBit := rune(p.mask.raw[revi])
		switch maskBit {
		case '0':
			if addr&(1<<i) == (1 << i) {
				//bit set in addr, keep addr bit
				rmask[revi] = '1'
			} else {
				rmask[revi] = '0'
			}
		case '1':
			rmask[revi] = '1'
		case 'X':
			rmask[revi] = 'X'
			floatBits = append(floatBits, i)
		}
	}
	//fmt.Printf("=> %q\n", string(rmask))
	cntBits := make([]int, len(floatBits))
outer:
	for {
		//apply mask to rmask
		var faddr uint64
		fbit := 0
		for i := 0; i < len(rmask); i++ {
			revi := len(rmask) - 1 - i
			maskBit := rune(rmask[revi])
			switch maskBit {
			case '1':
				faddr |= (1 << i)
			case 'X':
				if cntBits[fbit] == 1 {
					faddr |= (1 << i)
				}
				fbit++
			}
		}
		p.mem[faddr] = value

		//fmt.Printf("bits %v => %d\n", cntBits, faddr)

		for i := 0; i < len(cntBits); i++ {
			if cntBits[i] < 1 {
				cntBits[i]++
				continue outer
			}
			cntBits[i] = 0
		}
		return
	}
}
