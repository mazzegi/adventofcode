package day_08

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"

	"github.com/pkg/errors"
)

const inputTest = `
be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`

func TestUniqueDigitCount(t *testing.T) {
	res, err := uniqueDigitCount(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 26
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestSumOfOutputsSingle(t *testing.T) {
	//s := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	s := "fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb"
	res, err := sumOfOutputs(s)
	testutil.CheckUnexpectedError(t, err)
	//exp := 5353
	exp := 9361
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestSumOfOutputsMulti(t *testing.T) {
	tests := []struct {
		in  string
		exp int
	}{
		{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe", 8394},
		{"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc", 9781},
		{"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg", 1197},
		{"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb", 9361},
		{"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea", 4873},
		{"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb", 8418},
		{"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe", 4548},
		{"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef", 1625},
		{"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb", 8717},
		{"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce", 4315},
	}

	for _, test := range tests {
		res, err := sumOfOutputs(test.in)
		testutil.CheckUnexpectedError(t, errors.Wrapf(err, "for %q", test.in))
		if test.exp != res {
			t.Fatalf("want %d, have %d (%q)", test.exp, res, test.in)
		}
	}
}

func TestSumOfOutputs(t *testing.T) {
	res, err := sumOfOutputs(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp := 61229
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
