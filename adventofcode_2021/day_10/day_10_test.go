package day_10

import (
	"adventofcode_2021/testutil"
	"testing"
)

const inputTest = `
[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

func TestSyntaxErrorScore(t *testing.T) {
	res, err := syntaxErrorScore(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 26397
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestMiddleCompletionScore(t *testing.T) {
	res, err := middleCompletionScore(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 288957
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
