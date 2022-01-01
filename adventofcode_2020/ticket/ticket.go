package ticket

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Ticket struct {
	Numbers []int
}

func ParseTicket(s string) (*Ticket, error) {
	sl := strings.Split(s, ",")
	t := &Ticket{}
	for _, sn := range sl {
		n, err := strconv.ParseInt(sn, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "parse-int")
		}
		t.Numbers = append(t.Numbers, int(n))
	}
	return t, nil
}

type Range struct {
	From, To int
}

func ParseRange(s string) (Range, error) {
	var f, t int
	_, err := fmt.Sscanf(s, "%d-%d", &f, &t)
	if err != nil {
		return Range{}, errors.Wrapf(err, "scan rule %q", s)
	}
	return Range{
		From: f,
		To:   t,
	}, nil
}

type Rule struct {
	Name   string
	Ranges []Range
}

func (r Rule) String() string {
	var rns []string
	for _, rn := range r.Ranges {
		rns = append(rns, fmt.Sprintf("[%d-%d]", rn.From, rn.To))
	}
	return fmt.Sprintf("%s: %s", r.Name, strings.Join(rns, ", "))
}

func ParseRule(s string) (*Rule, error) {
	fs := strings.Split(s, ":")
	if len(fs) != 2 {
		return nil, errors.Errorf("invalid rule %q", s)
	}
	if fs[0] == "" || fs[1] == "" {
		return nil, errors.Errorf("invalid rule %q", s)
	}
	rule := &Rule{
		Name: fs[0],
	}
	rssl := strings.Split(fs[1], " or ")
	for _, rs := range rssl {
		rn, err := ParseRange(strings.Trim(rs, " "))
		if err != nil {
			return nil, errors.Wrapf(err, "parse-range %q", rs)
		}
		rule.Ranges = append(rule.Ranges, rn)
	}
	return rule, nil
}

func (r *Rule) IsValid(n int) bool {
	for _, rn := range r.Ranges {
		if n >= rn.From && n <= rn.To {
			return true
		}
	}
	return false
}

func (r *Rule) IsValidForAll(ns []int) bool {
	for _, n := range ns {
		if !r.IsValid(n) {
			return false
		}
	}
	return true
}

type Rules []*Rule

func (rs Rules) IsValidNumber(n int) bool {
	for _, r := range rs {
		if r.IsValid(n) {
			return true
		}
	}
	return false
}

func (rs Rules) IsValidTicket(t *Ticket) ([]int, bool) {
	var nv []int
	valid := true
	for _, n := range t.Numbers {
		if !rs.IsValidNumber(n) {
			valid = false
			nv = append(nv, n)
		}
	}
	return nv, valid
}
