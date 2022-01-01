package main

import (
	"adventofcode_2018/marble"
)

func main() {
	// players := 9
	// lastMarble := 25

	players := 441
	lastMarble := 71032
	lastMarble *= 100

	g := marble.NewGame(players)
	for m := 0; m <= lastMarble; m++ {
		g.Add(m)
		//fmt.Printf("%s\n", g.Dump())
	}
	g.Highscore()
}
