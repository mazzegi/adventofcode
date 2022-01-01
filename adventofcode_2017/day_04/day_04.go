package day_04

import (
	"adventofcode_2017/readutil"
	"fmt"
	"sort"
	"strings"
)

func Part1() {
	res := ValidPassphraseCount(input)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res := ValidPassphraseCountExt(input)
	fmt.Printf("part2: result = %d\n", res)
}

func ValidPassphraseCount(in string) int {
	var cnt int
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		if containsNoDuplicateWords(line) {
			cnt++
		}
	}
	return cnt
}

func ValidPassphraseCountExt(in string) int {
	var cnt int
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		if containsNoDuplicateWords(line) && containsNoDuplicateAnagrams(line) {
			cnt++
		}
	}
	return cnt
}

func containsNoDuplicateWords(pp string) bool {
	words := strings.Fields(pp)
	occ := map[string]bool{}
	for _, word := range words {
		_, contains := occ[word]
		if contains {
			return false
		}
		occ[word] = true
	}

	return true
}

func containsNoDuplicateAnagrams(pp string) bool {
	words := strings.Fields(pp)

	anaNorm := func(s string) string {
		rs := []rune(s)
		sort.Slice(rs, func(i, j int) bool {
			return rs[i] < rs[j]
		})
		return string(rs)
	}

	occ := map[string]bool{}
	for _, word := range words {
		wordAnaNormed := anaNorm(word)
		_, contains := occ[wordAnaNormed]
		if contains {
			return false
		}
		occ[wordAnaNormed] = true
	}

	return true
}
