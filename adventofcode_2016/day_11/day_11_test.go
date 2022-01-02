package day_11

import (
	"adventofcode_2016/testutil"
	"testing"
)

func TestPart1MainFunc(t *testing.T) {
	res, err := minSteps(initTestState())
	testutil.CheckUnexpectedError(t, err)
	var exp int = 11
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

func TestList(t *testing.T) {
	makeState := func(pathValue int) *state {
		return &state{pathValue: pathValue}
	}

	ol := &orderedList{}
	ol.add(makeState(1))
	ol.add(makeState(4))
	ol.add(makeState(3))
	ol.add(makeState(2))
	ol.add(makeState(0))
	ol.add(makeState(8))
	log("first: %d", ol.first.state.pathValue)
	log("last : %d", ol.last.state.pathValue)
	log(ol.dump())

	var ts *state
	for ol.size() > 0 {
		ts = ol.mustTakeFirst()
		log("took %d", ts.pathValue)
		if ol.first != nil {
			log("first: %d", ol.first.state.pathValue)
		}
		if ol.last != nil {
			log("last : %d", ol.last.state.pathValue)
		}
		log(ol.dump())
	}

}
