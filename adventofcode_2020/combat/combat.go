package combat

import (
	"crypto/md5"
	"fmt"
)

type Deck struct {
	Cards []int
}

func (d *Deck) CopyCards(n int) []int {
	cc := make([]int, n)
	for i := 0; i < n; i++ {
		cc[i] = d.Cards[i]
	}
	return cc
}

func (d *Deck) Hash() string {
	bs := make([]byte, len(d.Cards))
	for i := 0; i < len(d.Cards); i++ {
		bs[i] = byte(d.Cards[i])
	}
	return fmt.Sprintf("%x", md5.Sum(bs))
}

func (d *Deck) IsEmpty() bool {
	return len(d.Cards) == 0
}

func (d *Deck) TakeFront() int {
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c
}

func (d *Deck) Push(cs ...int) {
	d.Cards = append(d.Cards, cs...)
}

func (d *Deck) Score() int {
	var sc int
	for i := len(d.Cards) - 1; i >= 0; i-- {
		sc += d.Cards[i] * (len(d.Cards) - i)
	}
	return sc
}

func Play(p1Cards, p2Cards []int) {
	d1 := &Deck{Cards: p1Cards}
	d2 := &Deck{Cards: p2Cards}

	for {
		c1 := d1.TakeFront()
		c2 := d2.TakeFront()
		if c1 == c2 {
			panic("cards equal")
		}

		if c1 > c2 {
			d1.Push(c1, c2)
		} else {
			d2.Push(c2, c1)
		}

		if d1.IsEmpty() || d2.IsEmpty() {
			break
		}
	}

	if d1.IsEmpty() {
		fmt.Printf("player 2 wins with a score of %d\n", d2.Score())
	} else {
		fmt.Printf("player 1 wins with a score of %d\n", d1.Score())
	}
}

func RecursivePlay(level int, p1Cards, p2Cards []int) int {
	d1 := &Deck{Cards: p1Cards}
	d2 := &Deck{Cards: p2Cards}
	d1Hashes := map[string]bool{}
	d2Hashes := map[string]bool{}
	d1Hashes[d1.Hash()] = true
	d2Hashes[d2.Hash()] = true

	fmt.Printf("level %d: play %v %v\n", level, d1.Cards, d2.Cards)

	for {
		c1 := d1.TakeFront()
		c2 := d2.TakeFront()
		if c1 == c2 {
			panic("cards equal")
		}

		//check recursion
		if len(d1.Cards) >= c1 && len(d2.Cards) >= c2 {
			fmt.Printf("level %d: play rec. (%d, %d) %v %v\n", level, c1, c2, d1.Cards, d2.Cards)
			winner := RecursivePlay(level+1, d1.CopyCards(c1), d2.CopyCards(c2))
			if winner == 1 {
				d1.Push(c1, c2)
			} else if winner == 2 {
				d2.Push(c2, c1)
			} else {
				panic("arrg")
			}
		} else {
			if c1 > c2 {
				d1.Push(c1, c2)
			} else {
				d2.Push(c2, c1)
			}
		}

		d1h := d1.Hash()
		d2h := d2.Hash()
		_, co1 := d1Hashes[d1h]
		_, co2 := d1Hashes[d2h]
		if co1 || co2 {
			fmt.Printf("level %d: recursion detected: player 1 wins with a score of %d\n", level, d1.Score())
			return 1
		}
		d1Hashes[d1h] = true
		d2Hashes[d2h] = true

		if d1.IsEmpty() || d2.IsEmpty() {
			break
		}
	}

	if d1.IsEmpty() {
		fmt.Printf("level %d: player 2 wins with a score of %d\n", level, d2.Score())
		return 2
	} else {
		fmt.Printf("level %d: player 1 wins with a score of %d\n", level, d1.Score())
		return 1
	}
}
