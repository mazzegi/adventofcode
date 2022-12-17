package day_15

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest = `
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest, 10)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 26
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest, 20)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 56000011
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestIntervalJoin(t *testing.T) {
	tests := []struct {
		in        []Interval
		join      Interval
		exp       []Interval
		exclusive bool
	}{
		{
			in: []Interval{
				{1, 2},
				{4, 6},
			},
			join: Interval{8, 9},
			exp: []Interval{
				{1, 2},
				{4, 6},
				{8, 9},
			},
		},
		{
			in: []Interval{
				{1, 2},
				{4, 6},
			},
			join: Interval{5, 9},
			exp: []Interval{
				{1, 2},
				{4, 9},
			},
		},
		{
			in: []Interval{
				{1, 2},
				{4, 6},
			},
			join: Interval{2, 8},
			exp: []Interval{
				{1, 8},
			},
		},
		{
			in: []Interval{
				{1, 2},
				{4, 6},
				{10, 12},
				{110, 113},
			},
			join: Interval{5, 11},
			exp: []Interval{
				{1, 2},
				{4, 12},
				{110, 113},
			},
			exclusive: true,
		},
		{
			in: []Interval{
				{1, 2},
				{4, 6},
				{10, 12},
				{110, 113},
			},
			join: Interval{0, 114},
			exp: []Interval{
				{0, 114},
			},
		},
		{
			in: []Interval{
				{1, 2},
				{4, 6},
				{10, 12},
				{110, 113},
				{120, 143},
			},
			join: Interval{1, 113},
			exp: []Interval{
				{1, 113},
				{120, 143},
			},
		},
		{
			in: []Interval{
				{1, 2},
				{4, 6},
				{10, 12},
				{110, 113},
				{120, 143},
			},
			join: Interval{8, 201},
			exp: []Interval{
				{1, 2},
				{4, 6},
				{8, 201},
			},
		},
		{
			in: []Interval{
				{1, 2},
				{4, 6},
				{10, 12},
				{110, 113},
				{120, 143},
			},
			join: Interval{8, 120},
			exp: []Interval{
				{1, 2},
				{4, 6},
				{8, 143},
			},
		},
		{
			in: []Interval{
				{1, 2},
				{4, 6},
				{10, 12},
				{110, 113},
				{120, 143},
			},
			join: Interval{8, 119},
			exp: []Interval{
				{1, 2},
				{4, 6},
				{8, 143},
			},
		},
	}

	skipNonExclusive := false
	for i, test := range tests {
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			if skipNonExclusive && !test.exclusive {
				t.Skip()
				return
			}

			is := NewIntervals(test.in...)
			is.Join(test.join)
			res := is.Slice()
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}
