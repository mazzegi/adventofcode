package main

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2015/rpg"
)

func newBoss() *rpg.Character {
	return rpg.NewCharacter("boss", 8, 2, 100)
}

func main() {
	for dam := 4; dam <= 9; dam++ {
		playerWon := false
		for arm := 0; arm <= 10; arm++ {
			p := rpg.NewCharacter("player", dam, arm, 100)
			res := play(p, newBoss())
			if res == 1 && !playerWon {
				//fmt.Printf("first win with dam=%d: arm=%d => player won\n", dam, arm)
				fmt.Printf("last boss one: dam=%d: arm=%d => boss won\n", dam, arm-1)
				playerWon = true
				rpg.NewStore().CheckExpenciest(dam, arm-1)
			} else {
				//fmt.Printf("dam=%d, arm=%d => boss won\n", dam, arm)
			}
		}
	}
}

func play(c1 *rpg.Character, c2 *rpg.Character) int {
	for {
		c1.Attack(c2)
		if c2.HitPoints <= 0 {
			return 1
		}
		c2.Attack(c1)
		if c1.HitPoints <= 0 {
			return 2
		}
	}
}

func playTest() {
	boss := rpg.NewCharacter("boss", 7, 2, 12)
	player := rpg.NewCharacter("player", 5, 5, 8)
	for {
		player.Attack(boss)
		fmt.Printf("%s vs. %s\n", player, boss)
		if boss.HitPoints <= 0 {
			break
		}
		boss.Attack(player)
		fmt.Printf("%s vs. %s\n", boss, player)
		if player.HitPoints <= 0 {
			break
		}
	}
}
