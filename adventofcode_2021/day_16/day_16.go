package day_16

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2021/day_16/bits"
	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := sumOfVersions(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := evalMessage(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

func sumOfVersions(in string) (int, error) {
	in = strings.Trim(in, " \r\n\t")
	msg, err := bits.ParseHex(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-hex")
	}
	sv := msg.SumOfVersion()
	return sv, nil
}

func evalMessage(in string) (int, error) {
	in = strings.Trim(in, " \r\n\t")
	msg, err := bits.ParseHex(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-hex")
	}
	sv, err := msg.Eval()
	return sv, err
}
