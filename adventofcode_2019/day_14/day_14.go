package day_14

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
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

func part1MainFunc(in string) (int, error) {
	ls := readutil.ReadLines(in)
	var reacts []reaction
	for _, l := range ls {
		react, err := parseReaction(l)
		if err != nil {
			return 0, err
		}
		reacts = append(reacts, react)
	}

	mustFindReactionFor := func(name string) reaction {
		for _, r := range reacts {
			if r.output.name == name {
				return r
			}
		}
		fatal("found no reaction for %q", name)
		return reaction{}
	}

	stock := map[string]int{}
	var makeChem func(name string, amount int) (oreCost int)
	makeChem = func(name string, amount int) (oreCost int) {
		log("make-chem %q: %d", name, amount)
		if name == "ORE" {
			return amount
		}
		// try to find in stock
		if stockAmount, ok := stock[name]; ok {
			take := mathutil.Min(amount, stockAmount)
			stock[name] -= take
			amount -= take
			if take > 0 {
				log("make-chem %q: take %d from stock", name, take)
			}
		}
		if amount == 0 {
			return 0
		}

		ra := mustFindReactionFor(name)
		var reactRuns int
		if amount%ra.output.amount == 0 {
			reactRuns = amount / ra.output.amount
		} else {
			reactRuns = amount/ra.output.amount + 1
		}
		log("make-chem %q: need to make %d: reaction produces %d - run %d times", name, amount, ra.output.amount, reactRuns)

		defer func() {
			leftQty := reactRuns*ra.output.amount - amount
			if leftQty > 0 {
				stock[name] += leftQty
				log("make-chem %q: put %d on stock", name, leftQty)
			}
		}()

		oreCost = 0
		for _, in := range ra.input {
			log("make-chem %q: produce %d of %q", name, in.amount, in.name)
			chemCost := makeChem(in.name, reactRuns*in.amount)
			oreCost += chemCost

		}
		log("make-chem %q: costs are %d", name, reactRuns*oreCost)
		return oreCost
	}

	oreCost := makeChem("FUEL", 1)
	return oreCost, nil
}

func part2MainFunc(in string) (int, error) {
	ls := readutil.ReadLines(in)
	var reacts []reaction
	for _, l := range ls {
		react, err := parseReaction(l)
		if err != nil {
			return 0, err
		}
		reacts = append(reacts, react)
	}

	mustFindReactionFor := func(name string) reaction {
		for _, r := range reacts {
			if r.output.name == name {
				return r
			}
		}
		fatal("found no reaction for %q", name)
		return reaction{}
	}

	stock := map[string]int{}
	var makeChem func(name string, amount int) (oreCost int)
	makeChem = func(name string, amount int) (oreCost int) {
		// log("make-chem %q: %d", name, amount)
		if name == "ORE" {
			return amount
		}
		// try to find in stock
		if stockAmount, ok := stock[name]; ok {
			take := mathutil.Min(amount, stockAmount)
			stock[name] -= take
			amount -= take
			if take > 0 {
				// log("make-chem %q: take %d from stock", name, take)
			}
		}
		if amount == 0 {
			return 0
		}

		ra := mustFindReactionFor(name)
		var reactRuns int
		if amount%ra.output.amount == 0 {
			reactRuns = amount / ra.output.amount
		} else {
			reactRuns = amount/ra.output.amount + 1
		}
		// log("make-chem %q: need to make %d: reaction produces %d - run %d times", name, amount, ra.output.amount, reactRuns)

		defer func() {
			leftQty := reactRuns*ra.output.amount - amount
			if leftQty > 0 {
				stock[name] += leftQty
				// log("make-chem %q: put %d on stock", name, leftQty)
			}
		}()

		oreCost = 0
		for _, in := range ra.input {
			// log("make-chem %q: produce %d of %q", name, in.amount, in.name)
			chemCost := makeChem(in.name, reactRuns*in.amount)
			oreCost += chemCost

		}
		// log("make-chem %q: costs are %d", name, reactRuns*oreCost)
		return oreCost
	}

	//	baseCost := makeChem("FUEL", 1)

	fuelCount := 0
	oreAvailable := 1000000000000
	makeChunk := 10
	for {
		if oreAvailable < 100000000 {
			makeChunk = 1
		}

		oreCost := makeChem("FUEL", makeChunk)
		if oreCost > oreAvailable {
			break
		}
		fuelCount += makeChunk
		oreAvailable -= oreCost
	}

	return fuelCount, nil
}

func parseChemicalAmount(s string) (chemicalAmount, error) {
	s = strings.TrimSpace(s)
	var ca chemicalAmount
	_, err := fmt.Sscanf(s, "%d %s", &ca.amount, &ca.name)
	if err != nil {
		return chemicalAmount{}, err
	}
	return ca, nil
}

func parseReaction(s string) (reaction, error) {
	left, right, ok := strings.Cut(s, "=>")
	if !ok {
		return reaction{}, errors.Errorf("invalid reaction format %q", s)
	}
	var react reaction
	insl := strings.Split(left, ",")
	for _, ins := range insl {
		ca, err := parseChemicalAmount(ins)
		if err != nil {
			return reaction{}, err
		}
		react.input = append(react.input, ca)
	}
	if len(react.input) == 0 {
		return reaction{}, errors.Errorf("inputs are empty")
	}
	caOut, err := parseChemicalAmount(right)
	if err != nil {
		return reaction{}, err
	}
	react.output = caOut
	return react, nil
}

type chemicalAmount struct {
	name   string
	amount int
}

type reaction struct {
	input  []chemicalAmount
	output chemicalAmount
}
