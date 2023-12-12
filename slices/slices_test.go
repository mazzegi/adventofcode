package slices

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

func TestInsert(t *testing.T) {
	ns := []int{1, 2, 3, 4, 6}
	ins := Insert(ns, 5, 4)
	exns := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(ins, exns) {
		t.Fatalf("want %v, have %v", exns, ins)
	}

	ns = []int{1, 2, 3, 4, 6}
	ins = Insert(ns, 5, 0)
	exns = []int{5, 1, 2, 3, 4, 6}
	if !reflect.DeepEqual(ins, exns) {
		t.Fatalf("want %v, have %v", exns, ins)
	}

	ns = []int{1, 2, 3, 4, 6}
	ins = Insert(ns, 5, 5)
	exns = []int{1, 2, 3, 4, 6, 5}
	if !reflect.DeepEqual(ins, exns) {
		t.Fatalf("want %v, have %v", exns, ins)
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		ns1   []int
		ns2   []int
		equal bool
	}{
		{
			ns1:   []int{1, 2, 3, 4},
			ns2:   []int{1, 2, 3, 4},
			equal: true,
		},
		{
			ns1:   []int{},
			ns2:   []int{},
			equal: true,
		},
		{
			ns1:   []int{1, 2, 4, 3},
			ns2:   []int{1, 2, 3, 4},
			equal: false,
		},
		{
			ns1:   []int{1, 2, 3, 4, 5},
			ns2:   []int{1, 2, 3, 4},
			equal: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Equal(test.ns1, test.ns2)
			testutil.Assert(t, test.equal, res)
		})
	}
}
