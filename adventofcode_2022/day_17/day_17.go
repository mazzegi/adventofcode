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
	res, err := part1MainFunc(input, 2022)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input, 1000000000000)
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

func (rc Rock) Height() int {
	return len(rc)
}

func part1MainFunc(in string, cnt int) (int, error) {
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
	//cnt := 2022

	occ := set.New[grid.Point]()
	maxY := 0

	tr := func(dx int) string {
		if dx == -1 {
			return "left"
		}
		return "right"
	}
	_ = tr

	collidesWithRockAt := func(rock Rock, xr, yr int) bool {
		for x := 0; x < rock.Width(); x++ {
			for iy := 0; iy < rock.Height(); iy++ {
				y := rock.Height() - 1 - iy
				if rock[y][x] == '#' {
					ax := xr + x
					ay := yr + iy
					if occ.Contains(grid.Pt(ax, ay)) {
						return true
					}
				}
			}
		}
		return false
	}

	for i := 0; i < cnt; i++ {
		rock := rocks[rockIdx%len(rocks)]
		xpos := 2
		ypos := maxY + 3
		dumpWithRock(occ, width, rock, xpos, ypos)
		//log("new rock begins falling (%d,%d)", xpos, ypos)
		for {
			dx := flows[flowIdx]
			flowIdx++
			if flowIdx >= len(flows) {
				flowIdx = 0
			}
			nx := xpos + dx
			if nx >= 0 && nx+rock.Width() <= width {
				//check if new pos collides with occ
				if !collidesWithRockAt(rock, nx, ypos) {
					xpos = nx
				}
			}
			dumpWithRock(occ, width, rock, xpos, ypos)

			canFall := ypos > 0 && !collidesWithRockAt(rock, xpos, ypos-1)

			if canFall {
				ypos--
				dumpWithRock(occ, width, rock, xpos, ypos)
				continue
			}
			for y := 0; y < len(rock); y++ {
				ri := len(rock) - 1 - y
				for x, r := range rock[ri] {
					if r == '#' {
						if occ.Contains(grid.Pt(xpos+x, ypos+y)) {
							fatal("not that way")
						}
						occ.Insert(grid.Pt(xpos+x, ypos+y))
						if ypos+y+1 > maxY {
							maxY = ypos + y + 1
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

func part2MainFunc(in string, cnt int) (int, error) {
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

	width := 7
	occ := set.New[grid.Point]()
	maxY := 0
	collidesWithRockAt := func(rock Rock, xr, yr int) bool {
		for x := 0; x < rock.Width(); x++ {
			for iy := 0; iy < rock.Height(); iy++ {
				y := rock.Height() - 1 - iy
				if rock[y][x] == '#' {
					ax := xr + x
					ay := yr + iy
					if occ.Contains(grid.Pt(ax, ay)) {
						return true
					}
				}
			}
		}
		return false
	}

	type result struct {
		flowIdx int
		rockIdx int
		xRest   int
	}
	type mark struct {
		iter int
		ypos int
	}
	results := map[result]mark{}

	flowIdx := 0
	rockIdx := 0

	for i := 0; i < cnt; i++ {
		rock := rocks[rockIdx]
		rockIdx++
		if rockIdx >= len(rocks) {
			rockIdx = 0
		}
		xpos := 2
		ypos := maxY + 3
		dumpWithRock(occ, width, rock, xpos, ypos)
		//log("new rock begins falling (%d,%d)", xpos, ypos)
		for {
			dx := flows[flowIdx]
			flowIdx++
			if flowIdx >= len(flows) {
				flowIdx = 0
			}
			nx := xpos + dx
			if nx >= 0 && nx+rock.Width() <= width {
				//check if new pos collides with occ
				if !collidesWithRockAt(rock, nx, ypos) {
					xpos = nx
				}
			}
			dumpWithRock(occ, width, rock, xpos, ypos)

			canFall := ypos > 0 && !collidesWithRockAt(rock, xpos, ypos-1)

			if canFall {
				ypos--
				dumpWithRock(occ, width, rock, xpos, ypos)
				continue
			}
			for y := 0; y < len(rock); y++ {
				ri := len(rock) - 1 - y
				for x, r := range rock[ri] {
					if r == '#' {
						if occ.Contains(grid.Pt(xpos+x, ypos+y)) {
							fatal("not that way")
						}
						occ.Insert(grid.Pt(xpos+x, ypos+y))
						if ypos+y+1 > maxY {
							maxY = ypos + y + 1
						}
					}
				}
			}
			dump(occ, width)
			break
		}
		res := result{
			flowIdx: flowIdx,
			rockIdx: rockIdx,
			xRest:   xpos,
		}
		if mark, ok := results[res]; ok {
			log("result: iter %d, maxy %d: %v - had also at %d, %d", i, maxY, res, mark.iter, mark.ypos)
			//
			dy := maxY - mark.ypos
			period := i - mark.iter
			front := i - period
			n := (cnt - front) / period
			left := (cnt - front) % period
			//left1 := cnt - (front + n*period)
			leftY, _ := part1MainFunc(in, left)
			//_ = left1

			toty := mark.ypos + n*dy + leftY
			maxY = toty - 1

			break
		}
		results[res] = mark{i, maxY}
	}

	return maxY, nil
}

const skipDump = true

func dump(occ *set.Set[grid.Point], width int) {
	if skipDump {
		return
	}
	log("")
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

func dumpWithRock(occ *set.Set[grid.Point], width int, rock Rock, xr, yr int) {
	if skipDump {
		return
	}
	log("")
	maxY := yr + rock.Height()
	rockContains := func(cx, cy int) bool {
		for x := 0; x < rock.Width(); x++ {
			for iy := 0; iy < rock.Height(); iy++ {
				y := rock.Height() - 1 - iy
				if rock[y][x] == '#' {
					ax := xr + x
					ay := yr + iy
					if ax == cx && ay == cy {
						return true
					}
				}
			}
		}
		return false
	}

	var sl []string
	for y := 0; y <= maxY; y++ {
		rsl := "|"
		for x := 0; x < width; x++ {
			if rockContains(x, y) {
				rsl += "@"
				continue
			}

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
