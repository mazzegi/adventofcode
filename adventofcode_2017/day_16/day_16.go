package day_16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := danceOrder(input, 16)
	errutil.ExitOnErr(err)
	log("part1: result = %q", res)
}

func Part2() {
	loops := 1000000000
	res, err := after(input, 16, loops, true)
	errutil.ExitOnErr(err)
	log("part2: result = %q", res)
}

type move interface{}

type spin struct {
	size int
}

type exchange struct {
	pos1, pos2 int
}

type partner struct {
	p1, p2 byte
}

func parseMove(s string) (move, error) {
	switch {
	case strings.HasPrefix(s, "s"):
		n, err := strconv.ParseInt(s[1:], 10, 8)
		if err != nil {
			return nil, err
		}
		return spin{int(n)}, nil
	case strings.HasPrefix(s, "x"):
		sl := strings.Split(s[1:], "/")
		if len(sl) != 2 {
			return nil, errors.Errorf("invalid exchange %q", s)
		}
		pos1, err := strconv.ParseInt(sl[0], 10, 8)
		if err != nil {
			return nil, err
		}
		pos2, err := strconv.ParseInt(sl[1], 10, 8)
		if err != nil {
			return nil, err
		}
		return exchange{int(pos1), int(pos2)}, nil
	case strings.HasPrefix(s, "p"):
		if len(s) != 4 {
			return nil, errors.Errorf("invalid partner %q", s)
		}
		return partner{s[1], s[3]}, nil
	default:
		return nil, errors.Errorf("invalid move %q", s)
	}
}

func parseMoves(in string) ([]move, error) {
	in = readutil.ReadString(in)
	sl := strings.Split(in, ",")
	var mvs []move
	for _, s := range sl {
		mv, err := parseMove(s)
		if err != nil {
			return nil, err
		}
		mvs = append(mvs, mv)
	}
	if len(mvs) == 0 {
		return nil, errors.Errorf("no data")
	}
	return mvs, nil
}

func danceOrder(inMoves string, inCount int) (string, error) {
	mvs, err := parseMoves(inMoves)
	if err != nil {
		return "", err
	}

	prgs := make([]byte, inCount)
	ba := byte('a')
	for i := 0; i < inCount; i++ {
		prgs[i] = ba + byte(i)
	}
	log("initial: %q", string(prgs))

	mustFind2 := func(b1, b2 byte) (pos1, pos2 int) {
		pos1, pos2 = -1, -1
		for i, b := range prgs {
			if b == b1 {
				pos1 = i
			}
			if b == b2 {
				pos2 = i
			}
			if pos1 >= 0 && pos2 >= 0 {
				return pos1, pos2
			}
		}
		panic("didn't find 2")
	}

	for _, mv := range mvs {
		switch mv := mv.(type) {
		case spin:
			tail := prgs[len(prgs)-mv.size:]
			prgs = append(tail, prgs[:len(prgs)-mv.size]...)
		case exchange:
			prgs[mv.pos1], prgs[mv.pos2] = prgs[mv.pos2], prgs[mv.pos1]
		case partner:
			pos1, pos2 := mustFind2(mv.p1, mv.p2)
			prgs[pos1], prgs[pos2] = prgs[pos2], prgs[pos1]
		default:
			return "", errors.Errorf("invalid move %T", mv)
		}
	}

	return string(prgs), nil
}

func after(inMoves string, inCount int, loops int, useCaching bool) (string, error) {
	mvs, err := parseMoves(inMoves)
	if err != nil {
		return "", err
	}

	prgs := make([]byte, inCount)
	ba := byte('a')
	for i := 0; i < inCount; i++ {
		prgs[i] = ba + byte(i)
	}
	log("initial: %q", string(prgs))

	mustFind2 := func(b1, b2 byte) (pos1, pos2 int) {
		pos1, pos2 = -1, -1
		for i, b := range prgs {
			if b == b1 {
				pos1 = i
			}
			if b == b2 {
				pos2 = i
			}
			if pos1 >= 0 && pos2 >= 0 {
				return pos1, pos2
			}
		}
		panic("didn't find 2")
	}

	//loops := 1000000000
	first := string(prgs)
	hashMap := map[string]int{}
	hashMap[string(prgs)] = 0

	for i := 0; i < loops; i++ {
		for _, mv := range mvs {
			switch mv := mv.(type) {
			case spin:
				tail := prgs[len(prgs)-mv.size:]
				prgs = append(tail, prgs[:len(prgs)-mv.size]...)
			case exchange:
				prgs[mv.pos1], prgs[mv.pos2] = prgs[mv.pos2], prgs[mv.pos1]
			case partner:
				pos1, pos2 := mustFind2(mv.p1, mv.p2)
				prgs[pos1], prgs[pos2] = prgs[pos2], prgs[pos1]
			default:
				return "", errors.Errorf("invalid move %T", mv)
			}
		}

		if useCaching {
			if string(prgs) == first {
				log("seen first again (%d): %q", i, string(prgs))
				period := i + 1
				i += (i + 1) * ((loops / period) - 1)
			}
		}
		// if pn, ok := hashMap[string(prgs)]; ok {
		// 	log("seen twice (%d / %d): %q", pn, i, string(prgs))
		// } else {
		// 	hashMap[string(prgs)] = i + 1
		// }

		// if i%10000 == 0 {
		// 	log("loop %d (%.2f)", i, float64(i)/float64(loops)*100.0)
		// }
	}

	return string(prgs), nil
}
