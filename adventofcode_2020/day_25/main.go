package main

import (
	"fmt"
	"time"

	"github.com/mazzegi/adventofcode/adventofcode_2020/pki"
)

func main() {
	//in := inputTest
	in := input
	//findEncKey(in)

	t0 := time.Now()
	//pk := in.CardPK
	lsCard, ok := pki.FindLoopSizeForSubj(7, in.CardPK)
	if !ok {
		fmt.Printf("found no loop size for card (%s)\n", time.Since(t0))
	} else {
		fmt.Printf("card loop-size: %d (%s)\n", lsCard, time.Since(t0))
	}
	t0 = time.Now()
	lsDoor, ok := pki.FindLoopSizeForSubj(7, in.DoorPK)
	if !ok {
		fmt.Printf("found no loop size for door (%s)\n", time.Since(t0))
	} else {
		fmt.Printf("door loop-size: %d (%s)\n", lsDoor, time.Since(t0))
	}

	eCard := pki.Calc(in.DoorPK, lsCard)
	eDoor := pki.Calc(in.CardPK, lsDoor)
	fmt.Printf("e-card: %d\n", eCard)
	fmt.Printf("e-door: %d\n", eDoor)
}

func findEncKey(in Input) {
	maxLS := 10000
	for lsCard := 1; lsCard <= maxLS; lsCard++ {
		for lsDoor := 1; lsDoor <= maxLS; lsDoor++ {
			eCard := pki.Calc(in.DoorPK, lsCard)
			eDoor := pki.Calc(in.CardPK, lsDoor)
			if eCard == eDoor {
				fmt.Printf("found. ls-card=%d, ls-door=%d ek=%d\n", lsCard, lsDoor, eCard)
				return
			}
		}
	}
	fmt.Printf("found no ekey\n")
}

type Input struct {
	CardPK int
	DoorPK int
}

var inputTest = Input{
	CardPK: 5764801,
	DoorPK: 17807724,
}

var input = Input{
	CardPK: 9033205,
	DoorPK: 9281649,
}
