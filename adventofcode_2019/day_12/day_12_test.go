package day_12

import (
	"testing"

	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/testutil"
)

var inputTest1 = []grid.Point3D{
	grid.P3D(-1, 0, 2),
	grid.P3D(2, -10, -7),
	grid.P3D(4, -8, 8),
	grid.P3D(3, 5, -1),
}

var inputTest2 = []grid.Point3D{
	grid.P3D(-8, -10, 0),
	grid.P3D(5, 5, 10),
	grid.P3D(2, -7, 3),
	grid.P3D(9, -8, -3),
}

func TestPart1MainFunc(t *testing.T) {
	t.Run("test_1", func(t *testing.T) {
		res, err := part1MainFunc(inputTest1, 10)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 179
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})
	t.Run("test_2", func(t *testing.T) {
		res, err := part1MainFunc(inputTest2, 100)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 1940
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})
}

func TestPart2MainFunc(t *testing.T) {
	t.Run("test_1", func(t *testing.T) {
		t.Skip()
		res, err := part2MainFuncSuper(inputTest1)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 2772
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})
	t.Run("test_2", func(t *testing.T) {
		res, err := part2MainFuncSuper(inputTest2)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 4686774924
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})
}
