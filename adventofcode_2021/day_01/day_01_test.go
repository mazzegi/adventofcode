package day_01

import "testing"

const inTest = `
199
200
208
210
200
207
240
269
260
263
`

func TestIncCount(t *testing.T) {
	ic, err := incCount(inTest)
	if err != nil {
		t.Fatalf("failed but no err expected")
	}
	if ic != 7 {
		t.Fatalf("want %d, have %d", 7, ic)
	}
}

func TestIncWindowCount(t *testing.T) {
	ic, err := incWindowCount(inTest)
	if err != nil {
		t.Fatalf("failed but no err expected")
	}
	if ic != 5 {
		t.Fatalf("want %d, have %d", 7, ic)
	}
}
