package conway4

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Pos struct {
	X, Y, Z, W int
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d, %d, %d)", p.X, p.Y, p.Z, p.W)
}

func (p Pos) Neighbours() []Pos {
	var ns []Pos
	for x := p.X - 1; x <= p.X+1; x++ {
		for y := p.Y - 1; y <= p.Y+1; y++ {
			for z := p.Z - 1; z <= p.Z+1; z++ {
				for w := p.W - 1; w <= p.W+1; w++ {
					np := Pos{
						X: x,
						Y: y,
						Z: z,
						W: w,
					}
					if np == p {
						continue
					}
					ns = append(ns, np)
				}
			}
		}
	}
	return ns
}

type Grid struct {
	Current map[Pos]bool
	Next    map[Pos]bool
	Min     Pos
	Max     Pos
}

func NewGrid(initRowsRaw []string) (*Grid, error) {
	g := &Grid{
		Current: map[Pos]bool{},
		Next:    map[Pos]bool{},
	}
	for ri, row := range initRowsRaw {
		for ci, r := range row {
			switch r {
			case '.':
			case '#':
				p := Pos{
					X: ci,
					Y: ri,
					Z: 0,
					W: 0,
				}
				g.Current[p] = true
				g.AdaptMinMax(p)
			default:
				return nil, errors.Errorf("invalid pos char %q", string(r))
			}
		}
	}
	return g, nil
}

func (g *Grid) ActiveCount() int {
	return len(g.Current)
}

func (g *Grid) AdaptMinMax(p Pos) {
	if p.X < g.Min.X {
		g.Min.X = p.X
	}
	if p.X > g.Max.X {
		g.Max.X = p.X
	}

	if p.Y < g.Min.Y {
		g.Min.Y = p.Y
	}
	if p.Y > g.Max.Y {
		g.Max.Y = p.Y
	}

	if p.Z < g.Min.Z {
		g.Min.Z = p.Z
	}
	if p.Z > g.Max.Z {
		g.Max.Z = p.Z
	}

	if p.W < g.Min.W {
		g.Min.W = p.W
	}
	if p.W > g.Max.W {
		g.Max.W = p.W
	}
}

func (g *Grid) CalcNext() {
	//g.DumpCurrent()
	g.Next = map[Pos]bool{}
	for y := g.Min.Y - 1; y <= g.Max.Y+1; y++ {
		for x := g.Min.X - 1; x <= g.Max.X+1; x++ {
			for z := g.Min.Z - 1; z <= g.Max.Z+1; z++ {
				for w := g.Min.W - 1; w <= g.Max.W+1; w++ {

					pos := Pos{X: x, Y: y, Z: z, W: w}
					nposs := pos.Neighbours()
					var actNeighbourCount int
					for _, npos := range nposs {
						if _, ok := g.Current[npos]; ok {
							actNeighbourCount++
						}
					}
					_, currActive := g.Current[pos]
					if currActive {
						if actNeighbourCount == 2 || actNeighbourCount == 3 {
							//remains active
							g.Next[pos] = true
							g.AdaptMinMax(pos)
						} else {
							// goes not to next active
							//fmt.Printf("%s goes inactive (ns = %d)\n", pos.String(), actNeighbourCount)
						}
					} else {
						if actNeighbourCount == 3 {
							//goes to next active
							g.Next[pos] = true
							g.AdaptMinMax(pos)
							//fmt.Printf("%s goes active (ns = %d)\n", pos.String(), actNeighbourCount)
						} else {
							//remains inactive - goes not to next active
						}
					}
				}
			}
		}
	}
}

func (g *Grid) DoNext() {
	g.CalcNext()
	g.Current = g.Next
}

func (g *Grid) DumpCurrent() {
	var sl []string
	for p := range g.Current {
		sl = append(sl, p.String())
	}
	fmt.Printf("%s. Min (%s), Max (%s)\n", strings.Join(sl, "; "), g.Min, g.Max)
}

func (g *Grid) Dump() {
	for z := g.Min.Z; z <= g.Max.Z; z++ {
		for w := g.Min.W; w <= g.Max.W; w++ {
			g.DumpZSlice(z, w)
			fmt.Printf("\n")
		}
	}
}

func (g *Grid) DumpZSlice(z int, w int) {
	fmt.Printf("z=%d, w=%d\n", z, w)
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			p := Pos{X: x, Y: y, Z: z}
			if _, ok := g.Current[p]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
