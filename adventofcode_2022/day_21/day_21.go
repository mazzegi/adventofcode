package day_21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
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

//

type Evaler interface {
	Eval(Tree) int
	CanEval() bool
}

type ConstEvaler struct {
	Value int
}

func (e ConstEvaler) Eval(Tree) int {
	return e.Value
}

func (e ConstEvaler) CanEval() bool {
	return true
}

type Op string

const (
	Plus  Op = "+"
	Minus Op = "-"
	Mult  Op = "*"
	Div   Op = "/"
	Eq    Op = "="
)

func (o Op) IsValid() bool {
	switch o {
	case Plus, Minus, Mult, Div, Eq:
		return true
	default:
		return false
	}
}

type ExprEvaler struct {
	XID, YID string
	x, y     *int
	Op       Op
}

func (e ExprEvaler) CanEval() bool {
	return true
}

func (e ExprEvaler) Eval(tree Tree) int {
	if e.x == nil {
		v := tree.Eval(e.XID)
		e.x = &v
	}
	if e.y == nil {
		v := tree.Eval(e.YID)
		e.y = &v
	}

	switch e.Op {
	case Plus:
		return *e.x + *e.y
	case Minus:
		return *e.x - *e.y
	case Mult:
		return *e.x * *e.y
	case Div:
		return *e.x / *e.y
	case Eq:
		if *e.x == *e.y {
			return 1
		}
		return 0
	default:
		panic("invalid op " + e.Op)
	}
}

type Tree map[string]Evaler

func (tree Tree) Eval(id string) int {
	e, ok := tree[id]
	if !ok {
		fatal("no such id %q", id)
	}
	return e.Eval(tree)
}

func (tree Tree) CanEval(id string) bool {
	e, ok := tree[id]
	if !ok {
		fatal("no such id %q", id)
	}
	return e.CanEval()
}

func mustParseTree(in string) Tree {
	tree := Tree{}
	for _, line := range readutil.ReadLines(in) {
		id, exp, ok := strings.Cut(line, ":")
		if !ok {
			fatal("failed to parse %q", line)
		}
		id = strings.TrimSpace(id)
		exp = strings.TrimSpace(exp)
		if id == "" || exp == "" {
			fatal("failed to parse %q", line)
		}

		if n, err := strconv.ParseInt(exp, 10, 64); err == nil {
			tree[id] = ConstEvaler{int(n)}
			continue
		}

		sl := strings.Split(exp, " ")
		if len(sl) != 3 {
			fatal("failed to parse %q", line)
		}

		left := strings.TrimSpace(sl[0])
		op := Op(strings.TrimSpace(sl[1]))
		right := strings.TrimSpace(sl[2])
		if left == "" || right == "" {
			fatal("failed to parse %q", line)
		}
		if !op.IsValid() {
			fatal("failed to parse %q", line)
		}

		tree[id] = ExprEvaler{
			XID: left,
			YID: right,
			Op:  op,
		}
	}
	return tree
}

func part1MainFunc(in string) (int, error) {
	tree := mustParseTree(in)
	res := tree.Eval("root")
	return res, nil
}

type HumanEvaler struct {
}

func (e HumanEvaler) Eval(Tree) int {
	return 0
}

func (e HumanEvaler) CanEval() bool {
	return false
}

func part2MainFunc(in string) (int, error) {
	tree := mustParseTree(in)
	root := tree["root"]
	expRoot := root.(ExprEvaler)
	expRoot.Op = Eq
	tree["root"] = expRoot

	humn := HumanEvaler{}
	tree["humn"] = humn

	res := 0
	return res, nil
}

// func part2MainFunc(in string) (int, error) {
// 	tree := mustParseTree(in)
// 	root := tree["root"]
// 	expRoot := root.(ExprEvaler)
// 	expRoot.Op = Eq
// 	tree["root"] = expRoot
// 	test := 0
// 	for {
// 		tree["humn"] = ConstEvaler{test}
// 		if tree.Eval("root") == 1 {
// 			break
// 		}

// 		test++
// 		if test > 1000000 {
// 			log("no result")
// 			break
// 		}
// 	}

// 	return test, nil
// }
