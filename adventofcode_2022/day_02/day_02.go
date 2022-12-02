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

type Shape string

const (
	Rock     Shape = "rock"
	Paper    Shape = "paper"
	Scissors Shape = "scissors"
)

func (sh Shape) Score() int {
	switch sh {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		fatal("invalid shape %q", sh)
		return 0
	}
}

func (sh Shape) Defeats(osh Shape) bool {
	switch sh {
	case Rock:
		return osh == Scissors
	case Paper:
		return osh == Rock
	case Scissors:
		return osh == Paper
	default:
		fatal("invalid shape %q", sh)
		return false
	}
}

func (sh Shape) DefeatsAgainst() Shape {
	switch sh {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		fatal("invalid shape %q", sh)
		return ""
	}
}

func (sh Shape) LoosesAgainst() Shape {
	switch sh {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		fatal("invalid shape %q", sh)
		return ""
	}
}

func myRoundScore(opp Shape, me Shape) int {
	if opp == me {
		return 3
	}
	if opp.Defeats(me) {
		return 0
	}
	return 6
}

type shapeMapping map[string]Shape

func (sm shapeMapping) mustLookup(s string) Shape {
	sh, ok := sm[s]
	if !ok {
		fatal("lookup %q", s)
	}
	return sh
}

type shapePair struct {
	opp Shape
	me  Shape
}

func myScore(opp Shape, me Shape) int {
	return myRoundScore(opp, me) + me.Score()
}

func part1MainFunc(in string) (int, error) {
	mappingOpp := shapeMapping{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	mappingMe := shapeMapping{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
	var pairs []shapePair
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		sl := strings.Fields(line)
		if len(sl) != 2 {
			fatal("invalid line %q", line)
		}
		pairs = append(pairs, shapePair{
			opp: mappingOpp.mustLookup(strings.TrimSpace(sl[0])),
			me:  mappingMe.mustLookup(strings.TrimSpace(sl[1])),
		})
	}
	var score int
	for _, pair := range pairs {
		score += myScore(pair.opp, pair.me)
	}
	return score, nil
}

func part2MainFunc(in string) (int, error) {
	mappingOpp := shapeMapping{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	var pairs []shapePair
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		sl := strings.Fields(line)
		if len(sl) != 2 {
			fatal("invalid line %q", line)
		}
		opp := mappingOpp.mustLookup(strings.TrimSpace(sl[0]))
		outcome := strings.TrimSpace(sl[1])
		pair := shapePair{
			opp: opp,
		}
		switch outcome {
		case "X": // have to loose
			pair.me = opp.DefeatsAgainst()
		case "Y": // need draw
			pair.me = opp
		case "Z": //have to win
			pair.me = opp.LoosesAgainst()
		default:
			fatal("unknown outcome %q", outcome)
		}
		pairs = append(pairs, pair)
	}
	var score int
	for _, pair := range pairs {
		score += myScore(pair.opp, pair.me)
	}
	return score, nil
}
