package day_03

import (
	"reflect"
	"testing"

	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 4361
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 467835
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestNumberItemAdjacents(t *testing.T) {
	{
		ni := NumberItem{
			Number:    828,
			NumDigits: 3,
			Position:  grid.Pt(10, 10),
		}

		wantAdjs := []grid.Point{
			grid.Pt(9, 9),
			grid.Pt(10, 9),
			grid.Pt(11, 9),
			grid.Pt(12, 9),
			grid.Pt(13, 9),

			grid.Pt(9, 10),
			grid.Pt(13, 10),

			grid.Pt(9, 11),
			grid.Pt(10, 11),
			grid.Pt(11, 11),
			grid.Pt(12, 11),
			grid.Pt(13, 11),
		}
		adjs := NumberItemAdjacents(ni)
		grid.SortPoints(wantAdjs)
		grid.SortPoints(adjs)
		if !reflect.DeepEqual(wantAdjs, adjs) {
			t.Fatalf("points not equal")
		}
	}
	{
		ni := NumberItem{
			Number:    5,
			NumDigits: 1,
			Position:  grid.Pt(0, 0),
		}

		wantAdjs := []grid.Point{
			grid.Pt(-1, -1),
			grid.Pt(0, -1),
			grid.Pt(1, -1),

			grid.Pt(-1, 0),
			grid.Pt(1, 0),

			grid.Pt(-1, 1),
			grid.Pt(0, 1),
			grid.Pt(1, 1),
		}
		adjs := NumberItemAdjacents(ni)
		grid.SortPoints(wantAdjs)
		grid.SortPoints(adjs)
		if !reflect.DeepEqual(wantAdjs, adjs) {
			t.Fatalf("points not equal")
		}
	}

}
