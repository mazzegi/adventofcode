package day_04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
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

type Card struct {
	ID         int
	MatchCount int
}

func mustStringsToInts(sl []string) []int {
	var ns []int
	for _, s := range sl {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		errutil.FatalWhen(err)
		ns = append(ns, n)
	}
	return ns
}

func mustParseCard(s string) Card {
	scard, rest, ok := strings.Cut(s, ":")
	if !ok {
		panic("cannot cut: " + s)
	}
	var id int
	_, err := fmt.Sscanf(scard, "Card %d", &id)
	errutil.FatalWhen(err)

	swinning, smy, ok := strings.Cut(rest, "|")
	if !ok {
		panic("cannot cut: " + rest)
	}

	winningNumbers := mustStringsToInts(strings.Split(swinning, " "))
	myNumbers := mustStringsToInts(strings.Split(smy, " "))
	winningNumbersSet := set.New[int](winningNumbers...)

	card := Card{ID: id}
	for _, myn := range myNumbers {
		if winningNumbersSet.Contains(myn) {
			card.MatchCount++
		}
	}
	return card
}

func mustParseCards(sl []string) []Card {
	cs := make([]Card, len(sl))
	for i, s := range sl {
		cs[i] = mustParseCard(s)
	}
	return cs
}

func powN(v int, n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= v
	}
	return res
}

func part1MainFunc(in string) (int, error) {
	cards := mustParseCards(readutil.ReadLines(in))

	var sum int
	for _, card := range cards {
		if card.MatchCount > 0 {
			sum += powN(2, card.MatchCount-1)
		}
	}

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	cards := mustParseCards(readutil.ReadLines(in))
	var total int
	for _, card := range cards {
		total += collectSubsOf(card, cards) + 1
	}
	return total, nil
}

func collectSubsOf(card Card, cards []Card) int {
	var winCardsIdxs []int
	for idx := card.ID; idx < card.ID+card.MatchCount; idx++ {
		if idx < len(cards) {
			winCardsIdxs = append(winCardsIdxs, idx)
		}
	}
	cnt := 0
	for _, wi := range winCardsIdxs {
		wcard := cards[wi]
		cnt += collectSubsOf(wcard, cards) + 1
	}

	return cnt
}
