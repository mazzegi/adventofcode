package day_07

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/readutil"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := nameOfBottomProgram(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %s\n", res)
}

func Part2() {
	res, err := correctionWeight(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

type program struct {
	name       string
	weight     int
	childNames []string
	childs     []*program
	parent     *program
}

func parseProgram(s string) (*program, error) {
	sl := strings.Split(s, "->")
	if len(sl) == 0 || len(sl) > 2 {
		return nil, fmt.Errorf("invalid program def %q", s)
	}

	prgS := strings.Trim(sl[0], " \r\n\t")
	var name string
	var weight int
	_, err := fmt.Sscanf(prgS, "%s (%d)", &name, &weight)
	if err != nil {
		return nil, errors.Wrapf(err, "scan program %q", prgS)
	}

	prg := &program{
		name:   name,
		weight: weight,
	}

	if len(sl) == 2 {
		subSl := strings.Split(sl[1], ", ")
		for _, subS := range subSl {
			subS = strings.Trim(subS, " \r\n\t")
			if subS == "" {
				continue
			}
			prg.childNames = append(prg.childNames, subS)
		}
	}
	return prg, nil
}

func parsePrograms(in string) ([]*program, error) {
	var ps []*program
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		p, err := parseProgram(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-program %q", line)
		}
		ps = append(ps, p)
	}
	return ps, nil
}

type tree struct {
	prgs map[string]*program
}

func buildTree(ps []*program) (*tree, error) {
	t := &tree{
		prgs: map[string]*program{},
	}

	for _, p := range ps {
		t.prgs[p.name] = p
	}

	//link them
	for _, p := range ps {
		for _, childName := range p.childNames {
			child, ok := t.prgs[childName]
			if !ok {
				return nil, errors.Errorf("%q: child %q not found in tree", p.name, childName)
			}
			if child.parent != nil {
				return nil, errors.Errorf("%q: child %q has already a parent (%q)", p.name, child.name, child.parent.name)
			}
			child.parent = p
			p.childs = append(p.childs, child)
		}
	}

	return t, nil
}

func (t *tree) root() *program {
	for _, p := range t.prgs {
		if p.parent == nil {
			return p
		}
	}
	return nil
}

// func dist(ns []int) (value int, count int) {
// 	return
// }

func (p *program) aggWeight() (weight int, err error, corr int) {
	if len(p.childs) == 0 {
		return p.weight, nil, 0
	}
	// ref, err := p.childs[0].aggWeight()
	// if err != nil {
	// 	return 0, errors.Wrapf(err, "agg-weight of %q", p.childs[0].name)
	// }

	weights := make([]int, len(p.childs))
	weightDist := map[int][]*program{}
	for i, child := range p.childs {
		cw, err, corr := child.aggWeight()
		if err != nil {
			return 0, errors.Wrapf(err, "agg-weight of %q", child.name), corr
		}
		weights[i] = cw
		weightDist[cw] = append(weightDist[cw], child)

		// if cw != ref {
		// 	return 0, errors.Errorf("child %q doesn't match ref %d", child.name, ref)
		// }
	}
	var majWeight int
	var majWeightCount int
	for w, cs := range weightDist {
		if len(cs) > majWeightCount {
			majWeightCount = len(cs)
			majWeight = w
		}
	}

	if len(weightDist) == 1 {
		return p.weight + len(p.childs)*majWeight, nil, 0

	} else if len(weightDist) == 2 {
		for w, prgs := range weightDist {
			if len(prgs) == 1 {
				child := prgs[0]
				diff := majWeight - w
				corrWeight := child.weight + diff

				return 0, errors.Errorf("child %q (w = %d) doesn't match majority %d. It has to be corrected to %d",
					child.name, w, majWeight, corrWeight), corrWeight
			}
		}
		return 0, errors.Errorf("no count 1 weight in %v", weightDist), 0
	} else {
		return 0, errors.Errorf("cannot handle weight-dist %v", weightDist), 0
	}

	//return p.weight + len(p.childs)*weights[0], nil
}

func nameOfBottomProgram(in string) (string, error) {
	ps, err := parsePrograms(in)
	if err != nil {
		return "", errors.Wrap(err, "parse-programs")
	}
	if len(ps) == 0 {
		return "", errors.Errorf("no entries")
	}
	t, err := buildTree(ps)
	if err != nil {
		return "", errors.Wrap(err, "build-tree")
	}

	rt := t.root()
	if rt == nil {
		return "", errors.Errorf("no root program found")
	}

	return rt.name, nil
}

func correctionWeight(in string) (int, error) {
	ps, err := parsePrograms(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-programs")
	}
	if len(ps) == 0 {
		return 0, errors.Errorf("no entries")
	}
	t, err := buildTree(ps)
	if err != nil {
		return 0, errors.Wrap(err, "build-tree")
	}
	rt := t.root()
	if rt == nil {
		return 0, errors.Errorf("no root program found")
	}

	_, err, corr := rt.aggWeight()
	if err != nil {
		fmt.Printf("agg-error: %v\n", err)
	} else {
		fmt.Printf("all are balanced")
	}
	return corr, nil
}
