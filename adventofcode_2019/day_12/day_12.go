package day_12

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc64"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/euler"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/mathutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input, 1000)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFuncSuper(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func part1MainFunc(in []grid.Point3D, steps int) (int, error) {
	var moons []*moon
	for _, pos := range in {
		moons = append(moons, &moon{
			pos: pos,
			vel: grid.P3D(0, 0, 0),
		})
	}

	type pair struct {
		m1 *moon
		m2 *moon
	}
	var pairs []pair
	for i := 0; i < len(moons); i++ {
		m1 := moons[i]
		for j := i + 1; j < len(moons); j++ {
			m2 := moons[j]
			pairs = append(pairs, pair{m1, m2})
		}
	}

	for i := 0; i < steps; i++ {
		for _, p := range pairs {
			switch {
			case p.m1.pos.X < p.m2.pos.X:
				p.m1.vel.X++
				p.m2.vel.X--
			case p.m1.pos.X > p.m2.pos.X:
				p.m1.vel.X--
				p.m2.vel.X++
			}
			switch {
			case p.m1.pos.Y < p.m2.pos.Y:
				p.m1.vel.Y++
				p.m2.vel.Y--
			case p.m1.pos.Y > p.m2.pos.Y:
				p.m1.vel.Y--
				p.m2.vel.Y++
			}
			switch {
			case p.m1.pos.Z < p.m2.pos.Z:
				p.m1.vel.Z++
				p.m2.vel.Z--
			case p.m1.pos.Z > p.m2.pos.Z:
				p.m1.vel.Z--
				p.m2.vel.Z++
			}
		}
		//log("after %d steps", i+1)
		for _, m := range moons {
			m.pos = m.pos.Add(m.vel)
			//dumpMoon(m)
		}
	}

	var energy int
	for _, m := range moons {
		pot := mathutil.Abs(m.pos.X) + mathutil.Abs(m.pos.Y) + mathutil.Abs(m.pos.Z)
		kin := mathutil.Abs(m.vel.X) + mathutil.Abs(m.vel.Y) + mathutil.Abs(m.vel.Z)
		energy += pot * kin
	}

	return energy, nil
}

func part2MainFunc(in []grid.Point3D) (int, error) {
	var moons []*moon
	for _, pos := range in {
		moons = append(moons, &moon{
			pos: pos,
			vel: grid.P3D(0, 0, 0),
		})
	}

	type pair struct {
		m1 *moon
		m2 *moon
	}
	var pairs []pair
	for i := 0; i < len(moons); i++ {
		m1 := moons[i]
		for j := i + 1; j < len(moons); j++ {
			m2 := moons[j]
			pairs = append(pairs, pair{m1, m2})
		}
	}

	tab := crc64.MakeTable(crc64.ISO)
	hash := func() uint64 {
		buf := bytes.Buffer{}
		for _, m := range moons {
			err := binary.Write(&buf, binary.LittleEndian, m.arr())
			if err != nil {
				fatal("write binary: %v", err)
			}
		}
		return crc64.Checksum(buf.Bytes(), tab)
	}

	step := func() {
		for _, p := range pairs {
			switch {
			case p.m1.pos.X < p.m2.pos.X:
				p.m1.vel.X++
				p.m2.vel.X--
			case p.m1.pos.X > p.m2.pos.X:
				p.m1.vel.X--
				p.m2.vel.X++
			}
			switch {
			case p.m1.pos.Y < p.m2.pos.Y:
				p.m1.vel.Y++
				p.m2.vel.Y--
			case p.m1.pos.Y > p.m2.pos.Y:
				p.m1.vel.Y--
				p.m2.vel.Y++
			}
			switch {
			case p.m1.pos.Z < p.m2.pos.Z:
				p.m1.vel.Z++
				p.m2.vel.Z--
			case p.m1.pos.Z > p.m2.pos.Z:
				p.m1.vel.Z--
				p.m2.vel.Z++
			}
		}
		for _, m := range moons {
			m.pos = m.pos.Add(m.vel)
		}
	}

	var period uint64
	stepIdx := uint64(0)
	states := map[uint64]uint64{}
	for {
		h := hash()
		if when, ok := states[h]; ok {
			period = stepIdx - when
			break
		}
		states[h] = uint64(stepIdx)
		step()
		stepIdx++
	}
	return int(period), nil
}

func part2MainFuncExt(in []grid.Point3D) (int, error) {
	var moons []*moon
	for _, pos := range in {
		moons = append(moons, &moon{
			pos: pos,
			vel: grid.P3D(0, 0, 0),
		})
	}

	type pair struct {
		m1 *moon
		m2 *moon
	}
	var pairs []pair
	for i := 0; i < len(moons); i++ {
		m1 := moons[i]
		for j := i + 1; j < len(moons); j++ {
			m2 := moons[j]
			pairs = append(pairs, pair{m1, m2})
		}
	}

	tab := crc64.MakeTable(crc64.ISO)
	hash := func() uint64 {
		buf := bytes.Buffer{}
		for _, m := range moons {
			err := binary.Write(&buf, binary.LittleEndian, m.arr())
			if err != nil {
				fatal("write binary: %v", err)
			}
		}
		return crc64.Checksum(buf.Bytes(), tab)
	}
	_ = hash

	hashMoon := func(m *moon) uint64 {
		buf := bytes.Buffer{}
		err := binary.Write(&buf, binary.LittleEndian, m.arr())
		if err != nil {
			fatal("write binary: %v", err)
		}
		return crc64.Checksum(buf.Bytes(), tab)
	}

	step := func() {
		for _, p := range pairs {
			switch {
			case p.m1.pos.X < p.m2.pos.X:
				p.m1.vel.X++
				p.m2.vel.X--
			case p.m1.pos.X > p.m2.pos.X:
				p.m1.vel.X--
				p.m2.vel.X++
			}
			switch {
			case p.m1.pos.Y < p.m2.pos.Y:
				p.m1.vel.Y++
				p.m2.vel.Y--
			case p.m1.pos.Y > p.m2.pos.Y:
				p.m1.vel.Y--
				p.m2.vel.Y++
			}
			switch {
			case p.m1.pos.Z < p.m2.pos.Z:
				p.m1.vel.Z++
				p.m2.vel.Z--
			case p.m1.pos.Z > p.m2.pos.Z:
				p.m1.vel.Z--
				p.m2.vel.Z++
			}
		}
		for _, m := range moons {
			m.pos = m.pos.Add(m.vel)
		}
	}

	var period uint64
	stepIdx := uint64(0)
	//states := map[uint64]uint64{}

	periods := make([]int, len(moons))
	periodStarts := make([]int, len(moons))
	moonStates := make([]map[uint64]uint64, len(moons))
	for i, m := range moons {
		moonStates[i] = map[uint64]uint64{hashMoon(m): 0}
		periods[i] = -1
		periodStarts[i] = -1
	}

	for {
		step()
		stepIdx++
		for i, m := range moons {
			if periods[i] > -1 {
				continue
			}
			h := hashMoon(m)
			if when, ok := moonStates[i][h]; ok {
				log("moon-period [%d]: %d - %d = %d", i, stepIdx, when, stepIdx-when)
				periods[i] = int(stepIdx) - int(when)
				periodStarts[i] = int(when)
			}
			moonStates[i][h] = stepIdx
		}
		if stepIdx > 3000 {
			break
		}

		// h := hash()
		// if when, ok := states[h]; ok {
		// 	period = stepIdx - when
		// 	break
		// }
		// states[h] = uint64(stepIdx)
		// step()

	}
	return int(period), nil
}

func part2MainFuncSuper(in []grid.Point3D) (int, error) {
	var moons []*moon
	for _, pos := range in {
		moons = append(moons, &moon{
			pos: pos,
			vel: grid.P3D(0, 0, 0),
		})
	}

	type pair struct {
		m1 *moon
		m2 *moon
	}
	var pairs []pair
	for i := 0; i < len(moons); i++ {
		m1 := moons[i]
		for j := i + 1; j < len(moons); j++ {
			m2 := moons[j]
			pairs = append(pairs, pair{m1, m2})
		}
	}

	tab := crc64.MakeTable(crc64.ISO)
	hash := func() (x, y, z uint64) {
		bufX := bytes.Buffer{}
		bufY := bytes.Buffer{}
		bufZ := bytes.Buffer{}
		for _, m := range moons {
			binary.Write(&bufX, binary.LittleEndian, []int64{int64(m.pos.X), int64(m.vel.X)})
			binary.Write(&bufY, binary.LittleEndian, []int64{int64(m.pos.Y), int64(m.vel.Y)})
			binary.Write(&bufZ, binary.LittleEndian, []int64{int64(m.pos.Z), int64(m.vel.Z)})
		}
		return crc64.Checksum(bufX.Bytes(), tab),
			crc64.Checksum(bufY.Bytes(), tab),
			crc64.Checksum(bufZ.Bytes(), tab)
	}

	step := func() {
		for _, p := range pairs {
			switch {
			case p.m1.pos.X < p.m2.pos.X:
				p.m1.vel.X++
				p.m2.vel.X--
			case p.m1.pos.X > p.m2.pos.X:
				p.m1.vel.X--
				p.m2.vel.X++
			}
			switch {
			case p.m1.pos.Y < p.m2.pos.Y:
				p.m1.vel.Y++
				p.m2.vel.Y--
			case p.m1.pos.Y > p.m2.pos.Y:
				p.m1.vel.Y--
				p.m2.vel.Y++
			}
			switch {
			case p.m1.pos.Z < p.m2.pos.Z:
				p.m1.vel.Z++
				p.m2.vel.Z--
			case p.m1.pos.Z > p.m2.pos.Z:
				p.m1.vel.Z--
				p.m2.vel.Z++
			}
		}
		for _, m := range moons {
			m.pos = m.pos.Add(m.vel)
		}
	}

	stepIdx := uint64(0)
	statesX := map[uint64]uint64{}
	statesY := map[uint64]uint64{}
	statesZ := map[uint64]uint64{}

	type wave struct {
		period uint64
		start  uint64
	}
	var waveX, waveY, waveZ wave

	for {
		hx, hy, hz := hash()

		if waveX.period == 0 {
			if when, ok := statesX[hx]; ok {
				waveX.period = stepIdx - when
				waveX.start = when
				log("wave-X: start %d with period %d", waveX.start, waveX.period)
			}
		}
		if waveY.period == 0 {
			if when, ok := statesY[hy]; ok {
				waveY.period = stepIdx - when
				waveY.start = when
				log("wave-Y: start %d with period %d", waveY.start, waveY.period)
			}
		}
		if waveZ.period == 0 {
			if when, ok := statesZ[hz]; ok {
				waveZ.period = stepIdx - when
				waveZ.start = when
				log("wave-Z: start %d with period %d", waveZ.start, waveZ.period)
			}
		}
		if waveX.period > 0 &&
			waveY.period > 0 &&
			waveZ.period > 0 {
			break
		}

		statesX[hx] = uint64(stepIdx)
		statesY[hy] = uint64(stepIdx)
		statesZ[hz] = uint64(stepIdx)
		step()
		stepIdx++
	}
	return euler.SmallestMultipleOf(int(waveX.period), int(waveY.period), int(waveZ.period)), nil
}

type moon struct {
	pos grid.Point3D
	vel grid.Point3D
}

func (m moon) arr() []int64 {
	return []int64{
		int64(m.pos.X), int64(m.pos.Y), int64(m.pos.Z),
		int64(m.vel.X), int64(m.vel.Y), int64(m.vel.Z),
	}
}
