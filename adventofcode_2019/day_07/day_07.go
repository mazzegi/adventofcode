package day_07

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2019/intcode"
	"github.com/mazzegi/adventofcode/combi"
	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/slices"
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

func part1MainFunc(prg []int) (int, error) {
	var max int
	perms := combi.Permutations([]int{0, 1, 2, 3, 4})
	for _, perm := range perms {
		phases := [5]int{perm[0], perm[1], perm[2], perm[3], perm[4]}
		out := exec(prg, phases)
		if out > max {
			max = out
			//log("new max %d for phases %v", max, phases)
		}
	}
	return max, nil
}

func exec(prg []int, phases [5]int) int {
	in := 0
	amp1Prg := slices.Clone(prg)
	_, outs, err := intcode.Exec2(amp1Prg, []int{phases[0], in})
	errutil.ExitOnErr(err)
	in = outs[len(outs)-1]

	amp2Prg := slices.Clone(prg)
	_, outs, err = intcode.Exec2(amp2Prg, []int{phases[1], in})
	errutil.ExitOnErr(err)
	in = outs[len(outs)-1]

	amp3Prg := slices.Clone(prg)
	_, outs, err = intcode.Exec2(amp3Prg, []int{phases[2], in})
	errutil.ExitOnErr(err)
	in = outs[len(outs)-1]

	amp4Prg := slices.Clone(prg)
	_, outs, err = intcode.Exec2(amp4Prg, []int{phases[3], in})
	errutil.ExitOnErr(err)
	in = outs[len(outs)-1]

	amp5Prg := slices.Clone(prg)
	_, outs, err = intcode.Exec2(amp5Prg, []int{phases[4], in})
	errutil.ExitOnErr(err)
	in = outs[len(outs)-1]
	return in
}

func part2MainFunc(prg []int) (int, error) {
	var max int
	perms := combi.Permutations([]int{5, 6, 7, 8, 9})
	for _, perm := range perms {
		phases := [5]int{perm[0], perm[1], perm[2], perm[3], perm[4]}
		out := execC(prg, phases)
		if out > max {
			//log("new max %d => %v", max, phases)
			max = out
		}
	}
	return max, nil
}

func execC(prg []int, phases [5]int) int {
	a1 := newAmp(prg, phases[0])
	a2 := newAmp(prg, phases[1])
	a3 := newAmp(prg, phases[2])
	a4 := newAmp(prg, phases[3])
	a5 := newAmp(prg, phases[4])

	a1.wireTo(a2)
	a2.wireTo(a3)
	a3.wireTo(a4)
	a4.wireTo(a5)
	a5.wireTo(a1)

	a1.startWithInput(0)
	a2.start()
	a3.start()
	a4.start()
	a5.start()

	var lastA5Out int
outer:
	for {
		select {
		case v, ok := <-a1.outC:
			if !ok {
				continue
			}
			a2.inC <- v
		case v, ok := <-a2.outC:
			if !ok {
				continue
			}
			a3.inC <- v
		case v, ok := <-a3.outC:
			if !ok {
				continue
			}
			a4.inC <- v
		case v, ok := <-a4.outC:
			if !ok {
				continue
			}
			a5.inC <- v
		case v, ok := <-a5.outC:
			if !ok {
				break outer
			}
			lastA5Out = v
			a1.inC <- v
		}
	}
	return lastA5Out
}

type amp struct {
	prg     []int
	phase   int
	inC     chan int
	outC    chan int
	nextAmp *amp
}

func newAmp(prg []int, phase int) *amp {
	return &amp{
		prg:   slices.Clone(prg),
		phase: phase,
		inC:   make(chan int, 2),
		outC:  make(chan int, 2),
	}
}

func (a *amp) wireTo(oa *amp) {
	a.nextAmp = oa
}

func (a *amp) start() {
	go func() {
		intcode.ExecC(a.prg, a.inC, a.outC)
		close(a.outC)
	}()
	a.inC <- a.phase
}

func (a *amp) startWithInput(in int) {
	a.start()
	a.inC <- in
}

func (a *amp) sel() {
	v := <-a.outC
	a.nextAmp.inC <- v
}
