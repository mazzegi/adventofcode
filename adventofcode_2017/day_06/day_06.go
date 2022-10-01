package day_06

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2017/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/readutil"
	"github.com/mazzegi/adventofcode/adventofcode_2017/stringutil"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := distCyclesToRepeat(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := loopSize(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

type memory struct {
	banks []int
}

func (m *memory) cycle() {
	// find most
	mostIdx := 0
	most := m.banks[0]
	for i, v := range m.banks {
		if v > most {
			mostIdx = i
			most = v
		}
	}

	//reset mostIdx
	m.banks[mostIdx] = 0

	//
	amount := most
	idx := mostIdx + 1
	for amount > 0 {
		if idx >= len(m.banks) {
			idx = 0
		}
		m.banks[idx]++
		amount -= 1
		idx++
	}
}

func (m memory) hash() string {
	h := md5.New()
	for i, b := range m.banks {
		err := binary.Write(h, binary.LittleEndian, int64(b+i))
		if err != nil {
			fmt.Printf("hash error: %v", err)
		}
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func distCyclesToRepeat(in string) (int, error) {
	ns, err := readutil.ReadInts(in, " ")
	if err != nil {
		return 0, errors.Wrapf(err, "read-ints %q", in)
	}
	if len(ns) == 0 {
		return 0, errors.Errorf("no entries")
	}

	maxCycles := 1024 * 1024
	hashes := stringutil.NewSet()

	mem := &memory{banks: ns}
	hashes.Insert(mem.hash())
	for i := 0; i < maxCycles; i++ {
		mem.cycle()
		h := mem.hash()
		if hashes.Contains(h) {
			return i + 1, nil
		}
		hashes.Insert(h)
	}

	return 0, errors.Errorf("max-cycels exceeded %d", maxCycles)
}

func loopSize(in string) (int, error) {
	ns, err := readutil.ReadInts(in, " ")
	if err != nil {
		return 0, errors.Wrapf(err, "read-ints %q", in)
	}
	if len(ns) == 0 {
		return 0, errors.Errorf("no entries")
	}

	maxCycles := 1024 * 1024
	hashes := map[string]int{}
	mem := &memory{banks: ns}
	hashes[mem.hash()] = 0
	for i := 0; i < maxCycles; i++ {
		mem.cycle()
		h := mem.hash()
		if idx, ok := hashes[h]; ok {
			return i - idx, nil
		}
		hashes[h] = i
	}

	return 0, errors.Errorf("max-cycels exceeded %d", maxCycles)
}
