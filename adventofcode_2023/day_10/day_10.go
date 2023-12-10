package day_10

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

type Direction int

const (
	Na Direction = iota
	North
	East
	South
	West
)

func allDirs() []Direction {
	return []Direction{
		North,
		East,
		South,
		West,
	}
}

func (d Direction) opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic("What !!! ???")
	}
}

type Connector struct {
	Dir1 Direction
	Dir2 Direction
}

func (con Connector) Equals(ocon Connector) bool {
	return (con.Dir1 == ocon.Dir1 && con.Dir2 == ocon.Dir2) ||
		(con.Dir1 == ocon.Dir2 && con.Dir2 == ocon.Dir1)
}

func (con Connector) Symbol() rune {
	switch {
	case Connector{North, South}.Equals(con):
		return '|'
	case Connector{East, West}.Equals(con):
		return '-'
	case Connector{North, East}.Equals(con):
		return 'L'
	case Connector{North, West}.Equals(con):
		return 'J'
	case Connector{South, West}.Equals(con):
		return '7'
	case Connector{South, East}.Equals(con):
		return 'F'
	default:
		panic("wohoa!")
	}
}

func (c Connector) connects(d Direction) (toDirection Direction, connects bool) {
	if c.Dir1 == d {
		return c.Dir2, true
	} else if c.Dir2 == d {
		return c.Dir1, true
	}
	return Na, false
}

const connectorChars = "|-LJ7F"

func isConnector(r rune) bool {
	return strings.ContainsRune(connectorChars, r)
}

func mustParseConnector(r rune) Connector {
	switch r {
	case '|':
		return Connector{North, South}
	case '-':
		return Connector{East, West}
	case 'L':
		return Connector{North, East}
	case 'J':
		return Connector{North, West}
	case '7':
		return Connector{South, West}
	case 'F':
		return Connector{South, East}
	default:
		panic(fmt.Errorf("invalid connector"))
	}
}

// func connects(con Connector, at grid.Point, to grid.Point) bool {
// 	sub := at.Sub(to)
// }

func inDir(p grid.Point, dir Direction) grid.Point {
	switch dir {
	case North:
		return p.Add(grid.Pt(0, -1))
	case South:
		return p.Add(grid.Pt(0, 1))
	case East:
		return p.Add(grid.Pt(1, 0))
	case West:
		return p.Add(grid.Pt(-1, 0))
	default:
		panic("what!!??")
	}
}

func determineLoop(rows []string) (start grid.Point, loop []grid.Point, connS Connector) {
	inRange := func(p grid.Point) bool {
		return p.Y >= 0 && p.Y < len(rows) &&
			p.X >= 0 && p.X < len(rows[p.Y])

	}

	at := func(p grid.Point) rune {
		return rune(rows[p.Y][p.X])
	}

	//find the "S"
	spt := grid.Pt(-1, -1)
	foundS := false
findS:
	for ir, row := range rows {
		for ic, r := range row {
			if r == 'S' {
				spt = grid.Pt(ic, ir)
				foundS = true
				break findS
			}
		}
	}
	if !foundS {
		panic("where's S????")
	}

	// query first neighbour, which connect to s
	currPt := spt
	var currDir Direction
	found := false
	for _, d := range allDirs() {
		dp := inDir(spt, d)
		if !inRange(dp) {
			continue
		}
		r := at(dp)
		if !isConnector(r) {
			continue
		}
		con := mustParseConnector(r)
		if dirTo, ok := con.connects(d.opposite()); ok {
			if !found {
				connS.Dir1 = d
				currDir = dirTo
				currPt = dp
				found = true
			} else {
				// must be the second one
				connS.Dir2 = d
			}
		}
	}

	loop = []grid.Point{spt, currPt}
	for {
		dp := inDir(currPt, currDir)
		if !inRange(dp) {
			panic("scroum!")
		}
		if dp == spt {
			break
		}

		r := at(dp)
		if !isConnector(r) {
			panic("qwzuirew")
		}
		con := mustParseConnector(r)
		if dirTo, ok := con.connects(currDir.opposite()); ok {
			currDir = dirTo
			currPt = dp
		} else {
			panic("fhdsjkfhks")
		}
		loop = append(loop, currPt)
	}

	return spt, loop, connS
}

