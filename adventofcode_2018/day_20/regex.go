package day_20

import (
	"strings"

	"github.com/pkg/errors"
)

func (r *Regex) Dump(level int) string {
	var sl []string
	ident := strings.Repeat("  ", level)
	sl = append(sl, ident+"---")
	for _, elt := range r.Elts {
		sl = append(sl, ident+string(elt.Value))
		for _, br := range elt.Branches {
			//sl = append(sl, ident+ident+"----------")
			sl = append(sl, br.Dump(level+1))
		}
	}
	return strings.Join(sl, "\n")
}

type RegexElt struct {
	Value    rune
	Branches []*Regex
}

type Regex struct {
	Elts []*RegexElt
}

func ParseRegex(s string) (*Regex, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "^")
	s = strings.TrimSuffix(s, "$")

	rex := &Regex{}

	pos := 0
outer:
	for {
		if pos >= len(s) {
			break
		}
		r := rune(s[pos])
		switch r {
		case 'E', 'S', 'W', 'N':
			rex.Elts = append(rex.Elts, &RegexElt{Value: r})
		case '(':
			//find closing ) and pipe |
			c, err := findClosing(s[pos+1:])
			if err != nil {
				return nil, err
			}
			last := rex.Elts[len(rex.Elts)-1]
			brs := s[pos+1 : pos+1+c]
			brsl := splitSkipBrackets(brs, '|')
			for _, bs := range brsl {
				bs = strings.TrimSpace(bs)
				if bs == "" {
					continue
				}
				brex, err := ParseRegex(bs)
				if err != nil {
					return nil, errors.Wrapf(err, "parse regex %q", bs)
				}
				last.Branches = append(last.Branches, brex)
			}

			pos += c + 2
			continue outer
		}
		pos++
	}

	return rex, nil
}

func findClosing(s string) (int, error) {
	if s == "" {
		return -1, errors.Errorf("empty string")
	}
	nopen := 0
	for i, r := range s {
		if r == '(' {
			nopen++
		} else if r == ')' {
			if nopen == 0 {
				return i, nil
			}
			nopen--
		}
	}
	return -1, errors.Errorf("no closing bracket found")
}

//
func splitSkipBrackets(s string, sep rune) []string {
	if s == "" {
		return []string{}
	}
	sl := []string{""}
	nopen := 0
	for _, r := range s {
		if r == '(' {
			nopen++
		} else if r == ')' {
			nopen--
		}

		if nopen == 0 && r == sep {
			sl = append(sl, "")
			continue
		}
		sl[len(sl)-1] += string(r)
	}
	return sl
}
