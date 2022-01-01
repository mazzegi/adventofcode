package day_19

import (
	"adventofcode_2016/errutil"
	"fmt"
	"strings"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

// func fatal(pattern string, args ...interface{}) {
// 	panic(fmt.Sprintf(pattern+"\n", args...))
// }

func Part1() {
	res, err := winningElf(3004953)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := winningElfSteelAcross(3004953)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
type circBufferElt struct {
	value         int
	next          *circBufferElt
	prevOfAccross *circBufferElt
}

type circBuffer struct {
	first *circBufferElt
	last  *circBufferElt
	size  int
}

func (cb *circBuffer) addTail(value int) {
	cb.size++
	if cb.first == nil {
		elt := &circBufferElt{value, nil, nil}
		elt.next = elt
		cb.first = &circBufferElt{value, elt, nil}
		cb.last = elt
		return
	}

	elt := &circBufferElt{value, cb.first, nil}
	if cb.first.value == cb.last.value {
		cb.last = elt
		cb.first.next = elt
		return
	}

	cb.last.next = elt
	cb.last = elt
}

func (cb *circBuffer) removeNextOf(elt *circBufferElt) {
	next := elt.next
	non := next.next
	elt.next = non
	if cb.first == next {
		cb.first = non
	}
	if cb.last == next {
		cb.last = non
	}
	cb.size--
}

func (cb *circBuffer) dump() string {
	var sl []string
	elt := cb.first
	for {
		sl = append(sl, fmt.Sprintf("%d", elt.value))
		elt = elt.next
		if elt == cb.first {
			break
		}
	}
	return strings.Join(sl, ", ")
}

func winningElf(numElves int) (int, error) {
	cb := &circBuffer{}
	for i := 1; i <= numElves; i++ {
		cb.addTail(i)
	}

	curr := cb.first
	for cb.size > 1 {
		// steel
		cb.removeNextOf(curr)
		curr = curr.next
	}
	log("last: %d", cb.first.value)

	return 0, nil
}

func winningElfSteelAcross(numElves int) (int, error) {
	cb := &circBuffer{}
	for i := 1; i <= numElves; i++ {
		cb.addTail(i)
	}
	_ = cb.dump

	curr := cb.first
	prevOfSteal := curr
	for i := 0; i < cb.size/2-1; i++ {
		prevOfSteal = prevOfSteal.next
	}

	//log("curr = %d, pos = %d", curr.value, prevOfSteal.value)
	iter := 1
	for cb.size > 1 {
		// steal
		//log("curr = %d, pos = %d", curr.value, prevOfSteal.value)
		cb.removeNextOf(prevOfSteal)

		curr = curr.next

		if iter == 1 {
			prevOfSteal = prevOfSteal.next
		} else if iter%2 == 1 {
			prevOfSteal = prevOfSteal.next
		}

		iter++
		if iter%1000 == 0 {
			log("step %d", iter)
		}
		// log("curr = %d, pos = %d", curr.value, prevOfSteal.value)
	}
	log("last: %d", cb.first.value)

	return cb.first.value, nil
}

// func winningElfSteelAcross(numElves int) (int, error) {
// 	cb := &circBuffer{}
// 	for i := 1; i <= numElves; i++ {
// 		cb.addTail(i)
// 	}

// 	curr := cb.first

// 	//log("curr = %d, pos = %d", curr.value, prevOfSteal.value)

// 	for cb.size > 1 {
// 		// steal
// 		prevOfSteal := curr
// 		for i := 0; i < cb.size/2-1; i++ {
// 			prevOfSteal = prevOfSteal.next
// 		}

// 		log("curr = %d, pos = %d", curr.value, prevOfSteal.value)
// 		cb.removeNextOf(prevOfSteal)

// 		curr = curr.next
// 		// prevOfSteal = prevOfSteal.next
// 		// log("curr = %d, pos = %d", curr.value, prevOfSteal.value)
// 	}
// 	log("last: %d", cb.first.value)

// 	return cb.first.value, nil
// }

//

// type circBufferSlice struct {
// 	elts []int
// }

// func (cb *circBufferSlice) addTail(value int) {
// 	cb.elts = append(cb.elts, value)
// }

// func (cb *circBufferSlice) removeAt(idx int) {
// 	cb.elts = append(cb.elts[:idx], cb.elts[idx+1:]...)
// }

// func (cb *circBufferSlice) size() int {
// 	return len(cb.elts)
// }

// func winningElfSteelAcross(numElves int) (int, error) {
// 	cb := &circBufferSlice{}
// 	for i := 1; i <= numElves; i++ {
// 		cb.addTail(i)
// 	}

// 	currIdx := 0
// 	for cb.size() > 1 {
// 		// steel from size/2 further
// 		stealIdx := (currIdx + cb.size()/2) % cb.size()
// 		cb.removeAt(stealIdx)
// 		currIdx++
// 		if currIdx >= cb.size() {
// 			currIdx = 0
// 		}
// 		if cb.size()%1000 == 0 {
// 			log("size = %d", cb.size())
// 		}
// 	}
// 	log("last: %d", cb.elts[0])

// 	return 0, nil
// }
