package day_12

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/slices"
	"github.com/mazzegi/adventofcode/stringutil"
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

type Record struct {
	Pattern []rune
	Groups  []int
}

func (r Record) Clone() Record {
	return Record{
		Pattern: slices.Clone(r.Pattern),
		Groups:  slices.Clone(r.Groups),
	}
}

func mustParseRecord(s string) Record {
	ps, gs, ok := strings.Cut(s, " ")
	if !ok {
		panic("cut: " + s)
	}
	grps := stringutil.MustStringsToInts(strings.Split(gs, ","))
	return Record{
		Pattern: []rune(ps),
		Groups:  grps,
	}
}

func mustParseRecords(sl []string) []Record {
	recs := make([]Record, len(sl))
	for i, s := range sl {
		recs[i] = mustParseRecord(s)
	}
	return recs
}

func recordMatch(rec Record) bool {
	var patternGroups []int
	currGrp := 0
	flush := func() {
		if currGrp > 0 {
			patternGroups = append(patternGroups, currGrp)
			currGrp = 0
		}
	}

	for _, r := range rec.Pattern {
		switch r {
		case '#':
			currGrp++
		case '.':
			flush()
		default:
			panic("invalid pattern: " + string(rec.Pattern))
		}
	}
	flush()

	eq := slices.Equal(patternGroups, rec.Groups)
	return eq
}

func numMatchingArrangements(rec Record) int {
	var num int

	qCnt := strings.Count(string(rec.Pattern), "?")
	repl := slices.Repeat('.', qCnt)
	last := slices.Repeat('#', qCnt)

	for {
		prec := rec.Clone()
		ir := 0
		for i, r := range prec.Pattern {
			if r == '?' {
				prec.Pattern[i] = repl[ir]
				ir++
			}
		}
		match := recordMatch(prec)
		if match {
			num++
		}

		if slices.Equal(repl, last) {
			break
		}

		// rotate repl
		for i := 0; i < qCnt; i++ {
			if repl[i] == '.' {
				repl[i] = '#'
				break
			}
			repl[i] = '.'
		}
	}

	return num
}

func part1MainFunc(in string) (int, error) {
	recs := mustParseRecords(readutil.ReadLines(in))

	var sum int
	for _, rec := range recs {
		sum += numMatchingArrangements(rec)
	}

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	recs := mustParseRecords(readutil.ReadLines(in))
	for i, rec := range recs {
		sl := slices.Repeat(string(rec.Pattern), 5)
		npattern := []rune(strings.Join(sl, "?"))
		var ngrps []int
		for i := 0; i < 5; i++ {
			ngrps = append(ngrps, rec.Groups...)
		}
		recs[i] = Record{
			Pattern: npattern,
			Groups:  ngrps,
		}
	}

	var sum int
	for _, rec := range recs {
		sum += numMatchingArrangements(rec)
	}

	return 0, nil
}
