package day_11

import (
	"fmt"
	"math/big"
	"os"
	"runtime/pprof"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
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
	createProfile := false
	if createProfile {
		pf, _ := os.Create("profile")
		defer pf.Close()
		pprof.StartCPUProfile(pf)
		defer pprof.StopCPUProfile()
	}

	res, err := part2MainFunc(inputBig)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type Monkey struct {
	ID             int
	Items          []int
	Operation      func(int) int
	Test           func(int) bool
	ThrowToIfTrue  int
	ThrowToIfFalse int
	Activity       int
}

type BigMonkey struct {
	ID             int
	Items          []*big.Int
	Operation      func(*big.Int) *big.Int
	Test           func(*big.Int) bool
	ThrowToIfTrue  int
	ThrowToIfFalse int
	Activity       int
	cache          *cache
}

func dumpActivity(ms []*BigMonkey) string {
	var sl []string
	for _, m := range ms {
		sl = append(sl, fmt.Sprintf("Monkey %d: %d", m.ID, m.Activity))
	}
	return strings.Join(sl, "\n")
}

func dumpItems(ms []*BigMonkey) string {
	var sl []string
	for _, m := range ms {
		sl = append(sl, fmt.Sprintf("Monkey %d: %v", m.ID, m.Items))
	}
	return strings.Join(sl, "\n")
}

func part1MainFunc(monkeys []*Monkey) (int, error) {
	rounds := 20
	for r := 0; r < rounds; r++ {
		for _, mon := range monkeys {
			for _, item := range mon.Items {
				wl := mon.Operation(item)
				wl = wl / 3
				if mon.Test(wl) {
					monkeys[mon.ThrowToIfTrue].Items = append(monkeys[mon.ThrowToIfTrue].Items, wl)
				} else {
					monkeys[mon.ThrowToIfFalse].Items = append(monkeys[mon.ThrowToIfFalse].Items, wl)
				}
				mon.Activity++
			}
			mon.Items = []int{}
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Activity > monkeys[j].Activity
	})
	return monkeys[0].Activity * monkeys[1].Activity, nil
}

type cache struct {
	ops   map[string]*big.Int
	tests map[string]bool
}

func newCache() *cache {
	return &cache{
		ops:   map[string]*big.Int{},
		tests: map[string]bool{},
	}
}

// func (c *cache) findOp(n *big.Int) (*big.Int, bool) {
// 	if v, ok := c.ops[n.String()]; ok {
// 		return bigClone(v), true
// 	}
// 	return nil, false
// }

// func (c *cache) addOp(n *big.Int) {
// 	c.ops[n.String()] = bigClone(n)
// }

// func (c *cache) findTest(n *big.Int) (bool, bool) {
// 	if v, ok := c.tests[n.String()]; ok {
// 		return v, true
// 	}
// 	return false, false
// }

// func (c *cache) addTest(n *big.Int, v bool) {
// 	c.tests[n.String()] = v
// }

func part2MainFunc(monkeys []*BigMonkey) (int, error) {
	rounds := 10000
	//rounds := 800
	for r := 0; r < rounds; r++ {
		for _, mon := range monkeys {
			for _, item := range mon.Items {
				wl := mon.Operation(item)
				tres := mon.Test(wl)

				if tres {
					monkeys[mon.ThrowToIfTrue].Items = append(monkeys[mon.ThrowToIfTrue].Items, wl)
				} else {
					monkeys[mon.ThrowToIfFalse].Items = append(monkeys[mon.ThrowToIfFalse].Items, wl)
				}
				mon.Activity++
			}
			mon.Items = []*big.Int{}
		}
		log("\nafter round %d\n%s", r, dumpItems(monkeys))
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Activity > monkeys[j].Activity
	})
	return monkeys[0].Activity * monkeys[1].Activity, nil
}
