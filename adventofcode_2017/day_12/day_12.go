package day_12

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/readutil"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := group0Count(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := totalGroups(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type program struct {
	id    int
	peers []int
	group *int
}

//29 <-> 29, 142, 474, 552, 1089
func parseProgram(s string) (*program, error) {
	s = strings.ReplaceAll(s, " <-> ", ", ")
	ns, err := readutil.ReadInts(s, ",")
	if err != nil {
		return nil, errors.Wrapf(err, "read-ints %q", s)
	}
	if len(ns) < 2 {
		return nil, errors.Errorf("invalid prg %q", s)
	}
	return &program{
		id:    ns[0],
		peers: ns[1:],
	}, nil
}

func parsePrograms(in string) ([]*program, error) {
	var ps []*program
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		p, err := parseProgram(line)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if len(ps) == 0 {
		return nil, errors.Errorf("no data")
	}
	return ps, nil
}

func group0Count(in string) (int, error) {
	ps, err := parsePrograms(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse programs")
	}
	prgMap := map[int]*program{}
	for _, p := range ps {
		prgMap[p.id] = p
	}

	p0 := prgMap[0]
	visited := map[int]bool{}
	var visit func(p *program)
	visit = func(p *program) {
		if visited[p.id] {
			return
		}
		visited[p.id] = true
		for _, ppid := range p.peers {
			pp := prgMap[ppid]
			visit(pp)
		}
	}
	visit(p0)

	log("total visite: %d", len(visited))

	return 0, nil
}

func totalGroups(in string) (int, error) {
	ps, err := parsePrograms(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse programs")
	}
	prgMap := map[int]*program{}
	for _, p := range ps {
		prgMap[p.id] = p
	}

	findFirstWithoutGroup := func() (*program, bool) {
		for _, p := range ps {
			if p.group == nil {
				return p, true
			}
		}
		return nil, false
	}

	var visit func(p *program, group int, visited map[int]bool)
	visit = func(p *program, group int, visited map[int]bool) {
		if visited[p.id] {
			return
		}
		visited[p.id] = true
		p.group = &group
		for _, ppid := range p.peers {
			pp := prgMap[ppid]
			visit(pp, group, visited)
		}
	}

	var totalGroups int
	for {
		pg, ok := findFirstWithoutGroup()
		if !ok {
			break
		}
		pg.group = &pg.id
		visited := map[int]bool{}
		visit(pg, pg.id, visited)
		totalGroups++
	}

	log("total groups: %d", totalGroups)

	return totalGroups, nil
}
