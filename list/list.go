package list

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

func (n *Node[T]) Format() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", n.Value)
}

type List[T any] struct {
	First *Node[T]
	Last  *Node[T]
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) IsEmpty() bool {
	return l.First == nil
}

func (l *List[T]) IsValid(tn *Node[T]) bool {
	if tn == nil {
		return false
	}
	for n := l.First; n != nil; n = n.Next {
		if n == tn {
			return true
		}
	}
	return false
}

func (l *List[T]) PushFront(v T) *Node[T] {
	nn := &Node[T]{
		Value: v,
		Prev:  nil,
		Next:  l.First,
	}
	if l.First == nil {
		l.First = nn
		l.Last = nn
		return nn
	}
	l.First.Prev = nn
	l.First = nn
	return nn
}

func (l *List[T]) PushBack(v T) *Node[T] {
	nn := &Node[T]{
		Value: v,
		Prev:  l.Last,
		Next:  nil,
	}
	if l.Last == nil {
		l.First = nn
		l.Last = nn
		return nn
	}
	l.Last.Next = nn
	l.Last = nn
	return nn
}

func (l *List[T]) InsertBefore(v T, bn *Node[T]) *Node[T] {
	if !l.IsValid(bn) {
		panic("invalid node")
	}
	nn := &Node[T]{
		Value: v,
		Prev:  bn.Prev,
		Next:  bn,
	}
	if bn.Prev != nil {
		bn.Prev.Next = nn
	}
	bn.Prev = nn
	if bn == l.First {
		l.First = nn
	}
	return nn
}

func (l *List[T]) InsertAfter(v T, an *Node[T]) *Node[T] {
	if !l.IsValid(an) {
		panic("invalid node")
	}
	nn := &Node[T]{
		Value: v,
		Prev:  an,
		Next:  an.Next,
	}
	if an.Next != nil {
		an.Next.Prev = nn
	}
	an.Next = nn
	if an == l.Last {
		l.Last = nn
	}
	return nn
}

func (l *List[T]) Remove(rn *Node[T]) {
	if !l.IsValid(rn) {
		panic("invalid node")
	}
	if rn == l.First {
		if l.First.Next != nil {
			l.First.Next.Prev = nil
		}
		return
	}
	if rn == l.Last {
		if l.Last.Prev != nil {
			l.Last.Prev.Next = nil
		}
		return
	}
	if rn.Prev != nil {
		rn.Prev.Next = rn.Next
	}
	if rn.Next != nil {
		rn.Next.Prev = rn.Prev
	}
}

func (l *List[T]) Values() []T {
	var vs []T
	for n := l.First; n != nil; n = n.Next {
		vs = append(vs, n.Value)
	}
	return vs
}

func (l *List[T]) Each(do func(T)) {
	for n := l.First; n != nil; n = n.Next {
		do(n.Value)
	}
}

func (l *List[T]) Find(start *Node[T], match func(T) bool) *Node[T] {
	if start == nil {
		start = l.First
	}
	for n := start; n != nil; n = n.Next {
		if match(n.Value) {
			return n
		}
	}
	return nil
}

func (l *List[T]) Dump() string {
	var sl []string
	for n := l.First; n != nil; n = n.Next {
		sl = append(sl, fmt.Sprintf("[%s-%v-%s]", n.Prev.Format(), n.Value, n.Next.Format()))
	}
	return strings.Join(sl, ", ")
}
