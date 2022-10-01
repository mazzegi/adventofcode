package day_03

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2016/errutil"

	"github.com/pkg/errors"
)

func parse(in string) ([]Triple, error) {
	var ts []Triple
	sl := strings.Split(in, "\n")
	for _, s := range sl {
		s = strings.Trim(s, " \r\n\t")
		if s == "" {
			continue
		}
		var a, b, c int
		_, err := fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
		if err != nil {
			return nil, errors.Wrapf(err, "scan %q", s)
		}
		ts = append(ts, Triple{a, b, c})
	}
	return ts, nil
}

func Part1() {
	ts, err := parse(input)
	errutil.ExitOnErr(err)
	var validCount int
	var total int
	for _, t := range ts {
		total++
		if IsPossibleTriangle(t) {
			validCount++
		}
	}
	fmt.Printf("%d of %d are possible\n", validCount, total)
}

func Part2() {
	ts, err := parse(input)
	errutil.ExitOnErr(err)

	// transpose to columns
	var cts []Triple
	for i := 0; i < len(ts); i += 3 {
		cts = append(cts, Triple{ts[i][0], ts[i+1][0], ts[i+2][0]})
		cts = append(cts, Triple{ts[i][1], ts[i+1][1], ts[i+2][1]})
		cts = append(cts, Triple{ts[i][2], ts[i+1][2], ts[i+2][2]})
	}

	var validCount int
	var total int
	for _, t := range cts {
		total++
		if IsPossibleTriangle(t) {
			validCount++
		}
	}
	fmt.Printf("%d of %d are possible\n", validCount, total)

}

type Triple [3]int

func IsPossibleTriangle(t Triple) bool {
	sl := t[:]
	sort.Ints(sl)
	return sl[0]+sl[1] > sl[2]
}
