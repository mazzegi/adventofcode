package day_05

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/mathutil"
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

type Category string

type MappingRange struct {
	FromStartIdx int
	ToStartIdx   int
	Len          int
}

// func MapOne(mr MappingRange, val int) int {
// 	if val >= mr.FromStartIdx && val < mr.FromStartIdx+mr.Len {
// 		return mr.ToStartIdx + (val - mr.FromStartIdx)
// 	}
// 	return val
// }

// func MapMany(mr MappingRange, vals []int) []int {
// 	mvals := make([]int, len(vals))
// 	for i, v := range vals {
// 		mvals[i] = MapOne(mr, v)
// 	}
// 	return mvals
// }

type Mapping struct {
	FromCategory Category
	ToCategory   Category
	Ranges       []MappingRange
}

func (m *Mapping) MapOne(val int) int {
	for _, mr := range m.Ranges {
		if val >= mr.FromStartIdx && val < mr.FromStartIdx+mr.Len {
			return mr.ToStartIdx + (val - mr.FromStartIdx)
		}
	}
	return val
}

func (m *Mapping) MapMany(vals []int) []int {
	mvals := make([]int, len(vals))
	for i, v := range vals {
		mvals[i] = m.MapOne(v)
	}
	return mvals
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	if len(lines) < 8 {
		return 0, fmt.Errorf("too less lines (%d)", len(lines))
	}
	if !strings.HasPrefix(lines[0], "seeds:") {
		return 0, fmt.Errorf("invalid prefix of first line (want: seeds:)")
	}
	seeds := stringutil.MustStringsToInts(strings.Split(lines[0][6:], " "))

	var mappings []*Mapping
	var currMapping *Mapping
	flushCurr := func() {
		if currMapping == nil {
			return
		}
		mappings = append(mappings, currMapping)
	}

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if strings.HasSuffix(line, " map:") {
			flushCurr()

			catstr := line[:len(line)-5]
			from, to, ok := strings.Cut(catstr, "-to-")
			if !ok || from == "" || to == "" {
				return 0, fmt.Errorf("invalid mapping-category-string: %s", catstr)
			}
			currMapping = &Mapping{
				FromCategory: Category(from),
				ToCategory:   Category(to),
			}
			continue
		}
		if currMapping == nil {
			return 0, fmt.Errorf("no curr mapping at line %d", i)
		}
		rangeNums := stringutil.MustStringsToInts(strings.Split(line, " "))
		if len(rangeNums) != 3 {
			return 0, fmt.Errorf("invalid range line %q", line)
		}
		currMapping.Ranges = append(currMapping.Ranges, MappingRange{
			ToStartIdx:   rangeNums[0],
			FromStartIdx: rangeNums[1],
			Len:          rangeNums[2],
		})
	}
	flushCurr()

	mustFindMappingForSource := func(srcCat Category) *Mapping {
		for _, m := range mappings {
			if m.FromCategory == srcCat {
				return m
			}
		}
		errutil.FatalWhen(fmt.Errorf("found no mapping for source %q", srcCat))
		return nil
	}

	m := mustFindMappingForSource("seed")
	vals := slices.Clone(seeds)
	for {
		vals = m.MapMany(vals)
		if m.ToCategory == "location" {
			return mathutil.MinOfSlice(vals), nil
		}
		m = mustFindMappingForSource(m.ToCategory)
	}
}

type Range struct {
	StartIdx int
	Len      int
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	if len(lines) < 8 {
		return 0, fmt.Errorf("too less lines (%d)", len(lines))
	}
	if !strings.HasPrefix(lines[0], "seeds:") {
		return 0, fmt.Errorf("invalid prefix of first line (want: seeds:)")
	}
	seedRangesRaw := stringutil.MustStringsToInts(strings.Split(lines[0][6:], " "))
	if len(seedRangesRaw)%2 != 0 {
		return 0, fmt.Errorf("invalid len for range pairs %d", len(seedRangesRaw))
	}
	seedRanges := []Range{}
	for i := 0; i < len(seedRangesRaw); i += 2 {
		seedRanges = append(seedRanges, Range{
			StartIdx: seedRangesRaw[i],
			Len:      seedRangesRaw[i+1],
		})
	}
	log("generating seeds ...")
	var seeds []int
	for _, sr := range seedRanges {
		for i := sr.StartIdx; i < sr.StartIdx+sr.Len; i++ {
			seeds = append(seeds, i)
		}
	}
	log("generating seeds ...done")

	var mappings []*Mapping
	var currMapping *Mapping
	flushCurr := func() {
		if currMapping == nil {
			return
		}
		mappings = append(mappings, currMapping)
	}

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if strings.HasSuffix(line, " map:") {
			flushCurr()

			catstr := line[:len(line)-5]
			from, to, ok := strings.Cut(catstr, "-to-")
			if !ok || from == "" || to == "" {
				return 0, fmt.Errorf("invalid mapping-category-string: %s", catstr)
			}
			currMapping = &Mapping{
				FromCategory: Category(from),
				ToCategory:   Category(to),
			}
			continue
		}
		if currMapping == nil {
			return 0, fmt.Errorf("no curr mapping at line %d", i)
		}
		rangeNums := stringutil.MustStringsToInts(strings.Split(line, " "))
		if len(rangeNums) != 3 {
			return 0, fmt.Errorf("invalid range line %q", line)
		}
		currMapping.Ranges = append(currMapping.Ranges, MappingRange{
			ToStartIdx:   rangeNums[0],
			FromStartIdx: rangeNums[1],
			Len:          rangeNums[2],
		})
	}
	flushCurr()

	mustFindMappingForSource := func(srcCat Category) *Mapping {
		for _, m := range mappings {
			if m.FromCategory == srcCat {
				return m
			}
		}
		errutil.FatalWhen(fmt.Errorf("found no mapping for source %q", srcCat))
		return nil
	}

	m := mustFindMappingForSource("seed")
	vals := slices.Clone(seeds)
	for {
		log("map %q => %q", m.FromCategory, m.ToCategory)
		vals = m.MapMany(vals)
		if m.ToCategory == "location" {
			return mathutil.MinOfSlice(vals), nil
		}
		m = mustFindMappingForSource(m.ToCategory)
	}
}
