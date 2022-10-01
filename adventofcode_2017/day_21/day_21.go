package day_21

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := onAfter(input, 5)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := onAfter(input, 18)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

type sub struct {
	pixels [][]bool
}

func (s sub) hash() string {
	var sl []string
	for _, row := range s.pixels {
		var sr string
		for _, b := range row {
			if b {
				sr += "1"
			} else {
				sr += "0"
			}
		}
		sl = append(sl, sr)
	}
	return strings.Join(sl, ":")
}

func (s sub) size() int {
	return len(s.pixels)
}

func (s sub) clone() sub {
	var cs sub
	for _, row := range s.pixels {
		crow := make([]bool, len(row))
		copy(crow, row)
		cs.pixels = append(cs.pixels, crow)
	}
	return cs
}

func (s sub) sect(x0, y0 int, size int) sub {
	ss := sub{}
	for y := y0; y < y0+size; y++ {
		var sr []bool
		for x := x0; x < x0+size; x++ {
			sr = append(sr, s.pixels[y][x])
		}
		ss.pixels = append(ss.pixels, sr)
	}

	return ss
}

type rule struct {
	id  int
	in  sub
	out sub
}

// [../.# => ##./#../...]
func mustParseSub(in string) sub {
	var s sub
	var curr []bool
	for _, r := range in {
		switch r {
		case '.':
			curr = append(curr, false)
		case '#':
			curr = append(curr, true)
		case '/':
			s.pixels = append(s.pixels, curr)
			curr = []bool{}
		default:
			fatal("invalid sub char in %q", in)
		}
	}
	if len(curr) > 0 {
		s.pixels = append(s.pixels, curr)
	}

	sz := len(s.pixels)
	if sz == 0 {
		fatal("sub with no data %q", in)
	}

	for _, row := range s.pixels {
		if len(row) != sz {
			fatal("invalid sub sizes in %q", in)
		}
	}

	return s
}

func mustParseRule(s string) rule {
	sl := strings.Split(s, " => ")
	if len(sl) != 2 {
		fatal("invalid rule %q", s)
	}
	return rule{
		in:  mustParseSub(sl[0]),
		out: mustParseSub(sl[1]),
	}
}

func mustParseRules(in string) []rule {
	var rs []rule
	for i, line := range readutil.ReadLines(in) {
		r := mustParseRule(line)
		r.id = i
		rs = append(rs, r)
	}
	if len(rs) == 0 {
		fatal("no data")
	}
	return rs
}

func flipped(s sub) []sub {
	var fss []sub
	// at y
	fs := s.clone()
	for y := 0; y < s.size()/2; y++ {
		for x := 0; x < s.size(); x++ {
			fs.pixels[y][x], fs.pixels[s.size()-1-y][x] = fs.pixels[s.size()-1-y][x], fs.pixels[y][x]
		}
	}
	fss = append(fss, fs)

	// at x
	fs = s.clone()
	for x := 0; x < s.size()/2; x++ {
		for y := 0; y < s.size(); y++ {
			fs.pixels[y][x], fs.pixels[y][s.size()-1-x] = fs.pixels[y][s.size()-1-x], fs.pixels[y][x]
		}
	}
	fss = append(fss, fs)

	return fss
}

func rot90(s sub) sub {
	var c0 int
	if s.size() == 3 {
		c0 = 2
	} else if s.size() == 2 {
		c0 = 1
	} else {
		fatal("invalid size %d", s.size())
	}

	rs := s.clone()
	for y := 0; y < s.size(); y++ {
		for x := 0; x < s.size(); x++ {
			xr := -y + c0
			yr := x
			rs.pixels[yr][xr] = s.pixels[y][x]
		}
	}

	return rs
}

func rotated(s sub) []sub {
	var rss []sub

	rs := rot90(s)
	rss = append(rss, rs)

	rs = rot90(rs)
	rss = append(rss, rs)

	rs = rot90(rs)
	rss = append(rss, rs)

	return rss
}

