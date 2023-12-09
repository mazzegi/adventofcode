package day_22

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/slices"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input, 10007)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(input, 10007)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

type shuffleFunc func([]int) []int

func identity() shuffleFunc {
	return func(ns []int) []int {
		return slices.Clone(ns)
	}
}

func dealIntoNewStack() shuffleFunc {
	return func(ns []int) []int {
		return slices.Reverse(ns)
	}
}

func cut(n int) shuffleFunc {
	switch {
	case n > 0:
		return func(ns []int) []int {
			if n >= len(ns) {
				return slices.Clone(ns)
			}
			cns := slices.Clone(ns[:n])
			lns := slices.Clone(ns[n:])
			return append(lns, cns...)
		}
	case n < 0:
		return func(ns []int) []int {
			an := -n
			if an >= len(ns) {
				return slices.Clone(ns)
			}
			cns := slices.Clone(ns[len(ns)-an:])
			lns := slices.Clone(ns[:len(ns)-an])
			return append(cns, lns...)
		}
	default:
		return identity()
	}
}

func dealWithInc(n int) shuffleFunc {
	return func(ns []int) []int {
		dns := make([]int, len(ns))
		destIdx := 0
		for si := 0; si < len(ns); si++ {
			v := ns[si]
			dns[destIdx] = v
			destIdx += n
			if destIdx >= len(dns) {
				destIdx = destIdx - len(dns)
			}
		}
		return dns
	}
}

func shuffle(in string, numCards int) []int {
	ls := readutil.ReadLines(in)
	var sfuncs []shuffleFunc
	for _, l := range ls {
		sfuncs = append(sfuncs, mustParseShuffleFunc(l))
	}

	deck := make([]int, numCards)
	for i := 0; i < numCards; i++ {
		deck[i] = i
	}
	for _, sf := range sfuncs {
		deck = sf(deck)
	}

	return deck
}

func mustParseShuffleFunc(s string) shuffleFunc {
	switch {
	case s == "deal into new stack":
		return dealIntoNewStack()
	case strings.HasPrefix(s, "deal with increment "):
		inc, err := strconv.Atoi(strings.TrimPrefix(s, "deal with increment "))
		errutil.FatalWhen(err)
		return dealWithInc(inc)
	case strings.HasPrefix(s, "cut "):
		n, err := strconv.Atoi(strings.TrimPrefix(s, "cut "))
		errutil.FatalWhen(err)
		return cut(n)
	default:
		panic("oolap!")
	}
}

func part1MainFunc(in string, numCards int) (int, error) {
	deck := shuffle(in, numCards)
	return slices.Find(deck, 2019), nil
}

func part2MainFunc(in string, numCards int) (int, error) {
	return 0, nil
}
