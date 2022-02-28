package day_19

import "github.com/pkg/errors"

type Instr struct {
	Op   string
	Args []int
}

func NewMachine(ip int, initial [6]int, iss []Instr) *Machine {
	return &Machine{
		ip:    ip,
		ipval: 0,
		regs:  initial,
		iss:   iss,
	}
}

type Machine struct {
	ip    int
	ipval int
	regs  [6]int
	iss   []Instr
}

const (
	addr = "addr"
	addi = "addi"
	mulr = "mulr"
	muli = "muli"
	banr = "banr"
	bani = "bani"
	borr = "borr"
	bori = "bori"
	setr = "setr"
	seti = "seti"
	gtir = "gtir"
	gtri = "gtri"
	gtrr = "gtrr"
	eqir = "eqir"
	eqri = "eqri"
	eqrr = "eqrr"
)

func allOps() []string {
	return []string{
		addr,
		addi,
		mulr,
		muli,
		banr,
		bani,
		borr,
		bori,
		setr,
		seti,
		gtir,
		gtri,
		gtrr,
		eqir,
		eqri,
		eqrr,
	}
}

func gt(a, b int) int {
	if a > b {
		return 1
	}
	return 0
}

func eq(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}

func (m *Machine) Process(op string, args [3]int) error {
	switch op {
	case addr:
		m.regs[args[2]] = m.regs[args[0]] + m.regs[args[1]]
	case addi:
		m.regs[args[2]] = m.regs[args[0]] + args[1]
	case mulr:
		m.regs[args[2]] = m.regs[args[0]] * m.regs[args[1]]
	case muli:
		m.regs[args[2]] = m.regs[args[0]] * args[1]
	case banr:
		m.regs[args[2]] = m.regs[args[0]] & m.regs[args[1]]
	case bani:
		m.regs[args[2]] = m.regs[args[0]] & args[1]
	case borr:
		m.regs[args[2]] = m.regs[args[0]] | m.regs[args[1]]
	case bori:
		m.regs[args[2]] = m.regs[args[0]] | args[1]
	case setr:
		m.regs[args[2]] = m.regs[args[0]]
	case seti:
		m.regs[args[2]] = args[0]
	case gtir:
		m.regs[args[2]] = gt(args[0], m.regs[args[1]])
	case gtri:
		m.regs[args[2]] = gt(m.regs[args[0]], args[1])
	case gtrr:
		m.regs[args[2]] = gt(m.regs[args[0]], m.regs[args[1]])
	case eqir:
		m.regs[args[2]] = eq(args[0], m.regs[args[1]])
	case eqri:
		m.regs[args[2]] = eq(m.regs[args[0]], args[1])
	case eqrr:
		m.regs[args[2]] = eq(m.regs[args[0]], m.regs[args[1]])
	default:
		return errors.Errorf("invalid op %q", op)
	}
	return nil
}

func (m *Machine) Run() {
	for {
		if m.ipval < 0 || m.ipval >= len(m.iss) {
			return
		}
		if m.ipval == 1 {
			log("%v", m.regs)
		}
		is := m.iss[m.ipval]

		m.regs[m.ip] = m.ipval
		m.Process(is.Op, [3]int{is.Args[0], is.Args[1], is.Args[2]})
		m.ipval = m.regs[m.ip]
		m.ipval++
	}
}
