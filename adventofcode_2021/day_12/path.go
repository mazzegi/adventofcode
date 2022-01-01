package day_12

import (
	"fmt"
	"strings"
)

var dbg = []string{"start", "A", "b", "A", "c", "A", "end", "A"}

func checkDbg(p *Path) {
	if p.Is(dbg) {
		fmt.Printf("check-dbg\n")
	}
}

func (p *Path) Is(sl []string) bool {
	if len(p.values) != len(sl) {
		return false
	}
	for i, ev := range p.values {
		if ev != sl[i] {
			return false
		}
	}
	return true
}

type Path struct {
	values []string
}

func NewPath(vs ...string) *Path {
	p := &Path{
		values: vs,
	}
	return p
}

func (p *Path) Clone() *Path {
	cp := &Path{}
	for _, v := range p.values {
		cp.values = append(cp.values, v)
	}
	return cp
}

func (p *Path) Append(v string) {
	p.values = append(p.values, v)
}

func (p *Path) Count(v string) int {
	var c int
	for _, exv := range p.values {
		if exv == v {
			c++
		}
	}
	return c
}

func (p *Path) Format() string {
	return strings.Join(p.values, ",")
}

func (p *Path) Last() string {
	if len(p.values) > 0 {
		return p.values[len(p.values)-1]
	}
	return ""
}

func (p *Path) smallTwiceCount() int {
	var stc int
	for _, v := range p.values {
		if isSmall(v) && !isStart(v) && !isEnd(v) {
			if p.Count(v) > 1 {
				stc++
			}
		}
	}
	return stc
}
