package day_25

import (
	"fmt"
	"testing"
)

func TestDecodeSNAFU(t *testing.T) {
	tests := []struct {
		in  string
		exp int
	}{
		{"101=", 128},
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			res := DecodeSNAFU(test.in)
			if test.exp != res {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestEncodeSNAFU(t *testing.T) {
	tests := []struct {
		in  int
		exp string
	}{
		{3, "1="},
		{128, "101="},
		{1747, "1=-0-2"},
		{906, "12111"},
		{198, "2=0="},
		{11, "21"},
		{201, "2=01"},
		{31, "111"},
		{1257, "20012"},
		{32, "112"},
		{353, "1=-1="},
		{107, "1-12"},
		{7, "12"},
		{37, "122"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			res := EncodeSNAFU(test.in)
			if test.exp != res {
				t.Fatalf("want %q, have %q", test.exp, res)
			}
		})
	}
}

func TestInverseSNAFU(t *testing.T) {
	var s string
	var si string

	s = "1=-0-2"
	si = invSNAFU(s)

	n := DecodeSNAFU(s)
	ni := DecodeSNAFU(si)

	s = "1121-1110-1=0"
	si = invSNAFU(s)
	n = DecodeSNAFU(s)
	ni = DecodeSNAFU(si)

	_ = n
	_ = ni
}
