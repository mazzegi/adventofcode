package day_13

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2019/intcode"
	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

const (
	empty       = 0
	wall        = 1
	block       = 2
	horz_paddle = 3
	ball        = 4
)

func part1MainFunc(prg []int) (int, error) {
	out := intcode.NewIntChannelWriter(0)
	comp := intcode.NewComputer(prg, intcode.NewIntSliceReader([]int{}), out)

	go func() {
		comp.Exec()
		out.Close()
	}()

	tiles := map[grid.Point]int{}
	for {
		x, ok := out.Get()
		if !ok {
			break
		}
		y, ok := out.Get()
		if !ok {
			break
		}
		t, ok := out.Get()
		if !ok {
			break
		}
		tiles[grid.Pt(x, y)] = t
	}

	//count block tiles
	var blockCount int
	for _, t := range tiles {
		if t == block {
			blockCount++
		}
	}

	return blockCount, nil
}

func part2MainFunc(prg []int) (int, error) {
	prg[0] = 2
	//in := intcode.NewIntChannelReader(0)
	in := intcode.NewSignallingIntChannelReader(0)
	out := intcode.NewIntChannelWriter(0)
	comp := intcode.NewComputer(prg, in, out)

	go func() {
		comp.Exec()
		out.Close()
	}()

	var score int
	tiles := map[grid.Point]int{}
	var ballPos grid.Point
	var paddlePos grid.Point
outer:
	for {
		select {
		case x, ok := <-out.C:
			if !ok {
				break outer
			}
			y, ok := out.Get()
			if !ok {
				break outer
			}
			t, ok := out.Get()
			if !ok {
				break outer
			}
			if x == -1 && y == 0 {
				score = t
				log("score: %d", score)
				printTiles(tiles)
			} else {
				pt := grid.Pt(x, y)
				if t == ball {
					ballPos = pt
				} else if t == horz_paddle {
					paddlePos = pt
				}
				tiles[pt] = t
			}
		case _, ok := <-in.WantC:
			if !ok {
				break outer
			}
			//in.Provide(0)
			// follow the ball
			if paddlePos.X < ballPos.X {
				in.Provide(1)
			} else if paddlePos.X > ballPos.X {
				in.Provide(-1)
			} else {
				in.Provide(0)
			}
		}
	}

	//count block tiles
	var blockCount int
	for _, t := range tiles {
		if t == block {
			blockCount++
		}
	}

	return blockCount, nil
}

func printTiles(tiles map[grid.Point]int) {
	if len(tiles) == 0 {
		log("no tiles")
		return
	}
	var topLeft grid.Point
	var bottomRight grid.Point
	first := true
	for pt := range tiles {
		if first {
			topLeft = pt
			bottomRight = pt
			first = false
		}
		if pt.X < topLeft.X {
			topLeft.X = pt.X
		}
		if pt.X > bottomRight.X {
			bottomRight.X = pt.X
		}
		if pt.Y < topLeft.Y {
			topLeft.Y = pt.Y
		}
		if pt.Y > bottomRight.Y {
			bottomRight.Y = pt.Y
		}
	}
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		var row string
		for x := topLeft.X; x <= bottomRight.X; x++ {
			t, ok := tiles[grid.Pt(x, y)]
			if !ok {
				row += " "
				continue
			}

			switch t {
			case empty:
				row += " "
			case wall:
				row += "#"
			case block:
				row += "+"
			case horz_paddle:
				row += "_"
			case ball:
				row += "o"
			default:
				row += " "
			}
		}
		fmt.Println(row)
	}
}
