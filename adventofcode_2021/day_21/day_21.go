package day_21

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/intutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := practiceCoefficient(9, 4)
	//res, err := practiceCoefficient(4, 8)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := winUniverseCount(9, 4)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type detDice struct {
	next int
}

func newDetDice() *detDice {
	return &detDice{
		next: 1,
	}
}

func (d *detDice) roll() int {
	defer func() {
		d.next++
		if d.next > 100 {
			d.next = 1
		}
	}()
	return d.next
}

type player struct {
	pos   int
	score int
}

func (p *player) clone() *player {
	return &player{
		pos:   p.pos,
		score: p.score,
	}
}

func newPlayer(pos int) *player {
	return &player{
		pos:   pos,
		score: 0,
	}
}

func (p *player) move(steps int) {
	for i := 0; i < steps; i++ {
		p.pos++
		if p.pos > 10 {
			p.pos = 1
		}
	}
	p.score += p.pos
}

func practiceCoefficient(start1, start2 int) (int, error) {
	d := newDetDice()
	p1 := newPlayer(start1)
	p2 := newPlayer(start2)

	rolled := 0
	rollThree := func() int {
		v := d.roll()
		v += d.roll()
		v += d.roll()
		rolled += 3
		return v
	}

	var losing *player
	for {
		steps1 := rollThree()
		p1.move(steps1)
		if p1.score >= 1000 {
			losing = p2
			break
		}

		steps2 := rollThree()
		p2.move(steps2)
		if p2.score >= 1000 {
			losing = p1
			break
		}
	}

	c := losing.score * rolled
	return c, nil
}

func winUniverseCount(start1, start2 int) (int, error) {
	p1 := newPlayer(start1)
	p2 := newPlayer(start2)
	var p1w, p2w int

	hashMap := map[psHash]psWins{}
	playOne(hashMap, 0, p1, p2, 21, 1, &p1w, &p2w)

	log("p1W = %d, p2w = %d", p1w, p2w)

	winCount := intutil.MaxInt(p1w, p2w)
	return winCount, nil
}

var threeRollsResults = []int{
	3, // 1,1,1
	4, // 1,1,2
	5, // 1,1,3

	4, // 1,2,1
	5, // 1,2,2
	6, // 1,2,3

	5, // 1,3,1
	6, // 1,3,2
	7, // 1,3,3

	4, // 2,1,1
	5, // 2,1,2
	6, // 2,1,3

	5, // 2,2,1
	6, // 2,2,2
	7, // 2,2,3

	6, // 2,3,1
	7, // 2,3,2
	8, // 2,3,3

	5, // 3,1,1
	6, // 3,1,2
	7, // 3,1,3

	6, // 3,2,1
	7, // 3,2,2
	8, // 3,2,3

	7, // 3,3,1
	8, // 3,3,2
	9, // 3,3,3
}

// type totalResult struct {
// 	steps int
// 	count int
// }

// var threeRollsTotalResults = []totalResult{
// 	{3, 1},
// 	{4, 3},
// 	{5, 6},
// 	{6, 7},
// 	{7, 6},
// 	{8, 3},
// 	{9, 1},
// }

type psHash struct {
	pos1   int
	score1 int
	pos2   int
	score2 int
}

type psWins struct {
	p1Wins int
	p2Wins int
}

//p1W = 306621346123766, p2w = 166105651528183
//      444356092776315        341960390180808

func playOne(hashMap map[psHash]psWins, level int, player1, player2 *player, winScore int, pathCount int, player1Wins *int, player2Wins *int) {
	//log("%d: play-one: p1 (%d, %d); p2 (%d, %d)", level, player1.pos, player1.score, player2.pos, player2.score)

	trials := 0
	wins := 0
roll1:
	for i1, res1 := range threeRollsResults {
		p1 := player1.clone()

		p1.move(res1)
		if p1.score >= winScore {
			(*player1Wins) += 1

			trials++
			wins++
			continue roll1
		}

		//
	roll2:
		for i2, res2 := range threeRollsResults {
			p2 := player2.clone()
			trials++
			p2.move(res2)
			if p2.score >= winScore {
				//(*player2Wins) += res2.count * p2PathCount
				(*player2Wins) += 1
				wins++
				continue roll2
			}

			var p1Wins, p2Wins int
			hash := psHash{
				pos1:   p1.pos,
				score1: p1.score,
				pos2:   p2.pos,
				score2: p2.score,
			}
			if hWins, ok := hashMap[hash]; ok {
				p1Wins, p2Wins = hWins.p1Wins, hWins.p2Wins
			} else {
				playOne(hashMap, level+1, p1, p2, winScore, 1, &p1Wins, &p2Wins)
				hashMap[hash] = psWins{
					p1Wins: p1Wins,
					p2Wins: p2Wins,
				}
			}

			(*player1Wins) += p1Wins
			(*player2Wins) += p2Wins

			_ = i1
			_ = i2
		}
	}
	log("%d: play-one: %d trials, %d wins (%d, %d)", level, trials, wins, *player1Wins, *player2Wins)
}

