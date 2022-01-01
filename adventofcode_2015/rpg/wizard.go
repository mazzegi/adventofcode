package rpg

import (
	"sort"

	"github.com/mazzegi/wasa/errors"
)

type Wizard struct {
	Armor         int
	Mana          int
	HitPoints     int
	ActiveEffects map[string]*Effect
}

func NewWizard(mana int, hitPoints int) *Wizard {
	return &Wizard{
		Mana:          mana,
		HitPoints:     hitPoints,
		ActiveEffects: map[string]*Effect{},
	}
}

func (w *Wizard) DumpActiveEffects() []string {
	var aes []string
	for e := range w.ActiveEffects {
		aes = append(aes, e)
	}
	sort.Strings(aes)
	return aes
}

func (w *Wizard) IsActive(ef string) bool {
	_, ok := w.ActiveEffects[ef]
	return ok
}

func (w *Wizard) AvailableSpells() []string {
	var spells []string
	for _, e := range Effects {
		if e.Cost < w.Mana && !w.IsActive(e.Name) {
			spells = append(spells, e.Name)
		}
	}
	return spells
}

func (w *Wizard) CheapestAvailableSpell() string {
	var cheapest string
	var minCost int
	for _, e := range Effects {
		if e.Cost < w.Mana && !w.IsActive(e.Name) {
			if cheapest == "" || e.Cost < minCost {
				cheapest = e.Name
				minCost = e.Cost
			}
		}
	}
	if cheapest == "" {
		panic("no spell available")
	}
	return cheapest
}

func (w *Wizard) NecessarySpell(c *Character) string {
	if _, ok := w.ActiveEffects[Poison]; ok {
		if c.HitPoints <= 7 {
			return MagicMissile
		}
	}

	if w.Mana <= 250 && w.IsAvailable(Recharge) {
		return Recharge
	}
	if w.HitPoints >= 2 && w.HitPoints <= 10 && w.IsAvailable(Shield) {
		return Shield
	}
	if w.HitPoints <= 1 && w.IsAvailable(Drain) {
		return Drain
	}
	if w.IsAvailable(Poison) {
		return Poison
	}
	return MagicMissile
}

func (w *Wizard) IsAvailable(spell string) bool {
	as := w.AvailableSpells()
	for _, a := range as {
		if a == spell {
			return true
		}
	}
	return false
}

func (w *Wizard) ApplyAttackEffects(c *Character) {
	for name, eff := range w.ActiveEffects {
		eff.OnAttack(w, c)
		eff.Timer--
		if eff.Timer <= 0 {
			delete(w.ActiveEffects, name)
		}
	}
}

func (w *Wizard) ApplyDefendEffects(c *Character) {
	for name, eff := range w.ActiveEffects {
		eff.OnDefend(w, c)
		eff.Timer--
		if eff.Timer <= 0 {
			delete(w.ActiveEffects, name)
		}
	}
}

func (w *Wizard) Attack(spell string, c *Character) {
	if !w.IsAvailable(spell) {
		panic("not available")
	}
	e, err := Cast(spell)
	if err != nil {
		panic(err)
	}
	w.Mana -= e.Cost

	if e.Timer == 0 {
		e.OnAttack(w, c)
	} else {
		w.ActiveEffects[e.Name] = e
	}
}

func (w *Wizard) Defend(c *Character) {
	cdam := c.Damage - w.Armor
	if cdam < 1 {
		cdam = 1
	}
	w.HitPoints -= cdam
	w.Armor = 0
}

const (
	MagicMissile = "Magic Missile"
	Drain        = "Drain"
	Shield       = "Shield"
	Poison       = "Poison"
	Recharge     = "Recharge"
)

type Effect struct {
	Name     string
	Cost     int
	Timer    int
	OnAttack func(w *Wizard, c *Character)
	OnDefend func(w *Wizard, c *Character)
}

var Effects = []Effect{
	{
		Name:  "Magic Missile",
		Cost:  53,
		Timer: 0,
		OnAttack: func(w *Wizard, c *Character) {
			c.HitPoints -= 4
		},
		OnDefend: func(w *Wizard, c *Character) {},
	},
	{
		Name:  "Drain",
		Cost:  73,
		Timer: 0,
		OnAttack: func(w *Wizard, c *Character) {
			c.HitPoints -= 2
			w.HitPoints += 2
		},
		OnDefend: func(w *Wizard, c *Character) {},
	},
	{
		Name:     "Shield",
		Cost:     113,
		Timer:    6,
		OnAttack: func(w *Wizard, c *Character) {},
		OnDefend: func(w *Wizard, c *Character) {
			w.Armor = 7
		},
	},
	{
		Name:  "Poison",
		Cost:  173,
		Timer: 6,
		OnAttack: func(w *Wizard, c *Character) {
			c.HitPoints -= 3
		},
		OnDefend: func(w *Wizard, c *Character) {
			c.HitPoints -= 3
		},
	},
	{
		Name:  "Recharge",
		Cost:  229,
		Timer: 5,
		OnAttack: func(w *Wizard, c *Character) {
			w.Mana += 101
		},
		OnDefend: func(w *Wizard, c *Character) {
			w.Mana += 101
		},
	},
}

func Cast(spell string) (*Effect, error) {
	for _, e := range Effects {
		if e.Name == spell {
			return &e, nil
		}
	}
	return nil, errors.Errorf("unknown spell %q", spell)
}
