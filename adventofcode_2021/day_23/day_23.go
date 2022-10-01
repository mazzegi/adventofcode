package day_23

import (
	"fmt"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
)

func setupTest() *burrow {
	b := setupBurrow()
	b.init(
		&amphipod{0, atypA, location{idSideRoomA, 0}},
		&amphipod{1, atypA, location{idSideRoomD, 0}},
		&amphipod{0, atypB, location{idSideRoomA, 1}},
		&amphipod{1, atypB, location{idSideRoomC, 1}},
		&amphipod{0, atypC, location{idSideRoomB, 1}},
		&amphipod{1, atypC, location{idSideRoomC, 0}},
		&amphipod{0, atypD, location{idSideRoomB, 0}},
		&amphipod{1, atypD, location{idSideRoomD, 1}},
	)
	return b
}

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
	panic("fatal")
}

func Part1() {
	res, err := leastEnergy(setup())
	//res, err := leastEnergy(setupTest())
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(setup())
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//model
/*
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
*/

type atyp rune

const (
	atypA atyp = 'A'
	atypB atyp = 'B'
	atypC atyp = 'C'
	atypD atyp = 'D'
)

func (at atyp) String() string {
	return string(at)
}

var stepEnergy = map[atyp]int{
	atypA: 1,
	atypB: 10,
	atypC: 100,
	atypD: 1000,
}

type amphipod struct {
	id  int
	typ atyp
	loc location
}

func (a *amphipod) String() string {
	return fmt.Sprintf("%s%d (%s)", string(a.typ), a.id, a.loc.String())
}

func (a *amphipod) clone() *amphipod {
	return &amphipod{a.id, a.typ, a.loc}
}

type hallway struct {
	positions [11]*amphipod
}

type roomID string

const (
	idHallway   roomID = "hallway"
	idSideRoomA roomID = "sideRoomA"
	idSideRoomB roomID = "sideRoomB"
	idSideRoomC roomID = "sideRoomC"
	idSideRoomD roomID = "sideRoomD"
)

type sideRoom struct {
	roomID    roomID
	typ       atyp
	positions [2]*amphipod
	hallwayAt int
}

func (sr *sideRoom) complete() bool {
	for _, a := range sr.positions {
		if a == nil {
			return false
		}
		if a.typ != sr.typ {
			return false
		}
	}
	return true
}

func (sr sideRoom) containsOtherThan(at atyp) bool {
	for _, a := range sr.positions {
		if a != nil && a.typ != at {
			return true
		}
	}
	return false
}

type burrow struct {
	hallway   *hallway
	sideRooms [4]*sideRoom
	amphipods []*amphipod
	energy    int
}

func (b *burrow) complete() bool {
	for _, sr := range b.sideRooms {
		if !sr.complete() {
			return false
		}
	}
	return true
}

func (b *burrow) clone() *burrow {
	cb := setupBurrow()
	var cas []*amphipod
	for _, a := range b.amphipods {
		cas = append(cas, a.clone())
	}
	cb.init(cas...)
	//cb.energy = b.energy
	return cb
}

func (b *burrow) destinationRoomFor(a *amphipod) *sideRoom {
	for _, sr := range b.sideRooms {
		if sr.typ == a.typ {
			return sr
		}
	}
	fatal("no room for typ %q", a.typ)
	return nil
}

type location struct {
	roomID roomID
	pos    int
}

func (l location) String() string {
	return fmt.Sprintf("%s:%d", l.roomID, l.pos)
}

type move struct {
	amphipod *amphipod
	loc      location
	energy   int
}

func (mv move) String() string {
	return fmt.Sprintf("%s => %s [%d]", mv.amphipod.String(), mv.loc.String(), mv.energy)
}

var sideRoomExits = map[int]bool{
	2: true,
	4: true,
	6: true,
	8: true,
}

func (b *burrow) possibleMovesForAmphipodInHallway(a *amphipod) []move {
	var mvs []move

	dstRoom := b.destinationRoomFor(a)
	if dstRoom.containsOtherThan(a.typ) {
		return []move{}
	}

	// is way free to crossing point
	var dir int
	if dstRoom.hallwayAt > a.loc.pos {
		dir = 1
	} else {
		dir = -1
	}
	pos := a.loc.pos
	var steps int
	for {
		if pos == dstRoom.hallwayAt {
			break
		}
		if b.hallway.positions[pos] != nil && *(b.hallway.positions[pos]) != *a {
			return []move{}
		}
		pos += dir
		steps++
	}

	if dstRoom.positions[0] == nil {
		steps += 2
		energy := steps * stepEnergy[a.typ]
		mvs = append(mvs, move{
			a,
			location{dstRoom.roomID, 0},
			energy,
		})
	} else {
		steps += 1
		energy := steps * stepEnergy[a.typ]
		mvs = append(mvs, move{
			a,
			location{dstRoom.roomID, 1},
			energy,
		})
	}

	return mvs
}

func (b *burrow) possibleMovesForAmphipodInSideRoom(a *amphipod, sr *sideRoom) []move {
	if a.loc.pos == 0 && sr.positions[1] != nil {
		// exit is occupied
		return []move{}
	}
	// already in place?
	if sr.typ == a.typ {
		if a.loc.pos == 0 {
			// already where it belongs
			return []move{}
		} else {
			// at pos 1-  move to 0 if possible
			if sr.positions[0] == nil {
				return []move{
					{
						a,
						location{sr.roomID, 0},
						1,
					},
				}
			} else if sr.positions[0].typ == sr.typ {
				// side room complete
				return []move{}
			}
		}
	}

	var exitSteps int
	if a.loc.pos == 0 {
		exitSteps = 2
	} else {
		exitSteps = 1
	}

	// try in any direction
	var mvs []move
	pos := sr.hallwayAt + 1
	var steps int = 1
	for {
		if pos >= len(b.hallway.positions) {
			break
		}
		if b.hallway.positions[pos] != nil && *(b.hallway.positions[pos]) != *a {
			//blocked
			break
		}
		if _, isExit := sideRoomExits[pos]; !isExit {
			energy := (steps + exitSteps) * stepEnergy[a.typ]
			mvs = append(mvs, move{
				a,
				location{idHallway, pos},
				energy,
			})
		}
		pos++
		steps++
	}

	pos = sr.hallwayAt - 1
	steps = 1
	for {
		if pos < 0 {
			break
		}
		if b.hallway.positions[pos] != nil && *(b.hallway.positions[pos]) != *a {
			//blocked
			break
		}
		if _, isExit := sideRoomExits[pos]; !isExit {
			energy := (steps + exitSteps) * stepEnergy[a.typ]
			mvs = append(mvs, move{
				a,
				location{idHallway, pos},
				energy,
			})
		}
		pos--
		steps++
	}
	return mvs
}

func (b *burrow) possibleMovesForAmphipod(a *amphipod) []move {
	if a.loc.roomID == idHallway {
		return b.possibleMovesForAmphipodInHallway(a)
	} else {
		// in sideroom
		var sr *sideRoom
		switch a.loc.roomID {
		case idSideRoomA:
			sr = b.sideRooms[0]
		case idSideRoomB:
			sr = b.sideRooms[1]
		case idSideRoomC:
			sr = b.sideRooms[2]
		case idSideRoomD:
			sr = b.sideRooms[3]
		}
		return b.possibleMovesForAmphipodInSideRoom(a, sr)
	}
}

func (b *burrow) possibleMoves() []move {
	var mvs []move
	for _, a := range b.amphipods {
		amvs := b.possibleMovesForAmphipod(a)
		mvs = append(mvs, amvs...)
	}
	return mvs
}

func (b *burrow) init(as ...*amphipod) {
	b.amphipods = as
	for _, a := range as {
		b.move(move{
			amphipod: a,
			loc:      a.loc,
			energy:   0,
		})
	}
}

func (b *burrow) move(mv move) {
	//remove from previous
	switch mv.amphipod.loc.roomID {
	case idHallway:
		b.hallway.positions[mv.amphipod.loc.pos] = nil
	case idSideRoomA:
		b.sideRooms[0].positions[mv.amphipod.loc.pos] = nil
	case idSideRoomB:
		b.sideRooms[1].positions[mv.amphipod.loc.pos] = nil
	case idSideRoomC:
		b.sideRooms[2].positions[mv.amphipod.loc.pos] = nil
	case idSideRoomD:
		b.sideRooms[3].positions[mv.amphipod.loc.pos] = nil
	}

	//
	mv.amphipod.loc = mv.loc
	switch mv.loc.roomID {
	case idHallway:
		b.hallway.positions[mv.loc.pos] = mv.amphipod
	case idSideRoomA:
		b.sideRooms[0].positions[mv.loc.pos] = mv.amphipod
	case idSideRoomB:
		b.sideRooms[1].positions[mv.loc.pos] = mv.amphipod
	case idSideRoomC:
		b.sideRooms[2].positions[mv.loc.pos] = mv.amphipod
	case idSideRoomD:
		b.sideRooms[3].positions[mv.loc.pos] = mv.amphipod
	}
	b.energy += mv.energy
}

/*
#############
#.....D.....#
###B#.#C#D###
  #A#B#C#A#
  #########
*/

func (b *burrow) dump() string {
	hwAt := func(i int) string {
		v := b.hallway.positions[i]
		if v == nil {
			return "."
		}
		return v.typ.String()
	}

	srAt := func(sr *sideRoom, i int) string {
		v := sr.positions[i]
		if v == nil {
			return "."
		}
		return v.typ.String()
	}

	srA := b.sideRooms[0]
	srB := b.sideRooms[1]
	srC := b.sideRooms[2]
	srD := b.sideRooms[3]

	var sl []string
	sl = append(sl, "#############")

	shw := "#"
	for i := 0; i < len(b.hallway.positions); i++ {
		shw += hwAt(i)
	}
	shw += "#"
	sl = append(sl, shw)

	sr1 := "###"
	sr1 += srAt(srA, 1)
	sr1 += "#"
	sr1 += srAt(srB, 1)
	sr1 += "#"
	sr1 += srAt(srC, 1)
	sr1 += "#"
	sr1 += srAt(srD, 1)
	sr1 += "###"
	sl = append(sl, sr1)

	sr2 := "  #"
	sr2 += srAt(srA, 0)
	sr2 += "#"
	sr2 += srAt(srB, 0)
	sr2 += "#"
	sr2 += srAt(srC, 0)
	sr2 += "#"
	sr2 += srAt(srD, 0)
	sr2 += "#  "
	sl = append(sl, sr2)

	sl = append(sl, "  #########  ")

	return strings.Join(sl, "\n")
}

func setupBurrow() *burrow {
	return &burrow{
		hallway: &hallway{},
		sideRooms: [4]*sideRoom{
			{
				idSideRoomA,
				atypA,
				[2]*amphipod{nil, nil},
				2,
			},
			{
				idSideRoomB,
				atypB,
				[2]*amphipod{nil, nil},
				4,
			},
			{
				idSideRoomC,
				atypC,
				[2]*amphipod{nil, nil},
				6,
			},
			{
				idSideRoomD,
				atypD,
				[2]*amphipod{nil, nil},
				8,
			},
		},
	}
}

func setup() *burrow {
	b := setupBurrow()
	b.init(
		&amphipod{0, atypA, location{idSideRoomC, 0}},
		&amphipod{1, atypA, location{idSideRoomD, 0}},
		&amphipod{0, atypB, location{idSideRoomA, 1}},
		&amphipod{1, atypB, location{idSideRoomD, 1}},
		&amphipod{0, atypC, location{idSideRoomB, 1}},
		&amphipod{1, atypC, location{idSideRoomC, 1}},
		&amphipod{0, atypD, location{idSideRoomA, 0}},
		&amphipod{1, atypD, location{idSideRoomB, 0}},
	)
	return b
}

func (b *burrow) copyMove(mv move) move {
	cmv := mv
	for _, a := range b.amphipods {
		if a.typ == mv.amphipod.typ && a.id == mv.amphipod.id {
			cmv.amphipod = a
			break
		}
	}
	return cmv
}

func (b *burrow) hashAmpiphods() string {
	var s string
	for _, a := range b.amphipods {
		s += a.loc.String() + "-"
	}
	return s
}

//

func leastEnergy(b *burrow) (int, error) {
	log("start with\n%s", b.dump())

	hashMap := map[string]stepResult{}
	e, comp := stepOne(hashMap, 0, b)
	if !comp {
		log("never completed")
	} else {
		log("completed: less-energy: %d", e)
	}
	return e, nil
}

type stepResult struct {
	compl  bool
	energy int
}

func stepOne(hashMap map[string]stepResult, level int, b *burrow) (leastEnergy int, completed bool) {
	//log("step one -----------\n%s", b.dump())
	var le int
	compl := false
	mvs := b.possibleMoves()
	//log("step-one: %d (pm = %d)", level, len(mvs))
	for i, mv := range mvs {
		if level < 2 {
			log("level %d - pm %d (%d)", level, i, len(mvs))
		}

		bc := b.clone()
		cmv := bc.copyMove(mv)
		bc.move(cmv)
		//log("option -----------\n%s", bc.dump())
		var res int
		var complstep bool
		if bc.complete() {
			res = bc.energy
			complstep = true
		} else {
			hash := bc.hashAmpiphods()
			if stepRes, ok := hashMap[hash]; ok {
				res, complstep = stepRes.energy, stepRes.compl
			} else {
				res, complstep = stepOne(hashMap, level+1, bc)
				hashMap[hash] = stepResult{complstep, res}
			}
			res += bc.energy
		}

		if complstep {
			if !compl {
				compl = true
				le = res
			} else if res < le {
				le = res
			}
		}
	}
	return le, compl
}

func part2MainFunc(b *burrow) (int, error) {
	return 0, nil
}
