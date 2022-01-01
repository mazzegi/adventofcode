package day_17

import (
	"adventofcode_2017/errutil"
	"fmt"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

const (
	inputStep         = 394
	inputInserts      = 2017
	inputInsertsPart2 = 50000000
)

func Part1() {
	res, err := valAfterLastInsert(inputStep, inputInserts)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := valAfter0(inputStep, inputInsertsPart2)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
type node struct {
	value int
	next  *node
}

type buffer struct {
	first *node
	curr  *node
}

func newBuffer(init int) *buffer {
	n := &node{
		value: init,
	}
	n.next = n
	return &buffer{
		first: n,
		curr:  n,
	}
}

func (b *buffer) advanceCurr(steps int) {
	for i := 0; i < steps; i++ {
		b.curr = b.curr.next
	}
}

func (b *buffer) insertAfterCurrent(v int) {
	n := &node{
		value: v,
	}
	n.next = b.curr.next
	b.curr.next = n
	b.curr = n
}

func (b *buffer) valAfter(v int) int {
	n := b.curr
	for {
		if n.value == v {
			break
		}
		n = n.next
		if n == b.curr {
			panic("circ closed")
		}
	}
	return n.next.value
}

func valAfterLastInsert(step int, inserts int) (int, error) {
	buf := newBuffer(0)
	for i := 1; i <= inserts; i++ {
		buf.advanceCurr(step)
		buf.insertAfterCurrent(i)
	}
	val := buf.curr.next.value

	return val, nil
}

func valAfter0(step int, inserts int) (int, error) {
	buf := newBuffer(0)
	for i := 1; i <= inserts; i++ {
		buf.advanceCurr(step)
		buf.insertAfterCurrent(i)

		if i%10000 == 0 {
			log("after step %d / %d => %.2f %%", i, inserts, float64(i)/float64(inserts)*100.0)
		}
	}
	val := buf.valAfter(0)

	return val, nil
}
