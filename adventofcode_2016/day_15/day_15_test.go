package day_15

import "testing"

func TestDropTime(t *testing.T) {
	ds := []*Disc{
		{5, 4},
		{2, 1},
	}
	dt := DropTime(ds)
	if dt != 5 {
		t.Fatalf("expect a drop-time of 5")
	}
}
