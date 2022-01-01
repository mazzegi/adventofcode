package day_15

import "fmt"

func Part1() {
	discs := Setup()
	dt := DropTime(discs)
	fmt.Printf("part1: the drop time is %d\n", dt)
}

func Part2() {
	discs := Setup()
	discs = append(discs, &Disc{
		numPositions: 11,
		currPosition: 0,
	})
	dt := DropTime(discs)
	fmt.Printf("part2: the drop time is %d\n", dt)
}

func Setup() []*Disc {
	ds := []*Disc{
		{17, 15},
		{3, 2},
		{19, 4},
		{13, 2},
		{7, 2},
		{5, 0},
	}
	return ds
}

type Disc struct {
	numPositions int
	currPosition int
}

func (d *Disc) MoveOne() {
	d.currPosition++
	if d.currPosition >= d.numPositions {
		d.currPosition = 0
	}
}

func (d *Disc) Move(steps int) {
	d.currPosition = (d.currPosition + steps) % d.numPositions
}

func (d *Disc) CanPass() bool {
	return d.currPosition == 0
}

func (d *Disc) CanPassIn(steps int) bool {
	return (d.currPosition+steps)%d.numPositions == 0
}

func CanPassAny(discs *[]*Disc) bool {
	// for _, d := range *discs {
	// 	d.Move(at)
	// }
	//
	for i, d := range *discs {
		if !d.CanPassIn(i + 1) {
			return false
		}
	}
	return true
}

func DropTime(discs []*Disc) int {
	max := 10000000
	for i := 0; i < max; i++ {
		if CanPassAny(&discs) {
			return i
		}
		for _, d := range discs {
			d.MoveOne()
		}
	}
	return -1
}
