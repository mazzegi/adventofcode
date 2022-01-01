package main

import (
	"adventofcode_2020/ints"
	"adventofcode_2020/jolt"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//in, err := readInput(inputTest2)
	in, err := readInput(input)
	if err != nil {
		panic(err)
	}
	if ints.HasDuplicate(in) {
		panic("duplicate")
	}
	fmt.Printf("no duplicates\n")

	min, max := ints.Min(in), ints.Max(in)
	if min > 0 {
		in = append(in, 0)
	}
	in = append(in, max+3)

	chain, ok := jolt.FindChain(in)
	if !ok {
		fmt.Printf("found no chain in %v\n", in)
		os.Exit(0)
	}
	fmt.Printf("found chain: %v\n", chain)
	diffs := ints.Diffs(chain)
	hist := ints.Hist(diffs)
	for k, v := range hist {
		fmt.Printf("%d -> %d times\n", k, v)
	}
	res := hist[1] * hist[3]
	fmt.Printf("result: %d\n", res)

	// fmt.Printf("**** Part2 ****\n")
	// cnt := jolt.Arrangements(in)
	// fmt.Printf("combi:    %d arrangements\n", cnt)

	// fmt.Printf("\n**** Part2 EX ****\n")
	// cnt = jolt.ArrangementsEx(in)
	// fmt.Printf("combi-ex: %d arrangements\n", cnt)
	ogs := jolt.Find1Groups(in)
	tot := int64(1)
	fmt.Printf("one-groups: ****\n")
	for _, og := range ogs {
		perms := og.Permutations()
		fmt.Printf("%v (%d)\n", og.Values(), perms)
		tot *= perms
	}
	fmt.Printf("tot: %d", tot)
	fmt.Printf("\n")
}

func readInput(s string) ([]int, error) {
	buf := bytes.NewBufferString(s)
	scanner := bufio.NewScanner(buf)
	in := []int{}
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		n, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			return nil, err
		}
		in = append(in, int(n))
	}
	return in, nil
}

var inputTest1 = `
16
10
15
5
1
11
7
19
6
12
4
`

var inputTest2 = `
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`

var input = `
80
87
10
122
57
142
134
59
113
139
101
41
138
112
46
96
43
125
36
54
133
17
42
98
7
114
78
67
77
28
149
58
20
105
31
19
18
27
40
71
117
66
21
72
146
90
97
94
123
1
119
30
84
61
91
118
2
29
104
73
13
76
24
148
68
111
131
83
49
8
132
9
64
79
124
95
88
135
3
51
39
6
60
108
14
35
147
89
34
65
50
145
128
`
