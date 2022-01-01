package rpg

import "fmt"

type Character struct {
	Name      string
	Damage    int
	Armor     int
	HitPoints int
}

func NewCharacter(name string, damage, armor, hitPoints int) *Character {
	return &Character{
		Name:      name,
		Damage:    damage,
		Armor:     armor,
		HitPoints: hitPoints,
	}
}

func (c Character) String() string {
	return fmt.Sprintf("%s: %d damage, %d armor, %d hit-points", c.Name, c.Damage, c.Armor, c.HitPoints)
}

func (c *Character) Attack(oc *Character) {
	dam := c.Damage - oc.Armor
	if dam < 1 {
		dam = 1
	}
	oc.HitPoints -= dam
}
