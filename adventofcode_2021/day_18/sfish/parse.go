package sfish

import (
	"strconv"

	"github.com/pkg/errors"
)

func Parse(s string) (*Pair, error) {
	par := &parser{
		data: []byte(s),
		pos:  0,
	}
	p, err := par.parsePair()
	return p, err
}

type parser struct {
	data []byte
	pos  int
}

func (p *parser) curr() byte {
	return p.data[p.pos]
}

func (p *parser) next() (byte, error) {
	p.pos++
	if p.pos >= len(p.data) {
		return 0, errors.Errorf("EOF")
	}
	return p.curr(), nil
}

func (p *parser) parsePair() (*Pair, error) {
	if p.curr() != '[' {
		return nil, errors.Errorf("at %d: expect %q, got %q", p.pos, "[", string(p.curr()))
	}
	n1, err := p.next()
	if err != nil {
		return nil, err
	}
	var e1 Element
	if n1 == '[' {
		p1, err := p.parsePair()
		if err != nil {
			return nil, err
		}
		e1 = p1
	} else {
		n, err := strconv.ParseInt(string(n1), 10, 8)
		if err != nil {
			return nil, err
		}
		e1 = Regular(n)
	}
	nc, err := p.next()
	if err != nil {
		return nil, err
	}
	if nc != ',' {
		return nil, errors.Errorf("at %d: expect %q, got %q", p.pos, ",", string(nc))
	}

	n2, err := p.next()
	if err != nil {
		return nil, err
	}
	var e2 Element
	if n2 == '[' {
		p2, err := p.parsePair()
		if err != nil {
			return nil, err
		}
		e2 = p2
	} else {
		n, err := strconv.ParseInt(string(n2), 10, 8)
		if err != nil {
			return nil, err
		}
		e2 = Regular(n)
	}

	nclose, err := p.next()
	if nclose != ']' {
		return nil, errors.Errorf("at %d: expect %q, got %q", p.pos, "]", string(nclose))
	}

	return &Pair{
		Left:  e1,
		Right: e2,
	}, nil
}
