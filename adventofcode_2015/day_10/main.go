package main

import (
	"fmt"
	"strings"
)

var input = `1321131112`

var inputTest1 = `1`
var inputTest2 = `11`
var inputTest3 = `21`
var inputTest4 = `1211`
var inputTest5 = `111221`

func main() {
	in := input
	// var sub1, sub2 string
	// calcSub := false
	// fmt.Printf("%q\n", in)
	// for i := 0; i < 40; i++ {
	// 	fmt.Printf("step %d\n", i+1)
	// 	if !calcSub {
	// 		in = next(in)
	// 		//fmt.Printf("%q\n", in)

	// 		ix := strings.Index(in, "111321")
	// 		if ix < 0 {
	// 			continue
	// 		}
	// 		sub1 = in[:ix+5]
	// 		sub2 = in[ix+5:]
	// 		// fmt.Printf("org : %q\n", in)
	// 		// fmt.Printf("sub1: %q\n", sub1)
	// 		// fmt.Printf("sub2: %q\n", sub2)
	// 		calcSub = true
	// 	} else {
	// 		fmt.Printf("start using subs\n")
	// 		//in = next(in)
	// 		sub1 = next(sub1)
	// 		sub2 = next(sub2)
	// 		// fmt.Printf("*** sub ***\n")
	// 		// fmt.Printf("%q\n", in)
	// 		// fmt.Printf("%q\n", sub1)
	// 		// fmt.Printf("%q\n", sub2)
	// 		// fmt.Printf("*** *** ***\n")
	// 		// if sub1+sub2 != in {
	// 		// 	fmt.Printf("örks\n")
	// 		// }
	// 	}
	// }
	//fmt.Printf("final length: %d (sub1+sub2 = %d)\n", len(in), len(sub1)+len(sub2))
	n := calc(in, 50)
	fmt.Printf("final length: %d\n", n)
}

func calc(s string, cnt int) int {
	i := 0
	for {
		fmt.Printf("step %d\n", i+1)
		s = next(s)
		i++
		if i == cnt {
			return len(s)
		}

		ix := strings.Index(s, "111321")
		if ix < 0 {
			continue
		}
		sub1 := s[:ix+5]
		sub2 := s[ix+5:]
		n1 := calc(sub1, cnt-i)
		n2 := calc(sub2, cnt-i)
		return n1 + n2
	}
}

func next(s string) string {
	var ns string
	var curr rune
	var currCount int
	flush := func() {
		if currCount > 0 {
			ns += fmt.Sprintf("%d%s", currCount, string(curr))
		}
	}
	for _, r := range s {
		if currCount == 0 {
			curr = r
			currCount++
			continue
		}
		if r == curr {
			currCount++
			continue
		}
		flush()
		curr = r
		currCount = 1
	}
	flush()
	return ns
}
