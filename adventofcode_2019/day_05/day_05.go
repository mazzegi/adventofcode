package day_05

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2019/intcode"
	"github.com/mazzegi/adventofcode/errutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func part1MainFunc(in []int) (int, error) {
	input := 1
	_, out, err := intcode.Exec2(in, input)
	errutil.ExitOnErr(err)
	fmt.Println(out)
	//fmt.Println(mod)
	return 0, nil
}

func part2MainFunc(in []int) (int, error) {
	input := 5
	_, out, err := intcode.Exec2(in, input)
	errutil.ExitOnErr(err)
	fmt.Println(out)
	//fmt.Println(mod)
	return 0, nil
}
