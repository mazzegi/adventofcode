package day_15

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
)

type genValues struct {
	AStart uint64
	AFac   uint64
	AMult  uint64
	BStart uint64
	BFac   uint64
	BMult  uint64
}

var inputGenValues = genValues{
	AStart: 679,
	BStart: 771,
	AFac:   16807,
	BFac:   48271,
	AMult:  4,
	BMult:  8,
}
var testGenValues = genValues{
	AStart: 65,
	BStart: 8921,
	AFac:   16807,
	BFac:   48271,
	AMult:  4,
	BMult:  8,
}

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

type generator struct {
	fac  uint64
	prev uint64
	mult uint64
}

func newGenerator(fac uint64, start uint64, mult uint64) *generator {
	return &generator{
		fac:  fac,
		prev: start,
		mult: mult,
	}
}

func (g *generator) nextIntern() uint64 {
	n := g.prev * g.fac
	v := n % 2147483647
	g.prev = v
	return v
}

func (g *generator) next() uint64 {
	for {
		v := g.nextIntern()
		if v%g.mult == 0 {
			return v
		}
	}
}

const Mio40 = 40e6
const Mio5 = 5e6

func Part1() {
	res, err := judgeCount(inputGenValues, Mio40)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := judgeCountMult(inputGenValues, Mio5)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func judgeCount(vs genValues, count int) (int, error) {
	genA := newGenerator(vs.AFac, vs.AStart, 1)
	genB := newGenerator(vs.BFac, vs.BStart, 1)

	var totalMatches int
	for i := 0; i < count; i++ {
		vA := genA.next()
		vB := genB.next()

		tvA := uint16(vA)
		tvB := uint16(vB)
		if tvA == tvB {
			totalMatches++
		}
		if i%100000 == 0 {
			log("%d steps (%d matches)", i, totalMatches)
		}
	}
	log("total matches: %d", totalMatches)

	return totalMatches, nil
}

func judgeCountMult(vs genValues, count int) (int, error) {
	genA := newGenerator(vs.AFac, vs.AStart, vs.AMult)
	genB := newGenerator(vs.BFac, vs.BStart, vs.BMult)

	var totalMatches int
	for i := 0; i < count; i++ {
		vA := genA.next()
		vB := genB.next()

		tvA := uint16(vA)
		tvB := uint16(vB)
		if tvA == tvB {
			totalMatches++
		}
		if i%100000 == 0 {
			log("%d steps (%d matches)", i, totalMatches)
		}
	}
	log("total matches: %d", totalMatches)

	return totalMatches, nil
}
