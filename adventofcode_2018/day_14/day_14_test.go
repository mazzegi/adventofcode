package day_14

import (
	"reflect"
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2018/testutil"
)

const inputTest = ""

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(3, 7, 9)
	testutil.CheckUnexpectedError(t, err)
	var exp string = "5158916779"
	if exp != res {
		t.Fatalf("want %v, have %v", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	var res int
	var err error
	var exp int

	res, err = part2MainFunc(3, 7, []int{5, 9, 4, 1, 4})
	testutil.CheckUnexpectedError(t, err)
	exp = 2018
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}

	res, err = part2MainFunc(3, 7, []int{9, 2, 5, 1, 0})
	testutil.CheckUnexpectedError(t, err)
	exp = 18
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}

	res, err = part2MainFunc(3, 7, []int{0, 1, 2, 4, 5})
	testutil.CheckUnexpectedError(t, err)
	exp = 5
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}

	res, err = part2MainFunc(3, 7, []int{5, 1, 5, 8, 9})
	testutil.CheckUnexpectedError(t, err)
	exp = 9
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func assertEqual(t *testing.T, v1, v2 any) {
	if reflect.DeepEqual(v1, v2) {
		return
	}
	t.Fatalf("want %v, have %v", v1, v2)
}

func TestDigits(t *testing.T) {
	var ds []int

	ds = digits(23573)
	assertEqual(t, []int{2, 3, 5, 7, 3}, ds)

	ds = digits(23)
	assertEqual(t, []int{2, 3}, ds)

	ds = digits(7)
	assertEqual(t, []int{7}, ds)
}
