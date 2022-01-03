package testutil

import "testing"

func Assert(t *testing.T, expr bool, want interface{}, have interface{}) {
	if expr {
		return
	}
	t.Fatalf("want %v, have %v", want, have)
}
