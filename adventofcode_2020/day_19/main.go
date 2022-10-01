package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2020/monster"
)

func main() {
	in := input
	rules, err := parseRawRules(in.RawRules)
	if err != nil {
		panic(err)
	}
	rule0, ok := rules[0]
	if !ok {
		panic("there's no rule 0")
	}

	var messages []string
	scanner := bufio.NewScanner(bytes.NewBufferString(in.RawMessages))
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		messages = append(messages, l)
	}

	// var matchCnt int
	// for _, msg := range messages {
	// 	matchMsg, match := rule0.Match(rules, msg)
	// 	if match {
	// 		if matchMsg == msg {
	// 			matchCnt++
	// 			fmt.Printf("%q matches\n", msg)
	// 		} else {
	// 			fmt.Printf("%q matches - but has extra (match = %q)\n", msg, matchMsg)
	// 		}
	// 	} else {
	// 		fmt.Printf("%q does NOT match\n", msg)
	// 	}
	// }
	// fmt.Printf("%d messages are matching rule 0\n", matchCnt)

	fmt.Printf("\n *** Part2 ***\n")
	fmt.Printf("update rules\n")

	sr8, err := monster.ParseRule("8: 42 | 42 8")
	if err != nil {
		panic(err)
	}
	sr11, err := monster.ParseRule("11: 42 31 | 42 11 31")
	if err != nil {
		panic(err)
	}
	rules[sr8.ID()] = sr8
	rules[sr11.ID()] = sr11

	matchCnt := 0
outer:
	for _, msg := range messages {
		matches, match := rule0.Match(0, rules, msg)
		if !match || len(matches) == 0 {
			fmt.Printf("%q does NOT match\n", msg)
			continue
		}

		for _, m := range matches {
			if m == msg {
				fmt.Printf("%q matches\n", msg)
				matchCnt++
				continue outer
			}
		}
		fmt.Printf("%q matches - but any has extra %v\n", msg, matches)

		// if match {
		// 	if matchMsg == msg {
		// 		matchCnt++
		// 		fmt.Printf("%q matches\n", msg)
		// 	} else {
		// 		fmt.Printf("%q matches - but has extra (match = %q)\n", msg, matchMsg)
		// 	}
		// } else {
		// 	fmt.Printf("%q does NOT match\n", msg)
		// }
	}
	fmt.Printf("%d messages are matching rule 0\n", matchCnt)

	//Test single
	// testMsg := "babbbbaabbbbbabbbbbbaabaaabaaa"
	// fmt.Printf("\n *** Part2 - Test %q\n", testMsg)
	// matchMsg, match := rule0.Match(0, rules, testMsg)
	// if match {
	// 	fmt.Printf("%q matches (mm = %q)\n", testMsg, matchMsg)
	// } else {
	// 	fmt.Printf("%q does NOT match\n", testMsg)
	// }
}

/*
babbbbaabbbbbabbbbbbaabaaabaaa -> NOT MATCH
babbbbaabbbbbabbbbbbaabaaabaaa
*/

func parseRawRules(s string) (map[int]monster.Rule, error) {
	rules := map[int]monster.Rule{}
	scanner := bufio.NewScanner(bytes.NewBufferString(s))
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		rule, err := monster.ParseRule(l)
		if err != nil {
			return nil, err
		}
		rules[rule.ID()] = rule
	}
	return rules, nil
}

type Input struct {
	RawRules    string
	RawMessages string
}

var inputPart2 = Input{
	RawRules: `
	42: 9 14 | 10 1
	9: 14 27 | 1 26
	10: 23 14 | 28 1
	1: "a"
	11: 42 31
	5: 1 14 | 15 1
	19: 14 1 | 14 14
	12: 24 14 | 19 1
	16: 15 1 | 14 14
	31: 14 17 | 1 13
	6: 14 14 | 1 14
	2: 1 24 | 14 4
	0: 8 11
	13: 14 3 | 1 12
	15: 1 | 14
	17: 14 2 | 1 7
	23: 25 1 | 22 14
	28: 16 1
	4: 1 1
	20: 14 14 | 1 15
	3: 5 14 | 16 1
	27: 1 6 | 14 18
	14: "b"
	21: 14 1 | 1 14
	25: 1 1 | 1 14
	22: 14 14
	8: 42
	26: 14 22 | 1 20
	18: 15 15
	7: 14 5 | 1 21
	24: 14 1
	`,
	RawMessages: `
	abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
	bbabbbbaabaabba
	babbbbaabbbbbabbbbbbaabaaabaaa
	aaabbbbbbaaaabaababaabababbabaaabbababababaaa
	bbbbbbbaaaabbbbaaabbabaaa
	bbbababbbbaaaaaaaabbababaaababaabab
	ababaaaaaabaaab
	ababaaaaabbbaba
	baabbaaaabbaaaababbaababb
	abbbbabbbbaaaababbbbbbaaaababb
	aaaaabbaabaaaaababaa
	aaaabbaaaabbaaa
	aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
	babaaabbbaaabaababbaabababaaab
	aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
	`,
}

