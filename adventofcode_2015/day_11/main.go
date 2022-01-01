package main

import (
	"fmt"
	"strings"
)

func main() {
	in, _ := inc("vzbxxyzz")
	var ok bool
	for {
		if isValidPwd(in) {
			fmt.Printf("the next is %q\n", in)
			break
		}
		in, ok = inc(in)
		if !ok {
			fmt.Printf("reached the end\n")
			break
		}
	}
}

func inc(s string) (string, bool) {
	bs := []byte(s)
	for i := len(bs) - 1; i >= 0; i-- {
		n := bs[i]
		if n < byte('z') {
			bs[i]++
			return string(bs), true
		}
		if i == 0 {
			return "", false
		}
		bs[i] = byte('a')
	}
	return "", false
}

func isValidPwd(s string) bool {
	if !contains3IncBy1(s) {
		return false
	}
	if strings.ContainsAny(s, "iol") {
		return false
	}
	if uniquePairs(s) < 2 {
		return false
	}
	return true
}

func contains3IncBy1(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i+1]+1 == s[i+2] {
			return true
		}
	}
	return false
}

func uniquePairs(s string) int {
	found := map[byte]bool{}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] && !found[s[i]] {
			found[s[i]] = true
			i += 1
		}
	}
	return len(found)
}
