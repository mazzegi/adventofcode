package dinner

import (
	"fmt"
	"strings"
)

func ParseRule(s string) (Rule, error) {
	s = strings.Trim(s, ".")
	s = strings.ReplaceAll(s, "would ", "")
	s = strings.ReplaceAll(s, "happiness units by sitting next to ", "")

	var gainLose string
	r := Rule{}
	_, err := fmt.Sscanf(s, "%s %s %d %s", &r.Who, &gainLose, &r.Score, &r.Next)
	if err != nil {
		return Rule{}, err
	}
	if gainLose == "lose" {
		r.Score = -r.Score
	}
	return r, nil
}

type Rule struct {
	Who   string
	Score int
	Next  string
}

func (r Rule) Pairing() string {
	return r.Who + ":" + r.Next
}

func (r Rule) String() string {
	return fmt.Sprintf("%s: %d next to %s", r.Who, r.Score, r.Next)
}

type Rules []Rule

type Arrangement struct {
	rules      Rules
	ruleLookup map[string]int
	positions  []string
}

func NewArrangement(rules Rules) *Arrangement {
	a := &Arrangement{
		rules:      rules,
		ruleLookup: map[string]int{},
	}
	placed := map[string]bool{}
	for _, r := range rules {
		a.ruleLookup[r.Pairing()] = r.Score

		if placed[r.Who] {
			continue
		}
		a.positions = append(a.positions, r.Who)
		placed[r.Who] = true
	}
	return a
}

func (a *Arrangement) Positions() []string {
	return a.positions
}

func (a *Arrangement) PairScore(who, next string) int {
	s, ok := a.ruleLookup[who+":"+next]
	if !ok {
		return 0
	}
	return s
}

func (a *Arrangement) Score() int {
	return a.scorePositions(a.positions)
}

func (a *Arrangement) scorePositions(positions []string) int {
	var score int
	for i := 0; i < len(positions); i++ {
		prev := i - 1
		if prev < 0 {
			prev = len(positions) - 1
		}
		next := i + 1
		if next >= len(positions) {
			next = 0
		}
		who := positions[i]
		prevP := positions[prev]
		nextP := positions[next]
		score += a.PairScore(who, prevP)
		score += a.PairScore(who, nextP)
	}
	return score
}

func (a *Arrangement) FindBest() {
	maxScore := a.Score()
	best := fmt.Sprintf("%v", a.positions)
	first := a.positions[0]
	rest := a.positions[1:]
	for perm := range Permutations(rest) {
		test := append([]string{first}, perm...)
		score := a.scorePositions(test)
		if score > maxScore {
			maxScore = score
			best = fmt.Sprintf("%v", test)
		}
	}
	fmt.Printf("best is %d, %s\n", maxScore, best)
}
