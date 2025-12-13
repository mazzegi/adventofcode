
package day_04

import (
	"fmt"	
	"time"
	
	"github.com/mazzegi/adventofcode/errutil"	
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

func part1MainFunc(in string) (int, error){
	return 0, nil
}

func part2MainFunc(in string) (int, error){
	return 0, nil
}
