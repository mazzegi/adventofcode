package day_09

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

func part1MainFunc(prg []int) (int, error) {
	inr := intcode.NewIntSliceReader([]int{1})
	outw := intcode.NewIntSliceWriter()
	com := intcode.NewComputer(prg, inr, outw)
	err := com.Exec()
	errutil.ExitOnErr(err)
	fmt.Println(outw.Values())
	return 0, nil
}

func part2MainFunc(prg []int) (int, error) {
	return 0, nil
}
