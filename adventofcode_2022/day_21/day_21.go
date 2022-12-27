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
	Eval(Tree) (int, bool)
	MustEvalTo(Tree, int)
	CanEval() bool
}

type ConstEvaler struct {
	Value int
}

func (e ConstEvaler) Eval(Tree) (int, bool) {
	return e.Value, true
}

func (e ConstEvaler) CanEval() bool {
	return true
}

func (e ConstEvaler) MustEvalTo(tree Tree, n int) {
	fatal("const evaler - cannot must eval to")
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

func (e *ExprEvaler) CanEval() bool {
	return true
}

func (e *ExprEvaler) MustEvalTo(tree Tree, n int) {
	if e.x != nil && e.y != nil {
		fatal("expr-evaler: all values are set - cannot must eval to")
	}
	if e.x == nil && e.y == nil {
		fatal("expr-evaler: all values are unset - cannot must eval to")
	}
	if e.y == nil {
		var mustY int
		switch e.Op {
		case Plus:
			mustY = n - *e.x
		case Minus:
			mustY = *e.x - n
		case Mult:
			mustY = n / (*e.x)
		case Div:
			mustY = (*e.x) / n
		case Eq:
			if n == 1 {
				mustY = *e.x
			} else {
				fatal("cannot must equal != 0")
			}
		default:
			panic("invalid op " + e.Op)
		}
		tree.MustEvalTo(e.YID, mustY)
	} else { // e.x == nil
		var mustX int
		switch e.Op {
		case Plus:
			mustX = n - *e.y
		case Minus:
			mustX = n + *e.y
		case Mult:
			mustX = n / (*e.y)
		case Div:
			mustX = n * (*e.y)
		case Eq:
			if n == 1 {
				mustX = *e.y
			} else {
				fatal("cannot must equal != 0")
			}
		default:
			panic("invalid op " + e.Op)
		}
		tree.MustEvalTo(e.XID, mustX)
	}
}

func (e *ExprEvaler) Eval(tree Tree) (int, bool) {
	allOk := true
	if e.x == nil {
		v, ok := tree.Eval(e.XID)
		if !ok {
			allOk = false
		} else {
			e.x = &v
		}
	}
	if e.y == nil {
		v, ok := tree.Eval(e.YID)
		if !ok {
			allOk = false
		} else {
			e.y = &v
		}
	}
	if !allOk {
		return 0, false
	}

	switch e.Op {
	case Plus:
		return *e.x + *e.y, true
	case Minus:
		return *e.x - *e.y, true
	case Mult:
		return *e.x * *e.y, true
	case Div:
		return *e.x / *e.y, true
	case Eq:
		if *e.x == *e.y {
			return 1, true
		}
		return 0, true
	default:
		panic("invalid op " + e.Op)
	}
}

type Tree map[string]Evaler

func (tree Tree) Eval(id string) (int, bool) {
	e, ok := tree[id]
	if !ok {
		fatal("no such id %q", id)
	}
	return e.Eval(tree)
}

func (tree Tree) MustEvalTo(id string, n int) {
	e, ok := tree[id]
	if !ok {
		fatal("no such id %q", id)
	}
	e.MustEvalTo(tree, n)
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

		tree[id] = &ExprEvaler{
			XID: left,
			YID: right,
			Op:  op,
		}
	}
	return tree
}

func part1MainFunc(in string) (int, error) {
	tree := mustParseTree(in)
	res, ok := tree.Eval("root")
	if !ok {
		fatal("failed to eval root")
	}
	return res, nil
}

type HumanEvaler struct {
	Val int
}

func (e *HumanEvaler) Eval(Tree) (int, bool) {
	return 0, false
}

func (e *HumanEvaler) CanEval() bool {
	return false
}

func (e *HumanEvaler) MustEvalTo(tree Tree, n int) {
	e.Val = n
}

func part2MainFunc(in string) (int, error) {
	tree := mustParseTree(in)
	root := tree["root"]
	expRoot := root.(*ExprEvaler)
	expRoot.Op = Eq
	tree["root"] = expRoot

	humn := &HumanEvaler{}
	tree["humn"] = humn

	_, ok := tree.Eval("root")
	if ok {
		fatal("eval root should fail here")
	}
	tree.MustEvalTo("root", 1)

	res := humn.Val
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
