package day_03

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
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

func mustCommonByte(s1, s2 []byte) byte {
	s2Contains := func(r byte) bool {
		for _, r2 := range s2 {
			if r2 == r {
				return true
			}
		}
		return false
	}
	for _, r1 := range s1 {
		if s2Contains(r1) {
			return byte(r1)
		}
	}
	fatal("no common byte s1=%q, s2=%q", string(s1), string(s2))
	return 0
}

func scoreByte(b byte) int {
	if b >= 97 {
		return int(b) - 97 + 1
	}
	return int(b) - 65 + 27
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var sum int
	for _, line := range lines {
		sz := len(line)
		if sz%2 != 0 {
			fatal("line-size is odd %q: %d", line, sz)
		}
		c1 := []byte(line[:sz/2])
		c2 := []byte(line[sz/2:])
		cb := mustCommonByte(c1, c2)
		cbs := scoreByte(cb)
		sum += cbs
	}
	return sum, nil
}

type group [3]string

func mkSet(bs []byte) map[byte]bool {
	set := map[byte]bool{}
	for _, b := range bs {
		set[b] = true
	}
	return set
}

func (g group) commonByte() byte {
	s1 := mkSet([]byte(g[0]))
	s2 := mkSet([]byte(g[1]))
	s3 := mkSet([]byte(g[2]))
	for b := range s1 {
		if _, ok := s2[b]; !ok {
			continue
		}
		if _, ok := s3[b]; !ok {
			continue
		}
		return b
	}
	fatal("no common byte in group %v", g)
	return 0
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	if len(lines)%3 != 0 {
		fatal("num lines must be a multiple of 3")
	}
	var sum int
	for i := 0; i < len(lines); i += 3 {
		g := group{
			lines[i],
			lines[i+1],
			lines[i+2],
		}
		cb := g.commonByte()
		cbs := scoreByte(cb)
		sum += cbs
	}
	return sum, nil
}
