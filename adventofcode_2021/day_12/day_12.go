package day_12

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/stringutil"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := pathCount(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := pathCountV2(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

type connection struct {
	fromCave string
	toCave   string
}

func parseConnection(s string) (*connection, error) {
	sl := strings.Split(s, "-")
	if len(sl) != 2 {
		return nil, errors.Errorf("invalid cave-id in %q", s)
	}
	if sl[0] == "" || sl[1] == "" {
		return nil, errors.Errorf("invalid cave-id in %q", s)
	}
	return &connection{
		fromCave: sl[0],
		toCave:   sl[1],
	}, nil
}

func parseConnections(in string) ([]*connection, error) {
	var cs []*connection
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		con, err := parseConnection(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse connection %q", line)
		}
		cs = append(cs, con)
	}
	return cs, nil
}

func isSmall(id string) bool {
	return unicode.IsLower(rune(id[0]))
}

func isBig(id string) bool {
	return unicode.IsUpper(rune(id[0]))
}

func isStart(id string) bool {
	return id == "start"
}

func isEnd(id string) bool {
	return id == "end"
}

type cave struct {
	id          string
	connections *stringutil.OrderedSet
}

func (c cave) isStart() bool {
	return isStart(c.id)
}

func (c cave) isEnd() bool {
	return isEnd(c.id)
}

func (c cave) isSmall() bool {
	return isSmall(c.id)
}

func (c cave) isBig() bool {
	return isBig(c.id)
}

type graph struct {
	start *cave
	end   *cave
	caves map[string]*cave
}

func (g *graph) ensureCave(id string) *cave {
	c, ok := g.caves[id]
	if !ok {
		c = &cave{
			id:          id,
			connections: stringutil.NewOrderedSet(),
		}
		g.caves[id] = c
	}
	return c
}

func buildGraph(connections []*connection) (*graph, error) {
	g := &graph{
		caves: map[string]*cave{},
	}
	for _, con := range connections {
		c1 := g.ensureCave(con.fromCave)
		c2 := g.ensureCave(con.toCave)

		c1.connections.Insert(c2.id)
		c2.connections.Insert(c1.id)

		if c1.isStart() {
			g.start = c1
		}
		if c1.isEnd() {
			g.end = c1
		}
		if c2.isStart() {
			g.start = c2
		}
		if c2.isEnd() {
			g.end = c2
		}
	}
	if g.start == nil {
		return nil, errors.Errorf("graph has no start")
	}
	if g.end == nil {
		return nil, errors.Errorf("graph has no end")
	}

	return g, nil
}

func pathCount(in string) (int, error) {
	cons, err := parseConnections(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-connections")
	}
	if len(cons) == 0 {
		return 0, errors.Errorf("no entries")
	}
	g, err := buildGraph(cons)
	if err != nil {
		return 0, errors.Wrap(err, "build-graph")
	}

	start := g.start
	var total int
	for _, con := range start.connections.Values() {
		path := stringutil.NewOrderedSet(start.id)
		tPaths := traverseToEnd(g, con, path)
		total += len(tPaths)
	}

	return total, nil
}

func traverseToEnd(g *graph, cid string, path *stringutil.OrderedSet) []*stringutil.OrderedSet {
	var subPaths []*stringutil.OrderedSet

	if isEnd(cid) {
		subPaths = append(subPaths, stringutil.NewOrderedSet(cid))
		return subPaths
	}

	c := g.caves[cid]
	for _, con := range c.connections.Values() {
		subPath := stringutil.NewOrderedSet(path.Values()...)
		subPath.Insert(cid)

		use := isBig(con) || !subPath.Contains(con)
		if !use {
			continue
		}
		tPaths := traverseToEnd(g, con, subPath)
		for _, tp := range tPaths {
			subPath.Insert(tp.Values()...)
			subPaths = append(subPaths, subPath)
		}
	}

	return subPaths
}

//

func pathCountV2(in string) (int, error) {
	cons, err := parseConnections(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-connections")
	}
	if len(cons) == 0 {
		return 0, errors.Errorf("no entries")
	}
	g, err := buildGraph(cons)
	if err != nil {
		return 0, errors.Wrap(err, "build-graph")
	}

	start := g.start
	var total int
	for _, con := range start.connections.Values() {
		path := NewPath(start.id)
		tPaths := traverseToEndV2(g, con, path)
		total += len(tPaths)
		for _, tp := range tPaths {
			fmt.Printf("%s\n", tp.Format())
		}
	}

	return total, nil
}

func traverseToEndV2(g *graph, cid string, path *Path) []*Path {
	var subPaths []*Path

	if isEnd(cid) {
		subPath := path.Clone()
		subPath.Append(cid)
		checkDbg(subPath)
		subPaths = append(subPaths, subPath)
		return subPaths
	}

	c := g.caves[cid]
	for _, con := range c.connections.Values() {
		subPath := path.Clone()
		subPath.Append(cid)
		checkDbg(subPath)

		var use bool
		if isBig(con) {
			use = true
		} else {
			vcnt := subPath.Count(con)
			stc := subPath.smallTwiceCount()
			switch {
			case isStart(con) || isEnd(con):
				if vcnt > 0 {
					use = false
				} else {
					use = true
				}
			default:
				if vcnt == 0 {
					use = true
				} else {
					if stc == 0 {
						use = true
					} else {
						use = false
					}
				}
			}
		}
		if !use {
			continue
		}

		tPaths := traverseToEndV2(g, con, subPath)
		for _, tp := range tPaths {
			checkDbg(tp)
			subPaths = append(subPaths, tp)
		}
	}

	return subPaths
}
