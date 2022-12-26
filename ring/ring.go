package ring

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

type Ring[T any] struct {
	First      *Node[T]
	CheckValid bool
}

func New[T any]() *Ring[T] {
	return &Ring[T]{}
}

func (r *Ring[T]) IsValid(n *Node[T]) bool {
	if n == nil {
		return false
	}
	_, ok := r.FindNode(func(en *Node[T]) bool {
		return n == en
	})
	return ok
}

func (r *Ring[T]) checkValid(n *Node[T]) {
	if !r.CheckValid {
		return
	}
	if !r.IsValid(n) {
		panic("invalid node")
	}
}

func (r *Ring[T]) Insert(v T) *Node[T] {
	if r.First == nil {
		nn := &Node[T]{
			Value: v,
		}
		r.First = nn
		nn.Next = nn
		nn.Prev = nn
		return nn
	}
	return r.InsertBefore(r.First, v)
}

func (r *Ring[T]) InsertBefore(bn *Node[T], v T) *Node[T] {
	r.checkValid(bn)
	nn := &Node[T]{
		Value: v,
		Prev:  bn.Prev,
		Next:  bn,
	}
	bn.Prev.Next = nn
	bn.Prev = nn
	return nn
}

func (r *Ring[T]) Remove(rn *Node[T]) {
	r.checkValid(rn)
	rn.Prev.Next = rn.Next
	rn.Next.Prev = rn.Prev

	if rn == r.First {
		if rn.Next != r.First {
			r.First = rn.Next
		} else {
			r.First = nil
		}
	}
}

func (r *Ring[T]) EachNode(fnc func(n *Node[T])) {
	if r.First == nil {
		return
	}
	curr := r.First
	for {
		fnc(curr)
		curr = curr.Next
		if curr == r.First {
			break
		}
	}
}

func (r *Ring[T]) FindNode(match func(n *Node[T]) bool) (*Node[T], bool) {
	if r.First == nil {
		return nil, false
	}
	curr := r.First
	for {
		if match(curr) {
			return curr, true
		}
		curr = curr.Next
		if curr == r.First {
			break
		}
	}
	return nil, false
}

func (r *Ring[T]) Dump() string {
	var sl []string
	r.EachNode(func(n *Node[T]) {
		sl = append(sl, fmt.Sprintf("[%s-%v-%s]", n.Prev.Format(), n.Value, n.Next.Format()))
	})
	return strings.Join(sl, ", ")
}

func (r *Ring[T]) ShiftLeft(n *Node[T]) {
	r.checkValid(n)
	prev := n.Prev
	next := n.Next

	prev.Prev.Next = n
	n.Prev = prev.Prev

	n.Next = prev
	prev.Prev = n

	prev.Next = next
	next.Prev = prev
}

func (r *Ring[T]) ShiftRight(n *Node[T]) {
	r.checkValid(n)
	prev := n.Prev
	next := n.Next

	next.Next.Prev = n
	n.Next = next.Next

	n.Prev = next
	next.Next = n

	next.Prev = prev
	prev.Next = next
}
