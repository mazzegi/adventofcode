package day_10

import (
	"adventofcode_2021/errutil"
	"adventofcode_2021/readutil"
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := syntaxErrorScore(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := middleCompletionScore(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

type bracket rune

type brackets []bracket

const (
	bInvalid bracket = ' '
	b1Open   bracket = '('
	b1Close  bracket = ')'
	b2Open   bracket = '['
	b2Close  bracket = ']'
	b3Open   bracket = '{'
	b3Close  bracket = '}'
	b4Open   bracket = '<'
	b4Close  bracket = '>'
)

func (b bracket) isValid() bool {
	switch b {
	case b1Open, b1Close,
		b2Open, b2Close,
		b3Open, b3Close,
		b4Open, b4Close:
		return true
	default:
		return false
	}
}

func (b bracket) isOpen() bool {
	switch b {
	case b1Open,
		b2Open,
		b3Open,
		b4Open:
		return true
	default:
		return false
	}
}

func (b bracket) isClose() bool {
	switch b {
	case b1Close,
		b2Close,
		b3Close,
		b4Close:
		return true
	default:
		return false
	}
}

func matchingClose(b bracket) bracket {
	switch b {
	case b1Open:
		return b1Close
	case b2Open:
		return b2Close
	case b3Open:
		return b3Close
	case b4Open:
		return b4Close
	default:
		return bInvalid
	}
}

func matches(bopen, bclose bracket) bool {
	return bclose == matchingClose(bopen)
}

var scores = map[bracket]int{
	b1Close: 3,
	b2Close: 57,
	b3Close: 1197,
	b4Close: 25137,
}

func parseBrackets(s string) (brackets, error) {
	var bs brackets
	for _, r := range s {
		b := bracket(r)
		if !b.isValid() {
			return nil, errors.Errorf("%q is not a valid bracket", string(r))
		}
		bs = append(bs, b)
	}
	return bs, nil
}

func parse(in string) ([]brackets, error) {
	var bss []brackets
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		bs, err := parseBrackets(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse brackets %q", line)
		}
		if len(bs) == 0 {
			return nil, errors.Errorf("no entries")
		}
		bss = append(bss, bs)
	}
	return bss, nil
}

func syntaxErrorScore(in string) (int, error) {
	bss, err := parse(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse")
	}
	if len(bss) == 0 {
		return 0, errors.Errorf("no entries")
	}

	var total int
	for _, bs := range bss {
		if ib, _, found := findFirstIllegalBracket(bs); found {
			score, ok := scores[ib]
			if !ok {
				return 0, errors.Errorf("no score for bracket %q", string(ib))
			}
			total += score
		}
	}

	return total, nil
}

type stack struct {
	brackets brackets
}

func (s *stack) push(b bracket) {
	s.brackets = append(s.brackets, b)
}

func (s *stack) isEmpty() bool {
	return len(s.brackets) == 0
}

func (s *stack) last() bracket {
	if len(s.brackets) > 0 {
		return s.brackets[len(s.brackets)-1]
	}
	return bInvalid
}

func (s *stack) pop() {
	if len(s.brackets) == 0 {
		return
	}
	s.brackets = append(brackets{}, s.brackets[:len(s.brackets)-1]...)
}

func findFirstIllegalBracket(bs brackets) (invalid bracket, completion brackets, isInvalid bool) {
	openStack := &stack{}
	for _, b := range bs {
		if openStack.isEmpty() {
			if b.isOpen() {
				openStack.push(b)
				continue
			} else {
				return b, nil, true
			}
		}

		if b.isOpen() {
			openStack.push(b)
			continue
		}
		// its a close - check if it matches the last open
		if openStack.isEmpty() {
			return b, nil, true
		}
		lastOpen := openStack.last()
		if !matches(lastOpen, b) {
			return b, nil, true
		}
		openStack.pop()
	}

	var compl brackets
	for !openStack.isEmpty() {
		b := openStack.last()
		openStack.pop()
		compl = append(compl, matchingClose(b))
	}

	return bInvalid, compl, false
}

func middleCompletionScore(in string) (int, error) {
	bss, err := parse(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse")
	}
	if len(bss) == 0 {
		return 0, errors.Errorf("no entries")
	}

	var complScores []int
	for _, bs := range bss {
		_, compl, found := findFirstIllegalBracket(bs)
		if found {
			continue
		}
		cs, err := complScore(compl)
		if err != nil {
			return 0, errors.Wrap(err, "compl. score")
		}
		complScores = append(complScores, cs)
	}
	if len(complScores)%2 == 0 {
		return 0, errors.Errorf("promise of odd violated")
	}
	sort.Ints(complScores)
	middle := complScores[len(complScores)/2]

	return middle, nil
}

var complScores = map[bracket]int{
	b1Close: 1,
	b2Close: 2,
	b3Close: 3,
	b4Close: 4,
}

func complScore(compl brackets) (int, error) {
	var value int
	for _, b := range compl {
		value *= 5
		add, ok := complScores[b]
		if !ok {
			return 0, errors.Errorf("no compl. score for %q", string(b))
		}
		value += add
	}
	return value, nil
}
