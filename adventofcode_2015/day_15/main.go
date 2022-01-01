package main

import (
	"adventofcode_2015/comb"
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

var inputTest = `
Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
`

var input = `
Sprinkles: capacity 5, durability -1, flavor 0, texture 0, calories 5
PeanutButter: capacity -1, durability 3, flavor 0, texture 0, calories 1
Frosting: capacity 0, durability -1, flavor 4, texture 0, calories 6
Sugar: capacity -1, durability 0, flavor 0, texture 2, calories 8
`

func ParseIngredient(s string) (Ingredient, error) {
	var ing Ingredient
	_, err := fmt.Sscanf(s, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
		&ing.Name, &ing.Capacity, &ing.Durability, &ing.Flavor, &ing.Texture, &ing.Calories)
	if err != nil {
		return Ingredient{}, err
	}
	ing.Name = strings.TrimSuffix(ing.Name, ":")
	return ing, nil
}

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func main() {
	in := input
	scanner := bufio.NewScanner(bytes.NewBufferString(in))
	var ings []Ingredient
	for scanner.Scan() {
		s := strings.Trim(scanner.Text(), " \r\n\t")
		if s == "" {
			continue
		}
		ing, err := ParseIngredient(s)
		if err != nil {
			panic(err)
		}
		ings = append(ings, ing)
	}
	fmt.Printf("scanned %d ingredients\n", len(ings))

	// res := mix(ings, []int{44, 56})
	// fmt.Printf("makes %d\n", res)
	var maxScore int
	var maxWeights []int
	for ws := range comb.FixedSumInts(100, len(ings)) {
		score, cals := mix(ings, ws)
		if cals == 500 && score > maxScore {
			maxScore = score
			maxWeights = ws
		}
	}
	fmt.Printf("max-score: %d, weights: %v\n", maxScore, maxWeights)
}

func mix(ings []Ingredient, weights []int) (score int, cals int) {
	if len(ings) != len(weights) {
		panic("fook!")
	}
	cals = 0
	var cap, dur, fla, tex int
	for i, ing := range ings {
		w := weights[i]
		cap += w * ing.Capacity
		dur += w * ing.Durability
		fla += w * ing.Flavor
		tex += w * ing.Texture
		cals += w * ing.Calories
	}
	cap = cutNeg(cap)
	dur = cutNeg(dur)
	fla = cutNeg(fla)
	tex = cutNeg(tex)
	score = cap * dur * fla * tex

	return
}

func cutNeg(n int) int {
	if n < 0 {
		return 0
	}
	return n
}
