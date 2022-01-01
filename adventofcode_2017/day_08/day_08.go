package day_08

import (
	"adventofcode_2017/errutil"
	"fmt"

	"github.com/pkg/errors"
)

func Part1() {
	res, _, err := largestRegisterValue(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	_, he, err := largestRegisterValue(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", he)
}

//

func largestRegisterValue(in string) (largest int, highestEver int, err error) {
	inss, err := parseInstructions(in)
	if err != nil {
		return 0, 0, errors.Wrap(err, "parse-instructions")
	}
	if len(inss) == 0 {
		return 0, 0, errors.Errorf("no entries")
	}
	cpu := NewCPU(inss)
	cpu.Run()

	lgst := cpu.LargestRegisterValue()
	return lgst, cpu.highestEver, nil
}
