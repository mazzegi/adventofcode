package day_23

import (
	"adventofcode_2016/assembunny"
	"adventofcode_2016/errutil"
	"fmt"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	n, err := RegAValue(input, map[string]int{
		"a": 7,
	})
	errutil.ExitOnErr(err)
	fmt.Printf("part1: regA = %d\n", n)
}

const part1only = true

func Part2() {
	if part1only {
		return
	}

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

func calc(a int) int {
	var c int
	var d int

	b := a - 1
	for b > 0 {
		d = a
		a = 0

		for d > 0 {
			c = b
			for c > 0 {
				a++
				c--
			}
			d--
		}

		b--
		c = b
		d = c

		for d > 0 {
			d--
			c++
		}
		log("b = %d, c = %d", b, c)

		//c = b+d => c = b+b = 2*b
		//toggle c
	}
	log("a = %d, b = %d, c = %d,d = %d", a, b, c, d)

	c = 78

	for c > 0 {
		d = 70

		for d > 0 {
			a++
			d--
		}
		c--
	}

	return a
}
