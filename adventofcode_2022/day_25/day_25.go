package day_25

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatalIfErr(err error) {
	if err == nil {
		return
	}
	fatal("err not nil: %v", err)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %s", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func part1MainFunc(in string) (string, error) {
	lines := readutil.ReadLines(in)
	var sum int
	for _, s := range lines {
		v := DecodeSNAFU(s)
		sum += v
	}

	es := EncodeSNAFU(sum)
	return es, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
