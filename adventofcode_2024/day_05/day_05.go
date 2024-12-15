package day_05

import (
	"fmt"
	"slices"
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
	center string
	lower  []string
	upper  []string
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
			return sep, nil
		}
		// not done but also no action - sth wrong
		if len(insertedNumbers) == 0 {
			return nil, fmt.Errorf("no actions during last run")
		}
	}

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
	numbers := ruleNumbers[1:]
	sep, err := separate(center, numbers, rules)
	if err != nil {
		return 0, fmt.Errorf("separate: %w", err)
	}
	_ = sep
	_ = updates

	return 0, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
