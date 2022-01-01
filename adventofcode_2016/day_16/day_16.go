package day_16

import (
	"adventofcode_2016/errutil"
	"fmt"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := checksum([]byte("11101000110010100"), 272)
	errutil.ExitOnErr(err)
	log("part1: result = %q", res)
}

func Part2() {
	res, err := checksum([]byte("11101000110010100"), 35651584)
	errutil.ExitOnErr(err)
	log("part2: result = %q", res)
}

//
func clone(bs []byte) []byte {
	cbs := make([]byte, len(bs))
	copy(cbs, bs)
	return cbs
}

func reverse(bs []byte) {
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}
}

func dragon(a []byte) []byte {
	b := clone(a)
	reverse(b)
	for i, x := range b {
		if x == '0' {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}
	res := clone(a)
	res = append(res, '0')
	res = append(res, b...)
	return res
}

func cs(a []byte) []byte {
	if len(a)%2 != 0 {
		fatal("cannot calc checksum of odd len")
	}
	var csb []byte
	for i := 0; i < len(a); i += 2 {
		if a[i] == a[i+1] {
			csb = append(csb, '1')
		} else {
			csb = append(csb, '0')
		}
	}
	if len(csb)%2 == 0 {
		//even
		return cs(csb)
	}

	return csb
}

func checksum(initial []byte, size int) (string, error) {
	a := initial
	for len(a) < size {
		a = dragon(a)
	}
	a = a[:size]
	res := cs(a)
	return string(res), nil
}
