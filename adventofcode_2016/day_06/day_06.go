package day_06

import (
	"fmt"
	"strings"
)

func Part1() {
	msg := ErrorCorrectedMessage(input)
	fmt.Printf("part1: ecm: %q\n", msg)
}

func Part2() {
	msg := ErrorCorrectedMessage2(input)
	fmt.Printf("part2: ecm: %q\n", msg)
}

func ErrorCorrectedMessage(in string) string {
	sl := strings.Split(in, "\n")
	var msgs []string
	var msgLen int
	for _, s := range sl {
		s = strings.Trim(s, " \r\t\n")
		if s == "" {
			continue
		}
		if len(msgs) == 0 {
			msgLen = len(s)
		} else if len(s) != msgLen {
			fmt.Printf("WARN: inconsistent msg-len %q", s)
			continue
		}
		msgs = append(msgs, s)
	}

	type Frequencies map[rune]int
	mostFreqent := func(f *Frequencies) rune {
		var mf rune
		var mfCount int
		for r, count := range *f {
			if count > mfCount {
				mf = r
				mfCount = count
			}
		}
		return mf
	}

	positions := map[int]*Frequencies{}

	for _, msg := range msgs {
		for i, r := range msg {
			fs, ok := positions[i]
			if !ok {
				fs = &Frequencies{}
				positions[i] = fs
			}
			(*fs)[r]++
		}
	}

	var ecm string
	for i := 0; i < msgLen; i++ {
		r := mostFreqent(positions[i])
		ecm += string(r)
	}

	return ecm
}

func ErrorCorrectedMessage2(in string) string {
	sl := strings.Split(in, "\n")
	var msgs []string
	var msgLen int
	for _, s := range sl {
		s = strings.Trim(s, " \r\t\n")
		if s == "" {
			continue
		}
		if len(msgs) == 0 {
			msgLen = len(s)
		} else if len(s) != msgLen {
			fmt.Printf("WARN: inconsistent msg-len %q", s)
			continue
		}
		msgs = append(msgs, s)
	}

	type Frequencies map[rune]int
	leastFreqent := func(f *Frequencies) rune {
		if f == nil || len(*f) == 0 {
			return ' '
		}

		first := true
		var lf rune
		var lfCount int
		for r, count := range *f {
			if first {
				lf = r
				lfCount = count
				first = false
			} else if count < lfCount {
				lf = r
				lfCount = count
			}
		}
		return lf
	}

	positions := map[int]*Frequencies{}

	for _, msg := range msgs {
		for i, r := range msg {
			fs, ok := positions[i]
			if !ok {
				fs = &Frequencies{}
				positions[i] = fs
			}
			(*fs)[r]++
		}
	}

	var ecm string
	for i := 0; i < msgLen; i++ {
		r := leastFreqent(positions[i])
		ecm += string(r)
	}

	return ecm
}
