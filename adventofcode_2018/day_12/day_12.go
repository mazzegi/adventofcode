package day_12

import (
	"adventofcode_2018/errutil"
	"bytes"
	"fmt"

	"github.com/mazzegi/scan"
	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(inputState, inputRules, 20)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(inputState, inputRules, 20)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type rule struct {
	Pattern []byte
	Prod    byte
}

func part1MainFunc(stateStr string, rulesStr string, gens int) (int, error) {
	rules, err := scan.Lines[rule]("{{pattern: []byte}} => {{prod: byte}}", scan.BuiltinFuncs(), bytes.NewBufferString(rulesStr))
	if err != nil {
		return 0, errors.Wrap(err, "scan rules")
	}
	_ = rules

	//state := []byte(stateStr)
	return 0, nil
}

func part2MainFunc(stateStr string, rulesStr string, gens int) (int, error) {
	return 0, nil
}
