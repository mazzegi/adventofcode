package intcode

import (
	"github.com/mazzegi/adventofcode/slices"
	"github.com/pkg/errors"
)

type IntReader interface {
	Read() (int, error)
}

type IntWriter interface {
	Write(int) error
}

func NewComputer(prg []int, inputReader IntReader, outputWriter IntWriter) *Computer {
	return &Computer{
		buffer:       slices.Clone(prg),
		inputReader:  inputReader,
		outputWriter: outputWriter,
	}
}

type Computer struct {
	buffer       []int
	inputReader  IntReader
	outputWriter IntWriter
}

const maxAddress int = 1e8

var appendChunk = slices.Repeat(0, 100)

func (com *Computer) ensureSize(size int) {
	if size < 0 || size > maxAddress {
		panic(errors.Errorf("ensure size 0 <= %d <= max-size", size))
	}
	for size > len(com.buffer) {
		com.buffer = append(com.buffer, appendChunk...)
	}
}

func (com *Computer) readFrom(pos int) int {
	if pos < 0 {
		panic(errors.Errorf("read from %d", pos))
	}
	com.ensureSize(pos + 1)
	return com.buffer[pos]
}

func (com *Computer) writeTo(pos int, value int) {
	if pos < 0 {
		panic(errors.Errorf("write to %d", pos))
	}
	com.ensureSize(pos + 1)
	com.buffer[pos] = value
}

func (com *Computer) mustReadInput() int {
	v, err := com.inputReader.Read()
	if err != nil {
		panic(err)
	}
	return v
}

func (com *Computer) mustWriteOutput(v int) {
	err := com.outputWriter.Write(v)
	if err != nil {
		panic(err)
	}
}

const (
	modePosition  int = 0
	modeImmediate int = 1
	modeRelative  int = 2
)

func (com *Computer) Exec() error {
	var relativeBase int = 0

	// address := func(v int, mode int) int {
	// 	switch mode {
	// 	case modePosition:
	// 		return v
	// 	case modeRelative:
	// 		return relativeBase + v
	// 	default:
	// 		panic("cannot address of immediate")
	// 	}
	// }

	value := func(v int, mode int) int {
		switch mode {
		case modePosition:
			return com.readFrom(v)
		case modeImmediate:
			return v
		case modeRelative:
			return com.readFrom(relativeBase + v)
		}
		return v
	}

	values := func(p1, p2 int, pm1, pm2 int) (int, int) {
		return value(p1, pm1), value(p2, pm2)
	}

	pos := 0
	for {
		oc := com.buffer[pos]
		code, mode1, mode2, mode3 := squeezeOpcode(oc)
		if mode3 != modePosition {
			return errors.Errorf("param-mode-3 is not 0 at pos %d, opcode %d", pos, oc)
		}
		switch code {
		case Halt:
			return nil
		case Input:
			pr := com.readFrom(pos + 1)
			com.writeTo(pr, com.mustReadInput())
			//com.writeTo(address(pr, mode1), com.mustReadInput())
			pos += 2
		case Output:
			pr := com.readFrom(pos + 1)
			v := value(pr, mode1)
			com.mustWriteOutput(v)
			pos += 2
		case Add, Mult, Less, Equal:
			v1, v2 := values(com.readFrom(pos+1), com.readFrom(pos+2), mode1, mode2)
			r := int(0)
			switch {
			case code == Add:
				r = v1 + v2
			case code == Mult:
				r = v1 * v2
			case code == Less && v1 < v2:
				r = 1
			case code == Equal && v1 == v2:
				r = 1
			}
			pr := com.readFrom(pos + 3)
			com.writeTo(pr, r)
			pos += 4
		case JmpIfTrue:
			v1, v2 := values(com.readFrom(pos+1), com.readFrom(pos+2), mode1, mode2)
			if v1 != 0 {
				pos = v2
			} else {
				pos += 3
			}
		case JmpIfFalse:
			v1, v2 := values(com.readFrom(pos+1), com.readFrom(pos+2), mode1, mode2)
			if v1 == 0 {
				pos = v2
			} else {
				pos += 3
			}
		case RelBaseOffset:
			v := value(com.readFrom(pos+1), mode1)
			relativeBase += v
			pos += 2
		default:
			return errors.Errorf("invalid opcode %d", code)
		}
	}
}
