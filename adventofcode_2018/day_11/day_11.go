package day_11

import (
	"adventofcode_2018/errutil"
	"fmt"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := coordOfLargestPowerLevelSquare(8141)
	errutil.ExitOnErr(err)
	log("part1: result = %s", res)
}

func Part2() {
	res, err := coordOfLargestPowerLevelFlexSquare(8141)
	errutil.ExitOnErr(err)
	log("part1: result = %s", res)
}

//
type point struct {
	x, y int
}

// x,y start with 1!
func powerLevel(pt point, serial int) int {
	rackID := pt.x + 10
	pl := rackID * pt.y
	pl += serial
	pl *= rackID

	pl /= 100
	pl = pl % 10
	pl -= 5
	return pl
}

func (pt point) String() string {
	return fmt.Sprintf("%d,%d", pt.x, pt.y)
}

func coordOfLargestPowerLevelSquare(serial int) (point, error) {
	xsize, ysize := 300, 300

	var maxPower int
	var maxPowerPoint point
	for x := 1; x <= xsize-3; x++ {
		for y := 1; y <= ysize-3; y++ {
			//
			var sqPow int
			for xs := x; xs < x+3; xs++ {
				for ys := y; ys < y+3; ys++ {
					sqPow += powerLevel(point{xs, ys}, serial)
				}
			}

			first := x == 1 && y == 1
			if first || sqPow > maxPower {
				maxPower = sqPow
				maxPowerPoint = point{x, y}
			}
			//
		}
	}
	log("max: point = %s; power = %d", maxPowerPoint, maxPower)

	return maxPowerPoint, nil
}

func coordOfLargestPowerLevelFlexSquare(serial int) (string, error) {
	xsize, ysize := 300, 300

	var maxPower int
	var maxPowerSize int
	var maxPowerPoint point
	for x := 1; x <= xsize; x++ {
		for y := 1; y <= ysize; y++ {
			for size := 1; size <= xsize-x+1; size++ {
				log("check: %s x %d", point{x, y}, size)
				//
				var sqPow int
				for xs := x; xs < x+size; xs++ {
					for ys := y; ys < y+size; ys++ {
						sqPow += powerLevel(point{xs, ys}, serial)
					}
				}

				first := x == 1 && y == 1
				if first || sqPow > maxPower {
					maxPower = sqPow
					maxPowerPoint = point{x, y}
					maxPowerSize = size
				}
				//
			}

		}
	}
	log("max: point = %s; power = %d; size = %d", maxPowerPoint, maxPower, maxPowerSize)

	return fmt.Sprintf("%s,%d", maxPowerPoint, maxPowerSize), nil
}
