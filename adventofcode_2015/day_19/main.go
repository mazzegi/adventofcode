package main

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	//in := inputTest
	in := input
	in.Replacements = map[string][]string{}
	scanner := bufio.NewScanner(bytes.NewBufferString(in.RawReplacements))
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		sl := strings.Split(l, " => ")
		if len(sl) != 2 {
			panic("invalid replacement")
		}
		rep := strings.Trim(sl[0], " ")
		with := strings.Trim(sl[1], " ")
		in.Replacements[rep] = append(in.Replacements[rep], with)
	}
	fmt.Printf("scanned %d replacements\n", len(in.Replacements))
	//cnt := countReplacements(in)
	//fmt.Printf("there are %d possible replacement-results\n", cnt)
	forge(in)
}

func findSubs(s string, sub string) []int {
	var subs []int
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i:], sub) {
			subs = append(subs, i)
		}
	}
	return subs
}

func replaceAt(s string, idx int, cnt int, with string) string {
	return s[:idx] + with + s[idx+cnt:]
}

func countReplacements(in Input) int {
	results := map[string]bool{}
	for rep, withs := range in.Replacements {
		subs := findSubs(in.Calibration, rep)
		for _, with := range withs {
			for _, sub := range subs {
				res := replaceAt(in.Calibration, sub, len(rep), with)
				results[res] = true
				fmt.Printf("%q via %s => %s at %d\n", res, rep, with, sub)
			}
		}

	}
	return len(results)
}

func forge(in Input) {
	var revOrdered []string
	revRep := map[string]string{}
	for rep, withs := range in.Replacements {
		for _, with := range withs {
			if _, ok := revRep[with]; ok {
				panic("found double reverse replacement")
			}
			revRep[with] = rep
			revOrdered = append(revOrdered, with)
		}
	}
	sort.Slice(revOrdered, func(i, j int) bool {
		return len(revOrdered[i]) > len(revOrdered[j])
	})

	maxIter := 1000000
	currIter := 0
	curr := in.Calibration
outer:
	for currIter < maxIter {
		//find largest substring, which may be reverse replaced
		for _, rr := range revOrdered {
			if strings.Contains(curr, rr) {
				rwith := revRep[rr]
				curr = strings.Replace(curr, rr, rwith, 1)
				currIter++
				if curr == "e" {
					break outer
				}
				continue outer
			}
		}
		fmt.Printf("stuck! left %q\n", curr)
		return
	}
	fmt.Printf("took %d replacements to reengenier\n", currIter)
}

type Input struct {
	RawReplacements string
	Replacements    map[string][]string
	Calibration     string
}

var inputTest = Input{
	RawReplacements: `
	e => H
	e => O
	H => HO
	H => OH
	O => HH
	`,
	Calibration: `HOHOHO`,
}

var input = Input{
	RawReplacements: `
	Al => ThF
	Al => ThRnFAr
	B => BCa
	B => TiB
	B => TiRnFAr
	Ca => CaCa
	Ca => PB
	Ca => PRnFAr
	Ca => SiRnFYFAr
	Ca => SiRnMgAr
	Ca => SiTh
	F => CaF
	F => PMg
	F => SiAl
	H => CRnAlAr
	H => CRnFYFYFAr
	H => CRnFYMgAr
	H => CRnMgYFAr
	H => HCa
	H => NRnFYFAr
	H => NRnMgAr
	H => NTh
	H => OB
	H => ORnFAr
	Mg => BF
	Mg => TiMg
	N => CRnFAr
	N => HSi
	O => CRnFYFAr
	O => CRnMgAr
	O => HP
	O => NRnFAr
	O => OTi
	P => CaP
	P => PTi
	P => SiRnFAr
	Si => CaSi
	Th => ThCa
	Ti => BP
	Ti => TiTi
	e => HF
	e => NAl
	e => OMg
	`,
	Calibration: `CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl`,
}