var inputTest0 = Input{
	RawRules: `
	0: 1 2
	1: "a"
	2: 1 3 | 3 1
	3: "b"
	`,
	RawMessages: `
	aaa
	aab
	aba
	abb	
	baa
	bab
	bba
	bbb
	`,
}

// match ababbb and abbbab
var inputTest = Input{
	RawRules: `
	0: 4 1 5
	1: 2 3 | 3 2
	2: 4 4 | 5 5
	3: 4 5 | 5 4
	4: "a"
	5: "b"
	`,
	RawMessages: `
	ababbb 
	bababa
	abbbab
	aaabbb
	aaaabbb
	`,
}

var input = Input{
	RawRules: `
	26: 97 126 | 123 57
	122: 84 97 | 92 123
	82: 97 138 | 123 130
	80: 131 97
	3: 123 107 | 97 66
	101: 1 123 | 95 97
	10: 97 138 | 123 107
	83: 123 27 | 97 12
	54: 97 126 | 123 107
	61: 97 50 | 123 114
	123: "a"
	86: 123 18 | 97 36
	70: 97 76 | 123 69
	34: 66 97 | 110 123
	68: 97 93 | 123 86
	131: 123 97 | 97 123
	126: 97 97 | 97 123
	93: 59 97 | 9 123
	51: 78 21
	94: 97 90 | 123 39
	49: 88 97 | 66 123
	104: 75 97 | 33 123
	114: 82 123 | 116 97
	45: 123 13 | 97 62
	36: 123 81 | 97 65
	58: 53 97 | 132 123
	135: 97 103 | 123 105
	15: 123 2 | 97 49
	85: 130 78
	124: 138 123 | 130 97
	105: 97 16 | 123 6
	66: 78 97 | 123 123
	76: 123 45 | 97 55
	138: 97 123
	20: 97 107 | 123 13
	38: 97 110
	16: 110 123 | 130 97
	97: "b"
	117: 13 97 | 66 123
	19: 123 138 | 97 131
	106: 123 35 | 97 13
	59: 97 5 | 123 74
	7: 97 21 | 123 66
	35: 97 123 | 123 78
	0: 8 11
	103: 97 71 | 123 128
	95: 123 13 | 97 138
	109: 123 136 | 97 60
	96: 123 43 | 97 52
	12: 6 97 | 137 123
	65: 123 138 | 97 138
	50: 123 7 | 97 26
	11: 42 31
	128: 21 123 | 107 97
	79: 35 123
	125: 123 130 | 97 126
	108: 123 85 | 97 26
	9: 23 123 | 20 97
	137: 123 126 | 97 57
	33: 97 28 | 123 43
	57: 123 97 | 123 123
	32: 26 97 | 80 123
	90: 123 44 | 97 4
	48: 110 97 | 130 123
	67: 57 123 | 13 97
	69: 97 67 | 123 30
	37: 88 123 | 107 97
	41: 111 123 | 72 97
	42: 119 97 | 58 123
	110: 78 123 | 97 97
	100: 97 131 | 123 57
	127: 24 97 | 41 123
	98: 32 97 | 134 123
	119: 123 22 | 97 129
	112: 13 97 | 107 123
	5: 138 97 | 126 123
	75: 123 121 | 97 10
	29: 110 123 | 13 97
	40: 97 89 | 123 46
	111: 97 5 | 123 44
	56: 123 133 | 97 96
	30: 57 123 | 62 97
	132: 98 123 | 70 97
	89: 107 97 | 13 123
	44: 97 131 | 123 13
	25: 15 123 | 108 97
	78: 123 | 97
	74: 110 123
	134: 123 113 | 97 54
	47: 131 123 | 21 97
	116: 130 97 | 35 123
	91: 115 97 | 102 123
	6: 57 78
	99: 123 68 | 97 120
	62: 123 123
	84: 123 110 | 97 62
	118: 78 97 | 97 123
	121: 97 62 | 123 107
	22: 25 97 | 61 123
	28: 97 131 | 123 110
	102: 117 123 | 29 97
	73: 126 97 | 130 123
	72: 97 124 | 123 47
	120: 56 97 | 109 123
	115: 97 3 | 123 38
	24: 97 14 | 123 101
	55: 130 97 | 131 123
	81: 107 123 | 126 97
	39: 123 112 | 97 79
	60: 106 123 | 100 97
	136: 97 64 | 123 48
	130: 97 97 | 123 97
	133: 97 74 | 123 34
	92: 123 130 | 97 110
	2: 97 138 | 123 88
	71: 110 97 | 131 123
	53: 123 91 | 97 17
	14: 97 125 | 123 87
	43: 97 62 | 123 21
	21: 97 97 | 123 123
	23: 130 123 | 66 97
	17: 97 122 | 123 40
	8: 42
	1: 13 123 | 57 97
	129: 97 94 | 123 135
	107: 78 78
	18: 97 37 | 123 51
	88: 97 97
	113: 57 97 | 126 123
	46: 97 13 | 123 138
	27: 19 123 | 73 97
	87: 107 123 | 118 97
	52: 97 130 | 123 126
	63: 123 77 | 97 127
	64: 97 118 | 123 88
	13: 123 97
	31: 99 123 | 63 97
	4: 57 97 | 131 123
	77: 123 104 | 97 83
	`,
	RawMessages: `
	bababbbbbbbaabbabbbbaaab
	baaaabbbabaaabbbaabbabbbabbbbbaa
	bbabbabaabbaabbaabbabbaa
	abbbabbbabaabababaaabbbb
	abbaaabaaaaaababababbaab
	babaabbaabaaaabbbbbbabababaaabbabbaabbbb
	aaabbaabbbbaaabbabbbabaa
	abbaaaaabbaabaaabbabaaaa
	aaaaaabababaabbabbababaabbbaaaaababbabbabaabbaba
	abbbbaabbababaaaabbaabbb
	aaaaabaaaabbaaabbbbabbaa
	abbbbaababbabbbabbababbb
	abbbbaababaabbbaababaabbaabaabbaaabaababbbabbbaaaabbbbaa
	bbbbabbbbabaaaabbababbabbabbbaaa
	abbbbaababbabbbaabaaabaabaababaaabababab
	abababaababababaabbaaaabbaabaaaaabbbaaba
	abbbababbbaaababbababaaaabbbabaabbbbabaa
	babbbaabbaaaaaaabaabbaaa
	aaaaaabaaabaabababbbaaba
	abaabbabbabaabbbaaabaaaabababaabbbbbabba
	aaaabaabaabaababbaaabaaa
	bbaaabbbababaabbabaaabbbbbaaababbbbbbbbaaaaaaaab
	bbabababbbbabaabbabbabaabababbabaabbbaaaabababbaabbaaaaaaaabaabbabbaaabbaababbbbbaabbaba
	aabbaabbbaaaaabbaaaaabbb
	abbaaaabbaabaaaaaabaaabb
	abbbbbbbaaababaabbaababaabbaabbbaababbabaabaabbbbabbbaaaabbbabba
	aaabababbbbabbabbaaaaaaaaaababba
	babaabbbabbbabbbbaaabbbb
	bbbbbaaaabbbbbababbbabba
	babaaaabaabbaababababbbbbaabbbba
	abaaaaabbbbbabababbabbaa
	abaabbbabaaaababaaabaaba
	abbbbbbbaabbbbbbabbbbbabaaababbbbaaaaabaaaaabbba
	aaababaabaaaabbbabbaaaaabbabaaabbaababbbaaaabbba
	babbaabbaaabbbbbabbbbbabababbaaa
	babaaaabbabababbbbaababb
	bbababaaababbbaabbbbabababbbbbabbbbabbabbaabbaaa
	abbbbbbaabbabaaaabaaabbaabbaabbbaaaabbba
	babaaaabbabababbbbabaaababaababb
	aabbbbbbbbbabbabbababbaa
	bbaaaaaaaaababaaabbbbbaaabbbabbbabbabaaaabaabaabbabaabbabbbbaabbbbaabaaaababbaabaababaab
	abbbbaabbbabaaabaabababb
	baabbbababbaaabaabbbabbbbabbbaaa
	aaababaaaabbbaaababbbbbababbbaba
	aaabbbbbabbaaababababbaa
	aaabbbabaaaabbaaaababbbbaabbaabaaaaaaabb
	bbabbabababbabbaabababbabbabbbaaaaaaaaaa
	bbaababaaabbbbbbabbbbaba
	abbbbbbaaabbaaabbbbaabbaaaaaaababbbbaabaaabbabab
	baaaaabbbabaaaaabbaaabbbbaabbabbaaaaaabb
	aaaaabaabbabbabaabbbaaba
	baabbaabaabbbbabaabbabab
	bbabbabababbababaabbabbaabababbb
	bababbabaabbbaaaabbaaaabbbabbabbbbbabbbb
	ababbbabbabbbabbaabababa
	bababbbababbbaabbaabbbabaabbbbabbbaabbbaaaaabbab
	bababaaaabaaaaabaaabbbabbaababab
	abbaabbababbaaaabbaaaaababbbabbbbbabbababaaaabaabaaabaababbabbbb
	aaaaaababaaaaabbbaaabaaa
	abbaaaaaabaaaaaabbbbabbbbaaaababbbbabbabababababbaaabbbbababbaba
	bbabbaaaabaaaaabbabbabbb
	bbbaaaaaabaaabaababaabaaaabbbaab
	abbaababbbbaabbabbbaabba
	baabaabaabaaaabaaaababbbbbaaaaaabbbbaaaaaabbabaaabbababababbaaba
	babbbabbbabbaaaaaaaaaabb
	bbaaababbbaabbabaaaabaaa
	bbbaaaaaabbbbbbbbbabaabb
	abaabbbbaaababbbabbbbabbaabbbabb
	bbbaaaaaaabbaaababbbabbbbaabbbba
	abbaababaaaaaababbbbbaaaabaabaaaababaaaa
	aabaababbbaabababaabaaabaaababbbabaababbababbabbbbbabbaa
	bbbbbabaaabbbbbbbabbbbaa
	aaaaabababaaaabbbaababab
	bbbbbaaabaaaaaabaaaaaaba
	bbaababaaaabaaaaabbbbaabbbaaabba
	aaaabababbabaaabbaaaabba
	ababbbababababbabbabbbab
	abbaaababababbaaaabbaaaababaaaabbabbbaba
	babbaabbabbabaaaababbbababbaabaa
	abaaabbabbabbababaababaa
	aaaaababaabbbbabbbbabaab
	bababbbaabaaaaabbaaaaaabbaaaabba
	ababaaabababaabbbbaaaabb
	bababbbababaabaaabababbb
	abbaabbaabbbabbbaababbaa
	baaababaabbaabbaaaabaaaaaaabbabbabbaaabaabbaabbb
	bbbaaaaabaaabababababaaabbbaaaaaaaabbbba
	bbbbaabbabbbbbbbbbbbabbbabababaabaaabbbb
	aaaaababaaaaababbaabbaba
	babbababaabbaaabbabbbbaa
	baaaaaaaaabbbbababaabaaa
	bbaaabbaabbbbbbaababbbababababaaaaaabababbabbbaaaaabbbba
	babbbabbababbbbababbbbbaabaaabbbaababaababbaaaababaaaaabaababbaa
	bbaaaaabbbaaabbbbbababbb
	aaaabababbabbabaababbaaa
	aaabbabbaaabbaabababbaaa
	abbaaaaabbabbbaaaaaaababaaabbaaaabbbaaaabbabbabb
	abbaabababbbabbbaaababba
	aabbbbbbababbbaabbaabbba
	baaabbaaabaabbbbababbbaaaaaaabababaabbbabbaabaaa
	abbbabbbbabbbaaaaaaabbbbbbaababbaabaaaaa
	aabaaababaabbaabbabababbabbbbbaa
	baabbbbbaaaabaabababbbabbbbaabbb
	abaaabaabbbbababaabaababbabbababaabbabaabaababab
	abababaaaaabbaabbbabaaaa
	aaaaabababaabbbbaaabbaaababbbbabbaaabbba
	bbababaabaaaaabbbbaaaaaa
	bbaabbbbabaabaaabbbaaaaabaabbaaabbbbabba
	aaabbabbbaaaaaababaaaabbaaaabaabbabbbbbbababbaab
	abaaabaaabbabbbaaaaaaaab
	aaabbabbbaabbbababaaabbabbbbbabaabbbabab
	aaaabaaaaabbbbbabbaababbaababbbabbaabbaabbaabaab
	bbabaaababaabababbabbaaaabbbabbbbaaaabab
	abaaabbbbaaaabbbbaaabbab
	aaaabababbabbbaaaaaaaabb
	baaaaaabbababaaaaabbbababbabbbaababbbaabbabbbbbbbaabaabb
	abbbbbababbaabbaabbbbabbbaaabbaaabbbaaabbbaabbbaabaabaaa
	abbaaaabaaabaaaaaaabaabb
	bbbbaabbbbbaaaaaaabbabaa
	aaabbaaabbabbabaabaaababbbbbbbbb
	ababbbbbbabaabaabbababbaaaaababaaaaabaabbbbabbabbaaaabaabaaabbaabababbab
	aaaaabaaabaaabbabbbaabbb
	abbaaabaabaababaabaabbabaabbbaababbbabaa
	aabbaaaaaaabbaaaabaababb
	abbaabbaaabaabbabaababab
	aabbbbbbbbaabbabaaabbabbaabbbabbbbbaabbb
	bbbababbaabbbaabaaabaaaababbababaaabbbabaaabaaabaaababaaababaaab
	aabaabababbaaaaabbababababbaabbaabbababaaabababa
	bbbaaabaabaaaaabbbaaaaabaaaaabaabaaaaaba
	bababbbaabbabaaaabbabbbb
	bababbabbaaabababbabaaabababbaba
	baabbbbbbabbabababaaaaba
	babababaabbbbbbbbaababaa
	abaaabbabbaabaabbaaabbaaababbbabbabaaabbaaabaaba
	bbaaaaabaaabbabbababbbbb
	bbabbaabaabaabbbaaaababbabaababbbbaaabbbabbbaabbabbbbbaa
	abbaaaaaabbbbaababbaaaaaabbaababbaabbbaa
	bbbbaabbaabbbbbaaabbaabababbbaabbbbabaab
	aababbabbbabaaaabaaabbbbbbabbbbb
	babbbabbabaabbbbbababaab
	bbbbbaaaaabbaababbabbaaabaaaabba
	bbaaabbaaaabbbabbaabaabb
	baabbbbbbbaabaabbbabbabababbaaba
	aabbabbaaaabbaabbabaabaaabbbaaab
	aabaaabababbbbbaababaaba
	baaabbaabbaaabbabaabbbba
	ababaabbbbababababbbbabbbbabbbaaaaabbabaaabaaaab
	abbbbaabbbbaabaabaaaaaaabbaaabbaabbabbbb
	aabaabbabaababbaaabaaaab
	abbbbaababbaaaababababab
	ababaabaaabbbbababbaaaabbaababbabaaaaabb
	babaaaabbbabbaabbbbaaaab
	aaababbaabaaaabbbaababbbaabbbbaaaabbbbaabbaababaaababaab
	abbaaaabaabbbababaabbaba
	bbbbaabbaaabbabbbabbbaabaababbbabbbbbabb
	babbabbaaaabbabbbababaababbbabab
	aaaaababbbabbabaaaaaaababbbbabaa
	aaababaaabaaababbaabbaabbbbaaabbbaabbbbbaaabaaaaaabaaaaabbabaaba
	bbbbbabaaabaaabaabaaabbaaabaabaaabaabaaa
	bbbaaaaaaaaabbbabaabbbaabbbabaaaabbabbbb
	aabbabbabbbaaabbabaaabbaaaabbaaaabbaaabb
	abaabbaaabaababaabbbbabbabbabbbaaaababaaaaabbbbabababaab
	bbaabababbabbbaaabbabaab
	abaaabbaaabaababbabbbaba
	aaaabababbababaababbaaba
	aaaaababbbbbaabaabaabaaa
	bbbababbbbbaaaaabbaaaababaabaabb
	baabaaaabbbbaaaaaabaababbaaaabba
	aabbaaababbbbabbaabaabbb
	baababbaaaaabbaaaababbba
	bbbaabbaabbaabababaaaaba
	aaaaaababaaaaaaaaabbbbbaaabaaaaa
	bbbbbbaaabaaabbbbbaaaaababbbabba
	babbababbabababbabbbabaaaaabbbbbaababbaabbbbbbbabbbbbabaabbaabab
	bbaabbababbbbaaaaaabbbaaabaaaaaaabababaaabbabbbababbbbbbbbabbbaaabaabbab
	aaaabababbabbaaaaaaaabaabbaaaaaaaabababa
	abaabbbbbbbbababbbbaaababbaababaabbaaaaabbabbbbbaaaaabba
	abbbbbaabbbabbbbbbabbaaaaababababaabaababaaabbbaabababab
	aabbaaaabbabbbbaaabbaabbabbbbaabababbabbbabbbaba
	bababbbabbabbaaabbaaabaa
	bbabbaaabbbbaabaaabbbbabaabbaaaaabbbaabbbbbbbabb
	ababbbabbbabbbbaabaaabbaabbbbbbaaaaaabbb
	aaabaaaababbababbbaabbbbaaababba
	bbaabbababaaaaaaababbbabbbabababbbaaabbbbababaababbbaaababaabbaa
	babbbaabbabbaaaabbaaaaba
	bbaababaabaaaaababaaabaaaaaabababaababbaaaababbabbbbabba
	abaaaaaabbaabaabaababaaa
	abbabaaaaaababababaaabbbaaabaaab
	aaaaabaaaaaaaabaaaabbaaa
	babababbbaaaaabbabbbbbaa
	aabbaaaaabbbabbbbabbabaa
	bbbbbaabbbaaabbbbaabbabb
	baaaabbbbaababbbaabbaaabbaabaabbababaabaaaabbbbabaabaababaabbaaa
	babbaaaabbabaabbbaabaaaabaabaaba
	abaaaaabaaaaabaabababbbaaaaabaaabbabaaba
	babaabbaabbbaabbabbbbaba
	abbbaabbababaaababbbbbaa
	bbaaaaabaababbaabaababab
	aabbbbbabbbbbaabbbabaaababbabbbaaabaaaab
	babbabbaabaaaabbbbabbaaabbababbbaabaabaa
	abbabaaabaaaaaaabaabbaabbbbbaaab
	baaaaaababababaaabaabbbabbbaabaabbbaaaabaabababb
	babbbaabbbbaabbabbbababbbbbbaababbbbabaa
	abababaababababbaaabaabb
	abbaababbbbbbbaabaaaaaababbaabbababaaabbaabbaabbbbbabaaababaaaba
	ababaaababaaaabbabaaabaaaabbaaaaabaaababbbaabbbaaabaabbb
	abbaaabaabbaabbabaaaabaa
	abaabbbabbbabbabaaaaaaab
	babbbbbbbbaabababbbbabaa
	aabaabbabababbbaabaababb
	abbaaabaabaaabaabbbaaaab
	bbbbaabbaabbaaabbabbbbaa
	bababbabbbaaababbaaababb
	bbbbbaabbabbbbbbaabbbbaaaaabbaba
	bbbaabaaaabbbabababbaabbbabbbbbbbbabbaaaababbbabaaabbaba
	bbbabaabbbbbabbababaababaababbaabaabaaabbabaabbabbbaaabbbabaababbaaabbbbababaabbbabbbabb
	abbaaabaaaabbaabbababbbbaaababbbaaababbaaaabbbaa
	aabbbaaaaaaabbabaabbaaabaaaabbbbaabaabbbaabaaabaaaaabbabbaaaababbababbbb
	bbaaabbbbbbabbabaabaabbb
	babaaaaaaaabbbbbaaabaaab
	abaabbbaabababbaaabaabababaaabababaabaaaaaaabbbaaababbab
	aaaabaababaabbabbaababbbbbaabbbb
	babbbbbaabbbbbbbbaaababb
	babbaaaabbbbaaaaabaabbabbbbababbaabbbbabaaabbaba
	babbbabbbabbbbbbbbbabaaa
	bababbbbabaaabaaaaabbbaa
	bbbaaabbaabaabababaabbababbabbbaaaabbabbbbaababbbaaaabba
	bbabababbabbabbaabbbbbaa
	bbababababbaababaaaabaabababbbbb
	aaabbabbbabbaabbbbaaabaa
	babbaaaababbaabbbaabaabb
	baaababaaabbaabbabbbabaa
	bbbaabaabbbaaaaabbbbbbba
	baaaaaabbbbbbaabbabaaaabaaababaaabbbbbaaababbabbbaababab
	bbabababaabbaabbaabbbbbbabbaaaababbbbbaaaaaaabba
	abaaaabbbabababbbabbbbaa
	babbabbbababaabbaabaaababbbabbabbababbaabbaabababbabbbbbbaaaaabb
	baabbbabbbabbaabaaaabbbb
	ababaaabaaabbaabbaabbaabababaabaababbabb
	abbbabbbbbababaabaabaabb
	babababbbbbaabaabbababbabbababbb
	babaabbaabababbabbbbaaaaabbaababbbababbaabbbbaaa
	bbaababababaabaabaababaa
	ababaaabbbbbbababababaaabaababbaababaaaa
	babbababbbaaababbabbaaabbbbababa
	aaabaaaabbbbaabbbbaaababbaaabbaaabbbababaaaabbab
	bbbaaabbabaaababbbabaaaa
	abbaabbaaaaabaabaaabbbba
	aabbbbbbaabaabababaaaaba
	baababbaaaaabbaaabbbaaab
	babbbaabaabbaaaabaaabaab
	babbbbbaaaabbbababbbaaaa
	bababbbbababbbaaabbbaabb
	abaaaaaaaaabababbbabababaaaabbba
	aabbbbbababbabbabbabaaaa
	bbbaaaaabbabbaabbbbbbaabbabaabaaabbbaaaaabbabaab
	abaabababaaababababbbbbbbaaaabaa
	aabaaabababbbabbaaaaaabb
	abbabaaabbbbbababaaabbaabbbbaaab
	abbbbbababaabbbabbbabbabbbabbabbbbaabbba
	abaabababbaabbabababaaaa
	aaaaabbbabbbabbaaabbbbaaababaaba
	baaaaaabaabbaaababbbbbbabababbbabbbbabaa
	bbaabbabaaabbbbbbaaabbba
	bbabbbbababaabaabaaaabaa
	abaaaaaaaabbbbabbaaaaabbbbbbbbabbaabbbaa
	bbbbaaaaababbbabaaabbaaa
	aaabbaaabbabbbbaaaabbaaabbbaabbbaabaabbb
	babababaaaaaaababbabbbaaababaabaaaaaabbb
	aaabbbabaaababaabbabaabb
	bbabbbbabaaaabbbbbbaabbb
	abbbbabbbaabaaabababaaba
	aabbbaaabbabbaaaaabababb
	bbbbbbaabbbaaabbbbbabaaa
	aaabbabbabaabbbabbabbaabbbbbaabaabbbaaabbbbbaaabbabbbbaa
	aabaaabaabaabbabaabbbbbaaabbbaaaaaabaabbbaabbbaa
	abaaaaaaabbbabbbbbbbbbaaaaaaaaaa
	ababbbaabaaaababbaaaaaaabaaaaaaaaaaabaabbbababbb
	bbabbbbbabaababababaaaabaabaaaaaaaabbaaaabababbaaabbaabbbabbaabababbabaaababbbba
	baaababababaaaaaaabaabbaaaababaaabbaabaa
	bbabbaaabbabbbaaabbaabbaababbbbb
	babbbbbbbabaabbbabbabaab
	bbaaaaabaabaababaabbabbabaabbaababaabbbbbbbbabbaabbbabbaaabaaaab
	babaabbbababbbaabbaaabba
	bbaaababbbabbbaabbabaabb
	abaaaabbbabaaaaaababbaba
	bbbaaabbbaabababbbababbabaabbaaaabaaaababaabbbbbbbbaabbb
	bbaaabbababbbbbbbbaaaabb
	ababbbabbaaaababbbabaaba
	abbaaaaaabaabbbbababbaaa
	aaaaaabbbbaabbbabaaabbba
	babbbaabaaabbbabbbbbaabbaaaaaaabbaabbaba
	babaaabaababbaaabbaaababbbbaabbbaabbaaabaaaaabbababbaaabaabaabbaaabaaabbabbbbaaa
	bababbbbababaaabaababaab
	bbbbbaaabbbabbabbbabbbbaaaabaaabbbaaaaba
	baabbbbbbbabbaabbbbaaaab
	bbbbaaaababbbbbababbabaa
	abbbbbbbbabbbaabbaaababbbbabbbbb
	babbbabbbbbbbabbbbbabaaaabbbbbaababbbbabbabbaaabaabbabaa
	babaabbaaabbabbbaabaabbb
	bbbbabbbbbbaaaaaabbaabaa
	babbabababbabaaaabababab
	baabbaababbbbaabababbaba
	aabbabbaaabbbbbbbaaabbbb
	aabbabbaabbaaaaababaabbaabaabaaaaabababa
	bbabaaabaaaaababbabbabbaaabbbabb
	abababbabbbbabbbabbbabba
	bababaaababaaaaabbbababbbababababababaaa
	abbbbbabbababbbaabaaabbbabbaaaabbbbbabbbaaaabaaaabbbbaba
	babbbbbabaaaaaaabbabababaabbaabaaabbabaa
	aabbaaababababaaabaaaaaabbaaababaaabaaab
	aaaaaabaabbbbaabbbabaabb
	abbabababbaaabbbbabaabab
	aabaaababababbbaabaaaaba
	aaaabaabbabababaabababbaaabbbaabbaaaabba
	babaabaaabbabababbbbbbba
	babaabbabbbbaaaaaabbabbaabbabbbabaabaaba
	aaaabaababaaaabbbaabbbba
	aaaababababbaabbaaaabbab
	aaabbabbbbbbaababaabbbabbbaaabbbaababbbbbbaabbba
	bbaaabbbaaabbaaabbbababbabbbbaabaaabbbbbbabbaabaabbbabbabaabbbbaaaaaaaaa
	aaaaabbbabaabbbbbaaababbbbbbabbaabbababababaaabb
	baabbbabaabaabbabbabaabb
	abaaabbbaababaaaaababbba
	bbababbaaaaabbaaaaaaaabaaaaabbab
	aabaababbabababbabbbbaaa
	abbaaaabbaaabbaaaaaabbba
	babaababbaabbbbabbaababa
	babaabbbaaaabaababaabbbabaaaaabbbbabbbbbbaabaaba
	abbbbaabababbbaabaabbaaa
	bbbabbababaaaaaaabbbbbbbabbabaab
	baababbbbaaaababbabbbbab
	bbbaabaabbabababaaabaaabaababbbbbbbbabba
	abbbaabbababbbbaabbaaabaaabaaaaababbaaabbbbbbababbbbbbaaabaabbbaabaabbabbbabbabababababb
	bbabbaababbaabbabaaaabaa
	bbabbaaaaabaababababbbbb
	bbabbaaabbabbabbbbbabaaaabbbbabb
	aabbbbbbbababaaabaaaabba
	baabaaaaaabbbaaabaaaaabbabaaabababbbabba
	abaabbababbaaaabbbbababbbabbababbbbaaaab
	abbaaaabaaaabbaabbabaaabababbaaa
	aaaaaabababbbbbbbabbbbbaaaabababbaabbbababbaabbb
	babababbaabbabbabaaaabab
	bbaaaaaaaababbabbbababbabbaabbabbbababbabaaaabaaaaaabbab
	abaabbbabbbaabaaabaabbaa
	abbbbbbbbabaabbababaaaabaaaaaabaabbabbbababababbaababbbbababaaaa
	babababaaaabbbbbbbbaaababababbbbabaaaaaaaaabbbba
	aabbbbbbbaaaaaabaabbaaaabababbababbbabbabaababaa
	babaaaaabaaaabbbbaaabbbb
	babaaaabbbbbaaaaaaabbbaa
	abbaaaaaabbaaaababaabbaa
	abbbbbbabbbaabaabbbbaaaaaaabbbabaaaaabba
	bbbbabbbbbbbabbaabbaabbbbbbabbbbaaabbbabaababbbaaabbbaabbaababba
	babbaaaaababbbbaaababbbb
	aaaaababbababbabababbbababaaabaaaaaaabba
	bbbbaabaabbbbbbbaaaabaabbbaabbabbbaababbababbaaa
	bbaabaabbaabbbbbbabbbaba
	bababbbaaaaaababbbaaababaaabbaaaabbbbbaa
	bbbbaabbaaaabbaaabbbbaabaabbabbaabbbabbbaabaabaa
	bbbaaababbbbaaaabbaaabaa
	bbbbbbaaabaaabbbabaaaaba
	abbbbbabbabaabaabbabbbab
	aaaababaababaaabbabaabaabbaaabaa
	abbabaaabbbbbaabbababbaa
	bbbbabbbababbbbaaabaabaa
	abbaabbaaabbbbabbaaabbbb
	bbbbabbbababbbbaababaabbaabbabaa
	abbabbbaabbbbabbababaaaa
	abaabbbbbabababbaaabbbaa
	bbbababaabaabaabbaabaababbbabbaa
	aabbabbbbaababbababaabbbabbbbaba
	aaabaaaaaabbaaabbbbabbbb
	aaaabbaaabbbbaabaaaabaaa
	baaaababaabbbababbbbbbaababbbaabbbabbbaaaabbbbababbbaaba
	bbbbabbbaabbaabbbabbbaba
	bbabbababaaabababbbaaaba
	abbbbbbabaabbaabaaaaaabb
	babbbbabaaabaabbaaabaaaabbababbbbabbabaabaabababbaabbaaababaabaa
	abbbbaabbaaaaabbabaabaab
	abbbaabbbaababbbbaaababababbabbb
	aabaabbaaabbaaaaabbbabba
	bbbabaaaababababbabbbbababbabaaaaabbbbbbaaaabbbbabaabbbbbbbbbaaa
	bababaaababaaaababaaaabbabbaaaaaabbabbaaabbbabba
	babbaabbbaaaababbbbabbba
	bbbabababbabbbbbabaaabbaaababbbbaaababbbaaaabaaabbbbbaababbabbaabbaaaabb
	baaaabbbabaaaaabbbbbaababbaabbbbaaabbaba
	aabbabbabbbababbbaaaaaaaaabababbbaabbaaa
	bbaaaaabbbbbaaaabaaabbaabababaabaaaaabbb
	aaabbaababaaaabbbbbababbbbbbbabaabbbbbbbbbabbbbb
	aaabbaaaabbbabbbababbaaa
	abababbabaabbaabbababbabaabaabbabbbbaabbbabbbbaa
	aaababbbaabaaabaabababbaababbbabbabbaaba
	bbbaabbaabbaabababbbbaaa
	abbbaabbaaababaabbbabaabbaabbababaaabbbb
	abbbbbabaabaaabababbbbab
	abbaaaababaabbbaabaabbaa
	aaabbabbaaabbaaaababbaab
	abbaaaabaaabbaaababababaabaaaabbbaabbabaaaaabaaa
	baabaaabbaababbababbbbbbababbaababbaaabb
	abbbbbabbaaaaabbbabbaabbabaaabbbabaababaabbaaaabbaaaaaaaabbbabaabbbaabaa
	bbbababbbbbbabababbaabbaabaaaaaababaabbbaaaabaabaabababb
	babbaabbaaaaabaaaaaababb
	babaaabbaaaaaababbabbabb
	bbbbababbabaabbbaabaaabaaaabaaab
	bababaaaabbaaaaaabbabaab
	bababbbabbabbbaabbaabaaa
	aaaabbaabaabbaabababbaaa
	bbbbbaaababbaaabbbbaabbb
	babbbabbbababbbbabbaaabb
	baababbbabbbbaabbbabbbbb
	abbbbaabbababbbabbaaaaaa
	abababaabbbaabaaaabbbbbabaaaabaa
	bbabbabababaabbabaaaabaa
	babbbabbbaabaaabaababbbbbaabaaab
	aaabbbabbaaaabbbbbabababbaabbbbbbbababbbbabbabbb
	bbabbaabbabbbbbabbabbaabbbabbbabbbabbabbaaabaaba
	bbababababbababaaababbaa
	baaababababbabbabaaaaabbbbbaabbbbaabbbba
	aabbbababaababbbbaababbaabbbbaababbbabaa
	babbababaaabababbbabbabaaaabababaaaaababbabbbaaaaaababbaaaabbabaaababaab
	babbabbaaabbbbbabbabbaabbbbabaaaaabaabaa
	bbaabaaababbbbaaabbbbbbbaaabbaabbaabaaaaabaaabbbbabbabababaaaaaababaabab
	abbbbabbbbaaababaaabbaba
	abbabaaaabaabbbbbbbaaaab
	abaaabaaaaababaabbabbbab
	bbbbbaaabbabbbaaabbbbabbaaaababbabbababb
	bbbaaaaaaabaababbbaaabaa
	abbbbabbabababbabaabbbaa
	abaaabbaababaaabbbbbaabaaaabaaab
	abaabbbabbabbaabbababbbbbbabbbab
	abaabbbabbbbababbabaaaba
	bbabbaababaaaaaabbaabbabbabaabbbabbabbaa
	baaaaaaaabababbababababababbabababbabaaabbaabbbb
	aabbbbbababbabbabaaaaaababbbbbabaaabbbbaaaaabbba
	aaabbaabbbabbaababbbabbbabbabbbbbbbbbbab
	babaabaabbbaaabbbbaabbba
	abaabbbaabbaaababbaabaaa
	bbbbabbbaabaaabaabbaaabababbababaabbaaaaaabbbbbbbaaababbaaaaaabbbbaaaababbaaaaaa
	aaaaababbabaabbaabaabbaa
	abaabababbbbbabababbabbaaaabbbba
	baaaababbbabbbbaaababbab
	bbbabbabaabbaabbabbaaabb
	abaaaababaabbbbabbbbbbbbaaaaaaaababaaaabbbaaaaaabbbbaaabbbbababb
	`,
}
