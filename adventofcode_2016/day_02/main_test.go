package main

import (
	"fmt"
	"testing"
)

func TestKeyPadNumber(t *testing.T) {
	tests := []struct {
		start int
		in    string
		num   int
	}{
		{
			start: 5,
			in:    "ULL",
			num:   1,
		},
		{
			start: 1,
			in:    "RRDDD",
			num:   9,
		},
		{
			start: 9,
			in:    "LURDL",
			num:   8,
		},
		{
			start: 8,
			in:    "UUUUD",
			num:   5,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			num, err := KeyPadNumber(test.start, test.in)
			if err != nil {
				t.Fatalf("failed, but no error was expected")
			}
			if num != test.num {
				t.Fatalf("want %d, have %d", test.num, num)
			}
		})
	}
}

func TestKeyCode(t *testing.T) {
	var in = `
		ULL
		RRDDD
		LURDL
		UUUUD
	`
	exp := "1985"
	res, err := KeyCode(in)
	if err != nil {
		t.Fatalf("failed, but no error was expected")
	}
	if res != exp {
		t.Fatalf("want %q, have %q", exp, res)
	}
}
