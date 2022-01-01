package monster

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Rule interface {
	ID() int
	Match(level int, rules map[int]Rule, msg string) ([]string, bool)
}

type CharRule struct {
	id   int
	char rune
}

type SubsRule struct {
	id   int
	subs [][]int
}

func (r SubsRule) String() string {
	var subS []string
	for _, sub := range r.subs {
		subS = append(subS, fmt.Sprintf("%v", sub))
	}
	return fmt.Sprintf("%d: %s", r.id, strings.Join(subS, " | "))
}

func ParseRule(s string) (Rule, error) {
	sl := strings.Split(s, ":")
	if len(sl) != 2 {
		return nil, errors.Errorf("invalid rule %q", s)
	}
	id, err := strconv.ParseInt(sl[0], 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "parse-int %q", sl[0])
	}
	rs := strings.Trim(sl[1], " ")
	if len(rs) < 1 {
		return nil, errors.Errorf("invalid rule %q", s)
	}
	if len(rs) == 3 && rune(rs[0]) == '"' && rune(rs[2]) == '"' {
		return CharRule{
			id:   int(id),
			char: rune(rs[1]),
		}, nil
	}

	// should be a subs rule
	rule := SubsRule{
		id: int(id),
	}
	conds := strings.Split(rs, "|")
	for _, sc := range conds {
		scsl := strings.Fields(sc)
		var scrs []int
		for _, scs := range scsl {
			rid, err := strconv.ParseInt(strings.Trim(scs, " "), 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "parse-int %q", scs)
			}
			scrs = append(scrs, int(rid))
		}
		if len(scrs) == 0 {
			return nil, errors.Errorf("subs: no rules %q", s)
		}
		rule.subs = append(rule.subs, scrs)
	}
	if len(rule.subs) == 0 {
		return nil, errors.Errorf("subs: no rules %q", s)
	}

	return rule, nil
}

func (r CharRule) ID() int {
	return r.id
}

func (r SubsRule) ID() int {
	return r.id
}

func (r CharRule) Match(level int, rules map[int]Rule, msg string) ([]string, bool) {
	if len(msg) < 1 {
		return nil, false
	}
	if r.char != rune(msg[0]) {
		return nil, false
	}
	return []string{string(r.char)}, true
}

func fmtLevel(level int) string {
	return strings.Repeat(" ", level*2)
}

func mustRule(rules map[int]Rule, id int) Rule {
	sr, ok := rules[id]
	if !ok {
		panic(fmt.Sprintf("no sub-rule with id %d", id))
	}
	return sr
}

func (r SubsRule) Match(level int, rules map[int]Rule, msg string) ([]string, bool) {
	//fmt.Printf("%scheck %s on %q\n", fmtLevel(level), r, msg)
	//var maxMatch string
	// var maxRIDs []int
	// var matches int
	var matches []string
outer:
	for _, srids := range r.subs {
		if len(srids) == 0 {
			continue
		}

		//step 0
		sr := mustRule(rules, srids[0])
		subMatches, match := sr.Match(level+1, rules, msg)
		if !match {
			continue outer
		}
		//

		for i := 1; i < len(srids); i++ {
			sr = mustRule(rules, srids[i])
			var newMatches []string
			for _, m := range subMatches {
				next := msg[len(m):]
				ms, match := sr.Match(level+1, rules, next)
				if !match {
					continue
				}
				for _, sm := range ms {
					newMatches = append(newMatches, m+sm)
				}
			}
			subMatches = newMatches
		}
		matches = append(matches, subMatches...)

		//fmt.Printf("%scheck %s on %q by %v\n", fmtLevel(level), r, msg, srids)
		// nextmsg := msg
		// var matchmsg string
		// var cumMatches []string

		// for _, srid := range srids {
		// 	sr, ok := rules[srid]
		// 	if !ok {
		// 		panic(fmt.Sprintf("no sub-rule with id %d", srid))
		// 	}
		// 	var ms []string
		// 	var match bool

		// 	ms, match = sr.Match(level+1, rules, nextmsg)
		// 	if !match {
		// 		continue outer
		// 	}
		// 	matchmsg += ms
		// 	nextmsg = nextmsg[len(ms):]
		// }
		// matches = append(matches, matchmsg)
		//fmt.Printf("%send-check %s on %q => match %q, rules %v\n", fmtLevel(level), r, msg, matchmsg, srids)
		//return matchmsg, true
		// matches++
		// if len(matchmsg) > len(maxMatch) {
		// 	maxMatch = matchmsg
		// 	maxRIDs = srids
		// }
	}
	if len(matches) > 0 {
		//fmt.Printf("%send-check %s on %q =>  matches = %v\n", fmtLevel(level), r, msg, matches)
		return matches, true
	}
	//fmt.Printf("%send-check %s on %q => NOT match\n", fmtLevel(level), r, msg)
	return nil, false
}

func runeAtIdx(s string, r rune, idx int) bool {
	if idx < 0 || idx >= len(s) {
		return false
	}
	return r == rune(s[idx])
}
