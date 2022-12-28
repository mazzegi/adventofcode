package day_24

import "github.com/mazzegi/adventofcode/grid"

func NewFinder(g *Grid) *Finder {
	return &Finder{
		grid:         g,
		currShortest: -1,
	}
}

type Finder struct {
	grid         *Grid
	currShortest int
	end          grid.Point
}

func (f *Finder) Find() int {
	start, end := f.grid.StartEnd()
	f.end = end
	gc := f.grid.Clone()
	res := f.findRec(gc, start, 0)
	return res
}

func (f *Finder) findRec(grid *Grid, pos grid.Point, iter int) int {
	if f.currShortest > 0 && iter >= f.currShortest {
		return -1
	}

	grid.Move()
	ons := grid.OpenNeighbours(pos)
	var min int
	for i, on := range ons {
		if on == f.end {
			f.currShortest = iter
			return iter
		}
		gc := grid.Clone()
		subRes := f.findRec(gc, on, iter+1)
		if subRes < 0 {
			continue
		}
		if i == 0 || subRes < min {
			min = subRes
		}
	}
	return min
}
