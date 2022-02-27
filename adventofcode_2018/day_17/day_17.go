package day_17

import (
	"adventofcode_2018/errutil"
	"bytes"
	"fmt"

	"github.com/mazzegi/scan"
	"github.com/pkg/errors"
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

//
type scanDesc struct {
	C1     string
	C1Val  int
	C2     string
	C2From int
	C2To   int
}

func parseScanDescs(in string) ([]scanDesc, error) {
	return scan.Lines[scanDesc]("{{c1:string}}={{c1val:int}}, {{c2:string}}={{c2from:int}}..{{c2to:int}}", scan.BuiltinFuncs(), bytes.NewBufferString(in))
}

//
type point struct {
	x, y int
}

func (p point) down() point {
	return point{p.x, p.y + 1}
}

func (p point) left() point {
	return point{p.x - 1, p.y}
}

func (p point) right() point {
	return point{p.x + 1, p.y}
}

func (p point) downLeft() point {
	return point{p.x - 1, p.y + 1}
}

func (p point) downRight() point {
	return point{p.x + 1, p.y + 1}
}

const (
	Sand    rune = '.'
	Clay    rune = '#'
	Water   rune = '~'
	Flow    rune = '|'
	Nirwana rune = 'x'
)

type EnvRow struct {
	elts []rune
}

type Env struct {
	rows   []*EnvRow
	spring point
}

func (e *Env) Put(x, y int, elt rune) {
	if y < 0 || y >= len(e.rows) {
		panic(errors.Errorf("index out of bounds y = %d (x = %d)", y, x))
	}
	row := e.rows[y]
	if x < 0 || x >= len(row.elts) {
		panic(errors.Errorf("index out of bounds x = %d (y = %d)", x, y))
	}
	if elt == Water {
		eltDown := e.eltAt(point{x, y}.down())
		if !(eltDown == Clay || eltDown == Water) {
			panic("cannot put water onto non clay/water")
		}
	}

	row.elts[x] = elt
}

func (e *Env) dump() {
	log("")
	for _, row := range e.rows {
		var sr string
		for _, e := range row.elts {
			sr += string(e)
		}
		log(sr)
	}
	log("")
}

func constructEnv(scans []scanDesc) (*Env, error) {
	var xmin, xmax, ymin, ymax int
	for i, sd := range scans {
		if sd.C2From >= sd.C2To {
			return nil, errors.Errorf("from (%d) >= to (%d) in scan %d", sd.C2From, sd.C2To, i)
		}
		if sd.C1 == "x" {
			if i == 0 {
				xmin, xmax = sd.C1Val, sd.C1Val
				ymin, ymax = sd.C2From, sd.C2To
				continue
			}
			if sd.C1Val < xmin {
				xmin = sd.C1Val
			}
			if sd.C1Val > xmax {
				xmax = sd.C1Val
			}
			if sd.C2From < ymin {
				ymin = sd.C2From
			}
			if sd.C2From > ymax {
				ymax = sd.C2From
			}
		} else if sd.C1 == "y" {
			if i == 0 {
				ymin, ymax = sd.C1Val, sd.C1Val
				xmin, xmax = sd.C2From, sd.C2To
				continue
			}
			if sd.C1Val < ymin {
				ymin = sd.C1Val
			}
			if sd.C1Val > ymax {
				ymax = sd.C1Val
			}
			if sd.C2From < xmin {
				xmin = sd.C2From
			}
			if sd.C2From > xmax {
				xmax = sd.C2From
			}
		} else {
			return nil, errors.Errorf("invalid c1 %q in scan %d", sd.C1, i)
		}
	}
	ymin = 0
	xmin--
	xmax++
	log("env: xmin=%d, xmax=%d, ymin=%d, ymax=%d", xmin, xmax, ymin, ymax)

	env := &Env{}
	for y := ymin; y <= ymax; y++ {
		row := &EnvRow{}
		for x := xmin; x <= xmax; x++ {
			row.elts = append(row.elts, Sand)
		}
		env.rows = append(env.rows, row)
	}

	for _, sd := range scans {
		if sd.C1 == "x" {
			for y := sd.C2From; y <= sd.C2To; y++ {
				env.Put(sd.C1Val-xmin, y, Clay)
			}
		} else {
			for x := sd.C2From; x <= sd.C2To; x++ {
				env.Put(x-xmin, sd.C1Val, Clay)
			}
		}
	}
	env.spring = point{500 - xmin, 0}

	return env, nil
}

func (e *Env) eltAt(pos point) rune {
	if pos.y < 0 || pos.y >= len(e.rows) {
		return Nirwana
	}
	row := e.rows[pos.y]
	if pos.x < 0 || pos.x >= len(row.elts) {
		return Sand
	}
	return row.elts[pos.x]
}

func (e *Env) canFlowTo(pos point) (bool, rune) {
	elt := e.eltAt(pos)
	return elt == Sand, elt
}

func (e *Env) count(elt rune) int {
	var cnt int
	for _, row := range e.rows {
		for _, e := range row.elts {
			if e == elt {
				cnt++
			}
		}
	}
	return cnt
}

func (e *Env) drop() bool {
	// creates a new flow at the spring
	pos := e.spring
	if ok, _ := e.canFlowTo(pos.down()); !ok {
		return false
	}
	pos = pos.down()

	mustFlow := false
	moveLeft := func() (ok bool, done bool) {
		dl := pos.left()
		ok, blocked := e.canFlowTo(dl)
		if ok {
			pos = dl
			return true, false
		}
		if blocked == Flow {
			e.Put(pos.x, pos.y, Flow)
			return false, true
		}
		if blocked == Clay && mustFlow {
			e.Put(pos.x, pos.y, Flow)
			return false, true
		}
		return false, false
	}

	moveRight := func() (ok bool, done bool) {
		dr := pos.right()
		ok, blocked := e.canFlowTo(dr)
		if ok {
			pos = dr
			return true, false
		}
		if blocked == Flow {
			e.Put(pos.x, pos.y, Flow)
			return false, true
		}
		if blocked == Clay || blocked == Water {
			if mustFlow {
				e.Put(pos.x, pos.y, Flow)
			} else {
				e.Put(pos.x, pos.y, Water)
			}
			return false, true
		}
		return false, false
	}

	var isMovingLeft bool
	for {
		dp := pos.down()
		ok, blocked := e.canFlowTo(dp)
		if ok {
			mustFlow = false
			pos = dp
			isMovingLeft = true
			continue
		}
		//if blocked == Nirwana || blocked == Flow {
		if blocked == Nirwana {
			e.Put(pos.x, pos.y, Flow)
			return true
		}
		if blocked == Water {
			// see, if theres, space for this drop
			wp := pos.down()
		outerleft:
			for {
				wp = wp.left()
				switch e.eltAt(wp) {
				case Sand:
					e.Put(wp.x, wp.y, Water)
					return true
				case Clay:
					break outerleft
				}
			}
			wp = pos.down()
		outerright:
			for {
				wp = wp.right()
				switch e.eltAt(wp) {
				case Sand:
					e.Put(wp.x, wp.y, Water)
					return true
				case Clay:
					break outerright
				}
			}
		}

		if blocked == Flow {
			if e.eltAt(pos.down().down()) == Flow {
				e.Put(pos.x, pos.y, Flow)
				return true
			}

			// see, if theres space for this drop to flow
			dok := false
			dr := pos.downRight()
			underDR := e.eltAt(dr.down())
			if underDR == Nirwana || underDR == Sand {
				e.Put(pos.x, pos.y, Flow)
				return true
			}

			if e.eltAt(dr) == Sand && (underDR == Clay || underDR == Water) {
				pos = dr
				isMovingLeft = false
				dok = true
				mustFlow = true
			}
			if !dok {
				dl := pos.downLeft()
				underDL := e.eltAt(dl.down())
				if underDL == Nirwana || underDL == Sand {
					e.Put(pos.x, pos.y, Flow)
					return true
				}
				if e.eltAt(dl) == Sand && (underDL == Clay || underDL == Water) {
					pos = dl
					isMovingLeft = true
					dok = true
					mustFlow = true
				}
			}
			if !dok {
				e.Put(pos.x, pos.y, Flow)
				return true
			}
			//check for nirwana put flow and return
		}

		if isMovingLeft {
			ok, done := moveLeft()
			if ok {
				continue
			}
			if done {
				return true
			}
			ok, done = moveRight()
			if ok {
				isMovingLeft = false
				continue
			}
			if done {
				return true
			}
		} else {
			ok, done := moveRight()
			if ok {
				continue
			}
			if done {
				return true
			}
		}
	}
}

func part1MainFunc(in string) (int, error) {
	scans, err := parseScanDescs(in)
	if err != nil {
		return 0, err
	}
	log("parsed %d scans", len(scans))
	env, err := constructEnv(scans)
	if err != nil {
		return 0, err
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				log("recovered: %v", r)
			}
		}()
		for env.drop() {
			//env.dump()
		}
	}()

	env.dump()

	wCnt := env.count(Water)
	fCnt := env.count(Flow)

	return wCnt + fCnt, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}

/*
......+.......
......|.....#.
.#..#||||...#.
.#..#~~#|.....
.#..#~~#|.....
.#~~~~~#|.....
.#~~~~~#|.....
.#######|.....
........|.....
...|||||||||..
...|#~~~~~#|..
...|#~~~~~#|..
...|#~~~~~#|..
...|#######|..
*/
