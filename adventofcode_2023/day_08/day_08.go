package day_08

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/euler"
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
	nodes := mustParseNodes(readutil.ReadLines(inNodes))
	mustFindNode := func(id string) Node {
		for _, n := range nodes {
			if n.ID == id {
				return n
			}
		}
		panic("not-found: " + id)
	}

	// lookup all, ending with "A"
	var itNodes []Node
	for _, n := range nodes {
		if strings.HasSuffix(n.ID, "A") {
			itNodes = append(itNodes, n)
		}
	}

	allEndWithZ := func() bool {
		for _, n := range itNodes {
			if !strings.HasSuffix(n.ID, "Z") {
				return false
			}
		}
		return true
	}
	_ = allEndWithZ

	type moment struct {
		id    string
		itSeq int
	}
	determineOffsetAndPeriod := func(startNode Node) (offset, period int, znodes []int) {
		var itSeq int = 0
		var hist []moment

		// iterate until we are ending for the same value of itSeq ate the same value for the node
		currNode := startNode
		hist = append(hist, moment{id: currNode.ID, itSeq: itSeq})
		for {
			switch inSeq[itSeq] {
			case 'L':
				currNode = mustFindNode(currNode.LeftID)
			case 'R':
				currNode = mustFindNode(currNode.RightID)
			default:
				panic("fooled")
			}

			itSeq++
			if itSeq >= len(inSeq) {
				itSeq = 0
			}

			cm := moment{id: currNode.ID, itSeq: itSeq}
			for hi, hm := range hist {
				if hm == cm {
					for zi, hn := range hist {
						if strings.HasSuffix(hn.id, "Z") {
							znodes = append(znodes, zi)
						}
					}

					return hi, len(hist) - hi, znodes
				}
			}
			hist = append(hist, cm)
		}
	}

	var periods []int
	for _, sn := range itNodes {
		o, p, zns := determineOffsetAndPeriod(sn)
		periods = append(periods, p)
		log("%s: %d, %d (%v)", sn.ID, o, p, zns)
	}

	probeN := func(startNode Node, num int) {
		var itSeq int = 0
		currNode := startNode
		cnt := 0
		itTotal := 0
		for {
			switch inSeq[itSeq] {
			case 'L':
				currNode = mustFindNode(currNode.LeftID)
			case 'R':
				currNode = mustFindNode(currNode.RightID)
			default:
				panic("fooled")
			}

			itTotal++
			itSeq++
			if itSeq >= len(inSeq) {
				itSeq = 0
			}

			if strings.HasSuffix(currNode.ID, "Z") {
				log("at %d", itTotal)
				cnt++
				if cnt >= num {
					return
				}
			}
		}
	}

	// for _, sn := range itNodes {
	// 	log("probe: %s", sn.ID)
	// 	probeN(sn, 3)
	// }
	_ = probeN
	sm := euler.SmallestMultipleOf(periods...)
	return sm, nil

	// i := 0
	// for {
	// 	if allEndWithZ() {
	// 		break
	// 	}
	// 	for j, n := range itNodes {
	// 		switch inSeq[i] {
	// 		case 'L':
	// 			itNodes[j] = mustFindNode(n.LeftID)
	// 		case 'R':
	// 			itNodes[j] = mustFindNode(n.RightID)
	// 		default:
	// 			panic("fooled")
	// 		}
	// 	}

	// 	stepsToAllZ++
	// 	i++
	// 	if i >= len(inSeq) {
	// 		i = 0
	// 	}

	// }

}
