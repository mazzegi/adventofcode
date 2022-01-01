package day_21

import (
	"adventofcode_2016/errutil"
	"adventofcode_2016/readutil"
	"bytes"
	"fmt"
	"strings"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := scrambled(input, "abcdefgh")
	errutil.ExitOnErr(err)
	log("part1: result = %q", res)
}

func Part2() {
	res, err := unscramble(input, "fbgdceah")
	errutil.ExitOnErr(err)
	log("part2: result = %q", res)
}

//
type operation interface {
	apply(s []byte) []byte
}

type swapPosition struct {
	x, y int
}

type swapLetter struct {
	x, y byte
}

type rotateLeft struct {
	x int
}

type rotateRight struct {
	x int
}

type rotateLetterPosBased struct {
	x byte
}

type reversePositions struct {
	x, y int
}

type movePosition struct {
	x, y int
}

func mustParseOp(in string) operation {
	var iop operation
	var err error
	switch {
	case strings.HasPrefix(in, "swap position"):
		op := swapPosition{}
		_, err = fmt.Sscanf(in, "swap position %d with position %d", &op.x, &op.y)
		iop = op
	case strings.HasPrefix(in, "swap letter"):
		op := swapLetter{}
		var xs, ys string
		_, err = fmt.Sscanf(in, "swap letter %s with letter %s", &xs, &ys)
		op.x = xs[0]
		op.y = ys[0]
		iop = op
	case strings.HasPrefix(in, "rotate left"):
		op := rotateLeft{}
		_, err = fmt.Sscanf(in, "rotate left %d step", &op.x)
		iop = op
	case strings.HasPrefix(in, "rotate right"):
		op := rotateRight{}
		_, err = fmt.Sscanf(in, "rotate right %d step", &op.x)
		iop = op
	case strings.HasPrefix(in, "rotate based on position of letter"):
		op := rotateLetterPosBased{}
		var xs string
		_, err = fmt.Sscanf(in, "rotate based on position of letter %s", &xs)
		op.x = xs[0]
		iop = op
	case strings.HasPrefix(in, "reverse positions"):
		op := reversePositions{}
		_, err = fmt.Sscanf(in, "reverse positions %d through %d", &op.x, &op.y)
		iop = op
	case strings.HasPrefix(in, "move position"):
		op := movePosition{}
		_, err = fmt.Sscanf(in, "move position %d to position %d", &op.x, &op.y)
		iop = op
	default:
		fatal("invalid op %q", in)
	}
	if err != nil {
		fatal("%v", err)
	}
	return iop
}

func mustParseOps(in string) []operation {
	var ops []operation
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		op := mustParseOp(line)
		ops = append(ops, op)
	}
	if len(ops) == 0 {
		fatal("no data")
	}
	return ops
}

//
func clone(bs []byte) []byte {
	cbs := make([]byte, len(bs))
	copy(cbs, bs)
	return cbs
}

func (o swapPosition) apply(in []byte) []byte {
	out := clone(in)
	out[o.x], out[o.y] = out[o.y], out[o.x]
	return out
}

func (o swapLetter) apply(in []byte) []byte {
	out := clone(in)
	ix := bytes.IndexByte(out, o.x)
	iy := bytes.IndexByte(out, o.y)
	out[ix], out[iy] = out[iy], out[ix]
	return out
}

func rotateOneLeft(in []byte) []byte {
	out := clone(in)
	out = append(out[1:], out[0])
	return out
}

func rotateOneRight(in []byte) []byte {
	out := clone(in)
	ix := len(out) - 1
	out = append([]byte{out[ix]}, out[:ix]...)
	return out
}

func (o rotateLeft) apply(in []byte) []byte {
	out := clone(in)
	for i := 0; i < o.x; i++ {
		out = rotateOneLeft(out)
	}
	return out
}

func (o rotateRight) apply(in []byte) []byte {
	out := clone(in)
	for i := 0; i < o.x; i++ {
		out = rotateOneRight(out)
	}
	return out
}

func (o rotateLetterPosBased) apply(in []byte) []byte {
	ix := bytes.IndexByte(in, o.x)
	if ix >= 4 {
		ix = ix + 2
	} else {
		ix = ix + 1
	}
	return rotateRight{ix}.apply(in)
}

func (o reversePositions) apply(in []byte) []byte {
	out := clone(in)
	if o.y < o.x {
		o.x, o.y = o.y, o.x
	}
	mid := (o.y - o.x) / 2
	for i := 0; i <= mid; i++ {
		out[o.x+i], out[o.y-i] = out[o.y-i], out[o.x+i]
	}

	return out
}

func (o movePosition) apply(in []byte) []byte {
	out := clone(in)
	ml := out[o.x]
	out = append(out[:o.x], out[o.x+1:]...)
	tail := clone(out[o.y:])
	out = append(out[:o.y], ml)
	out = append(out, tail...)
	return out
}

func scrambled(in string, start string) (string, error) {
	ops := mustParseOps(in)

	bs := []byte(start)
	for _, op := range ops {
		//bef := clone(bs)
		bs = op.apply(bs)
		//log("%T - %v: %q => %q", op, op, string(bef), string(bs))
	}

	return string(bs), nil
}

func unscramble(in string, res string) (string, error) {
	ops := mustParseOps(in)

	scramble := func(bs []byte) []byte {
		for _, op := range ops {
			bs = op.apply(bs)
		}
		return bs
	}

	//test := []byte("abcdefgh")
	test := []byte("aaaaaaaa")
	min := byte('a')
	max := byte('h')

	containsDupl := func(bs []byte) bool {
		m := map[byte]bool{}
		for _, b := range bs {
			if _, ok := m[b]; ok {
				return true
			}
			m[b] = true
		}
		return false
	}

	next := func() bool {
		for i := 0; i < len(test); i++ {
			if test[i] < max {
				test[i]++
				return true
			} else {
				test[i] = min
			}
		}
		return false
	}

	for {
		ok := next()
		if !ok {
			fatal("all translations tried")
		}
		if containsDupl(test) {
			continue
		}

		log("test: %q", string(test))
		bs := scramble(test)
		if string(bs) == res {
			return string(test), nil
		}
	}
}
