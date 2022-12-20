package day_17

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/set"
	"github.com/mazzegi/adventofcode/slices"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
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

const (
	left  = '<'
	right = '>'
)

type Rock [][]rune

var rocks = []Rock{
	{
		{'#', '#', '#', '#'},
	},
	{
		{'.', '#', '.'},
		{'#', '#', '#'},
		{'.', '#', '.'},
	},
	{
		{'.', '.', '#'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	},
	{
		{'#'},
		{'#'},
		{'#'},
		{'#'},
	},
	{
		{'#', '#'},
		{'#', '#'},
	},
}

func (rc Rock) Width() int {
	return len(rc[0])
}

func (rc Rock) bottomXs() []int {
	var bxs []int
	bl := rc[len(rc)-1]
	for x, r := range bl {
		if r == '#' {
			bxs = append(bxs, x)
		}
	}
	return bxs
}

func part1MainFunc(in string) (int, error) {
	flows := make([]int, len(in))
	for i, r := range in {
		switch r {
		case left:
			flows[i] = -1
		case right:
			flows[i] = 1
		default:
			fatal("invalid flow rune %q", string(r))
		}
	}

	flowIdx := 0
	rockIdx := 0
	width := 7
	cnt := 2022

	occ := set.New[grid.Point]()
	maxY := 0

	for i := 0; i < cnt; i++ {
		rock := rocks[rockIdx%len(rocks)]
		xpos := 2
		ypos := maxY + 3
		for {
			dx := flows[flowIdx]
			flowIdx++
			if flowIdx >= len(flows) {
				flowIdx = 0
			}
			nx := xpos + dx
			if nx > 0 && nx+rock.Width() <= width {
				xpos = nx
			}
			// down
			canFall := ypos > 0
			if canFall {
				bxs := rock.bottomXs()
				for _, bx := range bxs {
					if occ.Contains(grid.Pt(xpos+bx, ypos-1)) {
						canFall = false
						break
					}
				}
			}
			if canFall {
				ypos--
				continue
			}
			for y := 0; y < len(rock); y++ {
				ri := len(rock) - 1 - y
				for x, r := range rock[ri] {
					if r == '#' {
						occ.Insert(grid.Pt(xpos+x, ypos+y))
						if ypos+y > maxY {
							maxY = ypos + y
						}
					}
				}
			}
			dump(occ, width)
			break
		}

		rockIdx++
	}

	return maxY, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}

func dump(occ *set.Set[grid.Point], width int) {
	var maxY int
	for _, p := range occ.Values() {
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	var sl []string
	for y := 0; y <= maxY; y++ {
		rsl := "|"
		for x := 0; x < width; x++ {
			if occ.Contains(grid.Pt(x, y)) {
				rsl += "#"
			} else {
				rsl += "."
			}
		}
		rsl += "|"
		sl = append(sl, rsl)
	}
	sl = slices.Reverse(sl)
	for _, s := range sl {
		log(s)
	}
	log("+-------+")
}
