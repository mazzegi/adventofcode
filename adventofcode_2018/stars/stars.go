package stars

import (
	"fmt"
	"strings"
)

func Parse(s string) (*Star, error) {
	var px, py, vx, vy int
	_, err := fmt.Sscanf(s, "position=<%d, %d> velocity=<%d, %d>", &px, &py, &vx, &vy)
	if err != nil {
		return nil, err
	}
	return &Star{
		Pos: Vector{X: px, Y: py},
		Vel: Vector{X: vx, Y: vy},
	}, nil
}

type Vector struct {
	X, Y int
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

type Star struct {
	Pos Vector
	Vel Vector
}

func (s Star) String() string {
	return fmt.Sprintf("P: %s, V: %s", s.Pos, s.Vel)
}

type Galaxy struct {
	sts []*Star
	occ map[Vector]bool
	min Vector
	max Vector
}

func NewGalaxy(sts []*Star) *Galaxy {
	g := &Galaxy{
		sts: sts,
		occ: map[Vector]bool{},
	}
	g.calcAux()
	return g
}

func (g *Galaxy) calcAux() {
	occ := map[Vector]bool{}
	min := g.sts[0].Pos
	max := g.sts[0].Pos
	for _, s := range g.sts {
		occ[s.Pos] = true
		if s.Pos.X < min.X {
			min.X = s.Pos.X
		}
		if s.Pos.X > max.X {
			max.X = s.Pos.X
		}
		if s.Pos.Y < min.Y {
			min.Y = s.Pos.Y
		}
		if s.Pos.Y > max.Y {
			max.Y = s.Pos.Y
		}
	}
	g.min = min
	g.max = max
	g.occ = occ
}

func (g *Galaxy) Step() {
	for _, s := range g.sts {
		s.Pos.X += s.Vel.X
		s.Pos.Y += s.Vel.Y
	}
	g.calcAux()
}

func (g *Galaxy) Dump() string {
	var sl []string
	for y := g.min.Y; y <= g.max.Y; y++ {
		var srow string
		for x := g.min.X; x <= g.max.X; x++ {
			if _, ok := g.occ[Vector{X: x, Y: y}]; ok {
				srow += "#"
			} else {
				srow += "."
			}
		}
		sl = append(sl, srow)
	}
	return strings.Join(sl, "\n")
}

func (g *Galaxy) Score() int {
	return (g.max.X - g.min.X) * (g.max.Y - g.min.Y)
}
