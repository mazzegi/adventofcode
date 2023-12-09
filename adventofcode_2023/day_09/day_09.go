package day_09

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/stringutil"
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

func extrapolate(seq []int) []int {
	redseqs := [][]int{}
	redseqs = append(redseqs, seq)
	prevseq := seq

	for {
		var rseq []int
		allZero := true
		for i := 0; i < len(prevseq)-1; i++ {
			d := prevseq[i+1] - prevseq[i]
			if d != 0 {
				allZero = false
			}
			rseq = append(rseq, d)
		}
		redseqs = append(redseqs, rseq)
		prevseq = rseq
		if allZero {
			break
		}
	}

	//
	i := len(redseqs) - 1
	rseq := redseqs[i]
	rseq = append(rseq, 0)
	for {
		i--
		if i < 0 {
			break
		}
		nextreseq := redseqs[i]
		nextreseq = append(nextreseq, nextreseq[len(nextreseq)-1]+rseq[len(rseq)-1])
		rseq = nextreseq
	}

	return rseq
}

func part1MainFunc(in string) (int, error) {
	var seqs [][]int
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		seq := stringutil.MustStringsToInts(strings.Split(line, " "))
		seqs = append(seqs, seq)
	}

	var sum int
	for _, seq := range seqs {
		eseq := extrapolate(seq)
		sum += eseq[len(eseq)-1]
	}

	return sum, nil
}

func prepend(ns []int, n int) []int {
	return append([]int{n}, ns...)
}

func extrapolateBackwards(seq []int) []int {

	redseqs := [][]int{}
	redseqs = append(redseqs, seq)
	prevseq := seq

	for {
		var rseq []int
		allZero := true
		for i := 0; i < len(prevseq)-1; i++ {
			d := prevseq[i+1] - prevseq[i]
			if d != 0 {
				allZero = false
			}
			rseq = append(rseq, d)
		}
		redseqs = append(redseqs, rseq)
		prevseq = rseq
		if allZero {
			break
		}
	}

	i := len(redseqs) - 1
	rseq := redseqs[i]
	rseq = prepend(rseq, 0)
	for {
		i--
		if i < 0 {
			break
		}
		nextreseq := redseqs[i]
		nextreseq = prepend(nextreseq, nextreseq[0]-rseq[0])

		rseq = nextreseq
	}

	return rseq
}

func part2MainFunc(in string) (int, error) {
	var seqs [][]int
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		seq := stringutil.MustStringsToInts(strings.Split(line, " "))
		seqs = append(seqs, seq)
	}

	var sum int
	for _, seq := range seqs {
		eseq := extrapolateBackwards(seq)
		sum += eseq[0]
	}

	return sum, nil
}
