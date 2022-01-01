package ship

import (
	"fmt"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type Op rune

const (
	OpMvNorth   Op = 'N'
	OpMvSouth   Op = 'S'
	OpMvEast    Op = 'E'
	OpMvWest    Op = 'W'
	OpTurnLeft  Op = 'L'
	OpTurnRight Op = 'R'
	OpMvFwd     Op = 'F'
)

type Instruction struct {
	Op    Op
	Value int
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %d", string(i.Op), i.Value)
}

func ParseInstruction(s string) (Instruction, error) {
	if len(s) < 2 {
		return Instruction{}, errors.Errorf("invalid instruction %q", s)
	}
	op := rune(s[0])
	inst := Instruction{}
	switch Op(op) {
	case OpMvNorth, OpMvSouth, OpMvEast, OpMvWest, OpTurnLeft, OpTurnRight, OpMvFwd:
		inst.Op = Op(op)
	default:
		return Instruction{}, errors.Errorf("invalid op %q", string(op))
	}
	vs := s[1:]
	v, err := strconv.ParseInt(vs, 10, 64)
	if err != nil {
		return Instruction{}, errors.Wrap(err, "parse-int")
	}
	inst.Value = int(v)
	return inst, nil
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Pos struct {
	EW, NS int
}

func (p Pos) Distance(op Pos) int {
	return AbsInt(p.EW-op.EW) + AbsInt(p.NS-op.NS)
}

func (p Pos) String() string {
	var s string
	if p.EW >= 0 {
		s += fmt.Sprintf("%d east, ", p.EW)
	} else {
		s += fmt.Sprintf("%d west, ", AbsInt(p.EW))
	}

	if p.NS >= 0 {
		s += fmt.Sprintf("%d north", p.NS)
	} else {
		s += fmt.Sprintf("%d south", AbsInt(p.NS))
	}

	return s
}

type Dir int

const (
	DirEast  Dir = 0
	DirSouth Dir = 90
	DirWest  Dir = 180
	DirNorth Dir = 270
)

func (d Dir) Added(v int) Dir {
	if v >= 0 {
		nv := int(d) + v
		return Dir(nv % 360)
	} else {
		nv := 360 + int(d) + v
		return Dir(nv % 360)
	}

}

func (d Dir) String() string {
	switch d {
	case DirEast:
		return "east"
	case DirSouth:
		return "south"
	case DirWest:
		return "west"
	case DirNorth:
		return "north"
	default:
		return "unknown"
	}
}

type Ship struct {
	StartPos Pos
	Pos      Pos
	Dir      Dir
}

func New(pos Pos, dir Dir) *Ship {
	return &Ship{
		StartPos: pos,
		Pos:      pos,
		Dir:      dir,
	}
}

func (s *Ship) Distance() int {
	return s.Pos.Distance(s.StartPos)
}

func (s *Ship) Report() string {
	return fmt.Sprintf("at %q, facing %q, distance: %d", s.Pos.String(), s.Dir.String(), s.Distance())
}

func (s *Ship) Process(inst Instruction) {
	switch inst.Op {
	case OpMvNorth:
		s.Pos.NS += inst.Value
	case OpMvSouth:
		s.Pos.NS -= inst.Value
	case OpMvEast:
		s.Pos.EW += inst.Value
	case OpMvWest:
		s.Pos.EW -= inst.Value
	case OpTurnLeft:
		s.Dir = s.Dir.Added(-inst.Value)
	case OpTurnRight:
		s.Dir = s.Dir.Added(inst.Value)
	case OpMvFwd:
		switch s.Dir {
		case DirEast:
			s.Pos.EW += inst.Value
		case DirWest:
			s.Pos.EW += -inst.Value
		case DirNorth:
			s.Pos.NS += inst.Value
		case DirSouth:
			s.Pos.NS -= inst.Value
		}
	default:
		panic("unknown instruction")
	}
}

//
type ShipWithWaypoint struct {
	StartPos Pos
	Pos      Pos
	Waypoint Pos
}

func NewWithWaypoint(pos Pos, waypoint Pos) *ShipWithWaypoint {
	return &ShipWithWaypoint{
		StartPos: pos,
		Pos:      pos,
		Waypoint: waypoint,
	}
}

func (s *ShipWithWaypoint) Distance() int {
	return s.Pos.Distance(s.StartPos)
}

func (s *ShipWithWaypoint) Report() string {
	return fmt.Sprintf("at %q, waypoint: %q, distance: %d", s.Pos.String(), s.Waypoint.String(), s.Distance())
}

func (s *ShipWithWaypoint) RotateWaypoint(deg int) {
	rad := 2 * math.Pi * float64(deg) / 360.0
	ewn := float64(s.Waypoint.EW)*math.Cos(rad) - float64(s.Waypoint.NS)*math.Sin(rad)
	nsn := float64(s.Waypoint.EW)*math.Sin(rad) + float64(s.Waypoint.NS)*math.Cos(rad)
	s.Waypoint.EW = int(math.Round(ewn))
	s.Waypoint.NS = int(math.Round(nsn))
}

func (s *ShipWithWaypoint) Process(inst Instruction) {
	switch inst.Op {
	case OpMvNorth:
		s.Waypoint.NS += inst.Value
	case OpMvSouth:
		s.Waypoint.NS -= inst.Value
	case OpMvEast:
		s.Waypoint.EW += inst.Value
	case OpMvWest:
		s.Waypoint.EW -= inst.Value
	case OpTurnLeft:
		s.RotateWaypoint(inst.Value)
	case OpTurnRight:
		s.RotateWaypoint(-inst.Value)
	case OpMvFwd:
		s.Pos.EW += inst.Value * (s.Waypoint.EW)
		s.Pos.NS += inst.Value * (s.Waypoint.NS)
	default:
		panic("unknown instruction")
	}
}
