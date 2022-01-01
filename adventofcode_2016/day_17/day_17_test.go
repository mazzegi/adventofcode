package day_17

import (
	"adventofcode_2016/testutil"
	"testing"
)

func TestPart1MainFunc(t *testing.T) {
	res, err := shortestPath("ihgpwlah", 4, 4)
	testutil.CheckUnexpectedError(t, err)
	var exp string = "DDRRRD"
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}

	res, err = shortestPath("kglvqrro", 4, 4)
	testutil.CheckUnexpectedError(t, err)
	exp = "DDUDRLRRUDRD"
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}

	res, err = shortestPath("ulqzkmiv", 4, 4)
	testutil.CheckUnexpectedError(t, err)
	exp = "DRURDRUDDLLDLUURRDULRLDUUDDDRR"
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := longestPath("ihgpwlah", 4, 4)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 370
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}

	res, err = longestPath("kglvqrro", 4, 4)
	testutil.CheckUnexpectedError(t, err)
	exp = 492
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}

	res, err = longestPath("ulqzkmiv", 4, 4)
	testutil.CheckUnexpectedError(t, err)
	exp = 830
	if exp != res {
		t.Fatalf("want %q, have %q", exp, res)
	}
}

func TestOpenDoors(t *testing.T) {
	tests := []struct {
		passcode string
		path     string
		exp      dirs
	}{
		{
			passcode: "hijkl",
			path:     "",
			exp:      dirs{up, down, left},
		},
		{
			passcode: "hijkl",
			path:     "D",
			exp:      dirs{up, left, right},
		},
		{
			passcode: "hijkl",
			path:     "DR",
			exp:      dirs{},
		},
	}

	for i, test := range tests {
		res := openDoors(test.passcode, test.path)
		if res.String() != test.exp.String() {
			t.Fatalf("%d: want %q, have %q", i, test.exp.String(), res.String())
		}
	}
}
