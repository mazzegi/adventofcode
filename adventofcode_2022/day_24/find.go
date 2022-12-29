package day_24

import "github.com/mazzegi/adventofcode/grid"

func NewFinder(g *Grid, maxIter int) *Finder {
	return &Finder{
		grid:         g,
		currShortest: -1,
		maxIter:      maxIter,
		cache:        map[string]int{},
	}
}

type Finder struct {
	grid         *Grid
	currShortest int
	end          grid.Point
	maxIter      int
	cache        map[string]int
}

func (f *Finder) Find() int {
	start, end := f.grid.StartEnd()
	f.end = end
	gc := f.grid.Clone()
	res := f.findRec(gc, start, 1)
	return res
}

func (f *Finder) findRec(grid *Grid, pos grid.Point, iter int) int {
	hash := grid.Hash(pos, iter)
	if v, ok := f.cache[hash]; ok {
		return v
	}

	if iter >= f.maxIter {
		f.cache[hash] = -1
		return -1
	}
	if f.currShortest > 0 && iter >= f.currShortest {
		shortestPossible := iter + pos.ManhattenDistTo(f.end)
		if shortestPossible >= f.currShortest {
			f.cache[hash] = -1
			return -1
		}
	}
	grid.Move(1)

	ons := grid.OpenNeighbours(pos)
	min := -1
	for _, on := range ons {
		if on == f.end {
			if f.currShortest < 0 || iter < f.currShortest {
				f.currShortest = iter
			}
			min = iter
			break
		}
		gc := grid.Clone()
		subRes := f.findRec(gc, on, iter+1)
		if subRes < 0 {
			continue
		}
		if min < 0 || subRes < min {
			min = subRes
		}
	}

	f.cache[hash] = min
	return min
}
