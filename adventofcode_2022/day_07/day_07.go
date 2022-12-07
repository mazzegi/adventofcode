package day_07

import (
	"fmt"
	"sort"

	"github.com/mazzegi/adventofcode/adventofcode_2022/elvefs"
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

func part1MainFunc(in string) (int, error) {
	fs := elvefs.New()
	err := fs.Process(readutil.ReadLines(in))
	if err != nil {
		return 0, err
	}
	var sum int
	fs.Walk(func(fi elvefs.FileInfo) {
		if fi.Type != elvefs.TypeDir {
			return
		}
		if fi.TotalSize <= 100000 {
			sum += fi.TotalSize
		}
	})

	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	fs := elvefs.New()
	err := fs.Process(readutil.ReadLines(in))
	if err != nil {
		return 0, err
	}
	total := 70000000
	unusedNeeded := 30000000
	used := fs.Used()
	unused := total - used
	toFree := unusedNeeded - unused

	var candidates []elvefs.FileInfo
	fs.Walk(func(fi elvefs.FileInfo) {
		if fi.Type != elvefs.TypeDir {
			return
		}
		if fi.TotalSize >= toFree {
			candidates = append(candidates, fi)
		}
	})
	if len(candidates) == 0 {
		return 0, fmt.Errorf("no candidates found")
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].TotalSize < candidates[j].TotalSize
	})
	return candidates[0].TotalSize, nil
}
