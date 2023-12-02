package day_02

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
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

type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

func ValidateColor(cr Color) error {
	switch cr {
	case Red, Green, Blue:
		return nil
	default:
		return fmt.Errorf("invalid color %q", cr)
	}
}

func NewSubset() Subset {
	return Subset{Counts: make(map[Color]int)}
}

type Subset struct {
	Counts map[Color]int
}

func SubsetPower(subset Subset) int {
	rc := subset.Counts[Red]
	gc := subset.Counts[Green]
	bc := subset.Counts[Blue]
	return rc * gc * bc
}

type Game struct {
	ID      int
	Subsets []Subset
}

func GameMinSet(g *Game) Subset {
	mset := NewSubset()
	mset.Counts[Red] = 0
	mset.Counts[Green] = 0
	mset.Counts[Blue] = 0

	for _, subset := range g.Subsets {
		sr := subset.Counts[Red]
		sg := subset.Counts[Green]
		sb := subset.Counts[Blue]

		if sr > mset.Counts[Red] {
			mset.Counts[Red] = sr
		}
		if sg > mset.Counts[Green] {
			mset.Counts[Green] = sg
		}
		if sb > mset.Counts[Blue] {
			mset.Counts[Blue] = sb
		}
	}
	return mset
}

func MustParseGame(s string) *Game {
	bef, after, ok := strings.Cut(s, ":")
	if !ok {
		panic(fmt.Errorf("cut %q by ':'", s))
	}
	var gameID int
	_, err := fmt.Sscanf(bef, "Game %d", &gameID)
	errutil.FatalWhen(err)

	g := &Game{ID: gameID}
	setssl := strings.Split(after, ";")
	for _, sets := range setssl {
		subset := NewSubset()
		csl := strings.Split(sets, ",")
		for _, cs := range csl {
			// something like 2 green
			var cnt int
			var cr Color
			_, err := fmt.Sscanf(cs, "%d %s", &cnt, &cr)
			errutil.FatalWhen(err)
			err = ValidateColor(cr)
			errutil.FatalWhen(err)
			subset.Counts[cr] = cnt
		}
		g.Subsets = append(g.Subsets, subset)
	}
	return g
}

func MustParseGames(sl []string) []*Game {
	var gs []*Game
	for _, s := range sl {
		gs = append(gs, MustParseGame(s))
	}
	return gs
}

func isSubsetPossible(constraintMax map[Color]int, subset Subset) bool {
	for cr, cnt := range subset.Counts {
		var maxcnt int = 0
		if mc, ok := constraintMax[cr]; ok {
			maxcnt = mc
		}
		if cnt > maxcnt {
			return false
		}
	}
	return true
}

func isGamePossible(constraintMax map[Color]int, g *Game) bool {
	for _, subset := range g.Subsets {
		if !isSubsetPossible(constraintMax, subset) {
			return false
		}
	}
	return true
}

func part1MainFunc(in string) (int, error) {
	sl := readutil.ReadLines(in)
	gs := MustParseGames(sl)

	constraintMax := map[Color]int{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	var sumIDs int
	for _, g := range gs {
		if isGamePossible(constraintMax, g) {
			sumIDs += g.ID
		}
	}

	return sumIDs, nil
}

func part2MainFunc(in string) (int, error) {
	sl := readutil.ReadLines(in)
	gs := MustParseGames(sl)

	var sumPowers int
	for _, g := range gs {
		ms := GameMinSet(g)
		p := SubsetPower(ms)
		sumPowers += p
	}

	return sumPowers, nil
}
