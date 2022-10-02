package combi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	permsContains := func(ps [][]int, p []int) bool {
		for _, ep := range ps {
			if reflect.DeepEqual(ep, p) {
				return true
			}
		}
		return false
	}

	permsEqual := func(ps1, ps2 [][]int) bool {
		for _, p1 := range ps1 {
			if !permsContains(ps2, p1) {
				return false
			}
		}
		for _, p2 := range ps2 {
			if !permsContains(ps1, p2) {
				return false
			}
		}
		return true
	}

	tests := []struct {
		ns    []int
		perms [][]int
	}{
		{
			ns: []int{1, 2, 3},
			perms: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Permutations(test.ns)
			if !permsEqual(res, test.perms) {
				t.Fatalf("want %v, have %v", test.perms, res)
			}
		})
	}
}
