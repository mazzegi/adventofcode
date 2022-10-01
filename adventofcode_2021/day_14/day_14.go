package day_14

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func Part1() {
	// res, err := diffMostLeastCommonSpliced(inputTemplate, inputRules, 10)
	// errutil.ExitOnErr(err)
	// fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := diffMostLeastCommonSplicedHashed(inputTemplate, inputRules, 40)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

type runePair struct {
	r1 rune
	r2 rune
}

func rp(r1, r2 rune) runePair {
	return runePair{r1: r1, r2: r2}
}

type rule struct {
	pair   runePair
	insert rune
}

func parseRule(s string) (rule, error) {
	var sp, si string
	_, err := fmt.Sscanf(s, "%s -> %s", &sp, &si)
	if err != nil {
		return rule{}, errors.Wrapf(err, "scan-rule")
	}
	if len(sp) != 2 || len(si) != 1 {
		return rule{}, errors.Errorf("invalid rule")
	}
	rl := rule{
		pair: runePair{
			r1: []rune(sp)[0],
			r2: []rune(sp)[1],
		},
		insert: []rune(si)[0],
	}
	return rl, nil
}

func parseRules(in string) (map[runePair]rune, error) {
	ruleMap := map[runePair]rune{}
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		rl, err := parseRule(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-rule %q", line)
		}
		ruleMap[rl.pair] = rl.insert
	}
	return ruleMap, nil
}