// func playOne(hashMap map[psHash]psWins, level int, player1, player2 *player, winScore int, pathCount int, player1Wins *int, player2Wins *int) {
// 	//log("%d: play-one: p1 (%d, %d); p2 (%d, %d)", level, player1.pos, player1.score, player2.pos, player2.score)

// 	trials := 0
// 	wins := 0
// roll1:
// 	for i1, res1 := range threeRollsTotalResults {
// 		p1 := player1.clone()

// 		p1.move(res1.steps)
// 		if p1.score >= winScore {
// 			(*player1Wins) += res1.count * pathCount

// 			trials++
// 			wins++
// 			continue roll1
// 		}

// 		//
// 	roll2:
// 		for i2, res2 := range threeRollsTotalResults {
// 			p2 := player2.clone()
// 			trials++
// 			p2.move(res2.steps)
// 			if p2.score >= winScore {
// 				//(*player2Wins) += res2.count * p2PathCount
// 				(*player2Wins) += res1.count * res2.count * pathCount
// 				wins++
// 				continue roll2
// 			}

// 			var p1Wins, p2Wins int
// 			hash := psHash{
// 				pos1:   p1.pos,
// 				score1: p1.score,
// 				pos2:   p2.pos,
// 				score2: p2.score,
// 			}
// 			if hWins, ok := hashMap[hash]; ok {
// 				p1Wins, p2Wins = hWins.p1Wins, hWins.p2Wins
// 			} else {
// 				playOne(hashMap, level+1, p1, p2, winScore, 1, &p1Wins, &p2Wins)
// 				hashMap[hash] = psWins{
// 					p1Wins: p1Wins,
// 					p2Wins: p2Wins,
// 				}
// 			}

// 			(*player1Wins) += p1Wins * res1.count * res2.count * pathCount
// 			(*player2Wins) += p2Wins * res1.count * res2.count * pathCount

// 			_ = i1
// 			_ = i2
// 		}
// 	}
// 	log("%d: play-one: %d trials, %d wins (%d, %d)", level, trials, wins, *player1Wins, *player2Wins)
// }

// func playOne(level int, player1, player2 *player, winScore int, pathCount int, player1Wins *int, player2Wins *int) {
// 	//log("%d: play-one: p1 (%d, %d); p2 (%d, %d)", level, player1.pos, player1.score, player2.pos, player2.score)

// 	trials := 0
// 	wins := 0
// roll1:
// 	for i1, res1 := range threeRollsTotalResults {
// 		p1 := player1.clone()

// 		p1.move(res1.steps)
// 		if p1.score >= winScore {
// 			(*player1Wins) += res1.count * pathCount

// 			trials++
// 			wins++
// 			continue roll1
// 		}

// 		//
// 	roll2:
// 		for i2, res2 := range threeRollsTotalResults {
// 			p2 := player2.clone()
// 			trials++
// 			p2.move(res2.steps)
// 			if p2.score >= winScore {
// 				//(*player2Wins) += res2.count * p2PathCount
// 				(*player2Wins) += res1.count * res2.count * pathCount
// 				wins++
// 				continue roll2
// 			}

// 			playOne(level+1, p1, p2, winScore, res1.count*res2.count*pathCount, player1Wins, player2Wins)
// 			_ = i1
// 			_ = i2
// 		}
// 	}
// 	log("%d: play-one: %d trials, %d wins (%d, %d)", level, trials, wins, *player1Wins, *player2Wins)
// }

// func playOne(level int, player1, player2 *player, winScore int, player1Wins *int, player2Wins *int) {
// 	//log("%d: play-one: p1 (%d, %d); p2 (%d, %d)", level, player1.pos, player1.score, player2.pos, player2.score)

// 	trials := 0
// 	wins := 0
// roll1:
// 	for i1, res1 := range threeRollsResults {
// 		p1 := player1.clone()

// 		p1.move(res1)
// 		if p1.score >= winScore {
// 			(*player1Wins)++
// 			//log("p1-probe-wins %d: steps %d", i1, res1)
// 			trials++
// 			wins++
// 			continue roll1
// 		}
// 		//log("p1-probe-no-win %d: steps %d", i1, res1)

// 		//
// 	roll2:
// 		for i2, res2 := range threeRollsResults {
// 			p2 := player2.clone()
// 			trials++
// 			p2.move(res2)
// 			if p2.score >= winScore {
// 				(*player2Wins)++
// 				//log("p2-probe-wins %d: steps %d", i2, res2)
// 				wins++
// 				continue roll2
// 			}
// 			//log("p2-probe-no-win %d: steps %d", i2, res2)

// 			playOne(level+1, p1, p2, winScore, player1Wins, player2Wins)
// 			_ = i1
// 			_ = i2
// 		}
// 	}
// 	log("%d: play-one: %d trials, %d wins (%d, %d)", level, trials, wins, *player1Wins, *player2Wins)
// }
