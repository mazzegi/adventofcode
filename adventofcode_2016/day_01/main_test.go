package main

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {

	var inputTest1 = `R2, L3`
	var inputTest2 = `R2, R2, R2`
	var inputTest3 = `R5, L5, R5, R3`

	tests := []struct {
		in   string
		dist int
	}{
		{
			in:   inputTest1,
			dist: 5,
		},
		{
			in:   inputTest2,
			dist: 2,
		},
		{
			in:   inputTest3,
			dist: 12,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			d := distance(test.in)
			if d != test.dist {
				t.Fatalf("want %d, have %d", test.dist, d)
			}
		})
	}
}

func TestDistanceVisitedTwice(t *testing.T) {

	var inputTest1 = `R8, R4, R4, R8`

	tests := []struct {
		in   string
		dist int
	}{
		{
			in:   inputTest1,
			dist: 4,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			d := distanceOfFirstVisitedTwice(test.in)
			if d != test.dist {
				t.Fatalf("want %d, have %d", test.dist, d)
			}
		})
	}
}
