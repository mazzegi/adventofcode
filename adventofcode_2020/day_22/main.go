package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2020/combat"
)

func main() {
	//in := inputTest
	in := input
	scanner := bufio.NewScanner(bytes.NewBufferString(in))
	var player1Deck []int
	var player2Deck []int
	var player int
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "Player 1:") {
			player = 1
			continue
		} else if strings.HasPrefix(l, "Player 2:") {
			player = 2
			continue
		}
		n, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			panic(err)
		}
		switch player {
		case 1:
			player1Deck = append(player1Deck, int(n))
		case 2:
			player2Deck = append(player2Deck, int(n))
		}
	}
	fmt.Printf("player1: %v\n", player1Deck)
	fmt.Printf("player2: %v\n", player2Deck)

	combat.RecursivePlay(0, player1Deck, player2Deck)
}

var inputTest = `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

var input = `
Player 1:
7
1
9
10
12
4
38
22
18
3
27
31
43
33
47
42
21
24
50
39
8
6
16
46
11

Player 2:
49
41
40
35
44
29
30
19
14
2
34
17
25
5
15
32
20
48
45
26
37
28
36
23
13
`
