package intcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

func TestComputer(t *testing.T) {
	tests := []struct {
		prg    []int
		in     []int
		expOut []int
	}{
		{
			prg:    []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			in:     []int{},
			expOut: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			prg:    []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			in:     []int{},
			expOut: []int{1219070632396864},
		},
		{
			prg:    []int{104, 1125899906842624, 99},
			in:     []int{},
			expOut: []int{1125899906842624},
		},
		//from day5
		{
			prg:    []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			in:     []int{8},
			expOut: []int{1},
		},
		{
			prg:    []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			in:     []int{12},
			expOut: []int{0},
		},
		{
			prg:    []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			in:     []int{7},
			expOut: []int{1},
		},
		{
			prg:    []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			in:     []int{8},
			expOut: []int{0},
		},
		{
			prg:    []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			in:     []int{8},
			expOut: []int{1},
		},
		{
			prg:    []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			in:     []int{34234},
			expOut: []int{0},
		},
		{
			prg:    []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			in:     []int{7},
			expOut: []int{1},
		},
		{
			prg:    []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			in:     []int{8},
			expOut: []int{0},
		},
		// jumps
		{
			prg: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			in:     []int{4},
			expOut: []int{999},
		},
		{
			prg: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			in:     []int{8},
			expOut: []int{1000},
		},
		{
			prg: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			in:     []int{34},
			expOut: []int{1001},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			inr := NewIntSliceReader(test.in)
			outw := NewIntSliceWriter()
			com := NewComputer(test.prg, inr, outw)
			err := com.Exec()
			testutil.CheckUnexpectedError(t, err)
			if !reflect.DeepEqual(test.expOut, outw.Values()) {
				t.Fatalf("want %v, have %v", test.expOut, outw.Values())
			}
		})
	}
}
