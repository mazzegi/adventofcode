package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode"

	"github.com/mazzegi/adventofcode/adventofcode_2018/polymer"
)

func main() {
	//in := inputTest()
	in := input()
	t0 := time.Now()
	r := polymer.React(in)
	fmt.Printf("%d units remaining (%s)\n", len(r), time.Since(t0))

	fmt.Printf("*** Part2 ***\n")
	min := len(in)
	var minRm rune
	for r := 'a'; r <= 'z'; r++ {
		rin := strings.ReplaceAll(in, string(r), "")
		rin = strings.ReplaceAll(rin, string(unicode.ToUpper(r)), "")

		res := polymer.React(rin)
		if len(res) < min {
			min = len(res)
			minRm = r
		}
		fmt.Printf("removing %q => %d\n", string(r), len(res))
	}
	fmt.Printf("shortest is: %d (removeing %q)\n", min, string(minRm))
}

func inputTest() string {
	return `dabAcCaCBAcCcaDA`
}

func input() string {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(bs), " \r\n\t")
}