func part1MainFunc(in string) (int, error) {
	rows := readutil.ReadLines(in)
	_, loop, connS := determineLoop(rows)
	_ = connS
	return len(loop) / 2, nil
}

func part2MainFunc(in string) (int, error) {
	rows := readutil.ReadLines(in)
	start, loop, connS := determineLoop(rows)
	loopSet := set.New(loop...)

	// set everything not on the loop to "."
	for y, row := range rows {
		rs := []rune(row)
		for x := range row {
			p := grid.Pt(x, y)
			if !loopSet.Contains(p) {
				rs[x] = '.'
			} else if p == start {
				rs[x] = connS.Symbol()
			}
		}
		rows[y] = string(rs)
	}

	//nextWouldBeInside := false

	// prevOnLoop := 0
	// prevInside := false
	// inside := false
	var insideLoopCount int
	for _, row := range rows {
		cnt := insideCount(row)
		insideLoopCount += cnt

		// prevOnLoop = 0
		// prevInside = false
		// inside = false
		// for x := range row {
		// 	p := grid.Pt(x, y)
		// 	switch {
		// 	case onLoop(p):
		// 		prevOnLoop++
		// 		prevInside = false
		// 	default:
		// 		if (!inside && prevOnLoop == 1) || prevInside {
		// 			insideLoopCount++
		// 			prevInside = true
		// 			inside = true
		// 		} else {
		// 			prevInside = false
		// 			inside = false
		// 		}
		// 		prevOnLoop = 0
		// 	}
		// }
	}

	return insideLoopCount, nil
}

func insideCount(row string) int {
	var cnt int

	inside := false
	onEdge := false
	var edgeEntered rune

	assertNotOnEdge := func() {
		if onEdge {
			panic("on-edge")
		}
	}
	assertOnEdge := func() {
		if !onEdge {
			panic("not-on-edge")
		}
	}

	for _, r := range row {
		switch r {
		case '|':
			assertNotOnEdge()
			if !inside {
				inside = true
			} else {
				inside = false
			}
		case '-':
			assertOnEdge()
		case 'F':
			assertNotOnEdge()
			if !inside {
				onEdge = true
				edgeEntered = 'F'
			} else {
				onEdge = true
				edgeEntered = 'F'
			}
		case 'L':
			assertNotOnEdge()
			if !inside {
				onEdge = true
				edgeEntered = 'L'
			} else {
				onEdge = true
				edgeEntered = 'L'
			}
		case 'J':
			assertOnEdge()
			if !inside {
				if edgeEntered == 'F' {
					inside = true
					onEdge = false
					edgeEntered = rune(0)
				} else if edgeEntered == 'L' {
					onEdge = false
					edgeEntered = rune(0)
				} else {
					panic("dont want to be here")
				}
			} else {
				if edgeEntered == 'F' {
					inside = false
					onEdge = false
					edgeEntered = rune(0)
				} else if edgeEntered == 'L' {
					onEdge = false
					edgeEntered = rune(0)
				} else {
					panic("dont want to be here")
				}
			}
		case '7':
			assertOnEdge()
			if !inside {
				if edgeEntered == 'F' {
					inside = false
					onEdge = false
					edgeEntered = rune(0)
				} else if edgeEntered == 'L' {
					inside = true
					onEdge = false
					edgeEntered = rune(0)
				} else {
					panic("dont want to be here")
				}
			} else {
				if edgeEntered == 'L' {
					inside = false
					onEdge = false
					edgeEntered = rune(0)
				} else if edgeEntered == 'F' {
					onEdge = false
					edgeEntered = rune(0)
				} else {
					panic("dont want to be here")
				}
			}
		default:
			if inside {
				cnt++
			}
		}

	}

	return cnt
}
