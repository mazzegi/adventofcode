package day_11

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
	res, err := part1MainFunc(input, 0)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part1MainFunc(input, 1)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func part1MainFunc(prg []int, startColor int) (int, error) {

	in := intcode.NewIntChannelReader(1)
	out := intcode.NewIntChannelWriter(1)

	go func() {
		comp := intcode.NewComputer(prg, in, out)
		err := comp.Exec()
		log("computer finished with error: %v", err)
		out.Close()
	}()

	panels := map[grid.Point]int{}
	curr := grid.Pt(0, 0)
	currDir := grid.Pt(0, -1)

	rotateLeft := func(dir grid.Point) grid.Point {
		switch dir {
		case grid.Pt(0, -1):
			return grid.Pt(-1, 0)
		case grid.Pt(-1, 0):
			return grid.Pt(0, 1)
		case grid.Pt(0, 1):
			return grid.Pt(1, 0)
		case grid.Pt(1, 0):
			return grid.Pt(0, -1)
		default:
			fatal("cannot rotate %s", dir)
			return dir
		}
	}

	rotateRight := func(dir grid.Point) grid.Point {
		switch dir {
		case grid.Pt(0, -1):
			return grid.Pt(1, 0)
		case grid.Pt(1, 0):
			return grid.Pt(0, 1)
		case grid.Pt(0, 1):
			return grid.Pt(-1, 0)
		case grid.Pt(-1, 0):
			return grid.Pt(0, -1)
		default:
			fatal("cannot rotate %s", dir)
			return dir
		}
	}

	first := true
	for {
		var color int = 0
		if first {
			color = startColor
			first = false
		} else if pc, ok := panels[curr]; ok {
			color = pc
		}

		in.Provide(color)
		paint, ok := out.Get()
		if !ok {
			break
		}
		dir, ok := out.Get()
		if !ok {
			break
		}
		panels[curr] = paint
		switch dir {
		case 0: //left by 90 deg
			currDir = rotateLeft(currDir)
		case 1: //right by 90 deg
			currDir = rotateRight(currDir)
		default:
			fatal("invalid direction %d", dir)
		}
		curr = curr.Add(currDir)
	}
	paintedCount := len(panels)
	paintPanels(panels)
	return paintedCount, nil
}

func paintPanels(panels map[grid.Point]int) {
	if len(panels) == 0 {
		log("no panels")
		return
	}
	var topLeft grid.Point
	var bottomRight grid.Point
	first := true
	for pt := range panels {
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
			cr := 0
			if pcr, ok := panels[grid.Pt(x, y)]; ok {
				cr = pcr
			}
			switch cr {
			case 0:
				row += " "
			default:
				row += "#"
			}
		}
		fmt.Println(row)
	}
}

func part2MainFunc(prg []int) (int, error) {
	return 0, nil
}
