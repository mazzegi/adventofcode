package ring2

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func NewBuffer(vs []int) *Buffer {
	b := &Buffer{
		Lookup: map[int]*Elt{},
	}
	for _, v := range vs {
		b.Append(v)
	}
	return b
}

type Elt struct {
	Value int
	Prev  *Elt
	Next  *Elt
}

func (e Elt) String(curr bool) string {
	s := fmt.Sprintf("(%d)%d(%d)", e.Prev.Value, e.Value, e.Next.Value)
	if curr {
		s = "*" + s
	}
	return s
}

type Buffer struct {
	First  *Elt
	Curr   *Elt
	Max    *Elt
	Lookup map[int]*Elt
}

func (b *Buffer) String() string {
	if b.First == nil {
		return "<empty>"
	}
	curr := (b.First == b.Curr)
	sl := []string{b.First.String(curr)}
	compact := []string{fmt.Sprintf("%d", b.First.Value)}
	e := b.First
	for e.Next != b.First {
		e = e.Next
		curr = (e == b.Curr)
		sl = append(sl, e.String(curr))
		compact = append(compact, fmt.Sprintf("%d", e.Value))
	}

	return strings.Join(sl, ", ") + fmt.Sprintf(" => %q", strings.Join(compact, ""))
}

func (b *Buffer) Last() *Elt {
	if b.First == nil {
		return nil
	}
	return b.First.Prev
}

func (b *Buffer) Append(v int) {
	ne := &Elt{
		Value: v,
	}
	b.Lookup[ne.Value] = ne
	if b.First == nil {
		b.First = ne
		b.First.Prev = b.First
		b.First.Next = b.First
		return
	}
	last := b.Last()
	ne.Prev = last
	ne.Next = b.First
	last.Next = ne
	b.First.Prev = last.Next
}

func (b *Buffer) Insert(e *Elt, after *Elt) {
	e.Prev = after
	e.Next = after.Next

	e.Next.Prev = e
	e.Prev.Next = e
}

func (b *Buffer) Find(v int) (*Elt, bool) {
	e, ok := b.Lookup[v]
	return e, ok

	// e := b.First
	// for {
	// 	if e.Value == v {
	// 		return e, true
	// 	}
	// 	e = e.Next
	// 	if e == b.First {
	// 		break
	// 	}
	// }
	// return nil, false
}

func (b *Buffer) FindMax() *Elt {
	max := b.First
	e := b.First
	for {
		if e.Value > max.Value {
			max = e
		}
		e = e.Next
		if e == b.First {
			break
		}
	}
	return max
}

func (b *Buffer) Step() {
	if b.Curr == nil {
		b.Curr = b.First
	}
	var picked []*Elt
	pe := b.Curr.Next

	pickedVals := map[int]bool{}

	for len(picked) < 3 {
		picked = append(picked, pe)
		pickedVals[pe.Value] = true
		pe.Prev.Next = pe.Next
		pe.Next.Prev = pe.Prev
		if pe == b.First {
			b.First = pe.Next
		}
		//fmt.Printf("    picked %d: %s\n", pe.Value, b.String())
		pe = pe.Next
	}
	//fmt.Printf("     after pick: %s\n", b.String())

	var dest *Elt
	destVal := b.Curr.Value - 1
	for {
		if destVal < 1 {
			dest = b.FindMax()
			break
		}
		if !pickedVals[destVal] {
			if dv, ok := b.Find(destVal); ok {
				dest = dv
				break
			}
		}
		destVal--
	}
	//insert
	for _, pe := range picked {
		b.Insert(pe, dest)
		dest = pe
	}

	//new-current
	b.Curr = b.Curr.Next
}

func (b *Buffer) Hash() string {
	one, ok := b.Find(1)
	if !ok {
		panic("<found no 1>")
	}

	h := md5.New()
	e := one.Next
	for {
		io.WriteString(h, fmt.Sprintf("%d", e.Value))
		e = e.Next
		if e == one {
			break
		}
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (b *Buffer) NextToOne() string {
	one, ok := b.Find(1)
	if !ok {
		return "<found no 1>"
	}
	var sl []string
	e := one.Next
	for {
		sl = append(sl, fmt.Sprintf("%d", e.Value))
		e = e.Next
		if e == one {
			break
		}
	}

	return strings.Join(sl, " ")
}

func (b *Buffer) Dump() string {
	one := b.First
	var sl []string
	if one == b.Curr {
		sl = append(sl, fmt.Sprintf("*%d", one.Value))
	} else {
		sl = append(sl, fmt.Sprintf("%d", one.Value))
	}

	e := one.Next
	for {
		if e == b.Curr {
			sl = append(sl, fmt.Sprintf("*%d", e.Value))
		} else {
			sl = append(sl, fmt.Sprintf("%d", e.Value))
		}
		e = e.Next
		if e == one {
			break
		}
	}

	return strings.Join(sl, " ")
}

func (b *Buffer) TwoNextToOne() (string, int) {
	one, ok := b.Find(1)
	if !ok {
		return "<found no 1>", -1
	}
	var sl []string
	p := 1
	e := one.Next
	for len(sl) < 2 {
		sl = append(sl, fmt.Sprintf("%d", e.Value))
		p *= e.Value
		e = e.Next
		if e == one {
			break
		}
	}

	return strings.Join(sl, ","), p
}
