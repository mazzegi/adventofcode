package day_08

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := uniqueDigitCount(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := sumOfOutputs(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

const signals = "abcdefg"

type Entry struct {
	signals []string
	outputs []string
}

func ParseEntry(s string) (Entry, error) {
	s = strings.Trim(s, " \r\n\t")
	sl := strings.Split(s, "|")
	if len(sl) != 2 {
		return Entry{}, errors.Errorf("invalid entry %q", s)
	}
	sigs := strings.Fields(sl[0])
	outs := strings.Fields(sl[1])
	if len(sigs) != 10 {
		return Entry{}, errors.Errorf("invalid entry %q. len(sigs) != 10", s)
	}
	if len(outs) != 4 {
		return Entry{}, errors.Errorf("invalid entry %q. len(outs) != 4", s)
	}

	for _, s := range sigs {
		for _, r := range s {
			if !strings.ContainsRune(signals, r) {
				return Entry{}, errors.Errorf("invalid rune %q", string(r))
			}
		}
	}
	for _, s := range outs {
		for _, r := range s {
			if !strings.ContainsRune(signals, r) {
				return Entry{}, errors.Errorf("invalid rune %q", string(r))
			}
		}
	}

	return Entry{
		signals: sigs,
		outputs: outs,
	}, nil
}

func ParseEntries(in string) ([]Entry, error) {
	var es []Entry
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		e, err := ParseEntry(line)
		if err != nil {
			return nil, err
		}
		es = append(es, e)
	}
	return es, nil
}

const (
	segCount0 = 6
	segCount1 = 2
	segCount2 = 5
	segCount3 = 5
	segCount4 = 4
	segCount5 = 5
	segCount6 = 6
	segCount7 = 3
	segCount8 = 7
	segCount9 = 6
)

func uniqueDigitCount(in string) (int, error) {
	es, err := ParseEntries(in)
	if err != nil {
		return 0, err
	}

	isUnique := func(d string) bool {
		switch len(d) {
		case segCount1, segCount4, segCount7, segCount8:
			return true
		default:
			return false
		}
	}

	var total int
	for _, e := range es {
		for _, d := range e.outputs {
			if isUnique(d) {
				total++
			}
		}
	}

	return total, nil
}

func sumOfOutputs(in string) (int, error) {
	es, err := ParseEntries(in)
	if err != nil {
		return 0, err
	}

	var sum int
	for _, e := range es {
		mapping, err := deductSignals(e.signals)
		if err != nil {
			return 0, errors.Wrap(err, "deduct signals")
		}
		var digs []int
		for _, o := range e.outputs {
			n, err := decode(o, mapping)
			if err != nil {
				return 0, errors.Wrap(err, "decode output")
			}
			digs = append(digs, n)
		}
		var value int
		pow := 0
		for i := len(digs) - 1; i >= 0; i-- {
			value += digs[i] * int(math.Pow10(pow))
			pow++
		}
		sum += value
	}

	return sum, nil
}

//assume positions
//		0
//	1		2
//		3
//	4		5
//		6

type Position int

var displays = [][]Position{
	{0, 1, 2, 4, 5, 6},    // 0 => 6
	{2, 5},                // 1 => 2 *
	{0, 2, 3, 4, 6},       // 2 => 5
	{0, 2, 3, 5, 6},       // 3 => 5
	{1, 2, 3, 5},          // 4 => 4 *
	{0, 1, 3, 5, 6},       // 5 => 5
	{0, 1, 3, 4, 5, 6},    // 6 => 6
	{0, 2, 5},             // 7 => 3 *
	{0, 1, 2, 3, 4, 5, 6}, // 8 => 7 *
	{0, 1, 2, 3, 5, 6},    // 9 => 6
}

func displaysWith(size int) [][]Position {
	var ds [][]Position
	for _, d := range displays {
		if size == len(d) {
			ds = append(ds, d)
		}
	}
	return ds
}

func matchDisplay(poss []Position, disp []Position) bool {
	if len(poss) != len(disp) {
		return false
	}
	for i, pos := range poss {
		if pos != disp[i] {
			return false
		}
	}
	return true
}

func decode(s string, mapping map[rune]Position) (int, error) {
	var poss []Position
	for _, r := range s {
		pos, ok := mapping[r]
		if !ok {
			return 0, errors.Errorf("no mapping for rune %q", string(r))
		}
		poss = append(poss, pos)
	}
	sort.Slice(poss, func(i, j int) bool {
		return poss[i] < poss[j]
	})
	for num, disp := range displays {
		if matchDisplay(poss, disp) {
			return num, nil
		}
	}
	return 0, errors.Errorf("positions %v does not match any display", poss)
}

func intersection(poss1, poss2 []Position) []Position {
	var is []Position
	in2 := func(p Position) bool {
		for _, n := range poss2 {
			if p == n {
				return true
			}
		}
		return false
	}
	for _, p := range poss1 {
		if in2(p) {
			is = append(is, p)
		}
	}
	return is
}

type cand struct {
	r    rune
	poss []Position
}

func isValidMapping(sigs []string, mapping map[rune]Position) bool {
	for _, sig := range sigs {
		if _, err := decode(sig, mapping); err != nil {
			return false
		}
	}
	return true
}

func deductSignals(sigs []string) (map[rune]Position, error) {
	//mapping := map[rune]Position{}

	sort.SliceStable(sigs, func(i, j int) bool {
		return len(sigs[i]) < len(sigs[j])
	})

	constraints := map[rune][]Position{}
	for _, sig := range sigs {
		ds := displaysWith(len(sig))
		if len(ds) == 0 {
			return nil, errors.Errorf("no display with len %d", len(sig))
		}

		if len(ds) > 1 {
			continue
		}
		d := ds[0]
		for _, r := range sig {
			var is []Position
			if eposs, ok := constraints[r]; ok {
				is = intersection(d, eposs)
				if len(is) == 0 {
					return nil, errors.Errorf("impossible - intersection is empty")
				}
			} else {
				is = d
			}
			constraints[r] = is
		}
	}

	//dbg
	var cands []cand
	for r, poss := range constraints {
		cands = append(cands, cand{
			r:    r,
			poss: poss,
		})
	}
	sort.SliceStable(cands, func(i, j int) bool {
		if len(cands[i].poss) == len(cands[j].poss) {
			return cands[i].r < cands[j].r
		}
		return len(cands[i].poss) < len(cands[j].poss)
	})
	for _, cand := range cands {
		fmt.Printf("%q => %v\n", cand.r, cand.poss)
	}

	cand0 := cands[0]

	// for _, cand := range cands {
	for _, pos := range cand0.poss {
		candMapping := map[rune]Position{
			cand0.r: pos,
		}
		// remove pos from other
		rcands, err := removedPos(cands, cand0.r, pos)
		if err != nil {
			continue
		}
		err = mapCandidates(sigs, candMapping, rcands)
		if err == nil {
			// check fi all signals can be decoded
			if isValidMapping(sigs, candMapping) {
				fmt.Printf("found mapping\n")
				for r, pos := range candMapping {
					fmt.Printf("%q => %d\n", r, pos)
				}
				return candMapping, nil
			}
		}
	}
	// }

	return nil, errors.Errorf("not able to map candidates")
}

func mapCandidates(sigs []string, mapping map[rune]Position, cands []cand) error {
	for _, cand := range cands {
		for _, pos := range cand.poss {
			mapping[cand.r] = pos
			// remove pos from other
			rcands, err := removedPos(cands, cand.r, pos)
			if err != nil {
				continue
			}
			if len(rcands) == 0 {
				return nil
			}
			err = mapCandidates(sigs, mapping, rcands)
			if err == nil {
				if isValidMapping(sigs, mapping) {
					return nil
				}
			}
		}
	}
	return errors.Errorf("not able to empyt candidates")
}

func removedPos(cands []cand, remRune rune, rempos Position) ([]cand, error) {
	var rcs []cand
	for _, c := range cands {
		if c.r == remRune {
			continue
		}
		rcand := cand{
			r: c.r,
		}
		for _, p := range c.poss {
			if p != rempos {
				rcand.poss = append(rcand.poss, p)
			}
		}
		if len(rcand.poss) == 0 {
			return nil, errors.Errorf("no positions left for %q", string(rcand.r))
		}
		rcs = append(rcs, rcand)
	}
	return rcs, nil
}
