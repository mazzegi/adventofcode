package marble

import (
	"fmt"
	"strings"
)

type Marble struct {
	Value int
	Prev  *Marble
	Next  *Marble
}

type Player struct {
	id    int
	score int
}

type Game struct {
	currPlayer int
	players    []*Player
	First      *Marble
	Current    *Marble
}

func NewGame(players int) *Game {
	g := &Game{
		currPlayer: -1,
		players:    make([]*Player, players),
	}
	for p := 0; p < players; p++ {
		g.players[p] = &Player{
			id:    p + 1,
			score: 0,
		}
	}
	return g
}

func (g *Game) Highscore() {
	var max int
	var maxPlayer int
	for i := 0; i < len(g.players); i++ {
		p := g.players[i]
		if p.score > max {
			max = p.score
			maxPlayer = p.id
		}
	}
	fmt.Printf("highscore: %d by player %d\n", max, maxPlayer)
}

func (g *Game) Add(n int) {
	if g.First == nil {
		nm := &Marble{
			Value: n,
		}
		g.First = nm
		g.First.Prev = nm
		g.First.Next = nm
		g.Current = g.First
		return
	}

	g.currPlayer++
	if g.currPlayer >= len(g.players) {
		g.currPlayer = 0
	}

	if n%23 == 0 {
		g.Handle23(n)
		return
	}

	//
	nm := &Marble{
		Value: n,
	}
	n1 := g.Current.Next
	n2 := n1.Next
	nm.Prev = n1
	nm.Next = n2
	nm.Prev.Next = nm
	nm.Next.Prev = nm
	g.Current = nm
}

func (g *Game) Handle23(n int) {
	g.players[g.currPlayer].score += n
	//rotate 7 times prev
	rm := g.Current
	for i := 0; i < 7; i++ {
		rm = rm.Prev
	}
	g.players[g.currPlayer].score += rm.Value
	g.Current = rm.Next

	//remove
	prev := rm.Prev
	next := rm.Next
	prev.Next = next
	next.Prev = prev
}

func (g *Game) Dump() string {
	m := g.First
	ms := []string{}
	for {
		if m == g.Current {
			ms = append(ms, fmt.Sprintf("*%d", m.Value))
		} else {
			ms = append(ms, fmt.Sprintf(" %d", m.Value))
		}
		m = m.Next
		if m == g.First {
			break
		}
	}
	return fmt.Sprintf("%d: %s", g.currPlayer+1, strings.Join(ms, " "))
}
