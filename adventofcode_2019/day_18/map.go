package day_18

import (
	"fmt"
	"slices"
	"unicode"

	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/maps"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/pkg/errors"
)

// type Key struct {
// 	value rune
// 	pos   grid.Point
// }

// type Door struct {
// 	value rune
// 	pos   grid.Point
// }

func NewState(m *Map) *State {
	return &State{
		m: m,
		//collectedKeys: make(map[rune]struct{}),
		currPos:       m.entry,
		collectedKeys: []rune{},
	}
}

func cloneState(state *State) *State {
	return &State{
		m:       state.m,
		currPos: state.currPos,
		//collectedKeys: maps.Clone(state.collectedKeys),
		collectedKeys: slices.Clone(state.collectedKeys),
	}
}

type State struct {
	currPos grid.Point
	//collectedKeys map[rune]struct{}
	collectedKeys []rune
	m             *Map
	done          bool
}

func hashState(state *State) string {
	return fmt.Sprintf("%s_%s", state.currPos.String(), string(state.collectedKeys))
}

type Map struct {
	grid        [][]bool // true = wall; false = empty
	dimX, dimY  int
	keys        map[grid.Point]rune
	doors       map[grid.Point]rune
	entry       grid.Point
	allKeysHash string
}

func canAccess(state *State, pt grid.Point) bool {
	if pt.Y < 0 || pt.Y >= state.m.dimY || pt.X < 0 || pt.X >= state.m.dimX {
		return false
	}
	if state.m.grid[pt.Y][pt.X] { // wall!
		return false
	}
	if door, ok := state.m.doors[pt]; ok {
		// a door, see if we have the key
		relKey := unicode.ToLower(door)
		haveKey := slices.Contains(state.collectedKeys, relKey)
		return haveKey
	}
	return true
}

func accessibleNeighbours(state *State) []grid.Point {
	var nps []grid.Point
	for x := state.currPos.X - 1; x <= state.currPos.X+1; x++ {
		for y := state.currPos.Y - 1; y <= state.currPos.Y+1; y++ {
			pt := grid.Pt(x, y)
			if pt == state.currPos {
				continue
			}
			if canAccess(state, pt) {
				nps = append(nps, pt)
			}
		}
	}
	return nps
}

func moveTo(state *State, pt grid.Point) {
	state.currPos = pt
	// is there a key?
	if key, ok := state.m.keys[pt]; ok {
		if !slices.Contains(state.collectedKeys, key) {
			state.collectedKeys = append(state.collectedKeys, key)
			slices.Sort(state.collectedKeys)
		}

		if string(state.collectedKeys) == state.m.allKeysHash {
			state.done = true
		}
	}
}

//

func ParseMap(s string) (*Map, error) {
	m := &Map{
		keys:  make(map[grid.Point]rune),
		doors: make(map[grid.Point]rune),
	}
	allKeysM := map[rune]struct{}{}
	ls := readutil.ReadLines(s)
	for _, l := range ls {
		row := []bool{}
		y := len(m.grid)
		for x, r := range l {
			isWall := false
			switch {
			case r == '#':
				isWall = true
			case r >= 'a' && r <= 'z':
				if _, ok := allKeysM[r]; ok {
					return nil, fmt.Errorf("duplicate key %s", string(r))
				}
				allKeysM[r] = struct{}{}
				m.keys[grid.Pt(x, y)] = r
			case r >= 'A' && r <= 'Z':
				m.doors[grid.Pt(x, y)] = r
			case r == '@':
				m.entry = grid.Pt(x, y)
			default:
			}
			row = append(row, isWall)
		}
		if len(row) == 0 {
			continue
		}
		if len(m.grid) > 0 && len(row) != len(m.grid[0]) {
			return nil, errors.Errorf("invalid row size")
		}
		m.grid = append(m.grid, row)
	}
	if len(m.grid) == 0 {
		return nil, errors.Errorf("grid is empty")
	}
	m.dimY = len(m.grid)
	m.dimX = len(m.grid[0])

	allKeys := maps.Keys(allKeysM)
	slices.Sort(allKeys)
	m.allKeysHash = string(allKeys)
	return m, nil
}
