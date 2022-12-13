package day_13

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTestCheck = `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

var inputTest = []any{
	[]any{1, 1, 3, 1, 1},
	[]any{1, 1, 5, 1, 1},

	[]any{[]any{1}, []any{2, 3, 4}},
	[]any{[]any{1}, 4},

	[]any{9},
	[]any{[]any{8, 7, 6}},

	[]any{[]any{4, 4}, 4, 4},
	[]any{[]any{4, 4}, 4, 4, 4},

	[]any{7, 7, 7, 7},
	[]any{7, 7, 7},

	[]any{},
	[]any{3},

	[]any{[]any{[]any{}}},
	[]any{[]any{}},

	[]any{1, []any{2, []any{3, []any{4, []any{5, 6, 7}}}}, 8, 9},
	[]any{1, []any{2, []any{3, []any{4, []any{5, 6, 0}}}}, 8, 9},
}

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, inputTestCheck)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 13
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 140
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
