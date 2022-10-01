package day_10

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2017/testutil"
)

const inputTest = "3,4,1,5"

func listsEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, e1 := range l1 {
		if e1 != l2[i] {
			return false
		}
	}
	return true
}

func TestReverse(t *testing.T) {
	tests := []struct {
		pos    int
		length int
		list   []int
		result []int
	}{
		{
			pos:    2,
			length: 4,
			list:   []int{1, 2, 3, 4, 5, 6},
			result: []int{1, 2, 6, 5, 4, 3},
		},
		{
			pos:    0,
			length: 6,
			list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			result: []int{6, 5, 4, 3, 2, 1, 7, 8, 9},
		},
		{
			pos:    7,
			length: 6,
			list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, // 8 9 1 2 3 4
			result: []int{2, 1, 9, 8, 5, 6, 7, 4, 3}, // 4 3 2 1 9 8
		},
	}

	for i, test := range tests {
		rlist, err := reverseCopy(test.list, test.pos, test.length)
		testutil.CheckUnexpectedError(t, err)
		if !listsEqual(test.result, rlist) {
			t.Fatalf("%d: want %v, have %v", i, test.result, rlist)
		}
	}
}

func TestListProduct(t *testing.T) {
	res, err := listProduct(inputTest, 5)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 12
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestKnotHash(t *testing.T) {
	tests := []struct {
		in   string
		hash string
	}{
		{
			in:   "",
			hash: "a2582a3a0e66e6e86e3812dcb672a272",
		},
		{
			in:   "AoC 2017",
			hash: "33efeb34ea91902bb2f59c9920caa6cd",
		},
		{
			in:   "1,2,3",
			hash: "3efbe78a8d82f29979031a4aa0b16a9d",
		},
		{
			in:   "1,2,4",
			hash: "63960835bcdc130f0b66d7ff4f6a5a8e",
		},
	}

	for i, test := range tests {
		hash, err := knotHash(test.in)
		testutil.CheckUnexpectedError(t, err)
		if test.hash != hash {
			t.Fatalf("%d: want %v, have %v", i, test.hash, hash)
		}
	}
}
