package day_20

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/ring"
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

type Carrier struct {
	Index int
	Value int
}

func part1MainFunc(in string) (int, error) {
	ns, err := readutil.ReadInts(in, "\n")
	if err != nil {
		fatal("%v", err)
	}
	rg := ring.New[Carrier]()
	for ix, n := range ns {
		rg.Insert(Carrier{ix, n})
	}

	for ix := 0; ix < len(ns); ix++ {
		cn, ok := rg.FindNode(func(n *ring.Node[Carrier]) bool {
			return n.Value.Index == ix
		})
		if !ok {
			fatal("find carrier with index %d", ix)
		}
		if cn.Value.Value == 0 {
			continue
		}
		if cn.Value.Value > 0 {
			for i := 0; i < cn.Value.Value; i++ {
				rg.ShiftRight(cn)
			}
		} else { // < 0
			for i := 0; i < -cn.Value.Value; i++ {
				rg.ShiftLeft(cn)
			}
		}
	}
	cn, ok := rg.FindNode(func(n *ring.Node[Carrier]) bool {
		return n.Value.Value == 0
	})
	if !ok {
		fatal("find carrier zero ")
	}

	var v1, v2, v3 int
	for i := 0; i < 3000; i++ {
		cn = cn.Next
		switch i {
		case 999:
			v1 = cn.Value.Value
		case 1999:
			v2 = cn.Value.Value
		case 2999:
			v3 = cn.Value.Value
		}
	}

	sum := v1 + v2 + v3
	return sum, nil
}

func part2MainFunc(in string) (int, error) {
	decKey := 811589153
	ns, err := readutil.ReadInts(in, "\n")
	if err != nil {
		fatal("%v", err)
	}
	rg := ring.New[Carrier]()
	for ix, n := range ns {
		rg.Insert(Carrier{ix, n * decKey})
	}

	for i := 0; i < 10; i++ {
		for ix := 0; ix < len(ns); ix++ {
			cn, ok := rg.FindNode(func(n *ring.Node[Carrier]) bool {
				return n.Value.Index == ix
			})
			if !ok {
				fatal("find carrier with index %d", ix)
			}
			if cn.Value.Value == 0 {
				continue
			}
			if cn.Value.Value > 0 {
				cnt := cn.Value.Value % (len(ns) - 1)
				for i := 0; i < cnt; i++ {
					rg.ShiftRight(cn)
				}
			} else { // < 0
				cnt := (-cn.Value.Value) % (len(ns) - 1)
				for i := 0; i < cnt; i++ {
					rg.ShiftLeft(cn)
				}
			}
		}
	}

	cn, ok := rg.FindNode(func(n *ring.Node[Carrier]) bool {
		return n.Value.Value == 0
	})
	if !ok {
		fatal("find carrier zero ")
	}

	var v1, v2, v3 int
	for i := 0; i < 3000; i++ {
		cn = cn.Next
		switch i {
		case 999:
			v1 = cn.Value.Value
		case 1999:
			v2 = cn.Value.Value
		case 2999:
			v3 = cn.Value.Value
		}
	}

	sum := v1 + v2 + v3
	return sum, nil
}
