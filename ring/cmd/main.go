package main

import (
	"github.com/mazzegi/adventofcode/ring"
	"github.com/mazzegi/log"
)

const doshift = true

func main() {
	if doshift {
		shift()
		return
	}
	r := ring.New[int]()
	r.CheckValid = true

	r.Insert(0)
	log.Infof(r.Dump())

	r.Insert(1)
	log.Infof(r.Dump())

	n2 := r.Insert(2)
	log.Infof(r.Dump())

	n3 := r.Insert(3)
	log.Infof(r.Dump())

	r.Remove(n2)
	log.Infof(r.Dump())

	n44 := r.InsertBefore(n3, 44)
	log.Infof(r.Dump())

	r.ShiftLeft(n44)
	log.Infof(r.Dump())

	r.ShiftLeft(n44)
	log.Infof(r.Dump())

	r.ShiftRight(n44)
	log.Infof(r.Dump())

	r.ShiftRight(n44)
	log.Infof(r.Dump())
}

func shift() {
	r := ring.New[int]()
	r.CheckValid = true

	r.Insert(0)
	r.Insert(1)
	n2 := r.Insert(2)
	n44 := r.InsertBefore(n2, 44)
	r.Insert(3)
	log.Infof(r.Dump())

	// n := 7
	// for i := 0; i < n; i++ {
	// 	r.ShiftRight(n44)
	// }

	n := 7
	n = n%5 + 1
	for i := 0; i < n; i++ {
		r.ShiftLeft(n44)
	}

	log.Infof(r.Dump())
}

/*
2022-12-26T23:43:00.505 [default] [] [INFO ] [3-0-1], [0-1-44], [1-44-2], [44-2-3], [2-3-0]
2022-12-26T23:43:01.788 [default] [] [INFO ] [3-0-44], [0-44-1], [44-1-2], [1-2-3], [2-3-0]


2022-12-26T23:52:00.796 [default] [] [INFO ] [3-0-1], [0-1-44], [1-44-2], [44-2-3], [2-3-0]
2022-12-26T23:52:03.300 [default] [] [INFO ] [3-0-1], [0-1-2], [1-2-44], [2-44-3], [44-3-0]
*/
