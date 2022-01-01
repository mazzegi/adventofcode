package day_07

import (
	"fmt"
	"strings"
)

func Part1() {
	ips := Parse(input)
	var count int
	for _, ip := range ips {
		if IPSupportsTLS(ip) {
			count++
		}
	}
	fmt.Printf("1/TLS: %d support TLS\n", count)
}

func Part2() {
	ips := Parse(input)
	var count int
	for _, ip := range ips {
		if IPSupportsSSL(ip) {
			count++
		}
	}
	fmt.Printf("2:/SSL %d support SSL\n", count)
}

func Parse(in string) []string {
	var ips []string
	sl := strings.Split(in, "\n")
	for _, s := range sl {
		s = strings.Trim(s, " \r\n\t")
		if s == "" {
			continue
		}
		ips = append(ips, s)
	}
	return ips
}

type Token struct {
	Value    string
	Hypernet bool
}

func Tokens(ip string) []Token {
	var ts []Token
	var curr Token
	flush := func() {
		if curr.Value == "" {
			return
		}
		ts = append(ts, curr)
		curr = Token{}
	}

	for _, r := range ip {
		switch r {
		case '[':
			flush()
			curr.Hypernet = true
		case ']':
			flush()
			curr.Hypernet = false
		default:
			curr.Value += string(r)
		}
	}
	flush()

	return ts
}

func IPSupportsTLS(ip string) bool {
	ts := Tokens(ip)
	var abbaCount int
	for _, t := range ts {
		if t.Hypernet {
			if containsABBA(t.Value) {
				return false
			}
		} else if containsABBA(t.Value) {
			abbaCount++
		}
	}
	return abbaCount > 0
}

func IPSupportsSSL(ip string) bool {
	ts := Tokens(ip)
	var supernetABAs []string
	var hypernetABAs []string
	for _, t := range ts {
		abas := allABAs(t.Value)
		if len(abas) == 0 {
			continue
		}
		if t.Hypernet {
			hypernetABAs = append(hypernetABAs, abas...)
		} else {
			supernetABAs = append(supernetABAs, abas...)
		}
	}
	for _, snaba := range supernetABAs {
		for _, hnaba := range hypernetABAs {
			if isABACorresponding(snaba, hnaba) {
				return true
			}
		}
	}
	return false
}

func isABBA(s string) bool {
	if len(s) != 4 {
		return false
	}
	return s[0] == s[3] &&
		s[1] == s[2] &&
		s[0] != s[1]
}

func containsABBA(s string) bool {
	if len(s) < 4 {
		return false
	}
	for i := 0; i <= len(s)-4; i++ {
		sub := s[i : i+4]
		if isABBA(sub) {
			return true
		}
	}
	return false
}

func isABA(s string) bool {
	if len(s) != 3 {
		return false
	}
	return s[0] == s[2] &&
		s[0] != s[1]
}

func isABACorresponding(aba string, bab string) bool {
	if !isABA(aba) || !isABA(bab) {
		return false
	}
	return bab[0] == aba[1] && bab[1] == aba[0]
}

func allABAs(s string) []string {
	var abas []string
	if len(s) < 3 {
		return abas
	}
	for i := 0; i <= len(s)-3; i++ {
		sub := s[i : i+3]
		if isABA(sub) {
			abas = append(abas, sub)
		}
	}
	return abas
}
