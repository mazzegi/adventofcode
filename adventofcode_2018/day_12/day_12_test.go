package day_12

import (
	"adventofcode_2018/testutil"
	"testing"
)

const inputStateTest = "#..#.#..##......###...###"

const inputRulesTest = `
...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputStateTest, inputRulesTest, 20, 0)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 325
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputStateTest, inputRulesTest, 20)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