func diffMostLeastCommon(inTpl string, inRules string, steps int) (int, error) {
	tpl := strings.Trim(inTpl, " \r\n\t")
	if tpl == "" {
		return 0, errors.Errorf("template is empty")
	}
	rls, err := parseRules(inRules)
	if err != nil {
		return 0, errors.Wrap(err, "parse-rules")
	}
	if len(rls) == 0 {
		return 0, errors.Errorf("rules are empty")
	}

	poly := []rune(tpl)
	//fmt.Printf("Template: %s\n", string(poly))
	for step := 0; step < steps; step++ {
		// t0 := time.Now()
		// fmt.Printf("step %d (len = %d) ...\n", len(poly), step+1)
		var newPoly []rune
		for i := 0; i < len(poly)-1; i++ {
			newPoly = append(newPoly, poly[i])
			ir, ok := rls[rp(poly[i], poly[i+1])]
			if !ok {
				continue
			}
			newPoly = append(newPoly, ir)
		}
		newPoly = append(newPoly, poly[len(poly)-1])
		poly = newPoly
		//fmt.Printf("After %02d: %s\n", step+1, string(poly))
		//fmt.Printf("step %d ...done (len = %d) in %s\n", step+1, len(poly), time.Since(t0))
	}

	quants := map[rune]int{}
	for _, r := range poly {
		quants[r]++

	}

	var min int
	var max int
	first := true
	for _, v := range quants {
		if first {
			min, max = v, v
			first = false
			continue
		}
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return max - min, nil
}

//splice: 1000000

func mostLeastCommon(poly []rune) (most, least int) {
	quants := map[rune]int{}
	for _, r := range poly {
		quants[r]++
	}
	var min int
	var max int
	first := true
	for _, v := range quants {
		if first {
			min, max = v, v
			first = false
			continue
		}
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return max, min
}

func diffMostLeastCommonSpliced(inTpl string, inRules string, steps int) (int, error) {
	tpl := strings.Trim(inTpl, " \r\n\t")
	if tpl == "" {
		return 0, errors.Errorf("template is empty")
	}
	rls, err := parseRules(inRules)
	if err != nil {
		return 0, errors.Wrap(err, "parse-rules")
	}
	if len(rls) == 0 {
		return 0, errors.Errorf("rules are empty")
	}

	poly := []rune(tpl)
	quants := map[rune]int{}
	var leafCount int = 0
	splice(poly, rls, quants, steps, 0, &leafCount, 0, 0)

	for r, c := range quants {
		fmt.Printf("%s: %d\n", string(r), c)
	}

	//most least
	var min int
	var max int
	first := true
	for _, v := range quants {
		if first {
			min, max = v, v
			first = false
			continue
		}
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return max - min, nil
}

var splitAt = 1000000

func splice(poly []rune, rules map[runePair]rune, quants map[rune]int, times int, horz int, leafCount *int, left, right rune) {
	if len(poly) > splitAt+1 {
		//fmt.Printf("times: %d: (h = %d) len = %d, split\n", times, horz, len(poly))
		si := splitAt/2 + 1
		insert, ok := rules[rp(poly[si-1], poly[si])]

		poly1 := poly[:si]
		poly2 := poly[si:]

		splice(poly1, rules, quants, times, horz, leafCount, left, 0)
		splice(poly2, rules, quants, times, horz+1, leafCount, 0, right)

		if ok {
			if times-1 > 0 {
				polyIn := []rune{insert}
				splice(polyIn, rules, quants, times-1, horz, leafCount, poly1[len(poly1)-1], poly2[0])
			} else {
				quants[insert]++
			}
		}

		return
	}

	t0 := time.Now()
	//fmt.Printf("times: %d: (h = %d) len = %d, splice ...\n", times, horz, len(poly))
	var newPoly []rune
	if left != 0 {
		if ir, ok := rules[rp(left, poly[0])]; ok {
			newPoly = append(newPoly, ir)
		}
	}
	for i := 0; i < len(poly)-1; i++ {
		newPoly = append(newPoly, poly[i])
		ir, ok := rules[rp(poly[i], poly[i+1])]
		if !ok {
			continue
		}
		newPoly = append(newPoly, ir)
	}
	last := poly[len(poly)-1]
	newPoly = append(newPoly, last)
	if right != 0 {
		if ir, ok := rules[rp(last, right)]; ok {
			newPoly = append(newPoly, ir)
		}
	}

	//fmt.Printf("times: %d: (h = %d) len = %d, splice ... done in %s\n", times, horz, len(newPoly), time.Since(t0))

	if times > 1 {

		splice(newPoly, rules, quants, times-1, horz, leafCount, left, right)
	} else {
		*leafCount++
		fmt.Printf("spliced: times: %d: (h = %d) leaf-count: %d, in %s\n", times, horz, *leafCount, time.Since(t0))
		for _, r := range newPoly {
			quants[r]++
		}

	}
}

//// HASHED

func diffMostLeastCommonSplicedHashed(inTpl string, inRules string, steps int) (int, error) {
	tpl := strings.Trim(inTpl, " \r\n\t")
	if tpl == "" {
		return 0, errors.Errorf("template is empty")
	}
	rls, err := parseRules(inRules)
	if err != nil {
		return 0, errors.Wrap(err, "parse-rules")
	}
	if len(rls) == 0 {
		return 0, errors.Errorf("rules are empty")
	}

	poly := []rune(tpl)
	quants := map[rune]int{}
	hashLookup := newHashMap()
	spliceHashed(hashLookup, poly, rls, quants, steps, 0, 0)

	for r, c := range quants {
		fmt.Printf("%s: %d\n", string(r), c)
	}

	//most least
	var min int
	var max int
	first := true
	for _, v := range quants {
		if first {
			min, max = v, v
			first = false
			continue
		}
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return max - min, nil
}

//var splitHashedAt = 16

type hashKey struct {
	poly  string
	left  rune
	right rune
	times int
}

type hashMap struct {
	values map[hashKey]map[rune]int
	hits   int
}

func newHashMap() *hashMap {
	return &hashMap{
		values: map[hashKey]map[rune]int{},
		hits:   0,
	}
}

func (hm *hashMap) insert(poly string, left rune, right rune, times int, quants map[rune]int) {
	hm.values[hashKey{
		poly:  poly,
		left:  left,
		right: right,
		times: times,
	}] = quants
}

func (hm *hashMap) lookup(poly string, left rune, right rune, times int) (map[rune]int, bool) {
	q, ok := hm.values[hashKey{
		poly:  poly,
		left:  left,
		right: right,
		times: times,
	}]
	if !ok {
		return nil, false
	}
	hm.hits++
	return q, true
}

var splitHashedAt = 7

func spliceHashed(hashLookup *hashMap, poly []rune, rules map[runePair]rune, quants map[rune]int, times int, left, right rune) {
	t0 := time.Now()
	if hashedQuants, ok := hashLookup.lookup(string(poly), left, right, times); ok {
		fmt.Printf("hit hash-cache\n")
		for r, c := range hashedQuants {
			quants[r] += c
		}
		return
	}

	subQuants := map[rune]int{}
	defer func() {
		hashLookup.insert(string(poly), left, right, times, subQuants)
		for r, c := range subQuants {
			quants[r] += c
		}

		var total int
		for _, c := range quants {
			total += c
		}
		fmt.Printf("spliced: times: %d: cache-hits = %d, total = %d, in %s\n", times, hashLookup.hits, total, time.Since(t0))
	}()

	if len(poly) > splitHashedAt+1 {
		//fmt.Printf("times: %d: (h = %d) len = %d, split\n", times, horz, len(poly))
		si := splitHashedAt/2 + 1
		insert, ok := rules[rp(poly[si-1], poly[si])]

		poly1 := poly[:si]
		poly2 := poly[si:]

		spliceHashed(hashLookup, poly1, rules, subQuants, times, left, 0)
		spliceHashed(hashLookup, poly2, rules, subQuants, times, 0, right)

		if ok {
			subTimes := times - 1
			if subTimes > 0 {
				polyIn := []rune{insert}
				left := poly1[len(poly1)-1]
				right := poly2[0]
				// if hashedQuants, ok := hashLookup.lookup(string(polyIn), left, right, subTimes); ok {
				// 	fmt.Printf("hit hash-cache\n")
				// 	for r, c := range hashedQuants {
				// 		quants[r] += c
				// 	}
				// 	return
				// }

				//subQuants := map[rune]int{}
				spliceHashed(hashLookup, polyIn, rules, subQuants, subTimes, left, right)
				// for r, c := range subQuants {
				// 	quants[r] += c
				// }
				//hashLookup.insert(string(polyIn), left, right, subTimes, subQuants)
			} else {
				subQuants[insert]++
			}
		}

		return
	} else {

		//fmt.Printf("times: %d: (h = %d) len = %d, splice ...\n", times, horz, len(poly))
		var newPoly []rune
		if left != 0 {
			if ir, ok := rules[rp(left, poly[0])]; ok {
				newPoly = append(newPoly, ir)
			}
		}
		for i := 0; i < len(poly)-1; i++ {
			newPoly = append(newPoly, poly[i])
			ir, ok := rules[rp(poly[i], poly[i+1])]
			if !ok {
				continue
			}
			newPoly = append(newPoly, ir)
		}
		last := poly[len(poly)-1]
		newPoly = append(newPoly, last)
		if right != 0 {
			if ir, ok := rules[rp(last, right)]; ok {
				newPoly = append(newPoly, ir)
			}
		}

		//fmt.Printf("times: %d: (h = %d) len = %d, splice ... done in %s\n", times, horz, len(newPoly), time.Since(t0))

		if times > 1 {
			spliceHashed(hashLookup, newPoly, rules, subQuants, times-1, left, right)
		} else {
			//fmt.Printf("spliced: times: %d: (h = %d) leaf-count: %d, cache-hits = %d, in %s\n", times, horz, *leafCount, hashLookup.hits, time.Since(t0))
			for _, r := range newPoly {
				subQuants[r]++
			}

		}
	}
}
