package day_11_v2

import (
	"fmt"
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

}

func Part2() {
	createProfile := false
	if createProfile {
		pf, _ := os.Create("profile")
		defer pf.Close()
		pprof.StartCPUProfile(pf)
		defer pprof.StopCPUProfile()
	}

	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//

func MultBy(n int64) func(x int64) int64 {
	return func(x int64) int64 {
		return x * n
	}
}

func Add(n int64) func(x int64) int64 {
	return func(x int64) int64 {
		return x + n
	}
}

func Square(x int64) int64 {
	return x * x
}

func TestDivBy(n int64) func(x int64) bool {
	return func(x int64) bool {
		return x%n == 0
	}
}

type InputMonkey struct {
	Items          []int
	Operation      func(int64) int64
	TestDivBy      int
	ThrowToIfTrue  int
	ThrowToIfFalse int
}

type Item struct {
	ID       string
	Value    int64
	AtMonkey int
	History  []int
}

type Monkey struct {
	ID             int
	Items          []*Item
	Operation      func(int64) int64
	Test           func(int64) bool
	TestDivBy      int
	ThrowToIfTrue  int
	ThrowToIfFalse int
	Activity       int
}

func (m Monkey) ItemIDs() []string {
	var ids []string
	for _, item := range m.Items {
		ids = append(ids, item.ID)
	}
	sort.Strings(ids)
	return ids
}

func dumpActivity(ms []*Monkey) string {
	var sl []string
	for _, m := range ms {
		sl = append(sl, fmt.Sprintf("Monkey %d: %d", m.ID, m.Activity))
	}
	return strings.Join(sl, "\n")
}

func dumpItemHistory(ms []*Monkey) string {
	var sl []string
	for _, m := range ms {
		for _, item := range m.Items {
			sl = append(sl, fmt.Sprintf("%s: %v", item.ID, item.History))
		}
	}
	sort.Strings(sl)
	return strings.Join(sl, "\n")
}

func dumpItems(ms []*Monkey) string {
	var sl []string
	for _, m := range ms {
		sl = append(sl, fmt.Sprintf("Monkey %d (%d): %v", m.ID, m.Activity, m.ItemIDs()))
	}
	return strings.Join(sl, "\n")
}

func part2MainFunc(in []InputMonkey) (int, error) {
	//Setup monkeys
	monkeys := []*Monkey{}
	var pMod int64 = 1
	for mid, im := range in {
		m := &Monkey{
			ID:             mid,
			Operation:      im.Operation,
			Test:           TestDivBy(int64(im.TestDivBy)),
			TestDivBy:      im.TestDivBy,
			ThrowToIfTrue:  im.ThrowToIfTrue,
			ThrowToIfFalse: im.ThrowToIfFalse,
			Activity:       0,
		}
		for iid, iv := range im.Items {
			m.Items = append(m.Items, &Item{
				ID:       fmt.Sprintf("m%d:%d", mid, iid),
				Value:    int64(iv),
				AtMonkey: mid,
				History:  []int{mid},
			})
		}
		pMod *= int64(im.TestDivBy)
		monkeys = append(monkeys, m)
	}

	rounds := 10000
	//rounds := 800
	for r := 0; r < rounds; r++ {
		for _, mon := range monkeys {
			for _, item := range mon.Items {

				item.Value = mon.Operation(item.Value)
				item.Value = item.Value % pMod

				tres := mon.Test(item.Value)
				//tres := item.Value.Int64() == 0

				var moveTo int
				if tres {
					moveTo = mon.ThrowToIfTrue
				} else {
					moveTo = mon.ThrowToIfFalse
				}
				monkeys[moveTo].Items = append(monkeys[moveTo].Items, item)
				item.AtMonkey = moveTo
				item.History = append(item.History, moveTo)
				mon.Activity++
			}
			mon.Items = []*Item{}
		}
		// log("\nafter round %d\n%s", r, dumpItems(monkeys))
		// if r == 100 {
		// 	//log("after round 19\n%s", dumpActivity(monkeys))
		// 	log("%s", dumpItemHistory(monkeys))
		// }
		if r == 19 {
			log("%s", dumpActivity(monkeys))
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Activity > monkeys[j].Activity
	})
	return monkeys[0].Activity * monkeys[1].Activity, nil
}
