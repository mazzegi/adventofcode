package day_20

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/intutil"
	"adventofcode_2017/readutil"
	"fmt"
	"sort"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func fatalIfErr(err error) {
	if err == nil {
		return
	}
	fatal("ERROR: %v", err)
}

func Part1() {
	res, err := closestToOrigin(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := leftAfterCollisions(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
type vector struct {
	x, y, z int
}

func makeVector(x, y, z int) vector {
	return vector{x, y, z}
}

func dist(v1, v2 vector) int {
	return intutil.AbsInt(v2.x-v1.x) +
		intutil.AbsInt(v2.y-v1.y) +
		intutil.AbsInt(v2.z-v1.z)
}

func (v vector) length() int {
	return dist(v, makeVector(0, 0, 0))
}

type particle struct {
	id   int
	dead bool
	acc  vector
	vel  vector
	pos  vector
}

func (p *particle) kill() {
	p.dead = true
	p.acc = makeVector(0, 0, 0)
	p.vel = makeVector(0, 0, 0)
}

func (p *particle) step() {
	p.vel.x += p.acc.x
	p.vel.y += p.acc.y
	p.vel.z += p.acc.z

	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
	p.pos.z += p.vel.z
}

//p=<-1476,-1593,124>, v=<96,-58,141>, a=<-3,8,-10>
func mustParseParticle(s string) *particle {
	p := particle{}
	_, err := fmt.Sscanf(s, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
		&p.pos.x, &p.pos.y, &p.pos.z,
		&p.vel.x, &p.vel.y, &p.vel.z,
		&p.acc.x, &p.acc.y, &p.acc.z,
	)

	fatalIfErr(err)
	return &p
}

func mustParseParticles(in string) []*particle {
	var ps []*particle
	for i, line := range readutil.ReadLines(in) {
		p := mustParseParticle(line)
		p.id = i
		ps = append(ps, p)
	}
	if len(ps) == 0 {
		fatal("no data")
	}
	return ps
}

func closestToOrigin(in string) (int, error) {
	ps := mustParseParticles(in)

	// doStep := func() {
	// 	for _, p := range ps {
	// 		p.step()
	// 	}
	// }

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].acc.length() < ps[j].acc.length()
	})

	return ps[0].id, nil
}

func leftAfterCollisions(in string) (int, error) {
	ps := mustParseParticles(in)

	doStep := func() {
		for _, p := range ps {
			if p.dead {
				continue
			}
			p.step()
		}
	}

	resolveCollisions := func() int {
		posMap := map[vector][]*particle{}
		for _, p := range ps {
			if p.dead {
				continue
			}
			posMap[p.pos] = append(posMap[p.pos], p)
		}
		var rem int
		for _, pps := range posMap {
			if len(pps) > 1 {
				for _, pp := range pps {
					pp.kill()
					rem++
				}
			}
		}
		return rem
	}

	i := 0
	for {
		rem := resolveCollisions()
		if rem > 0 {
			log("%d: removed %d", i, rem)
		}
		doStep()
		i++
		if i%1000 == 0 {
			log("step %d", i)
		}
		if i >= 100000 {
			break
		}
	}

	var alive int
	for _, p := range ps {
		if p.dead {
			continue
		}
		alive++
	}
	return alive, nil
}
