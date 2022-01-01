package day_09

import (
	"fmt"
	"testing"
)

func TestDecompress(t *testing.T) {

	tests := []struct {
		in  string
		dec string
		len int
	}{
		{
			in:  "ADVENT",
			dec: "ADVENT",
			len: 6,
		},
		{
			in:  "A(1x5)BC",
			dec: "ABBBBBC",
			len: 7,
		},
		{
			in:  "(3x3)XYZ",
			dec: "XYZXYZXYZ",
			len: 9,
		},
		{
			in:  "A(2x2)BCD(2x2)EFG",
			dec: "ABCBCDEFEFG",
			len: 11,
		},
		{
			in:  "(6x1)(1x3)A",
			dec: "(1x3)A",
			len: 6,
		},
		{
			in:  "X(8x2)(3x3)ABCY",
			dec: "X(3x3)ABC(3x3)ABCY",
			len: 18,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			dec, err := Decompress(test.in)
			if err != nil {
				t.Fatalf("error where no error is expected: %v", err)
			}
			if dec != test.dec {
				t.Fatalf("decode: expect %q, got %q", test.dec, dec)
			}
			if len(dec) != test.len {
				t.Fatalf("decode-len: expect %d, got %d", test.len, len(dec))
			}
		})
	}

}

func TestDecompress2(t *testing.T) {

	tests := []struct {
		in  string
		len int
	}{
		{
			in:  "(3x3)XYZ",
			len: 9,
		},
		{
			in:  "X(8x2)(3x3)ABCY",
			len: 20,
		},
		{
			in:  "(27x12)(20x12)(13x14)(7x10)(1x12)A",
			len: 241920,
		},
		{
			in:  "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
			len: 445,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			decLen, err := Decompress2Len(test.in)
			if err != nil {
				t.Fatalf("error where no error is expected: %v", err)
			}

			if decLen != test.len {
				t.Fatalf("decode-len: expect %d, got %d", test.len, decLen)
			}
		})
	}

}
