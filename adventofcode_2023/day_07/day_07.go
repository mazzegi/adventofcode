package day_07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"golang.org/x/exp/maps"
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

const HandChars = "AKQJT98765432"

func CharRank(r rune) int {
	return strings.IndexRune(HandChars, r)
}

const HandCharsPart2 = "AKQT98765432J"

func CharRankPart2(r rune) int {
	return strings.IndexRune(HandCharsPart2, r)
}

type Hand string

func IsValidHand(s string) bool {
	if len(s) != 5 {
		return false
	}
	for _, r := range s {
		if !strings.ContainsRune(HandChars, r) {
			return false
		}
	}
	return true
}

type HandBid struct {
	Hand Hand
	Bid  int
}

func mustParseHandBit(s string) HandBid {
	h, b, ok := strings.Cut(s, " ")
	if !ok {
		errutil.FatalWhen(fmt.Errorf("cannot parse handbit %q", s))
	}
	if !IsValidHand(h) {
		errutil.FatalWhen(fmt.Errorf("invalid hand %q", s))
	}
	nb, err := strconv.ParseInt(b, 10, 64)
	errutil.FatalWhen(err)
	return HandBid{
		Hand: Hand(h),
		Bid:  int(nb),
	}
}

func mustParseHandBits(sl []string) []HandBid {
	var hbs []HandBid
	for _, s := range sl {
		hbs = append(hbs, mustParseHandBit(s))
	}
	return hbs
}

func ClassifyHand(h Hand) HandType {
	hg := map[rune]int{}
	for _, r := range h {
		hg[r]++
	}
	hgvals := maps.Values(hg)
	slices.Sort(hgvals)
	slices.Reverse(hgvals)
	switch {
	case len(hgvals) == 1 && hgvals[0] == 5:
		return FiveOfAKind
	case len(hgvals) == 2 && hgvals[0] == 4:
		return FourOfAKind
	case len(hgvals) == 2 && hgvals[0] == 3 && hgvals[1] == 2:
		return FullHouse
	case len(hgvals) == 3 && hgvals[0] == 3:
		return ThreeOfAKind
	case len(hgvals) == 3 && hgvals[0] == 2 && hgvals[1] == 2:
		return TwoPair
	case len(hgvals) == 4 && hgvals[0] == 2:
		return OnePair
	default:
		return HighCard
	}
}

type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func CompareHands(h1, h2 Hand) int {
	ty1 := ClassifyHand(h1)
	ty2 := ClassifyHand(h2)
	if ty1 < ty2 {
		return -1
	} else if ty1 > ty2 {
		return 1
	}
	//comp values
	for i := 0; i < 5; i++ {
		cr1 := CharRank(rune(h1[i]))
		cr2 := CharRank(rune(h2[i]))

		if cr1 > cr2 {
			return -1
		} else if cr1 < cr2 {
			return 1
		}
	}
	return 0
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	hbs := mustParseHandBits(lines)

	slices.SortFunc(hbs, func(a, b HandBid) int {
		return CompareHands(a.Hand, b.Hand)
	})

	wsum := 0
	for i, hb := range hbs {
		wsum += (i + 1) * hb.Bid
	}

	return wsum, nil
}

func ClassifyHandPart2(h Hand) HandType {
	var jokerCount int
	hg := map[rune]int{}
	for _, r := range h {
		if r == 'J' {
			jokerCount++
			continue
		}
		hg[r]++
	}

	hgvals := maps.Values(hg)
	slices.Sort(hgvals)
	slices.Reverse(hgvals)

	if jokerCount == 5 {
		return FiveOfAKind
	}

	// now add jokers to those with the greatest amount
	hgvals[0] += jokerCount

	switch {
	case len(hgvals) == 1 && hgvals[0] == 5:
		return FiveOfAKind
	case len(hgvals) == 2 && hgvals[0] == 4:
		return FourOfAKind
	case len(hgvals) == 2 && hgvals[0] == 3 && hgvals[1] == 2:
		return FullHouse
	case len(hgvals) == 3 && hgvals[0] == 3:
		return ThreeOfAKind
	case len(hgvals) == 3 && hgvals[0] == 2 && hgvals[1] == 2:
		return TwoPair
	case len(hgvals) == 4 && hgvals[0] == 2:
		return OnePair
	default:
		return HighCard
	}
}

func CompareHandsPart2(h1, h2 Hand) int {
	ty1 := ClassifyHandPart2(h1)
	ty2 := ClassifyHandPart2(h2)
	if ty1 < ty2 {
		return -1
	} else if ty1 > ty2 {
		return 1
	}
	//comp values
	for i := 0; i < 5; i++ {
		cr1 := CharRankPart2(rune(h1[i]))
		cr2 := CharRankPart2(rune(h2[i]))

		if cr1 > cr2 {
			return -1
		} else if cr1 < cr2 {
			return 1
		}
	}
	return 0
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	hbs := mustParseHandBits(lines)

	slices.SortFunc(hbs, func(a, b HandBid) int {
		return CompareHandsPart2(a.Hand, b.Hand)
	})

	wsum := 0
	for i, hb := range hbs {
		wsum += (i + 1) * hb.Bid
	}

	return wsum, nil
}
