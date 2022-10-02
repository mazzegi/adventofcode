package intcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

/*
2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
*/

func TestExec(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			out: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			in:  []int{1, 0, 0, 0, 99},
			out: []int{2, 0, 0, 0, 99},
		},
		{
			in:  []int{2, 3, 0, 3, 99},
			out: []int{2, 3, 0, 6, 99},
		},
		{
			in:  []int{2, 4, 4, 5, 99, 0},
			out: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			in:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			out: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("# %02d", i), func(t *testing.T) {
			out, err := Exec(test.in)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if !reflect.DeepEqual(out, test.out) {
				t.Fatalf("want %v, have %v", test.out, out)
			}
		})
	}
}

func TestSqueezeOpcode(t *testing.T) {
	code, p1, p2, p3 := squeezeOpcode(1002)
	testutil.Assert(t, 2, code)
	testutil.Assert(t, 0, p1)
	testutil.Assert(t, 1, p2)
	testutil.Assert(t, 0, p3)
}

func TestExec2(t *testing.T) {
	prg := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

	_, out, err := Exec2(prg, []int{45})
	testutil.CheckUnexpectedError(t, err)
	fmt.Println(out)
}
