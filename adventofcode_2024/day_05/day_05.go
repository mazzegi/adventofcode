package day_05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
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

type rule struct {
	low, high string
}

type update []string

func parseInput(in string) ([]rule, []update, error) {
	lines := readutil.ReadLines(in)
	if len(lines) < 10 {
		return nil, nil, fmt.Errorf("to less lines")
	}
	var rules []rule
	var updates []update
	for _, line := range lines {
		if strings.Contains(line, "|") {
			// a rule
			slow, shigh, ok := strings.Cut(line, "|")
			if !ok {
				return nil, nil, fmt.Errorf("invalid rule %q", line)
			}
			rules = append(rules, rule{
				low:  slow,
				high: shigh,
			})
		} else {
			// an update
			update := strings.Split(line, ",")
			updates = append(updates, update)
		}
	}

	return rules, updates, nil
}

type separation struct {
	center   string
	lower    []string
	upper    []string
	lowerSep *separation
	upperSep *separation
}

func (sep *separation) flatten() []string {
	var fsl []string

	if sep.lowerSep != nil {
		fsl = append(fsl, sep.lowerSep.flatten()...)
	} else if len(sep.lower) == 1 {
		fsl = append(fsl, sep.lower[0])
	}

	fsl = append(fsl, sep.center)
	if sep.upperSep != nil {
		fsl = append(fsl, sep.upperSep.flatten()...)
	} else if len(sep.upper) == 1 {
		fsl = append(fsl, sep.upper[0])
	}

	return fsl
}

func (sep *separation) isCenterOrInLower(num string) bool {
	if num == sep.center {
		return true
	}
	return slices.Contains(sep.lower, num)
}

func (sep *separation) isCenterOrInUpper(num string) bool {
	if num == sep.center {
		return true
	}
	return slices.Contains(sep.upper, num)
}

func separate(center string, numbers []string, rules []rule) (*separation, error) {

	sep := &separation{
		center: center,
	}

outer_loop:
	for {
		var insertedNumbers []string
	loop_numbers:
		for _, num := range numbers {
			for _, rule := range rules {
				if rule.low == num {
					if sep.isCenterOrInLower(rule.high) {
						sep.lower = append(sep.lower, num)
						insertedNumbers = append(insertedNumbers, num)
						continue loop_numbers
					}
				} else if rule.high == num {
					if sep.isCenterOrInUpper(rule.low) {
						sep.upper = append(sep.upper, num)
						insertedNumbers = append(insertedNumbers, num)
						continue loop_numbers
					}
				}
			}
		}
		//remove inserted numbers
		numbers = slices.DeleteFunc(numbers, func(num string) bool {
			return slices.Contains(insertedNumbers, num)
		})
		if len(numbers) == 0 {
			//done
			break outer_loop
		}
		// not done but also no action - sth wrong
		if len(insertedNumbers) == 0 {
			return nil, fmt.Errorf("no actions during last run")
		}
	}
	//
	var err error
	if len(sep.lower) > 1 {
		lowerCenter := sep.lower[0]
		lowerNumbers := slices.Clone(sep.lower[1:])
		sep.lowerSep, err = separate(lowerCenter, lowerNumbers, rules)
		if err != nil {
			return nil, fmt.Errorf("separate")
		}
	}
	if len(sep.upper) > 1 {
		upperCenter := sep.upper[0]
		upperNumbers := slices.Clone(sep.upper[1:])
		sep.upperSep, err = separate(upperCenter, upperNumbers, rules)
		if err != nil {
			return nil, fmt.Errorf("separate")
		}
	}

	return sep, nil
}

func part1MainFunc(in string) (int, error) {
	rules, updates, err := parseInput(in)
	if err != nil {
		return 0, fmt.Errorf("parse-input: %w", err)
	}
	// extract all numbers
	var ruleNumbers []string
	for _, rule := range rules {
		if !slices.Contains(ruleNumbers, rule.low) {
			ruleNumbers = append(ruleNumbers, rule.low)
		}
		if !slices.Contains(ruleNumbers, rule.high) {
			ruleNumbers = append(ruleNumbers, rule.high)
		}
	}
	//
	center := ruleNumbers[0]
	numbers := slices.Clone(ruleNumbers[1:])
	sep, err := separate(center, numbers, rules)
	if err != nil {
		return 0, fmt.Errorf("separate: %w", err)
	}

	orderedNumbers := sep.flatten()
	log("ordered: %v", orderedNumbers)
	findOrdinal := func(num string) (int, bool) {
		for i, on := range orderedNumbers {
			if on == num {
				return i, true
			}
		}
		return 0, false
	}

	// rule [74,64] gives invalid ordinals [19,10]
	//-> 64|55
	// -> 74|55
	// -> 74|12
	// -> 74|51
	// -> 74|57
	// -> 74|34
	// -> 74|48
	// -> 74|76
	// -> 74|28
	// -> 74|53 warum wird 74 dann in upper einsortiert? (sep. mit center 55)

	// 83|74 - 83 is in upper
	// 39|83 => 39???
	// wg 55|39

	// check rules against ordinals
	for _, rule := range rules {
		olow, ok := findOrdinal(rule.low)
		if !ok {
			return 0, fmt.Errorf("didnt find low value %q in ordinals", rule.low)
		}
		ohigh, ok := findOrdinal(rule.high)
		if !ok {
			return 0, fmt.Errorf("didnt find high value %q in ordinals", rule.high)
		}
		if !(olow < ohigh) {
			return 0, fmt.Errorf("rule [%s,%s] gives invalid ordinals [%d,%d]", rule.low, rule.high, olow, ohigh)
		}
	}

	isInOrder := func(u update) (bool, error) {
		for i := 1; i < len(u); i++ {
			nlow, ok := findOrdinal(u[i-1])
			if !ok {
				return false, fmt.Errorf("ordinal of %q not found", u[i-1])
			}
			nhigh, ok := findOrdinal(u[i])
			if !ok {
				return false, fmt.Errorf("ordinal of %q not found", u[i])
			}
			if nlow > nhigh {
				return false, nil
			}
		}
		return true, nil
	}

	midNum := func(u update) int {
		i := len(u) / 2
		num := u[i]
		n, _ := strconv.Atoi(num)
		return n
	}

	var sum int
	for _, u := range updates {
		ino, err := isInOrder(u)
		if err != nil {
			return 0, fmt.Errorf("isinorder: %w", err)
		}
		if ino {
			sum += midNum(u)
			log("pass: %v", u)
		} else {
			log("NOT pass: %v", u)
		}
	}

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
