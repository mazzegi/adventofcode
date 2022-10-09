package day_15

import (
	"fmt"
	"math"

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
	empty = 0
	wall  = 1
	oxsys = 2
)

const (
	north = 1
	south = 2
	west  = 3
	east  = 4
)

func directions() []int {
	return []int{north, east, south, west}
}

func oppositeDir(dir int) int {
	switch dir {
	case north:
		return south
	case south:
		return north
	case east:
		return west
	case west:
		return east
	default:
		fatal("invalid direction %d", dir)
		return 0
	}
}

var (
	dirNorth = grid.Pt(0, -1)
	dirSouth = grid.Pt(0, 1)
	dirWest  = grid.Pt(-1, 0)
	dirEast  = grid.Pt(1, 0)
)

func move(pos grid.Point, dir int) grid.Point {
	switch dir {
	case north:
		return pos.Add(dirNorth)
	case east:
		return pos.Add(dirEast)
	case south:
		return pos.Add(dirSouth)
	case west:
		return pos.Add(dirWest)
	default:
		fatal("invalid dir %d", dir)
		return grid.Point{}
	}
}

func part1MainFunc(prg []int) (int, error) {

	in := intcode.NewIntChannelReader(0)
	out := intcode.NewIntChannelWriter(0)
	comp := intcode.NewComputer(prg, in, out)
	go func() {
		comp.Exec()
		out.Close()
	}()

	droidPos := grid.Pt(0, 0)
	smap := map[grid.Point]int{}
	smap[droidPos] = empty

	var oxyAt grid.Point
	var exploreUndiscovered func()
	exploreUndiscovered = func() {
		for _, dir := range directions() {
			dir := dir
			probe := move(droidPos, dir)
			if _, ok := smap[probe]; ok {
				continue
			}

			// undiscovered
			in.Provide(dir)
			state, ok := out.Get()
			if !ok {
				fatal("unexpected close of output channel")
			}
			switch state {
			case 0: //wall
				smap[probe] = wall
				//printMap(smap, droidPos)
				continue
			case 1, 2: //moved
				if state == 1 {
					smap[probe] = empty
				} else {
					smap[probe] = oxsys
					oxyAt = probe
					log("found oxygen system at %s", probe)
				}
				droidPos = probe
				//printMap(smap, droidPos)
				exploreUndiscovered()
				//move one back and continue
				droidPos = move(droidPos, oppositeDir(dir))
				in.Provide(oppositeDir(dir))
				bstate, bok := out.Get()
				if !bok {
					fatal("unexpected close of output channel")
				}
				if bstate == 0 {
					fatal("step backwards should be possible")
				}
			default:
				fatal("unexpected state %d", state)
			}
		}
		//
	}

	exploreUndiscovered()
	printMap(smap, droidPos)

	dist := minDistDijkstra(smap, grid.Pt(0, 0), oxyAt)
	oxyTime := fillWithOxygen(smap, oxyAt)
	log("oxy-time: %d", oxyTime)

	return dist, nil
}

func fillWithOxygen(smap map[grid.Point]int, oxySource grid.Point) int {
	locs := map[grid.Point]bool{}
	for pos, val := range smap {
		if val == 1 { //wall
			continue
		}
		if pos == oxySource {
			locs[pos] = true // has oxy
		} else {
			locs[pos] = false // no oxy, but space
		}
	}

	emptyNeighbours := func(pos grid.Point) []grid.Point {
		var ens []grid.Point
		for _, dir := range directions() {
			cand := move(pos, dir)
			if containsOxy, ok := locs[cand]; ok && !containsOxy {
				ens = append(ens, cand)
			}
		}
		return ens
	}

	steps := 0
	for {
		var nextOxyCells []grid.Point
		for pos, containsOxy := range locs {
			if !containsOxy {
				continue
			}
			ens := emptyNeighbours(pos)
			nextOxyCells = append(nextOxyCells, ens...)
		}
		if len(nextOxyCells) == 0 {
			break
		}
		for _, noc := range nextOxyCells {
			locs[noc] = true
		}
		steps++
	}

	return steps
}

func part2MainFunc(prg []int) (int, error) {
	return 0, nil
}

func printMap(smap map[grid.Point]int, droid grid.Point) {
	fmt.Println("\n-----------------------------")
	if len(smap) == 0 {
		log("no tiles")
		return
	}
	var topLeft grid.Point
	var bottomRight grid.Point
	first := true
	for pt := range smap {
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
			if grid.Pt(x, y) == droid {
				row += "D"
				continue
			}
			if grid.Pt(x, y) == grid.Pt(0, 0) {
				row += "S"
				continue
			}

			t, ok := smap[grid.Pt(x, y)]
			if !ok {
				row += " "
				continue
			}

			switch t {
			case empty:
				row += "."
			case wall:
				row += "#"
			case oxsys:
				row += "X"
			default:
				row += " "
			}
		}
		fmt.Println(row)
	}
}

func minDistDijkstra(smap map[grid.Point]int, start, target grid.Point) int {
	type node struct {
		pos  grid.Point
		dist int
		prev *node
	}
	var startNode *node
	var targetNode *node
	nodes := map[grid.Point]*node{}
	for pos, val := range smap {
		if val == 1 { //wall
			continue
		}
		n := &node{
			pos:  pos,
			prev: nil,
		}
		if pos == start {
			n.dist = 0
			startNode = n
		} else {
			n.dist = math.MaxInt
		}

		if pos == target {
			targetNode = n
		}

		nodes[pos] = n
	}

	minDistNodeWith := func() *node {
		var minDist int
		var minDistNode *node = nil
		for _, n := range nodes {
			if minDistNode == nil || n.dist < minDist {
				minDist = n.dist
				minDistNode = n
			}
		}
		return minDistNode
	}

	neighbours := func(n *node) []*node {
		var nns []*node
		for _, dir := range directions() {
			if nn, ok := nodes[move(n.pos, dir)]; ok {
				nns = append(nns, nn)
			}
		}
		return nns
	}

	for {
		if len(nodes) == 0 {
			break
		}
		n := minDistNodeWith()
		delete(nodes, n.pos)
		for _, nn := range neighbours(n) {
			candDist := n.dist + 1
			if candDist < nn.dist {
				nn.dist = candDist
				nn.prev = n
			}
		}
	}

	//
	var dist int
	tn := targetNode
	for {
		if tn.prev == nil {
			fatal("no prev for node %s", tn.pos)
		}
		tn = tn.prev
		dist++
		if tn.pos == startNode.pos {
			break
		}
	}

	return dist
}
