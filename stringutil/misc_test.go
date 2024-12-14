package stringutil

import "testing"

func TestReverse(t *testing.T) {
	s := "XMAS"
	rs := Reverse(s)
	if "SAMX" != rs {
		t.Fatalf("wrong")
	}
}
