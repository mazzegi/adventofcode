package day_16

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest1 = "12345678"
const inputTest2 = "80871224585914546619083218645595"
const inputTest3 = "19617804207202209144916044189917"
const inputTest4 = "69317163492948606335995924319873"

func TestPart1MainFunc(t *testing.T) {
	t.Run("test_01", func(t *testing.T) {
		res, err := part1MainFunc(inputTest1, 4)
		testutil.CheckUnexpectedError(t, err)
		exp := "01029498"
		if exp != res {
			t.Fatalf("want %v, have %v", exp, res)
		}
	})
	t.Run("test_02", func(t *testing.T) {
		//t.Skip()
		res, err := part1MainFunc(inputTest2, 100)
		testutil.CheckUnexpectedError(t, err)
		exp := "24176176"
		if exp != res {
			t.Fatalf("want %v, have %v", exp, res)
		}
	})
	t.Run("test_03", func(t *testing.T) {
		//t.Skip()
		res, err := part1MainFunc(inputTest3, 100)
		testutil.CheckUnexpectedError(t, err)
		exp := "73745418"
		if exp != res {
			t.Fatalf("want %v, have %v", exp, res)
		}
	})
	t.Run("test_04", func(t *testing.T) {
		//t.Skip()
		res, err := part1MainFunc(inputTest4, 100)
		testutil.CheckUnexpectedError(t, err)
		exp := "52432133"
		if exp != res {
			t.Fatalf("want %v, have %v", exp, res)
		}
	})
}

func TestPart2MainFunc(t *testing.T) {
	t.Run("test_02", func(t *testing.T) {
		//t.Skip()
		res, err := part2MainFunc(inputTest2, 100)
		testutil.CheckUnexpectedError(t, err)
		exp := "84462026"
		if exp != res {
			t.Fatalf("want %v, have %v", exp, res)
		}
	})
	t.Run("test_03", func(t *testing.T) {
		t.Skip()
		res, err := part2MainFunc(inputTest3, 100)
		testutil.CheckUnexpectedError(t, err)
		exp := "78725270"
		if exp != res {
			t.Fatalf("want %v, have %v", exp, res)
		}
	})
	t.Run("test_04", func(t *testing.T) {
		t.Skip()
		res, err := part2MainFunc(inputTest4, 100)
		testutil.CheckUnexpectedError(t, err)
		exp := "53553731"
		if exp != res {
			t.Fatalf("want %v, have %v", exp, res)
		}
	})
}
