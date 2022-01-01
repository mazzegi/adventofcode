package day_23

import (
	"adventofcode_2016/assembunny"
	"adventofcode_2016/errutil"
	"fmt"

	"github.com/pkg/errors"
)

func Part1() {
	n, err := RegAValue(input, map[string]int{
		"a": 7,
	})
	errutil.ExitOnErr(err)
	fmt.Printf("part1: regA = %d\n", n)
}

func Part2() {
	n, err := RegAValue(input, map[string]int{
		"a": 12,
	})
	errutil.ExitOnErr(err)
	fmt.Printf("part2: regA = %d\n", n)
}

var registers = []string{"a", "b", "c", "d"}

func RegAValue(in string, initValues map[string]int) (int, error) {
	iss, err := assembunny.ParseInstructions(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-instructions")
	}
	comp := assembunny.NewComputer(registers)
	for reg, val := range initValues {
		comp.SetReg(reg, val)
	}

	comp.Execute(iss, nil)

	return comp.ValueOf("a"), nil
}