func onAfter(in string, steps int) (int, error) {
	rs := mustParseRules(in)
	_ = rs

	hashMap := map[string]rule{}
	addToMap := func(r rule) {
		hash := r.in.hash()
		if er, ok := hashMap[hash]; ok {
			if er.id != r.id {
				fatal("same hash on different rules")
			}
		} else {
			hashMap[hash] = r
		}

	}

	for _, r := range rs {
		//log("variations for: %s", r.in.hash())
		addToMap(r)
		for _, rs := range rotated(r.in) {
			rr := rule{
				id:  r.id,
				in:  rs,
				out: r.out,
			}
			addToMap(rr)
			//log("%s", rr.in.hash())
		}

		for _, fs := range flipped(r.in) {
			fr := rule{
				id:  r.id,
				in:  fs,
				out: r.out,
			}
			addToMap(fr)
			//log("%s", fr.in.hash())
			for _, rs := range rotated(fs) {
				rr := rule{
					id:  r.id,
					in:  rs,
					out: r.out,
				}
				addToMap(rr)
				//log("%s", rr.in.hash())
			}
		}

	}

	grid := sub{
		pixels: [][]bool{
			{false, true, false},
			{false, false, true},
			{true, true, true},
		},
	}
	dumpGrid := func() {
		var sl []string
		for _, row := range grid.pixels {
			var sr string
			for _, b := range row {
				if b {
					sr += "#"
				} else {
					sr += "."
				}
			}
			sl = append(sl, sr)
		}
		ds := strings.Join(sl, "\n")
		log("*** grid (size = %d) ***\n%s", grid.size(), ds)
	}
	dumpGrid()

	for i := 0; i < steps; i++ {
		if grid.size()%2 == 0 {
			newGridSubs := [][]sub{}
			for y := 0; y < grid.size(); y += 2 {
				var subsRow []sub
				for x := 0; x < grid.size(); x += 2 {
					sub := grid.sect(x, y, 2)
					if r, ok := hashMap[sub.hash()]; ok {
						subsRow = append(subsRow, r.out.clone())
					} else {
						fatal("found no rule for %q", sub.hash())
					}
				}
				newGridSubs = append(newGridSubs, subsRow)
			}

			newGrid := sub{
				pixels: make([][]bool, len(newGridSubs)*3),
			}
			for ys, subsRow := range newGridSubs {
				for y := 0; y < 3; y++ {
					//new row
					row := make([]bool, len(newGridSubs)*3)
					for xs, sub := range subsRow {
						for x := 0; x < 3; x++ {
							row[xs*3+x] = sub.pixels[y][x]
						}
					}

					newGrid.pixels[ys*3+y] = row
				}
			}
			grid = newGrid
		} else if grid.size()%3 == 0 {
			newGridSubs := [][]sub{}
			for y := 0; y < grid.size(); y += 3 {
				var subsRow []sub
				for x := 0; x < grid.size(); x += 3 {
					sub := grid.sect(x, y, 3)
					if r, ok := hashMap[sub.hash()]; ok {
						subsRow = append(subsRow, r.out.clone())
					} else {
						fatal("found no rule for %q", sub.hash())
					}
				}
				newGridSubs = append(newGridSubs, subsRow)
			}

			newGrid := sub{
				pixels: make([][]bool, len(newGridSubs)*4),
			}
			for ys, subsRow := range newGridSubs {
				for y := 0; y < 4; y++ {
					//new row
					row := make([]bool, len(newGridSubs)*4)
					for xs, sub := range subsRow {
						for x := 0; x < 4; x++ {
							row[xs*4+x] = sub.pixels[y][x]
						}
					}

					newGrid.pixels[ys*4+y] = row
				}
			}
			grid = newGrid
		}
		dumpGrid()
	}

	var on int
	for _, row := range grid.pixels {
		for _, b := range row {
			if b {
				on++
			}
		}
	}

	return on, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
