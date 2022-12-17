package main

import (
	"github.com/mazzegi/adventofcode/list"
	"github.com/mazzegi/log"
)

func main() {
	l := list.New[int]()
	log.Infof(l.Dump())
	n2 := l.PushBack(2)
	l.PushBack(4)
	log.Infof(l.Dump())
	n1 := l.PushFront(1)
	l.PushFront(99)
	log.Infof(l.Dump())
	l.Remove(n1)
	log.Infof(l.Dump())

	l.InsertAfter(14, n2)
	log.Infof(l.Dump())

	log.Infof("%v", l.Values())
}
