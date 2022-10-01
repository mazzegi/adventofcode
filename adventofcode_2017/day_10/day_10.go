package day_10

import (
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"

	"fmt"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := listProduct(input, 256)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := knotHash(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %s\n", res)
}

func reverse(list []int, pos int, length int) {
	for i := 0; i < length/2; i++ {
		start := pos + i
		if start >= len(list) {
			start = start - len(list)
		}
		end := pos + length - 1 - i
		if end >= len(list) {
			end = end - len(list)
		}
		list[start], list[end] = list[end], list[start]
	}
}

func reverseCopy(list []int, pos int, length int) ([]int, error) {
	if length > len(list) {
		return nil, errors.Errorf("length %d is geater than list-size %d", length, len(list))
	}
	if pos < 0 || pos >= len(list) {
		return nil, errors.Errorf("invalid pos %d. list-size is %d", pos, len(list))
	}

	rlist := make([]int, len(list))
	copy(rlist, list)

	reverse(rlist, pos, length)

	return rlist, nil
}

func listProduct(in string, listSize int) (int, error) {
	lengths, err := readutil.ReadInts(in, ",")
	if err != nil {
		return 0, errors.Wrapf(err, "read-ints %q", in)
	}
	if len(lengths) == 0 {
		return 0, errors.Errorf("no entries")
	}

	//prepare list
	list := make([]int, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = i
	}

	//iterate lenghts
	pos := 0
	skip := 0
	for _, length := range lengths {
		list, err = reverseCopy(list, pos, length)
		if err != nil {
			return 0, errors.Wrap(err, "reverse copy")
		}
		pos += length + skip
		if pos >= len(list) {
			pos = pos - len(list)
		}
		skip++
	}

	return list[0] * list[1], nil
}

func knotHash(in string) (string, error) {
	in = strings.Trim(in, " \r\n\t")
	var lengths []int
	for _, b := range []byte(in) {
		lengths = append(lengths, int(b))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	//prepare list
	var listSize int = 256
	list := make([]int, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = i
	}

	//iterate lenghts
	var err error
	pos := 0
	skip := 0
	rounds := 64
	for round := 0; round < rounds; round++ {
		for il, length := range lengths {
			list, err = reverseCopy(list, pos, length)
			if err != nil {
				return "", errors.Wrapf(err, "reverse copy (in=%q, round=%d, il=%d)", in, round, il)
			}
			pos += length + skip
			if pos >= len(list) {
				pos = pos % len(list)
			}
			skip++
		}
	}

	//
	dense := make([]byte, 16)
	for di := 0; di < 16; di++ {
		var d byte
		for i := di * 16; i < di*16+16; i++ {
			d ^= byte(list[i])
		}
		dense[di] = d
	}
	hash := fmt.Sprintf("%x", dense)

	return hash, nil
}
