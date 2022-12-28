package day_23

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
)

func P(x, y int) grid.Point {
	return grid.Pt(x, y)
}

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatalIfErr(err error) {
	if err == nil {
		return
	}
	fatal("err not nil: %v", err)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//

func Adjacent(p grid.Point) []grid.Point {
	var aps []grid.Point
	for x := p.X - 1; x <= p.X+1; x++ {
		for y := p.Y - 1; y <= p.Y+1; y++ {
			ap := P(x, y)
			if ap == p {
				continue
			}
			aps = append(aps, ap)
		}
	}
	return aps
}

func PosAt(p grid.Point, dir Direction) grid.Point {
	switch dir {
	case N:
		return p.Add(P(0, -1))
	case NE:
		return p.Add(P(1, -1))
	case E:
		return p.Add(P(1, 0))
	case SE:
		return p.Add(P(1, 1))
	case S:
		return p.Add(P(0, 1))
	case SW:
		return p.Add(P(-1, 1))
	case W:
		return p.Add(P(-1, 0))
	case NW:
		return p.Add(P(-1, -1))
	default:
		fatal("invalid direction %q", dir)
	}
	return grid.Point{}
}

type Direction string

const (
	None Direction = "None"
	N    Direction = "N"
	NE   Direction = "NE"
	E    Direction = "E"
	SE   Direction = "SE"
	S    Direction = "S"
	SW   Direction = "SW"
	W    Direction = "W"
	NW   Direction = "NW"
)

func (dir Direction) Prev() Direction {
	switch dir {
	case N:
		return NW
	case NE:
		return N
	case E:
		return NE
	case SE:
		return E
	case S:
		return SE
	case SW:
		return S
	case W:
		return SW
	case NW:
		return W
	default:
		fatal("invalid direction %q", dir)
	}
	return None
}

func (dir Direction) Next() Direction {
	switch dir {
	case N:
		return NE
	case NE:
		return E
	case E:
		return SE
	case SE:
		return S
	case S:
		return SW
	case SW:
		return W
	case W:
		return NW
	case NW:
		return N
	default:
		fatal("invalid direction %q", dir)
	}
	return None
}

type Elve struct {
	ID       int
	Pos      grid.Point
	Proposal grid.Point
	Proposed bool
}

type Grid struct {
	elves map[grid.Point]*Elve
}

func (g *Grid) Bounds() (min, max grid.Point) {
	first := true
	for _, e := range g.elves {
		p := e.Pos
		if first {
			min, max = p, p
			first = false
		}

		if p.X < min.X {
			min.X = p.X
		}
		if p.X > max.X {
			max.X = p.X
		}

		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	return
}

func (g *Grid) TilesCount() (open, occupied int) {
	min, max := g.Bounds()
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			if _, ok := g.elves[P(x, y)]; !ok {
				open++
			} else {
				occupied++
			}
		}
	}
	return
}

func mustParseGrid(in string) *Grid {
	g := &Grid{
		elves: map[grid.Point]*Elve{},
	}
	nextID := 1
	for y, line := range readutil.ReadLines(in) {
		for x, r := range line {
			if r == '#' {
				p := P(x, y)
				g.elves[p] = &Elve{
					ID:       nextID,
					Pos:      p,
					Proposal: p,
					Proposed: false,
				}
				nextID++
			}
		}
	}
	return g
}

func (g *Grid) NoElveInEachOf(poss ...grid.Point) bool {
	for _, pos := range poss {
		if _, ok := g.elves[pos]; ok {
			return false
		}
	}
	return true
}

func (g *Grid) Propose(order []Direction) {
outer:
	for _, e := range g.elves {
		pe := e.Pos
		e.Proposed = false
		if g.NoElveInEachOf(Adjacent(pe)...) {
			continue outer
		}

		for _, dir := range order {
			if g.NoElveInEachOf(PosAt(pe, dir), PosAt(pe, dir.Prev()), PosAt(pe, dir.Next())) {
				e.Proposal = PosAt(pe, dir)
				e.Proposed = true
				continue outer
			}
		}
		//log("never should come here")
	}
}

func (g *Grid) Move() int {
	proposedPoints := map[grid.Point][]int{}
	for _, e := range g.elves {
		if !e.Proposed {
			continue
		}
		proposedPoints[e.Proposal] = append(proposedPoints[e.Proposal], e.ID)
	}

	notMoveElves := set.New[int]()
	for _, ppids := range proposedPoints {
		if len(ppids) > 1 {
			notMoveElves.Insert(ppids...)
		}
	}

	moveCount := 0
	for _, e := range g.elves {
		if !e.Proposed {
			continue
		}
		if notMoveElves.Contains(e.ID) {
			continue
		}
		delete(g.elves, e.Pos)
		e.Pos = e.Proposal
		if _, ok := g.elves[e.Pos]; ok {
			fatal("panic!!!")
		}
		g.elves[e.Pos] = e
		moveCount++
	}
	return moveCount
}

func part1MainFunc(in string) (int, error) {
	g := mustParseGrid(in)
	open, occ := g.TilesCount()
	log("open=%d, occ=%d", open, occ)

	dump(g, -1)
	order := []Direction{N, S, W, E}
	for i := 0; i < 10; i++ {
		g.Propose(order)
		g.Move()
		dump(g, i)
		//order = append(order[1:], order[0])
		neworder := append([]Direction{}, order[1:]...)
		neworder = append(neworder, order[0])
		order = neworder
	}

	open, occ = g.TilesCount()
	log("open=%d, occ=%d", open, occ)
	return open, nil
}

func part2MainFunc(in string) (int, error) {
	g := mustParseGrid(in)

	dump(g, -1)
	order := []Direction{N, S, W, E}
	var noMove int
	i := 0
	for {
		g.Propose(order)
		mcnt := g.Move()
		if mcnt == 0 {
			noMove = i + 1
			break
		}
		dump(g, i)
		//order = append(order[1:], order[0])
		neworder := append([]Direction{}, order[1:]...)
		neworder = append(neworder, order[0])
		order = neworder
		i++
	}
	return noMove, nil
}

const skipDump = true

func dump(g *Grid, iter int) {
	if skipDump {
		return
	}
	if iter == -1 {
		log("== initial state ==")
	} else {
		log("== end of round %d ==", iter+1)
	}

	min, max := g.Bounds()
	for y := min.Y; y <= max.Y; y++ {
		var s string
		for x := min.X; x <= max.X; x++ {
			if _, ok := g.elves[P(x, y)]; !ok {
				s += "."
			} else {
				s += "#"
			}
		}
		log(s)
	}
}
