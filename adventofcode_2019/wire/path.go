package wire

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Step struct {
	Dir   rune
	Count int
}

func (s Step) String() string {
	return fmt.Sprintf("%s:%d", string(s.Dir), s.Count)
}

func ParseStep(s string) (Step, error) {
	if len(s) < 2 {
		return Step{}, errors.Errorf("invalid step %q", s)
	}
	dir := rune(s[0])
	count, err := strconv.ParseInt(s[1:], 10, 64)
	if err != nil {
		return Step{}, errors.Wrapf(err, "parse-int %q", s[1:])
	}
	return Step{
		Dir:   dir,
		Count: int(count),
	}, nil
}

type Point struct {
	X, Y int
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (pt Point) ManhattanDist() int {
	return AbsInt(pt.X) + AbsInt(pt.Y)
}

type Path struct {
	Steps  []Step
	Lookup map[Point]int
}

func (p *Path) String() string {
	sl := make([]string, len(p.Steps))
	for i := 0; i < len(p.Steps); i++ {
		sl[i] = p.Steps[i].String()
	}
	return strings.Join(sl, ", ")
}

func ParsePath(sp string) (*Path, error) {
	p := &Path{
		Lookup: map[Point]int{},
	}
	sl := strings.Split(sp, ",")
	for _, s := range sl {
		step, err := ParseStep(strings.Trim(s, " "))
		if err != nil {
			return nil, errors.Wrap(err, "parse-step")
		}
		p.Steps = append(p.Steps, step)
	}
	err := p.BuildLookup()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Path) BuildLookup() error {
	pt := Point{
		X: 0,
		Y: 0,
	}
	p.Lookup[pt] = 0
	var totSteps int
	for _, step := range p.Steps {
		for i := 0; i < step.Count; i++ {
			switch step.Dir {
			case 'L':
				pt.X -= 1
			case 'R':
				pt.X += 1
			case 'U':
				pt.Y += 1
			case 'D':
				pt.Y -= 1
			default:
				return errors.Errorf("invalid direction %q", string(step.Dir))
			}
			totSteps++
			p.Lookup[pt] = totSteps
		}
	}
	return nil
}

func (p *Path) Intersections(op *Path) []Point {
	var is []Point
	for pt, steps := range p.Lookup {
		if pt.X == 0 && pt.Y == 0 {
			continue
		}
		if osteps, ok := op.Lookup[pt]; ok {
			is = append(is, pt)
			fmt.Printf("intersection: (%d, %d) => (%d, %d)\n", pt.X, pt.Y, steps, osteps)
		}
	}
	return is
}

func (p *Path) BestIntersection(op *Path) int {
	var is []Point
	var bestPt Point
	var bestSteps int
	for pt, steps := range p.Lookup {
		if pt.X == 0 && pt.Y == 0 {
			continue
		}
		if osteps, ok := op.Lookup[pt]; ok {
			is = append(is, pt)
			fmt.Printf("intersection: (%d, %d) => (%d, %d)\n", pt.X, pt.Y, steps, osteps)
			if bestSteps == 0 || steps+osteps < bestSteps {
				bestPt = pt
				bestSteps = steps + osteps
			}
		}
	}
	fmt.Printf("best: (%d, %d) => %d\n", bestPt.X, bestPt.Y, bestSteps)
	return bestSteps
}
