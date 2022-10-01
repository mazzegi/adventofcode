package day_03

import (
	"fmt"
	"strconv"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func Part1() {
	r, err := CalcRate(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: rate=%v, power=%d \n", r, r.Power())
}

func Part2() {
	r, err := CalcLifeSupportRating(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: rate=%v, value=%d \n", r, r.Value())
}

type Rate struct {
	gamma   int
	epsilon int
}

func (r Rate) Power() int {
	return r.gamma * r.epsilon
}

type LifeSupportRating struct {
	oxyGen   int
	co2Scrub int
}

func (r LifeSupportRating) Value() int {
	return r.oxyGen * r.co2Scrub
}

type BinCode []bool

func (bc BinCode) Int() int {
	var s string
	for _, b := range bc {
		if b {
			s += "1"
		} else {
			s += "0"
		}
	}
	n, _ := strconv.ParseUint(s, 2, 64)
	return int(n)
}

func ParseBinCode(s string) (BinCode, error) {
	var bc BinCode
	for _, r := range s {
		switch r {
		case '0':
			bc = append(bc, false)
		case '1':
			bc = append(bc, true)
		default:
			return nil, errors.Errorf("invalid rune in bin-code %q", string(r))
		}
	}
	return bc, nil
}

func ParseBinCodes(in string) ([]BinCode, error) {
	var bcs []BinCode
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		bc, err := ParseBinCode(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-bin-code %q", line)
		}
		if len(bcs) > 0 {
			refLen := len(bcs[0])
			if len(bc) != refLen {
				return nil, errors.Errorf("inconsistent lens. have %d, want %d", len(bc), refLen)
			}
		}

		bcs = append(bcs, bc)
	}
	return bcs, nil
}

func CalcRate(in string) (Rate, error) {
	bcs, err := ParseBinCodes(in)
	if err != nil {
		return Rate{}, errors.Wrap(err, "parse-bin-codes")
	}
	if len(bcs) == 0 {
		return Rate{}, errors.Errorf("no bin codes")
	}

	refLen := len(bcs[0])
	oneCounts := make([]int, refLen)
	zeroCounts := make([]int, refLen)
	for _, bc := range bcs {
		for i, b := range bc {
			if b {
				oneCounts[i]++
			} else {
				zeroCounts[i]++
			}
		}
	}

	gammaBC := make(BinCode, refLen)
	epsBC := make(BinCode, refLen)
	for i := 0; i < refLen; i++ {
		if oneCounts[i] >= zeroCounts[i] {
			gammaBC[i] = true
		} else {
			epsBC[i] = true
		}
	}

	return Rate{
		gamma:   gammaBC.Int(),
		epsilon: epsBC.Int(),
	}, nil
}

func MostCommon(bcs []BinCode, bitIdx int) bool {
	var oneCount int
	var zeroCount int
	for _, bc := range bcs {
		if bc[bitIdx] {
			oneCount++
		} else {
			zeroCount++
		}
	}
	if oneCount >= zeroCount {
		return true
	}
	return false
}

func LeastCommon(bcs []BinCode, bitIdx int) bool {
	var oneCount int
	var zeroCount int
	for _, bc := range bcs {
		if bc[bitIdx] {
			oneCount++
		} else {
			zeroCount++
		}
	}
	if zeroCount <= oneCount {
		return false
	}
	return true
}

func FilterBinCodesMostCommon(bcs []BinCode, bitIdx int) []BinCode {
	var fbcs []BinCode
	mc := MostCommon(bcs, bitIdx)
	for _, bc := range bcs {
		if bc[bitIdx] == mc {
			fbcs = append(fbcs, bc)
		}
	}
	return fbcs
}

func FilterBinCodesLeastCommon(bcs []BinCode, bitIdx int) []BinCode {
	var fbcs []BinCode
	mc := LeastCommon(bcs, bitIdx)
	for _, bc := range bcs {
		if bc[bitIdx] == mc {
			fbcs = append(fbcs, bc)
		}
	}
	return fbcs
}

func CalcOxygenGeneratorRating(bcs []BinCode) (int, error) {
	if len(bcs) == 0 {
		return 0, errors.Errorf("empy bin-codes")
	}
	refLen := len(bcs[0])

	for i := 0; i < refLen; i++ {
		bcs = FilterBinCodesMostCommon(bcs, i)
		if len(bcs) == 1 {
			return bcs[0].Int(), nil
		}
	}
	return 0, errors.Errorf("more than one left after whole iteration")
}

func CalcCO2ScrubberRating(bcs []BinCode) (int, error) {
	if len(bcs) == 0 {
		return 0, errors.Errorf("empy bin-codes")
	}
	refLen := len(bcs[0])

	for i := 0; i < refLen; i++ {
		bcs = FilterBinCodesLeastCommon(bcs, i)
		if len(bcs) == 1 {
			return bcs[0].Int(), nil
		}
	}
	return 0, errors.Errorf("more than one left after whole iteration")
}

func CalcLifeSupportRating(in string) (LifeSupportRating, error) {
	bcs, err := ParseBinCodes(in)
	if err != nil {
		return LifeSupportRating{}, errors.Wrap(err, "parse-bin-codes")
	}
	lsr := LifeSupportRating{}
	lsr.oxyGen, err = CalcOxygenGeneratorRating(bcs)
	if err != nil {
		return LifeSupportRating{}, errors.Wrap(err, "calc-oxy-gen")
	}
	lsr.co2Scrub, err = CalcCO2ScrubberRating(bcs)
	if err != nil {
		return LifeSupportRating{}, errors.Wrap(err, "calc-co2-scrub")
	}
	return lsr, nil
}
