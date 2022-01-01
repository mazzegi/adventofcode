package cond

import (
	"fmt"
	"sort"
	"strings"
)

type Condition struct {
	ID       string
	Requires []string
	Done     bool
}

func (c *Condition) IsAvailable(cs Conditions) bool {
	if c.Done {
		return false
	}
	for _, req := range c.Requires {
		if rc, ok := cs[req]; ok {
			if !rc.Done {
				return false
			}
		} else {
			panic(fmt.Sprintf("no such condition %q", req))
		}
	}
	return true
}

type Conditions map[string]*Condition

func (cs Conditions) Available() []*Condition {
	var acs []*Condition
	for _, c := range cs {
		if c.IsAvailable(cs) {
			acs = append(acs, c)
		}
	}
	return acs
}

func (cs Conditions) AllDone() bool {
	for _, c := range cs {
		if !c.Done {
			return false
		}
	}
	return true
}

func Evaluate(conds Conditions) {
	var seq []string
	for !conds.AllDone() {
		acs := conds.Available()
		if len(acs) == 0 {
			panic("not all done but no available")
		}
		sort.Slice(acs, func(i, j int) bool {
			return acs[i].ID < acs[j].ID
		})
		acs[0].Done = true
		seq = append(seq, acs[0].ID)
	}
	fmt.Printf("the sequence is %q\n", strings.Join(seq, ""))
}

var a = int(byte('A'))

func StepDuration(cid string, offset int) int {
	return offset + int(cid[0]) - a + 1
}

type Worker struct {
	Cond  *Condition
	Start int
	Dur   int
}

func EvaluateParallel(conds Conditions, workers int, stepOffset int) {
	var seq []string
	sec := 0
	busyWorkers := map[string]Worker{}

	dumpWorkers := func() string {
		var sl []string
		for _, w := range busyWorkers {
			sl = append(sl, fmt.Sprintf("%s", w.Cond.ID))
		}
		return strings.Join(sl, ", ")
	}

	nowAvailable := func() []*Condition {
		cs := conds.Available()
		ncs := []*Condition{}
		//check if workers are currently working
		for _, c := range cs {
			if _, busy := busyWorkers[c.ID]; !busy {
				ncs = append(ncs, c)
			}
		}
		sort.Slice(ncs, func(i, j int) bool {
			return ncs[i].ID < ncs[j].ID
		})
		return ncs
	}

	for !conds.AllDone() {
		//check workers to be finished
		for _, w := range busyWorkers {
			if sec-w.Start >= w.Dur {
				//this one finished
				w.Cond.Done = true
				seq = append(seq, w.Cond.ID)
				delete(busyWorkers, w.Cond.ID)
			}
		}

		acs := nowAvailable()
		for _, ac := range acs {
			if workers-len(busyWorkers) > 0 {
				w := Worker{
					Cond:  ac,
					Start: sec,
					Dur:   StepDuration(ac.ID, stepOffset),
				}
				busyWorkers[w.Cond.ID] = w
			} else {
				break
			}
		}
		fmt.Printf("%d: %s => %q\n", sec, dumpWorkers(), strings.Join(seq, ""))
		sec++
	}
	fmt.Printf("the sequence is %q\n", strings.Join(seq, ""))
}
