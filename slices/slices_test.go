package slices

import (
	"reflect"
	"testing"
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
