package day_12

import (
	"testing"

	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 21
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 525152
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

const inputTestMatches = `
#.#.### 1,1,3
.#...#....###. 1,1,3
.#.###.#.###### 1,3,1,6
####.#...#... 4,1,1
#....######..#####. 1,6,5
.###.##....# 3,2,1
`

func TestRecordMatch(t *testing.T) {
	recs := mustParseRecords(readutil.ReadLines(inputTestMatches))
	for _, rec := range recs {
		if !recordMatch(rec) {
			t.Fatalf("should match")
		}
	}
}
