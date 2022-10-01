package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2020/conway4"
)

func main() {
	scanner := bufio.NewScanner(bytes.NewBufferString(input))
	var rawRows []string
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		rawRows = append(rawRows, l)
	}
	g, err := conway4.NewGrid(rawRows)
	if err != nil {
		panic(err)
	}
	fmt.Printf("*** Cycle 0 ***\n")
	g.Dump()

	numCycles := 6
	for i := 0; i < numCycles; i++ {
		g.DoNext()
		fmt.Printf("\n*** Cycle %d ***\n", i+1)
		g.Dump()
	}
	fmt.Printf("after %d cycles %d cubes are active\n", numCycles, g.ActiveCount())
}

var inputTest = `
.#.
..#
###
`

var input = `
.#.####.
.#...##.
..###.##
#..#.#.#
#..#....
#.####..
##.##..#
#.#.#..#
`
