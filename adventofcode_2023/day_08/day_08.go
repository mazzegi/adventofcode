package day_08

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(inputSeq, inputNodes)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(inputSeq, inputNodes)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

type Node struct {
	ID      string
	LeftID  string
	RightID string
}

func mustParseNode(s string) Node {
	s = strings.ReplaceAll(s, "= (", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, ",", "")

	sl := strings.Split(s, " ")
	if len(sl) != 3 {
		panic("invalid node-string: " + s)
	}
	return Node{
		ID:      sl[0],
		LeftID:  sl[1],
		RightID: sl[2],
	}
}

func mustParseNodes(sl []string) []Node {
	var ns []Node
	for _, s := range sl {
		ns = append(ns, mustParseNode(s))
	}
	return ns
}

func part1MainFunc(inSeq string, inNodes string) (int, error) {
	nodes := mustParseNodes(readutil.ReadLines(inNodes))
	mustFindNode := func(id string) Node {
		for _, n := range nodes {
			if n.ID == id {
				return n
			}
		}
		panic("not-found: " + id)
	}

	var stepsToZZZ int
	i := 0
	node := mustFindNode("AAA")

	for {
		var nextID string
		switch inSeq[i] {
		case 'L':
			nextID = node.LeftID
		case 'R':
			nextID = node.RightID
		default:
			panic("fooled")
		}
		stepsToZZZ++
		i++
		if i >= len(inSeq) {
			i = 0
		}

		if nextID == "ZZZ" {
			break
		}
		node = mustFindNode(nextID)
	}

	return stepsToZZZ, nil
}

func part2MainFunc(inSeq string, inNodes string) (int, error) {
	return 0, nil
}
