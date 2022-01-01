package day_11

import (
	"adventofcode_2016/testutil"
	"testing"
)

func initTestEnv() *environment {
	env := &environment{}
	env.floors = append(env.floors, &container{
		chips:      makeSet("hydrogen", "lithium"),
		generators: makeSet(),
	})
	env.floors = append(env.floors, &container{
		chips:      makeSet(),
		generators: makeSet("hydrogen"),
	})
	env.floors = append(env.floors, &container{
		chips:      makeSet(),
		generators: makeSet("lithium"),
	})
	env.floors = append(env.floors, &container{
		chips:      makeSet(),
		generators: makeSet(),
	})
	return env
}

func TestPart1MainFunc(t *testing.T) {
	res, err := minSteps(initTestEnv())
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

// func TestPart2MainFunc(t *testing.T) {
// 	res, err := part2MainFunc(inputTest)
// 	testutil.CheckUnexpectedError(t, err)
// 	var exp int = -42
// 	if exp != res {
// 		t.Fatalf("want %d, have %d", exp, res)
// 	}
// }
