package day_05

import (
	"fmt"
	"slices"
	"sort"
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

func sortNumbersWithRules(numbers []string, rules []rule) ([]string, error) {
	if len(numbers) < 2 {
		return numbers, nil
	}
	center := numbers[0]
	rest := slices.Clone(numbers[1:])
	sep, err := separate(center, rest, rules)
	if err != nil {
		return nil, fmt.Errorf("separate: %w", err)
	}
	return sep.flatten(), nil
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

	findOrdinal := func(num string, sorted []string) (int, bool) {
		for i, on := range sorted {
			if on == num {
				return i, true
			}
		}
		return 0, false
	}

	isInOrder := func(u update) (bool, error) {
		sorted, err := sortNumbersWithRules(slices.Clone(u), rules)
		if err != nil {
			return false, fmt.Errorf("sort: %w", err)
		}

		for i := 1; i < len(u); i++ {
			nlow, ok := findOrdinal(u[i-1], sorted)
			if !ok {
				return false, fmt.Errorf("ordinal of %q not found", u[i-1])
			}
			nhigh, ok := findOrdinal(u[i], sorted)
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

	findOrdinal := func(num string, sorted []string) (int, bool) {
		for i, on := range sorted {
			if on == num {
				return i, true
			}
		}
		return 0, false
	}

	isInOrder := func(u update) (bool, []string, error) {
		sorted, err := sortNumbersWithRules(slices.Clone(u), rules)
		if err != nil {
			return false, nil, fmt.Errorf("sort: %w", err)
		}

		for i := 1; i < len(u); i++ {
			nlow, ok := findOrdinal(u[i-1], sorted)
			if !ok {
				return false, nil, fmt.Errorf("ordinal of %q not found", u[i-1])
			}
			nhigh, ok := findOrdinal(u[i], sorted)
			if !ok {
				return false, nil, fmt.Errorf("ordinal of %q not found", u[i])
			}
			if nlow > nhigh {
				return false, sorted, nil
			}
		}
		return true, sorted, nil
	}

	midNum := func(u update) int {
		i := len(u) / 2
		num := u[i]
		n, _ := strconv.Atoi(num)
		return n
	}

	var sum int
	for _, u := range updates {
		ino, sorted, err := isInOrder(u)
		if err != nil {
			return 0, fmt.Errorf("isinorder: %w", err)
		}
		if ino {
			continue
		}

		// not in order - sort
		us := slices.Clone(u)
		sort.Slice(us, func(i, j int) bool {
			o1, ok := findOrdinal(us[i], sorted)
			if !ok {
				panic("find-ordinal of " + us[i])
			}
			o2, ok := findOrdinal(us[j], sorted)
			if !ok {
				panic("find-ordinal of " + us[j])
			}
			return o1 < o2
		})
		//sort.Slice()
		sum += midNum(us)

		// if ino {
		// 	sum += midNum(u)
		// 	log("pass: %v", u)
		// } else {
		// 	log("NOT pass: %v", u)
		// }
	}

	return sum, nil
}
