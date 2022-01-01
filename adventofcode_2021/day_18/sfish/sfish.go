package sfish

import "fmt"

type Element interface {
	String() string
	Magnitude() int
	Clone() Element
}

type Regular int

func (r Regular) String() string {
	return fmt.Sprintf("%d", r)
}

func (r Regular) Magnitude() int {
	return int(r)
}

func (r Regular) Clone() Element {
	return Regular(r)
}

type Pair struct {
	Left  Element
	Right Element
}

func Add(p1, p2 *Pair) *Pair {
	rp := &Pair{
		Left:  p1.Clone(),
		Right: p2.Clone(),
	}
	Reduce(rp)
	return rp
}

func (p *Pair) String() string {
	return fmt.Sprintf("[%s,%s]", p.Left.String(), p.Right.String())
}

func (p *Pair) Magnitude() int {
	return 3*p.Left.Magnitude() + 2*p.Right.Magnitude()
}

func (p *Pair) Clone() Element {
	cp := &Pair{
		Left:  p.Left.Clone(),
		Right: p.Right.Clone(),
	}
	return cp
}
