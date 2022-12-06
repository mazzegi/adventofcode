package day_06

import (
	"fmt"
	"testing"
)

func TestPart1MainFunc(t *testing.T) {
	tests := []struct {
		in  string
		exp int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			res, _ := part1MainFunc(test.in)
			if test.exp != res {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}

func TestPart2MainFunc(t *testing.T) {
	tests := []struct {
		in  string
		exp int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("#%02d", i), func(t *testing.T) {
			res, _ := part2MainFunc(test.in)
			if test.exp != res {
				t.Fatalf("want %d, have %d", test.exp, res)
			}
		})
	}
}
