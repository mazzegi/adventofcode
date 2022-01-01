package main

import (
	"adventofcode_2015/rpg"
	"fmt"
)

func newBoss() *rpg.Character {
	return rpg.NewCharacter("boss", 10, 0, 71)
}

func newTestBoss() *rpg.Character {
	return rpg.NewCharacter("test-boss", 8, 0, 14)
}

func newWizard() *rpg.Wizard {
	return rpg.NewWizard(500, 50)
}

func newTestWizard() *rpg.Wizard {
	return rpg.NewWizard(250, 10)
}

func main() {
	// w := newTestWizard()
	// b := newTestBoss()
	w := newWizard()
	b := newBoss()

	dump := func() {
		fmt.Printf("* wizard: %d hp, %d mana %v\n* boss  : %d hp\n", w.HitPoints, w.Mana, w.DumpActiveEffects(), b.HitPoints)
	}

	finished := func() bool {
		if w.HitPoints <= 0 {
			fmt.Printf("boss wins\n")
			return true
		}
		if b.HitPoints <= 0 {
			fmt.Printf("wizard wins\n")
			return true
		}
		return false
	}

	dump()
	fmt.Printf("\n")
	for {
		w.ApplyAttackEffects(b)
		if finished() {
			break
		}
		spell := w.NecessarySpell(b)
		if spell == "" {
			spell = w.CheapestAvailableSpell()
		}
		fmt.Printf("wizard casts %q and attacks\n", spell)
		w.Attack(spell, b)
		dump()
		fmt.Printf("\n")
		if finished() {
			break
		}

		fmt.Printf("boss attacks\n")
		w.ApplyDefendEffects(b)
		if finished() {
			break
		}
		w.Defend(b)
		dump()
		fmt.Printf("\n")
		if finished() {
			break
		}
	}

	// spells := []string{
	// 	rpg.Recharge,
	// 	rpg.Shield,
	// 	rpg.Drain,
	// 	rpg.Poison,
	// 	rpg.MagicMissile,
	// }

	// for _, spell := range spells {

	// 	fmt.Printf("wizard casts %q\n", spell)
	// 	w.Attack(spell, b)
	// 	dump()
	// 	fmt.Printf("boss attacks\n")
	// 	w.Defend(b)
	// 	dump()
	// 	fmt.Printf("\n")
	// }
	// finished()

	// dump()
	// w.Attack(rpg.Poison, b)
	// dump()
	// w.Defend(b)
	// dump()
	// w.Attack(rpg.MagicMissile, b)
	// dump()
	// w.Defend(b)
	// dump()

}
