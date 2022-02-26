package day_13

import (
	"adventofcode_2018/errutil"
	"adventofcode_2018/readutil"
	"fmt"
	"sort"
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

const (
	turnLeft     = "left"
	turnStraight = "straight"
	turnRight    = "right"
)

const (
	faceLeft  rune = '<'
	faceRight rune = '>'
	faceUp    rune = '^'
	faceDown  rune = 'v'
)

func nextTurn(turn string) string {
	switch turn {
	case turnLeft:
		return turnStraight
	case turnStraight:
		return turnRight
	case turnRight:
		return turnLeft
	default:
		return turnLeft
	}
}

type pos struct {
	x, y int
}

type car struct {
	pos      pos
	lastTurn string
	face     rune
	crashed  bool
}

type cars []*car

func (cs cars) uncrashed() cars {
	var ucs cars
	for _, c := range cs {
		if !c.crashed {
			ucs = append(ucs, c)
		}
	}
	return ucs
}

func (cs cars) sort() {
	sort.Slice(cs, func(i, j int) bool {
		c1, c2 := cs[i], cs[j]
		if c1.pos.y < c2.pos.y {
			return true
		} else if c1.pos.y > c2.pos.y {
			return false
		}
		return c1.pos.x < c2.pos.x
	})
}

func (cs cars) collision() (pos, bool) {
	for i1, c1 := range cs {
		for i2, c2 := range cs {
			if i2 <= i1 {
				continue
			}
			if c1.crashed || c2.crashed {
				continue
			}
			if c1.pos == c2.pos {
				c1.crashed = true
				c2.crashed = true
				return c1.pos, true
			}
		}
	}
	return pos{}, false
}

//
type row []rune

type track struct {
	rows []row
	cars cars
}

func parseTrack(in string) *track {
	t := &track{}
	lines := readutil.ReadLinesUntrimmed(in)
	for iy, line := range lines {
		row := row(line)
		for ix, r := range row {
			switch r {
			case faceLeft:
				t.cars = append(t.cars, &car{
					pos:  pos{ix, iy},
					face: faceLeft,
				})
				row[ix] = '-'
			case faceRight:
				t.cars = append(t.cars, &car{
					pos:  pos{ix, iy},
					face: faceRight,
				})
				row[ix] = '-'
			case faceUp:
				t.cars = append(t.cars, &car{
					pos:  pos{ix, iy},
					face: faceUp,
				})
				row[ix] = '|'
			case faceDown:
				t.cars = append(t.cars, &car{
					pos:  pos{ix, iy},
					face: faceDown,
				})
				row[ix] = '|'
			}
		}
		t.rows = append(t.rows, row)
	}
	return t
}

func (t *track) move(c *car) {
	switch c.face {
	case faceLeft:
		c.pos.x--
	case faceRight:
		c.pos.x++
	case faceUp:
		c.pos.y--
	case faceDown:
		c.pos.y++
	}

	if c.pos.y < 0 || c.pos.y >= len(t.rows) {
		log("invalid position %d,%d", c.pos.x, c.pos.y)
		return
	}
	row := t.rows[c.pos.y]
	if c.pos.x < 0 || c.pos.x >= len(row) {
		log("invalid position %d,%d", c.pos.x, c.pos.y)
		return
	}

	switch row[c.pos.x] {
	case '+':
		turn := nextTurn(c.lastTurn)
		c.lastTurn = turn
		switch {
		case c.face == faceLeft && turn == turnLeft:
			c.face = faceDown
		case c.face == faceLeft && turn == turnRight:
			c.face = faceUp
		case c.face == faceRight && turn == turnLeft:
			c.face = faceUp
		case c.face == faceRight && turn == turnRight:
			c.face = faceDown
		case c.face == faceUp && turn == turnLeft:
			c.face = faceLeft
		case c.face == faceUp && turn == turnRight:
			c.face = faceRight
		case c.face == faceDown && turn == turnLeft:
			c.face = faceRight
		case c.face == faceDown && turn == turnRight:
			c.face = faceLeft
		}
	case '/':
		switch c.face {
		case faceLeft:
			c.face = faceDown
		case faceRight:
			c.face = faceUp
		case faceUp:
			c.face = faceRight
		case faceDown:
			c.face = faceLeft
		}

	case '\\':
		switch c.face {
		case faceLeft:
			c.face = faceUp
		case faceRight:
			c.face = faceDown
		case faceUp:
			c.face = faceLeft
		case faceDown:
			c.face = faceRight
		}
	}
}

func part1MainFunc(in string) (int, error) {
	t := parseTrack(in)
	for {
		t.cars.sort()
		for _, c := range t.cars {
			t.move(c)
			if pos, ok := t.cars.collision(); ok {
				log("collision at %d,%d", pos.x, pos.y)
				return 0, nil
			}
		}
	}
}

func part2MainFunc(in string) (int, error) {
	t := parseTrack(in)
	for {
		t.cars.sort()
		for _, c := range t.cars {
			t.move(c)
			t.cars.collision()
		}
		t.cars = t.cars.uncrashed()
		if len(t.cars) > 1 {
			continue
		}
		if len(t.cars) == 1 {
			log("left: %d,%d", t.cars[0].pos.x, t.cars[0].pos.y)
		} else {
			log("no car left")
		}
		return 0, nil
	}
}
