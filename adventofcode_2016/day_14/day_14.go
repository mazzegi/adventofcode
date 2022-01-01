package day_14

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func Part1() {
	salt := "cuanljph"
	kis := keyIndexes(salt, 64, hashSalted)
	fmt.Printf("last: %d\n", kis[63])
}

func Part2() {
	salt := "cuanljph"
	kis := keyIndexes(salt, 64, hash2016)
	fmt.Printf("last: %d\n", kis[63])
}

func hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func hashSalted(salt string, idx int) string {
	return hash(fmt.Sprintf("%s%d", salt, idx))
}

func hash2016(salt string, idx int) string {
	h := hashSalted(salt, idx)
	for i := 0; i < 2016; i++ {
		h = hash(h)
	}
	return h
}

func tripletChars(s string) []byte {
	var tcs []byte
	var last byte
	var lastCount int
	first := true
	for _, c := range []byte(s) {
		if first {
			last = c
			lastCount = 1
			first = false
			continue
		}
		if c != last {
			last = c
			lastCount = 1
			continue
		}
		lastCount++
		if lastCount == 3 {
			tcs = append(tcs, last)
		}
	}

	return tcs
}

func keyIndexes(salt string, num int, hashFunc func(string, int) string) []int {
	var keyIdxs []int
	type candidate struct {
		char    byte
		atIdx   int
		dropped bool
	}

	var candidates []candidate
	var idx int = 0
	for {
		//h :=  hashSalted(salt, idx)
		h := hashFunc(salt, idx)
		tcs := tripletChars(h)
		if len(tcs) == 0 {
			idx++
			continue
		}

		// check candidates
		for ic, cand := range candidates {
			if cand.dropped {
				continue
			}

			dropCand := false
			if idx-cand.atIdx > 1000 {
				dropCand = true
			} else if strings.Contains(h, strings.Repeat(string(cand.char), 5)) {
				keyIdxs = append(keyIdxs, cand.atIdx)
				fmt.Printf("found key at %d (%d total)\n", cand.atIdx, len(keyIdxs))
				if len(keyIdxs) == num {
					sort.Ints(keyIdxs)
					return keyIdxs
				}
				dropCand = true
			}

			if dropCand {
				candidates[ic].dropped = true
			}
		}

		candidates = append(candidates, candidate{
			char:    tcs[0],
			atIdx:   idx,
			dropped: false,
		})
		idx++
	}
}
