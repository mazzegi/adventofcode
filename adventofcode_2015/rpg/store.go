package rpg

import "fmt"

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

func NewItem(name string, cost, damage, armor int) *Item {
	return &Item{
		Name:   name,
		Cost:   cost,
		Damage: damage,
		Armor:  armor,
	}
}

type Store struct {
	Weapons     map[string]*Item
	Armor       map[string]*Item
	DamageRings map[string]*Item
	ArmorRings  map[string]*Item
}

func NewStore() *Store {
	s := &Store{
		Weapons:     map[string]*Item{},
		Armor:       map[string]*Item{},
		DamageRings: map[string]*Item{},
		ArmorRings:  map[string]*Item{},
	}
	s.AddWeapon(NewItem("Dagger", 8, 4, 0))
	s.AddWeapon(NewItem("Shortsword", 10, 5, 0))
	s.AddWeapon(NewItem("Warhammer", 25, 6, 0))
	s.AddWeapon(NewItem("Longsword", 40, 7, 0))
	s.AddWeapon(NewItem("Greataxe", 74, 8, 0))

	s.AddArmor(NewItem("NoArmor", 0, 0, 0))
	s.AddArmor(NewItem("Leather", 13, 0, 1))
	s.AddArmor(NewItem("Chainmail", 31, 0, 2))
	s.AddArmor(NewItem("Splintmail", 53, 0, 3))
	s.AddArmor(NewItem("Bandedmail", 75, 0, 4))
	s.AddArmor(NewItem("Platemail", 102, 0, 5))

	s.AddDamageRing(NewItem("Damage +1", 25, 1, 0))
	s.AddDamageRing(NewItem("Damage +2", 50, 2, 0))
	s.AddDamageRing(NewItem("Damage +3", 100, 3, 0))

	s.AddArmorRing(NewItem("Defense +1", 20, 0, 1))
	s.AddArmorRing(NewItem("Defense +2", 40, 0, 2))
	s.AddArmorRing(NewItem("Defense +3", 80, 0, 3))

	return s
}

type RingPair struct {
	R1 *Item
	R2 *Item
}

func (rp RingPair) Cost() int {
	var c int
	if rp.R1 != nil {
		c += rp.R1.Cost
	}
	if rp.R2 != nil {
		c += rp.R2.Cost
	}
	return c
}

func (rp RingPair) NumRings() int {
	var c int
	if rp.R1 != nil {
		c++
	}
	if rp.R2 != nil {
		c++
	}
	return c
}

func (rp RingPair) Items() []*Item {
	var is []*Item
	if rp.R1 != nil {
		is = append(is, rp.R1)
	}
	if rp.R2 != nil {
		is = append(is, rp.R2)
	}
	return is
}

func (s *Store) ringsForDamage(dam int) []RingPair {
	if dam == 0 {
		return []RingPair{{}}
	}

	var rps []RingPair
	for _, r1 := range s.DamageRings {
		if r1.Damage == dam {
			rps = append(rps, RingPair{
				R1: r1,
			})
			continue
		}
		for _, r2 := range s.DamageRings {
			if r2.Name == r1.Name {
				continue
			}
			if r1.Damage+r2.Damage == dam {
				rps = append(rps, RingPair{
					R1: r1,
					R2: r2,
				})
			}
		}
	}
	return rps
}

type ArmorItems struct {
	Armor    *Item
	RingPair RingPair
}

func (ais ArmorItems) Cost() int {
	var c int
	if ais.Armor != nil {
		c += ais.Armor.Cost
	}
	c += ais.RingPair.Cost()
	return c
}

func (ais ArmorItems) Items() []*Item {
	var is []*Item
	if ais.Armor != nil {
		is = append(is, ais.Armor)
	}
	is = append(is, ais.RingPair.Items()...)
	return is
}

func (s *Store) armorItems(arm int, usedRings int) []ArmorItems {
	if arm == 0 {
		return []ArmorItems{{}}
	}

	var ais []ArmorItems
	for _, ai := range s.Armor {
		if ai.Armor == arm {
			ais = append(ais, ArmorItems{
				Armor: ai,
			})
			continue
		}
		if usedRings == 2 {
			continue
		}
		armLeft := arm - ai.Armor
		for _, r1 := range s.ArmorRings {
			if r1.Armor == armLeft {
				ais = append(ais, ArmorItems{
					Armor: ai,
					RingPair: RingPair{
						R1: r1,
					},
				})
				continue
			}
			if usedRings == 1 {
				continue
			}
			for _, r2 := range s.ArmorRings {
				if r2.Name == r1.Name {
					continue
				}
				if r2.Armor+r1.Armor == armLeft {
					ais = append(ais, ArmorItems{
						Armor: ai,
						RingPair: RingPair{
							R1: r1,
							R2: r2,
						},
					})
				}
			}
		}
	}

	return ais
}

func (s *Store) CheckCheapest(dam int, arm int) {
	//buy one weapon
	minCost := -1
	var minItems []*Item
	for _, w := range s.Weapons {
		if w.Damage > dam {
			continue
		}
		var cost int
		damLeft := dam - w.Damage
		cost += w.Cost

		drps := s.ringsForDamage(damLeft)
		if len(drps) == 0 {
			continue
		}

		for _, drp := range drps {
			ais := s.armorItems(arm, drp.NumRings())
			if len(ais) == 0 {
				continue
			}
			for _, ai := range ais {
				tcost := cost + drp.Cost() + ai.Cost()
				if minCost < 0 || tcost < minCost {
					minCost = tcost
					minItems = []*Item{w}
					minItems = append(minItems, drp.Items()...)
					minItems = append(minItems, ai.Items()...)
				}
			}
		}
	}
	fmt.Printf("    *** min-cost: %d\n", minCost)
	for _, i := range minItems {
		fmt.Printf("    %s\n", i.Name)
	}
}

func (s *Store) CheckExpenciest(dam int, arm int) {
	//buy one weapon
	maxCost := 0
	var maxItems []*Item
	for _, w := range s.Weapons {
		if w.Damage > dam {
			continue
		}
		var cost int
		damLeft := dam - w.Damage
		cost += w.Cost

		drps := s.ringsForDamage(damLeft)
		if len(drps) == 0 {
			continue
		}

		for _, drp := range drps {
			ais := s.armorItems(arm, drp.NumRings())
			if len(ais) == 0 {
				continue
			}
			for _, ai := range ais {
				tcost := cost + drp.Cost() + ai.Cost()
				if tcost > maxCost {
					maxCost = tcost
					maxItems = []*Item{w}
					maxItems = append(maxItems, drp.Items()...)
					maxItems = append(maxItems, ai.Items()...)
				}
			}
		}
	}
	fmt.Printf("    *** max-cost: %d\n", maxCost)
	for _, i := range maxItems {
		fmt.Printf("    %s\n", i.Name)
	}
}

func (s *Store) AddWeapon(i *Item) {
	s.Weapons[i.Name] = i
}

func (s *Store) AddArmor(i *Item) {
	s.Armor[i.Name] = i
}

func (s *Store) AddDamageRing(i *Item) {
	s.DamageRings[i.Name] = i
}

func (s *Store) AddArmorRing(i *Item) {
	s.ArmorRings[i.Name] = i
}
