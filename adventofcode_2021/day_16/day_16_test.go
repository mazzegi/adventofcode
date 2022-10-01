package day_16

import (
	"fmt"
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

func TestSumOfVerions(t *testing.T) {
	tests := []struct {
		in   string
		exp  int
		skip bool
	}{
		{
			in:   "D2FE28",
			exp:  6,
			skip: true,
		},
		{
			in:   "38006F45291200",
			exp:  1 + 6 + 2,
			skip: true,
		},
		{
			in:   "EE00D40C823060",
			exp:  7 + 2 + 4 + 1,
			skip: false,
		},
		{
			in:   "8A004A801A8002F478",
			exp:  16,
			skip: true,
		},
		{
			in:   "620080001611562C8802118E34",
			exp:  12,
			skip: true,
		},
		{
			in:   "C0015000016115A2E0802F182340",
			exp:  23,
			skip: true,
		},
		{
			in:   "A0016C880162017C3686B18A3D4780",
			exp:  31,
			skip: true,
		},
	}

	doSkip := false
	for i, test := range tests {
		if doSkip && test.skip {
			t.Logf("skip %d", i+1)
			continue
		}
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res, err := sumOfVersions(test.in)
			testutil.CheckUnexpectedError(t, err)
			if res != test.exp {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}

func TestEval(t *testing.T) {
	tests := []struct {
		in   string
		exp  int
		skip bool
	}{
		{
			in:   "C200B40A82",
			exp:  3,
			skip: true,
		},
		{
			in:   "04005AC33890",
			exp:  54,
			skip: true,
		},
		{
			in:   "880086C3E88112",
			exp:  7,
			skip: false,
		},
		{
			in:   "CE00C43D881120",
			exp:  9,
			skip: true,
		},
		{
			in:   "D8005AC2A8F0",
			exp:  1,
			skip: true,
		},
		{
			in:   "F600BC2D8F",
			exp:  0,
			skip: true,
		},
		{
			in:   "9C005AC2F8F0",
			exp:  0,
			skip: true,
		},
		{
			in:   "9C0141080250320F1802104A08",
			exp:  1,
			skip: true,
		},
	}

	doSkip := false
	for i, test := range tests {
		if doSkip && test.skip {
			t.Logf("skip %d", i+1)
			continue
		}
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res, err := evalMessage(test.in)
			testutil.CheckUnexpectedError(t, err)
			if res != test.exp {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}
